package distype

import "strconv"

type Permissions string

func (p Permissions) Parse() uint64 {
	v, _ := strconv.ParseUint(string(p), 10, 64)
	return v
}

func (p Permissions) String() string {
	return string(p)
}

func (p Permissions) Has(permission Permission) bool {
	return p.Parse()&permission != 0
}

type Permission = uint64

const (
	PermissionCreateInstantInvite Permission = 0x00000001 // TODO
)
