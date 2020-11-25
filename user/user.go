package user

import (
	"errors"
	"math/rand"
	"time"
)

type Info struct {
	UserID   uint
	UserName string
	Sex      uint8
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
	var prefixNames = [...]string{"小丸子", "小樱桃", "金刚", "抠脚", "肌肉"}
	var suffixNames = [...]string{"姑娘", "朋友"}
	rand.Seed(time.Now().Unix())
	return prefixNames[rand.Intn(len(prefixNames))] + suffixNames[rand.Intn(len(suffixNames))]
}

func GetUserList() []Info {
	var UserList = []Info{
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
	return UserList
}
