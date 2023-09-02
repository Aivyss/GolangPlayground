package dto

type Account struct {
	Id     int    `json:"id"`
	UserId string `json:"userId"`
	Name   string `json:"name"`
}

type AccountDb struct {
	Id       int
	UserId   string
	Name     string
	Password string
}
