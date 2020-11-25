package main

import (
	"fmt"
	"go-web-study/user"
)

func main() {

	id := user.GetUserID()
	userName, ok := user.GetUserName(100)
	if ok != nil {
		println(ok.Error())
		return
	}
	userRandName := user.GetUserRandName()
	fmt.Println(id, userName, userRandName)
	userList := user.GetUserList()
	fmt.Printf("%#v\n", userList)
	for _, v := range userList {
		fmt.Println(v.UserID, v.UserName, v.Sex)
	}
}
