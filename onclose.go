/**
* @program: web-rtc-server
*
* @description:
*
* @author: lemo
*
* @create: 2022-08-03 22:12
**/

package main

import (
	"log"

	"github.com/lemonyxk/kitty/v2/socket/websocket/server"
)

func Close(conn server.Conn) {
	log.Println(conn.FD(), "close")

	DeleteUserByFD(conn.FD())

	conn.Server().GetConnections(func(conn server.Conn) {
		_ = Server.Json(conn.FD(), Json{
			Event: "/UserList",
			Data:  GetAllUsers(),
		})
	})
}
