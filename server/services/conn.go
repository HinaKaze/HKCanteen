package chat

import (
	"errors"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	WSReadWait     = 30 * time.Second
	WSWriteWait    = 10 * time.Second
	WSPingDuration = (7 * WSReadWait) / 10
)

type Connection struct {
	WS        *websocket.Conn
	SendChan  chan []byte
	CloseChan chan []byte
}

func (c *Connection) Start() {
	defer func() {
		c.WS.SetWriteDeadline(time.Now().Add(WSWriteWait))
		c.WS.WriteMessage(websocket.CloseMessage, []byte{})
		c.WS.Close()
	}()
	go c.Read()
	c.WS.SetReadLimit(1024)
	pingTicker := time.Tick(WSPingDuration)
	for {
		select {
		case m := <-c.SendChan:
			c.SendText(m)
		case <-pingTicker:
			c.Ping([]byte{})
		case <-c.CloseChan:
			break
		}
	}
}

func (c *Connection) send(messageType int, message []byte) {
	err := c.WS.SetWriteDeadline(time.Now().Add(WSWriteWait))
	if err != nil {
		panic(err.Error())
	}
	err = c.WS.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		panic(err.Error())
	}
}

func (c *Connection) SendText(message []byte) {
	c.send(websocket.TextMessage, message)
	return
}

func (c *Connection) Ping(message []byte) {
	c.send(websocket.PingMessage, []byte{})
	return
}

func (c *Connection) Pong(message []byte) {
	c.send(websocket.PongMessage, []byte{})
	return
}

func (c *Connection) Read() {
	defer func() {
		c.WS.WriteMessage(websocket.CloseMessage, []byte{})
		hub.Unregister <- c
		c.WS.Close()
	}()
	for {
		err := c.WS.SetReadDeadline(time.Now().Add(WSReadWait))
		if err != nil {
			panic(err.Error())
		}
		mType, message, err := c.WS.ReadMessage()
		if err != nil {
			panic(err.Error())
		}
		switch mType {
		case websocket.CloseMessage:
			log.Printf("WebSocket remote host close this connection [%s]\n", c.WS.RemoteAddr().String())
			break
		case websocket.TextMessage:
			hub.BroadCast <- message
		case websocket.PingMessage:
			c.Ping([]byte{})
		case websocket.PongMessage:
			c.Pong([]byte{})
		default:
			log.Printf("WebSocket drop some wrong message [%s] with wrong message [%d]\n", string(message), mType)
		}
	}
}

func CreateConnection(wsconn *websocket.Conn) (conn Connection) {
	if wsconn == nil {
		panic(errors.New("web socket is nil"))
	}
	return Connection{
		WS:        wsconn,
		SendChan:  make(chan []byte),
		CloseChan: make(chan []byte),
	}
}
