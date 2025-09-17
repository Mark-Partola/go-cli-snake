package socket

const (
	StartGameCommand     = "StartGame"
	InvalidCommand       = "InvalidCommand"
	UnprocessableCommand = "UnprocessableCommand"
	InternalServerError  = "InternalServerError"
)

type Request[P any] struct {
	Type    string `json:"type"`
	Payload P      `json:"payload"`
}

type Response struct {
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}

type StartGameRequestPayload struct {
	Type     string `json:"type"`
	UserName string `json:"username"`
}

type StartGameResponsePayload struct {
	RoomId string `json:"roomId"`
}

type StartGameRequest = Request[StartGameRequestPayload]
