package socket

import (
	"crypto/rand"
	"quiz-server/internal/system/logger"
)

func Process(request *Request[any]) Response {
	log := logger.Get("process")

	switch request.Type {
	case StartGameCommand:
		{
			log.Info("start game")

			return Response{
				Type: StartGameCommand,
				Payload: StartGameResponsePayload{
					RoomId: rand.Text(),
				},
			}
		}
	default:
		log.Warn("unknown command: %s", request.Type)
		return Response{
			Type: InvalidCommand,
		}
	}
}
