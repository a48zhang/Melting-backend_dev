package gameroom

import (
	"main/service/ws"
	"math/rand"
	"time"
)

type room struct {
	ID         int
	LastUpdate time.Time
	Users      []int
	Roles      map[int]string
}

var roomMap = make(map[int]room)

func Entrypoint(data ws.WsData, ws *ws.Service, _ chan int) {
	for {
		select {
		case data := <-ws.DataChan:
			switch data.Service {
			case "new":
				// TODO
			}

		}
	}
}

// TODO
func newRoom(ws.WsData, *ws.Service, chan int) {
	rand.Seed(time.Now().Unix())
	id := rand.Int()
	for _, ok := roomMap[id]; ok; {
		if roomMap[id].LastUpdate.Before(time.Now().Add(-time.Minute * 10)) {
			delete(roomMap, id)
			break
		}
		id = rand.Int()
	}
	roomMap[id] = room{
		ID:         id,
		LastUpdate: time.Now(),
	}

}
