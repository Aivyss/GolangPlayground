package repository

import (
	"com.playground/dto"
	dbPkg "com.playground/sqlx"
	"com.playground/sqlx/domain"
	"errors"
	"fmt"
)

var AccountRepository = struct {
	Save         func(signup *dto.Signup) int
	FindByUserId func(userId string) (domain.AccountDb, error)
}{
	Save: func(signup *dto.Signup) int {
		var lastId int64 = -1
		fmt.Println("=====[Account Save]=====")
		fmt.Println(signup.Password)
		fmt.Println(signup.Id)
		fmt.Println(signup.Name)

		tx := dbPkg.Db.MustBegin()
		err := tx.QueryRow(`
			INSERT INTO ACCOUNT (
				USER_ID,
				PASSWORD,
				USER_NAME
			) VALUES ($1, $2, $3) RETURNING ACCOUNT_ID
		`, signup.Id, signup.Password, signup.Name).Scan(&lastId)

		if err != nil {
			_ = tx.Rollback()
		} else {
			_ = tx.Commit()
		}

		return int(lastId)
	},
	FindByUserId: func(userId string) (domain.AccountDb, error) {
		result := new(domain.AccountDb)
		fmt.Println("=====[Account Find]=====")
		fmt.Println(userId)

		// select one
		err := dbPkg.Db.Get(result, `SELECT * FROM ACCOUNT WHERE USER_ID = $1`, userId)

		// select list : dbPkg.Db.Select

		if err == nil {
			return *result, nil
		} else {
			return domain.AccountDb{}, errors.New("check again your ID or Password")
		}
	},
}
