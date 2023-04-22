package echo

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"main/service/ws"
)

func echo(data ws.WsData, r *websocket.Conn, _ chan int) {
	fmt.Println("recv: ", data)
	data.BinData, _ = json.Marshal(data)
	r.WriteMessage(websocket.TextMessage, data.BinData)
}

func Echo(r *websocket.Conn) {
	h := ws.SetWsHandler(1, r, ws.JSONListener, ws.HandlerWrapper(echo))
	h.Start()
}
