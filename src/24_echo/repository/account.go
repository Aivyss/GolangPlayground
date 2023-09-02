package repository

import (
	dto2 "com.playground/dto"
	"errors"
	"fmt"
)

var memoryPersistence = make(map[string]dto2.AccountDb)

var AccountRepository = struct {
	Save         func(signup *dto2.Signup) int
	FindByUserId func(userId string) (dto2.AccountDb, error)
}{
	Save: func(signup *dto2.Signup) int {
		fmt.Println("=====[Account Save]=====")
		fmt.Println(signup.Password)
		fmt.Println(signup.Id)
		fmt.Println(signup.Name)
		mapSize := 0
		_, ok := memoryPersistence[signup.Id]

		if !ok {
			for range memoryPersistence {
				mapSize += 1
			}

			memoryPersistence[signup.Id] = dto2.AccountDb{
				Id:       mapSize,
				Name:     signup.Name,
				Password: signup.Password,
				UserId:   signup.Id,
			}
		}

		return mapSize
	},
	FindByUserId: func(userId string) (dto2.AccountDb, error) {
		fmt.Println("=====[Account Find]=====")
		fmt.Println(userId)
		data, ok := memoryPersistence[userId]

		if ok {
			return data, nil
		} else {
			return dto2.AccountDb{}, errors.New("check again your ID or Password")
		}
	},
}
