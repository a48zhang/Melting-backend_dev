package wsrouter

import (
	"fmt"
	"github.com/gorilla/websocket"
	"main/service/echo"
	"main/service/ws"
	"strconv"
)

// RouteWsService Routes websocket to the correct place.
func RouteWsService(r *websocket.Conn) func(*websocket.Conn) {
	for {
		var data ws.WsData
		r.ReadJSON(&data)
		if data.Service != "router" {
			r.WriteJSON(ws.WsData{
				Typ:     websocket.TextMessage,
				Service: "router",
				Message: "Not routed",
				BinData: nil,
			})
			continue
		}
		i, _ := strconv.Atoi(data.Message)
		if v, ok := routePath[i]; ok {
			r.WriteJSON(ws.WsData{
				Typ:     websocket.TextMessage,
				Service: "router",
				Message: "routed successfully to service " + data.Message,
				BinData: nil,
			})
			return v
		} else {
			r.WriteJSON(ws.WsData{
				Typ:     websocket.TextMessage,
				Service: "router",
				Message: "Service " + data.Message + " not found ",
				BinData: nil,
			})
		}
	}
}

var routePath map[int]func(*websocket.Conn)

func Init() {
	// Insert WebSocket services here
	routePath = make(map[int]func(*websocket.Conn))
	routePath[1] = echo.Echo
	fmt.Println(routePath[1])
	// 1 -- echo
}
