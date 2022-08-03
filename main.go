package main

import (
	"log"
	"os"
	"time"

	"github.com/json-iterator/go"
	"github.com/lemonyxk/kitty/v2"
	"github.com/lemonyxk/kitty/v2/socket"
	"github.com/lemonyxk/kitty/v2/socket/websocket/server"
	"github.com/lemonyxk/utils/v3"
)

var Server = &WebSocket{}

func main() {

	Server = &WebSocket{}

	Server.Addr = "0.0.0.0:8667"

	Server.HeartBeatTimeout = time.Second * 30

	Server.OnOpen = Open

	Server.OnClose = Close

	Server.OnError = Error

	// handle unknown proto
	Server.OnUnknown = func(conn server.Conn, message []byte, next server.Middle) {
		var j = jsoniter.Get(message)
		var route = j.Get("event").ToString()
		var data = j.Get("data").ToString()
		if route == "" {
			return
		}

		next(&socket.Stream[server.Conn]{Conn: conn, Event: route, Data: []byte(data)})
	}

	Server.OnSuccess = func() {
		log.Println("server started success")
	}

	var wsServerRouter = kitty.NewWebSocketServerRouter()

	Router(wsServerRouter)

	Server.SetRouter(wsServerRouter).Start()

	utils.Signal.ListenKill().Done(func(sig os.Signal) {
		log.Println("received signal:", sig.String())
	})
}
