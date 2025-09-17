package main

import (
	"fmt"
	"net/http"
	"quiz-server/internal/api/socket"
	"quiz-server/internal/system/config"
	"quiz-server/internal/system/logger"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func init() {
	cfg := config.Setup()
	log := logger.Setup(cfg.Env)
	log.Info("setting up, env: %s", cfg.Env)
}

func handleSocketConnections(w http.ResponseWriter, r *http.Request) {
	log := logger.Get("websocket")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("upgrade:", err)
		return
	}
	defer ws.Close()

	log.Info("new connection %s", ws.RemoteAddr())
	client := socket.NewClient(ws)
	err = handle(client)
	if err != nil {
		log.Error("handle:", err)
		client.Write(socket.Response{
			Type: socket.InternalServerError,
		})
	}
}

func handle(client *socket.Client) error {
	for {
		request, err := client.Read()
		if err != nil {
			return err
		}

		response := socket.Process(request)
		err = client.Write(response)
		if err != nil {
			return err
		}
	}
}

func main() {
	cfg := config.Get()
	log := logger.Get("setup")

	http.HandleFunc("/ws", handleSocketConnections)

	address := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Info("server started on %s", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Error("listen: ", err)
	}
}
