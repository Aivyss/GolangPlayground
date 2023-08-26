package repository

import (
	"com.playground/23_gin/dto"
	"errors"
	"fmt"
	"sync"
)

type AccountRepository struct{}

var once sync.Once
var a *AccountRepository
var memoryPersistence = make(map[string]dto.AccountDb)

func AccountRepo() *AccountRepository {
	once.Do(func() { a = &AccountRepository{} })

	return a
}

func (a AccountRepository) Save(signup *dto.Signup) int {
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

		memoryPersistence[signup.Id] = dto.AccountDb{
			Id:       mapSize,
			Name:     signup.Name,
			Password: signup.Password,
			UserId:   signup.Id,
		}
	}

	return mapSize
}

func (a AccountRepository) FindByUserId(userId string) (dto.AccountDb, error) {
	fmt.Println("=====[Account Find]=====")
	fmt.Println(userId)
	data, ok := memoryPersistence[userId]

	if ok {
		return data, nil
	} else {
		return dto.AccountDb{}, errors.New("check again your ID or Password")
	}
}
