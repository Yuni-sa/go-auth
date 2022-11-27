package main

type User struct {
	Id       string `json:"id_number"`
	Username string `json:"username"`
	Password string `json:"-"`
}
