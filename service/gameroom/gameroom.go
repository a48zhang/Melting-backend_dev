package gameroom

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
)

func clientListener(r *websocket.Conn, ch chan []byte) {
	for {
		_, b, err := r.ReadMessage()
		if err != nil {
			fmt.Println(err)
			close(ch)
			return
		}
		ch <- b
	}
}

type tt struct {
	Message string `json:"message"`
}

func HandleGameRoom(r *websocket.Conn) {
	message := make(chan []byte, 4096)
	go clientListener(r, message)
	for data := range message {
		x := tt{}
		json.Unmarshal(data, &x)
		fmt.Println(x.Message)
	}

}
