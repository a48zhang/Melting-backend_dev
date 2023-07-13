package ws

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"reflect"
	"runtime"
	"sync"
)

// Start It blocks. When this function end, the websocket may be closed safely.
func (h *Service) Start() {
	h.Waiter.Add(2)
	go h.ClientListener(h)
	go h.ClientResponser(h)
	h.Status = StatusWorking
	h.Waiter.Wait()
	h.Status = StatusStopped
}

// Abort Close the channel to stop handler goroutines.
func (h *Service) Abort() {
	if h.Status == StatusClosing {
		return
	}
	close(h.DataChan)
	h.Status = StatusClosing
}

func NewWsHandler(status int, r *websocket.Conn, listener HandlerFunc, responser HandlerFunc) *Service {
	if listener == nil {
		// use default listener
		listener = JSONListener
	}
	if responser == nil {
		// use default responser
		responser = HandlerWrapper(func(ws WsData, r *Service, _ chan int) {
			r.Conn.WriteJSON(WsData{Service: "Server", Message: "[Error] Service don't have a responser."})
		})
	}
	return &Service{
		Status:          status,
		Conn:            r,
		DataChan:        make(chan WsData, 32),
		ClientListener:  listener,  // ClientListener you may use a HandlerWrapper with certain func, or a custom HandlerFunc.
		ClientResponser: responser, // ClientResponser you may use a JSONListener.
		Waiter:          &sync.WaitGroup{},
		context:         make(map[string]string),
	}
}

// JSONListener Read everything from connection and write to channel.
func JSONListener(ws *Service) {
	defer ws.Waiter.Done()
	for {
		var b WsData
		err := ws.Conn.ReadJSON(&b)
		if err != nil {
			ws.Abort()
			return
		}
		// If connection is closing, stop reading.
		if ws.Status == StatusClosing {
			return
		} else {
			ws.DataChan <- b
		}
	}
}

func PanicDetect(cherr chan int, ws *Service, foo ResponseFunc) {
	if r := recover(); r != nil {
		// Return a panic message to client. Including function name and error message.
		err := fmt.Sprint("[Panic] Websocket Connection has been aborted", runtime.FuncForPC(reflect.ValueOf(foo).Pointer()).Name(), r)
		log.Println(err)
		ws.Conn.WriteJSON(WsData{
			Typ:     500,
			Service: "Server Error",
			Message: err,
			Data:    "",
		})
		cherr <- -1
		return
	}
}

// HandlerWrapper return a ClientResponser. Data from client will be handled by foo() asynchronously.
func HandlerWrapper(foo ResponseFunc) HandlerFunc {
	return func(ws *Service) {
		defer ws.Waiter.Done()
		// cherr is a channel to handle error from foo().
		cherr := make(chan int, 8)

		for {
			select {

			// Run a goroutine to handle data from client.
			case data := <-ws.DataChan:
				go func() {
					defer PanicDetect(cherr, ws, foo)
					foo(data, ws, cherr)
				}()

			// Handle error from foo().
			case errno := <-cherr:
				if errno == -1 {
					ws.Abort()
					return
				}
			}
		}
	}
}
