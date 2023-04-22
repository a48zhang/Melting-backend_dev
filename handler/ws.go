package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"main/service/wsrouter"
	"net/http"
)

func NewWebSocket(r *gin.Context) {
	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}
	c, err := upgrader.Upgrade(r.Writer, r.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	wsrouter.RouteWsService(c)(c)
	defer c.Close()
}
