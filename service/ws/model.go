package ws

import (
	"github.com/gorilla/websocket"
	"sync"
)

type WsData struct {
	Typ     int
	Service string `json:"service"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type HandlerFunc func(*Service)

// ResponseFunc Args: data from listener, ws connection, error channel.
type ResponseFunc func(WsData, *Service, chan int)

type Service struct {
	Status          int
	Conn            *websocket.Conn
	DataChan        chan WsData
	ClientListener  HandlerFunc
	ClientResponser HandlerFunc
	Waiter          *sync.WaitGroup
	Lock            *sync.Mutex

	context map[string]string
}

func (h *Service) Get(key string) string {
	h.Lock.Lock()
	defer h.Lock.Unlock()
	return h.context[key]
}

func (h *Service) Set(key, value string) {
	h.Lock.Lock()
	defer h.Lock.Unlock()
	h.context[key] = value
}

const StatusNotRouted = 101
const StatusStopped = -1
const StatusRouted = 0
const StatusWorking = 1
const StatusClosing = -3
