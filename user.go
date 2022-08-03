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

import hash "github.com/lemonyxk/structure/v3/map"

type User struct {
	Name string `json:"name"`
	FD   int64  `json:"fd"`
}

var nameMap hash.SyncHash[string, int64]
var fdMap hash.SyncHash[int64, *User]

func AddUser(name string, user *User) {
	nameMap.Set(name, user.FD)
	fdMap.Set(user.FD, user)
}

func DeleteUser(name string) {
	var fd = nameMap.Get(name)
	if fd == 0 {
		return
	}
	nameMap.Delete(name)
	fdMap.Delete(fd)
}

func DeleteUserByFD(fd int64) {
	var user = fdMap.Get(fd)
	if user == nil {
		return
	}
	fdMap.Delete(fd)
	nameMap.Delete(user.Name)
}

func GetUserByName(name string) *User {
	var fd = nameMap.Get(name)
	if fd == 0 {
		return nil
	}
	return fdMap.Get(fd)
}

func GetUserByFD(fd int64) *User {
	return fdMap.Get(fd)
}

func GetAllUsers() []*User {
	var users []*User
	fdMap.Range(func(fd int64, user *User) bool {
		users = append(users, user)
		return true
	})
	return users
}
