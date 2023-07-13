package auth

import (
	"main/service/ws"
	"strconv"
)

// Login is a function to handle login request. Require a token in data.Data.
func Login(data ws.WsData, h *ws.Service, _ chan int) {
	c := h.Conn
	if data.Message == "login" {
		_, m, err := Parsetoken(data.Data)
		if err != nil {
			c.WriteJSON(ws.WsData{
				Service: "login",
				Message: "login failed",
			})
			return
		}
		h.Set("uid", strconv.Itoa(m.UID))
		c.WriteJSON(ws.WsData{
			Service: "login",
			Message: "login successfully",
			Data:    strconv.Itoa(m.UID),
		})
	}
}
