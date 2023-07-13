package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"main/service/ws"
	"main/service/wsrouter"
	"net/http"
)

// NewWebSocket is a handler to upgrade http connection to websocket, and route the connection to wsrouter.
func NewWebSocket(r *gin.Context) {
	// upgrade to websocket
	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}
	c, err := upgrader.Upgrade(r.Writer, r.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close()
	// use default router
	ws.NewWsHandler(ws.StatusNotRouted, c, ws.JSONListener, wsrouter.RouterResponser).Start()
}
