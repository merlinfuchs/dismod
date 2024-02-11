package disway

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
	"github.com/merlinfuchs/dismod/distype"
	"github.com/merlinfuchs/dismod/disutil"
)

const apiVersion = "10"

type Shard struct {
	sync.RWMutex

	Log      disutil.Logger
	Dispatch func(t distype.EventType, d interface{})

	Token      string
	ShardCount int
	ShardID    int
	Gateway    string

	// used to make sure gateway websocket writes do not happen concurrently
	wsMutex   sync.Mutex
	ws        *websocket.Conn
	listening chan struct{}

	ResumeGateway     string
	Sequence          atomic.Int64
	SessionID         string
	LastHeartbeatAck  time.Time
	LastHeartbeatSent time.Time
}

func NewShard(token string) *Shard {
	return &Shard{
		Log:              disutil.DefaultLogger,
		Token:            token,
		LastHeartbeatAck: time.Now().UTC(),
	}
}

func (s *Shard) GatewayURL() string {
	var base string
	if s.ResumeGateway != "" {
		base = s.ResumeGateway
	} else {
		base = s.Gateway
	}

	if !strings.HasSuffix(base, "/") {
		base += "/"
	}

	return base + "?v=" + apiVersion + "&encoding=json"
}

func (s *Shard) Open() error {
	s.Lock()
	defer s.Unlock()

	s.listening = make(chan struct{})

	header := http.Header{}
	header.Add("accept-encoding", "zlib")

	var err error
	s.ws, _, err = websocket.DefaultDialer.Dial(s.GatewayURL(), header)
	if err != nil {
		return err
	}

	s.ws.SetCloseHandler(func(code int, text string) error {
		return nil
	})

	defer func() {
		// because of this, all code below must set err to the error
		if err != nil {
			s.ws.Close()
			s.ws = nil
		}
	}()

	var hello *distype.HelloData
	hello, err = s.readHello()
	if err != nil {
		return err
	}

	if err = s.identifyOrResume(); err != nil {
		return err
	}

	if err = s.readReadyOrResumed(); err != nil {
		return err
	}

	s.dispatch(distype.EventTypeConnect, &distype.ConnectEvent{})

	go s.listen(s.ws)
	go s.heartbeat(time.Duration(hello.HeartbeatInterval) * time.Millisecond)

	return nil
}

func (s *Shard) Close() error {
	return s.CloseWithCode(websocket.CloseNormalClosure)
}

func (s *Shard) CloseWithCode(closeCode int) (err error) {
	s.Lock()
	defer s.Unlock()

	if s.listening != nil {
		s.Log(disutil.LogInfo, "closing listening channel")
		close(s.listening)
		s.listening = nil
	}

	if s.ws != nil {
		s.Log(disutil.LogInfo, "sending close frame")
		// To cleanly close a connection, a client should send a close
		// frame and wait for the server to close the connection.
		err := s.sendMessage(websocket.CloseMessage, websocket.FormatCloseMessage(closeCode, ""))
		if err != nil {
			s.Log(disutil.LogInfo, "error closing websocket, %w", err)
		}

		time.Sleep(1 * time.Second)

		s.Log(disutil.LogInfo, "closing gateway websocket")
		err = s.ws.Close()
		if err != nil {
			s.Log(disutil.LogInfo, "error closing websocket, %w", err)
		}

		s.ws = nil
	}

	s.Log(disutil.LogDebug, "emit disconnect event")
	s.dispatch(distype.EventTypeDisconnect, &distype.DisconnectEvent{})

	return
}

func (s *Shard) reconnect() {
	wait := time.Duration(1)

	for {
		s.Log(disutil.LogInfo, "trying to reconnect to gateway")

		err := s.Open()
		if err == nil {
			s.Log(disutil.LogInfo, "successfully reconnected to gateway")
			return
		}

		// Certain race conditions can call reconnect() twice. If this happens, we
		// just break out of the reconnect loop
		if err == ErrWSAlreadyOpen {
			s.Log(disutil.LogInfo, "Websocket already exists, no need to reconnect")
			return
		}

		s.Log(disutil.LogError, "error reconnecting to gateway, %s", err)

		<-time.After(wait * time.Second)
		wait *= 2
		if wait > 600 {
			wait = 600
		}
	}
}

func (s *Shard) readHello() (*distype.HelloData, error) {
	mt, m, err := s.ws.ReadMessage()
	if err != nil {
		return nil, err
	}

	e, err := s.handleMessage(mt, m)
	if err != nil {
		return nil, err
	}

	if e.Operation != distype.GatewayOpcodeHello {
		return nil, fmt.Errorf("expected Op 10 (hello), got %d instead", e.Operation)
	}

	return e.Data.(*distype.HelloData), nil
}

func (s *Shard) readReadyOrResumed() error {
	mt, m, err := s.ws.ReadMessage()
	if err != nil {
		return err
	}

	e, err := s.handleMessage(mt, m)
	if err != nil {
		return err
	}

	if e.Operation != distype.GatewayOpcodeDispatch {
		return fmt.Errorf("expected Op 0 (dispatch), got %d instead", e.Operation)
	}

	if e.Type != distype.EventTypeReady && e.Type != distype.EventTypeResumed {
		return fmt.Errorf("expected event type READY or RESUMED, got %s instead", e.Type)
	}

	if e.Type == distype.EventTypeReady {
		d := e.Data.(*distype.ReadyEvent)
		s.SessionID = d.SessionID
		s.ResumeGateway = d.ResumeGatewayURL
	}

	return nil
}

func (s *Shard) identifyOrResume() error {
	if s.SessionID == "" {
		fmt.Println("identify")
		return s.identify()
	} else {
		fmt.Println("resume")
		return s.resume()
	}
}

func (s *Shard) identify() error {
	compress := true
	err := s.sendOP(distype.GatewayOpcodeIdentify, distype.IdentifyData{
		Token:    s.Token,
		Compress: &compress,
		Properties: distype.ConnectionProperties{
			OS:      runtime.GOOS,
			Browser: "dismod/disway",
		},
		Intents: 1, // TODO
		Shard:   &[2]int{s.ShardID, s.ShardCount},
	})
	if err != nil {
		return fmt.Errorf("error sending identify packet, %w", err)
	}
	return nil
}

func (s *Shard) resume() error {
	err := s.sendOP(distype.GatewayOpcodeResume, distype.ResumeData{
		Token:     s.Token,
		SessionID: s.SessionID,
		Sequence:  s.Sequence.Load(),
	})
	if err != nil {
		return fmt.Errorf("error sending resume packet, %w", err)
	}
	return nil
}

func (s *Shard) listen(ws *websocket.Conn) {
	for {
		mt, m, err := ws.ReadMessage()
		if err != nil {
			s.RLock()
			sameConnection := s.ws == ws
			s.RUnlock()

			if sameConnection {
				s.Log(disutil.LogWarn, "error reading from websocket, reconnecting: %w", err)
				s.Close()
				s.reconnect()
			}
			return
		}

		_, err = s.handleMessage(mt, m)
		if err != nil {
			s.Log(disutil.LogError, "failed to handle message: %w", err)
		}

		select {
		case <-s.listening:
			return
		default:
			continue
		}
	}
}

func (s *Shard) heartbeat(heartbeatInterval time.Duration) {
	ticker := time.NewTicker(heartbeatInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			sequence := s.Sequence.Load()

			err := s.sendOP(distype.GatewayOpcodeHeartbeat, distype.HeartbeatData(sequence))
			if err != nil {
				s.Log(disutil.LogError, "error sending heartbeat, %w", err)
				s.Close()
				s.reconnect()
				return
			}

			s.Lock()
			s.LastHeartbeatSent = time.Now().UTC()
			s.Unlock()

			if time.Now().UTC().Sub(s.LastHeartbeatAck) > 5*heartbeatInterval {
				s.Log(disutil.LogError, "heartbeat ACK not received in time, closing and reconnecting shard")
				s.Close()
				s.reconnect()
				return
			}
		case <-s.listening:
			return
		}
	}
}

func (s *Shard) handleMessage(messageType int, msg []byte) (*distype.GatewayEvent, error) {
	var reader io.Reader = bytes.NewReader(msg)

	if messageType == websocket.BinaryMessage {
		z, err := zlib.NewReader(reader)
		if err != nil {
			s.Log(disutil.LogError, "error uncompressing websocket message, %w", err)
			return nil, err
		}

		defer func() {
			err := z.Close()
			if err != nil {
				s.Log(disutil.LogError, "error closing zlib, %w", err)
			}
		}()

		reader = z
	}

	var e *distype.GatewayEvent
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&e); err != nil {
		s.Log(disutil.LogError, "error decoding websocket message, %w", err)
		return e, err
	}

	s.Log(disutil.LogDebug, "Op: %d, Seq: %d, Type: %s\n\n", e.Operation, e.Sequence, e.Type)

	switch e.Operation {
	case distype.GatewayOpcodeDispatch:
		s.Sequence.Store(e.Sequence)
		s.dispatch(e.Type, e.Data)
		return e, nil
	case distype.GatewayOpcodeHeartbeat:
		err := s.sendOP(distype.GatewayOpcodeHeartbeat, e.Data)
		if err != nil {
			s.Log(disutil.LogError, "error sending heartbeat, %w", err)
			return e, err
		}

		return e, nil
	case distype.GatewayOpcodeReconnect:
		s.Log(disutil.LogInfo, "received reconnect opcode, closing and reconnecting shard")
		s.CloseWithCode(websocket.CloseServiceRestart)
		s.reconnect()
		return e, nil
	case distype.GatewayOpcodeInvalidSession:
		s.Log(disutil.LogInfo, "sending identify packet to gateway in response to Op9")
		if !*e.Data.(*bool) {
			s.SessionID = ""
			s.ResumeGateway = ""
		}
		s.CloseWithCode(websocket.ClosePolicyViolation)
		s.reconnect()
		return e, nil
	case distype.GatewayOpcodeHello:
		return e, nil
	case distype.GatewayOpcodeHeartbeatAck:
		s.Lock()
		s.LastHeartbeatAck = time.Now().UTC()
		s.Unlock()
		s.Log(disutil.LogDebug, "received heartbeat ACK")
		return e, nil
	default:
		s.Log(disutil.LogError, "received unexpected opcode: %d", e.Operation)
		return e, fmt.Errorf("received unexpected opcode: %d", e.Operation)
	}
}

func (s *Shard) dispatch(t distype.EventType, d interface{}) {
	if s.Dispatch != nil {
		s.Dispatch(t, d)
	}
}

func (s *Shard) sendOP(op distype.GatewayOpcode, data interface{}) error {
	body, err := json.Marshal(distype.GatewayEvent{
		Operation: op,
		Data:      data,
	})
	if err != nil {
		return err
	}

	return s.sendMessage(websocket.TextMessage, body)
}

func (s *Shard) sendMessage(messageType int, msg []byte) error {
	s.wsMutex.Lock()
	defer s.wsMutex.Unlock()

	return s.ws.WriteMessage(messageType, msg)
}
