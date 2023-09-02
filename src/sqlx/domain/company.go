package domain

type CompanyRole string

const (
	Admin   CompanyRole = "ADMIN"
	Manager             = "MANAGER"
	None                = "NONE"
)

type CompanyDb struct {
	Id   int    `db:"company_id"`
	Name string `db:"company_name"`
}

type CompanyRoleDb struct {
	Role      CompanyRole `db:"role"`
	AccountId int         `db:"account_id"`
	CompanyId int         `db:"company_id"`
}
