package main

import (
	"fmt"
	"net/http"
	"quiz-server/internal/system/config"
	"quiz-server/internal/system/logger"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleSocketConnections(w http.ResponseWriter, r *http.Request) {
	log := logger.Get("websocket")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("upgrade:", err)
		return
	}
	defer ws.Close()

	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Error("read:", err)
			break
		}

		log.Info("received: %s", message)

		if err := ws.WriteMessage(messageType, message); err != nil {
			log.Error("write:", err)
			break
		}
		log.Info("sent: %s", message)
	}
}

func main() {
	cfg := config.Setup()
	log := logger.Setup(cfg.Env)
	log.Info("setting up, env: %s", cfg.Env)

	http.HandleFunc("/ws", handleSocketConnections)

	address := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Info("server started on %s", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Error("listen: ", err)
	}
}
