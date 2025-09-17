package socket

import (
	"encoding/json"
	"quiz-server/internal/system/logger"

	"github.com/gorilla/websocket"
)

type Client struct {
	log  logger.Logger
	conn *websocket.Conn
}

func NewClient(conn *websocket.Conn) *Client {
	return &Client{
		log:  logger.Get("client"),
		conn: conn,
	}
}

func (c *Client) Read() (*Request[any], error) {
	_, message, err := c.conn.ReadMessage()
	if err != nil {
		c.log.Error("read:", err)
		return nil, err
	}

	var request Request[any]
	if err := json.Unmarshal(message, &request); err != nil {
		c.log.Error("unmarshal:", err)
		return nil, err
	}

	return &request, nil
}

func (c *Client) Write(response Response) error {
	message, err := json.Marshal(response)
	if err != nil {
		c.log.Error("marshal:", err)
		return err
	}

	if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
		c.log.Error("write:", err)
		return err
	}

	return nil
}
