package websockets

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

//Socket represents a websocket connection
type Socket struct {
	ws *websocket.Conn
}

//Execute opens a new websocket connection and executes the given function, after which the socket is closed
func Execute(w http.ResponseWriter, r *http.Request, f func(s Socket) error) error {
	var s Socket
	var err error
	s.ws, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	if err := f(s); err != nil {
		return err
	}
	return s.ws.Close()
}

//Read waits for a JSON message and decodes it into the given value.
func (s Socket) Read(value interface{}) error {
	return s.ws.ReadJSON(value)
}

//Write sends the given value as JSON over the websocket.
func (s Socket) Write(value interface{}) error {
	return s.ws.WriteJSON(value)
}
