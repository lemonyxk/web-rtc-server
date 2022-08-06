/**
* @program: web-rtc-server
*
* @description:
*
* @author: lemo
*
* @create: 2022-08-03 22:10
**/

package main

import (
	"github.com/lemonyxk/kitty/v2/errors"
	"github.com/lemonyxk/kitty/v2/router"
	"github.com/lemonyxk/kitty/v2/socket"
	"github.com/lemonyxk/kitty/v2/socket/websocket/server"
	"github.com/lemonyxk/utils/v3"
)

func Before(stream *socket.Stream[server.Conn]) error {

	var user User

	_ = utils.Json.Decode(stream.Data, &user)

	if user.Name == "" {
		return errors.New("please login first")
	}

	return nil
}

func Router(r *router.Router[*socket.Stream[server.Conn]]) {
	r.Group().Handler(func(handler *router.Handler[*socket.Stream[server.Conn]]) {
		handler.Route("/login").Handler(Login)
	})

	r.Group().Before(Before).Handler(func(handler *router.Handler[*socket.Stream[server.Conn]]) {
		handler.Route("/createOffer").Handler(CreateOffer)
		handler.Route("/createAnswer").Handler(CreateAnswer)
		handler.Route("/addAnswer").Handler(AddAnswer)
		handler.Route("/endCall").Handler(EndCall)
		handler.Route("/requestAccount").Handler(RequestAccount)
	})
}
