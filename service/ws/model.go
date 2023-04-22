package ws

import (
	"github.com/gorilla/websocket"
	"sync"
)

type WsData struct {
	Typ     int
	Service string `json:"service"`
	Message string `json:"message"`
	BinData []byte
}

type HandlerFunc func(handler *WsHandler)

// ResponseFunc data from listener, ws connection, error channel.
type ResponseFunc func(WsData, *websocket.Conn, chan int)

type WsHandler struct {
	Status          int
	Conn            *websocket.Conn
	DataChan        chan WsData
	ClientListener  HandlerFunc
	ClientResponser HandlerFunc
	Waiter          *sync.WaitGroup
}

const StatusNotRouted = 101
const StatusStopped = -1
const StatusRouted = 0
const StatusWorking = 1
