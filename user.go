/**
* @program: web-rtc-server
*
* @description:
*
* @author: lemo
*
* @create: 2022-08-03 22:26
**/

package main

import "sync"

type User struct {
	Name string `json:"name"`
	FD   int64  `json:"fd"`
}

var nameMap = make(map[string]int64)
var fdMap = make(map[int64]*User)
var mux sync.RWMutex

func AddUser(name string, user *User) {
	mux.Lock()
	defer mux.Unlock()
	if user == nil {
		return
	}
	nameMap[name] = user.FD
	fdMap[user.FD] = user
}

func DeleteUser(name string) {
	mux.Lock()
	defer mux.Unlock()
	var fd = nameMap[name]
	delete(nameMap, name)
	delete(fdMap, fd)
}

func DeleteUserByFD(fd int64) {
	mux.Lock()
	defer mux.Unlock()
	var user = fdMap[fd]
	if user == nil {
		return
	}
	delete(fdMap, fd)
	delete(nameMap, user.Name)
}

func GetUserByName(name string) *User {
	mux.RLock()
	defer mux.RUnlock()
	var fd = nameMap[name]
	return fdMap[fd]
}

func GetUserByFD(fd int64) *User {
	mux.RLock()
	defer mux.RUnlock()
	return fdMap[fd]
}

func GetAllUsers() []*User {
	mux.RLock()
	defer mux.RUnlock()
	var users []*User
	for _, user := range fdMap {
		users = append(users, user)
	}
	return users
}
