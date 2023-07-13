package wsrouter

import (
	"main/service/auth"
	"main/service/echo"
	"main/service/ws"
)

// wsrouter is a service to route WebSocket services.
// New Websocket services should be registered here.
// A service requires a function to handle data from client, as `func(ws.WsData, *ws.Service, chan int)`

func Init() {
	routeFunc = make(map[string]func(ws.WsData, *ws.Service, chan int))

	// Insert WebSocket services below
	// please use the following format:
	routeFunc["echo"] = echo.EchoFunc
	// 1 -- echo

	routeFunc["login"] = auth.Login
	// 2 -- login

	// Insert WebSocket services here
}

func RouterResponser(h *ws.Service) {
	defer h.Waiter.Done()
	// cherr is a channel to handle error from foo().
	cherr := make(chan int, 8)
	register := make(map[string]func(ws.WsData, *ws.Service, chan int))

	for {
		select {
		case data := <-h.DataChan:
			switch data.Service {

			case "router":
				if data.Message == "register" {
					routerRegister(h, register, data)
				} else {
					h.Conn.WriteJSON(ws.WsData{
						Service: "router",
						Message: "Unknown command " + string(data.Message),
					})
				}

			default:
				foo, ok := register[data.Service]
				if !ok {
					h.Conn.WriteJSON(ws.WsData{
						Service: "router",
						Message: "Service " + data.Service + " not found ",
					})
					continue
				}
				// Run a goroutine to handle data from client.
				go func() {
					defer ws.PanicDetect(cherr, h, foo)
					foo(data, h, cherr)
				}()
			}

		// Handle error from foo().
		case errno := <-cherr:
			if errno == -1 {
				h.Abort()
				return
			}
		}
	}
}

func routerRegister(h *ws.Service, register map[string]func(ws.WsData, *ws.Service, chan int), data ws.WsData) {
	ok := false
	register[data.Data], ok = routeFunc[data.Data]
	if ok {
		h.Conn.WriteJSON(ws.WsData{
			Service: "router",
			Message: "Registered successfully to service " + string(data.Data),
		})
	} else {
		h.Conn.WriteJSON(ws.WsData{
			Service: "router",
			Message: "Requested service " + string(data.Data) + " not found ",
		})
	}
}

var routeFunc map[string]func(ws.WsData, *ws.Service, chan int)
