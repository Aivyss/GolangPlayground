package dto

type LoginDto struct {
	Id       string `form:"id"`
	Password string `form:"password"`
}
