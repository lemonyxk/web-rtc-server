package main

import (
	"github.com/json-iterator/go"
	"github.com/lemonyxk/kitty/v2/errors"
	"github.com/lemonyxk/kitty/v2/socket/protocol"
	"github.com/lemonyxk/kitty/v2/socket/websocket/server"
)

type WebSocket struct {
	server.Server
}

type Json struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

func (w *WebSocket) Json(fd int64, pack Json) error {

	var conn = w.Server.GetConnection(fd)

	if conn == nil {
		return errors.ConnNotFount
	}

	data, err := jsoniter.Marshal(pack)
	if err != nil {
		return err
	}

	_, err = conn.Write(int(protocol.Text), data)

	return err
}
