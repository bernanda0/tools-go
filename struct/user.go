package data

import "fmt"

type User struct {
	username, password string
	age                uint8
	message            func(string) string
}

func TestUser() {
	userA := User{
		"bernanda",
		"123",
		16,
		func(name string) string {
			return "Hello " + name
		},
	}
	fmt.Println(userA.message(userA.username))
}
