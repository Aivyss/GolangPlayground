package repository

import (
	dto2 "com.playground/dto"
	"errors"
	"fmt"
	"sync"
)

type AccountRepository struct{}

var once sync.Once
var a *AccountRepository
var memoryPersistence = make(map[string]dto2.AccountDb)

func AccountRepo() *AccountRepository {
	once.Do(func() { a = &AccountRepository{} })

	return a
}

func (a AccountRepository) Save(signup *dto2.Signup) int {
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
}

func (a AccountRepository) FindByUserId(userId string) (dto2.AccountDb, error) {
	fmt.Println("=====[Account Find]=====")
	fmt.Println(userId)
	data, ok := memoryPersistence[userId]

	if ok {
		return data, nil
	} else {
		return dto2.AccountDb{}, errors.New("check again your ID or Password")
	}
}
