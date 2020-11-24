package user

import "errors"

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
