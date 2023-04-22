package ws

import (
	"github.com/gorilla/websocket"
	"sync"
)

// Start It blocks. When this function end, the websocket may be closed safely.
func (h *WsHandler) Start() {
	h.Waiter.Add(2)
	go h.ClientListener(h)
	go h.ClientResponser(h)
	h.Waiter.Wait()
	h.Status = StatusStopped
}

func (h *WsHandler) Abort() {
	close(h.DataChan)
	h.Status = StatusStopped
}

func SetWsHandler(status int, r *websocket.Conn, listener HandlerFunc, responser HandlerFunc) *WsHandler {
	if listener == nil {
		listener = JSONListener
	}
	if responser == nil {
		responser = HandlerWrapper(func(WsData, *websocket.Conn, chan int) {})
	}
	return &WsHandler{
		Status:          status,
		Conn:            r,
		DataChan:        make(chan WsData, 32),
		ClientListener:  listener,
		ClientResponser: responser,
		Waiter:          &sync.WaitGroup{},
	}
}

// JSONListener Read everything from connection and write to channel.
func JSONListener(ws *WsHandler) {
	defer ws.Waiter.Done()
	for {
		var b WsData
		err := ws.Conn.ReadJSON(&b)
		if err != nil {
			close(ws.DataChan)
			return
		}
		ws.DataChan <- b
	}
}

// HandlerWrapper return a ClientResponser. Data from client will be handled by foo() asynchronously.
func HandlerWrapper(foo ResponseFunc) HandlerFunc {
	return func(ws *WsHandler) {
		defer ws.Waiter.Done()
		cherr := make(chan int, 8)

		for {
			select {
			case data := <-ws.DataChan:
				go foo(data, ws.Conn, cherr)

			case errno := <-cherr:
				if errno == -1 {
					ws.Abort()
					return
				}
			}
		}
	}
}
