package main

import (
	"fmt"
	"go-web-study/user"
)

func main() {

	id := user.GetUserID()
	userName, ok := user.GetUserName(110)
	if ok != nil {
		println(ok.Error())
		return
	}
	fmt.Println(id, userName)
}
