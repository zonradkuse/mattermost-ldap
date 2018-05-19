package main

type UserData struct {
	Email    string `json:"email"`
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	State    string `json:"state"`
}

func NewUserData() (data UserData) {
	data.State = "active"

	return
}
