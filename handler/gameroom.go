package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"main/service/gameroom"
)

func NewRoom(r *gin.Context) {
	upgrader := websocket.Upgrader{}
	c, _ := upgrader.Upgrade(r.Writer, r.Request, nil)
	defer c.Close()
	gameroom.HandleGameRoom(c)
}
