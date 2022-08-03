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

func Open(conn server.Conn) {
	log.Println(conn.FD(), "open")
}
