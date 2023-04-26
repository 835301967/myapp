package main

import "fmt"

type User struct {
	ID     int64
	Name   string
	Avatar string
}

func GetUserInfo() *User {
	return &User{ID: 13746731, Name: "EDDYCJY", Avatar: "https://avatars0.githubusercontent.com/u/13746731"}
}

func do() string {
	str := new(string)
	*str = "EDDYCJY"
	return ""
}

func fmtString() string  {
	str := new(string)
	*str = "EDDYCJY"

	fmt.Println(str)
	return ""
}

func main() {
	_ = GetUserInfo()
	_  = do()
	_ = fmtString()
}
