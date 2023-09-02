package dto

type Login struct {
	Id       string `form:"id"`
	Password string `form:"password"`
}
