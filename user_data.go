package main

type UserData struct {
	Email    string `json:"email"`
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	State    string `json:"state"`
}

func NewUserData() (data UserData) {
	data.State = "active"

	return
}
