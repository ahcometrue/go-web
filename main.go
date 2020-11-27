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
	fmt.Printf("UserList在函数外的地址%p\n", userList)
	fmt.Printf("%#v\n", userList)
	for _, v := range userList {
		fmt.Println(v.UserID, v.UserName, v.Sex, v.Sex.SexToCNStr(), v.Sex.SexToENStr())
	}
	date := user.GetUserRegisterDate(100)
	fmt.Printf("%#v\n", date)
	//c := make(chan int)
	//fmt.Printf("%#v\n", c)
	//fmt.Printf("%T\n", c)
	//fmt.Println(c)


}
