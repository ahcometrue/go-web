package user

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type typeSex uint8

type Info struct {
	UserID   uint
	UserName string
	Sex      typeSex
}

func (s typeSex) SexToCNStr() string {
	sexStrMap := map[typeSex]string{
		1: "男",
		2: "女",
	}
	return sexStrMap[s]
}

func (s typeSex) SexToENStr() string {
	sexStrMap := map[typeSex]string{
		1: "male",
		2: "female",
	}
	return sexStrMap[s]
}



func GetUserID() int {
	return 100
}

func GetUserName(id int) (string, error) {
	names := map[int]string{
		100: "韩大大",
		200: "李大大梅",
	}
	if name, ok := names[id]; ok {
		return name, nil
	}
	return "", errors.New("math: square root of negative number")
}

func GetUserRandName() string {
	var prefixNames = []string{"小丸子", "小樱桃", "金刚", "抠脚", "肌肉"}
	var suffixNames = []string{"姑娘", "朋友"}
	rand.Seed(time.Now().Unix())
	return prefixNames[rand.Intn(len(prefixNames))] + suffixNames[rand.Intn(len(suffixNames))]
}

func GetUserList() []Info {
	var userList = []Info{
		{
			100,
			"韩",
			1,
		},
		{
			101,
			"李",
			2,
		},
	}
	fmt.Printf("UserList在函数内的地址%p\n", userList)
	return userList
}

func GetUserRegisterDate(userId int) string {
	var birthTimeMap = make(map[int]int64)
	birthTimeMap[100] = 1605671940
	birthTimeMap[200] = 1605671950
	birthTimeMap[300] = 1605671960
	return time.Unix(birthTimeMap[userId], 0).Format("2006-01-02 15:04:05")
}
