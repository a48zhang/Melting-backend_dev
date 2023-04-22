package gameroom

import (
	"github.com/gorilla/websocket"
	"main/service/ws"
)

func NewRoom(r *websocket.Conn) {
	ws.SetWsHandler(1, r, nil, nil).Start()
}
