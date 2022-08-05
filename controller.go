/**
* @program: web-rtc-server
*
* @description:
*
* @author: lemo
*
* @create: 2022-08-03 22:11
**/

package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"strconv"
	"time"

	"github.com/lemonyxk/kitty/v2/kitty"
	"github.com/lemonyxk/kitty/v2/socket"
	"github.com/lemonyxk/kitty/v2/socket/websocket/server"
	"github.com/lemonyxk/utils/v3"
)

func Login(stream *socket.Stream[server.Conn]) error {

	var user User
	_ = utils.Json.Decode(stream.Data, &user)

	if user.Name == "" {
		return nil
	}

	user.FD = stream.Conn.FD()

	AddUser(user.Name, &user)

	stream.Conn.Server().GetConnections(func(conn server.Conn) {
		_ = Server.Json(conn.FD(), Json{
			Event: "/UserList",
			Data:  GetAllUsers(),
		})
	})

	return Server.Json(stream.Conn.FD(), Json{
		Event: "/Login",
		Data:  nil,
	})
}

func CreateOffer(stream *socket.Stream[server.Conn]) error {

	var offer Offer
	_ = utils.Json.Decode(stream.Data, &offer)

	if offer.To == "" {
		return nil
	}

	var user = GetUserByName(offer.To)

	if user == nil {
		return nil
	}

	return Server.Json(user.FD, Json{
		Event: "/CreateOffer",
		Data:  offer,
	})
}

func CreateAnswer(stream *socket.Stream[server.Conn]) error {
	var answer Answer
	_ = utils.Json.Decode(stream.Data, &answer)

	if answer.To == "" {
		return nil
	}

	var user = GetUserByName(answer.To)

	if user == nil {
		return nil
	}

	return Server.Json(user.FD, Json{
		Event: "/CreateAnswer",
		Data:  answer,
	})
}

func RequestAccount(stream *socket.Stream[server.Conn]) error {
	username := "account"
	expired := time.Now().Unix() + 9999
	username = strconv.Itoa(int(expired)) + ":" + username
	key := secret

	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(username))
	var password = base64.StdEncoding.EncodeToString(h.Sum(nil))

	return Server.Json(stream.Conn.FD(), Json{
		Event: "/RequestAccount",
		Data:  kitty.M{"password": password, "account": username},
	})
}
