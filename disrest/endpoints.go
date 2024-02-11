package disrest

var APIVersion = "10"

var (
	EndpointDiscord    = "https://discord.com/"
	EndpointAPI        = EndpointDiscord + "api/v" + APIVersion + "/"
	EndpointGatewayBot = EndpointAPI + "/gateway/bot"
)
