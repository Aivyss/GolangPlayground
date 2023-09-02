package dto

type Account struct {
	Id     int    `json:"id"`
	UserId string `json:"userId"`
	Name   string `json:"name"`
}
