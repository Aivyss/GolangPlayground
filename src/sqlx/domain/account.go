package domain

type AccountDb struct {
	Id       int    `db:"account_id"`
	UserId   string `db:"user_id"`
	Name     string `db:"user_name"`
	Password string `db:"password"`
}
