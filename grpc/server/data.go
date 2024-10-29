package main

type User struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone"`
}

var users = map[string]*User{
	"1": {
		Id:    "1",
		Name:  "陈海龙",
		Phone: "17628689483",
	},
	"2": {
		Id:    "2",
		Name:  "刘宇",
		Phone: "15883430560",
	},
}
