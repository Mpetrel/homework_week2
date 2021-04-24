package model

type User struct {
	Id int `json:"id"`
	Nickname string `json:"nickname"`
	Email string `json:"email"`
	Password string `json:"password"`
}