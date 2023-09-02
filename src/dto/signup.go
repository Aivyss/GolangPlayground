package dto

type Signup struct {
	Id       string `form:"id"`
	Password string `form:"password"`
	Name     string `form:"name"`
}
