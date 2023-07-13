package echo

import (
	"github.com/gin-gonic/gin"
	"main/service/ws"
)

func EchoFunc(data ws.WsData, ws *ws.Service, _ chan int) {
	ws.Conn.WriteJSON(data)
	ws.Conn.WriteJSON(gin.H{"service": "echo", "message": "debug", "data": *ws})
}
