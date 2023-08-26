package service

import (
	"com.playground/23_gin/dto"
	"com.playground/23_gin/repository"
	"errors"
	"sync"
)

var once sync.Once
var a *AccountService

type AccountService struct{}

func AccountServe() *AccountService {
	once.Do(func() { a = &AccountService{} })

	return a
}

func (a AccountService) Signup(signup *dto.Signup) {
	repository.AccountRepo().Save(signup)
}

func (a AccountService) Login(login *dto.LoginDto) (dto.Account, error) {
	account, err := repository.AccountRepo().FindByUserId(login.Id)

	if err != nil {
		return dto.Account{}, err
	} else {
		if account.Password == login.Password {
			return dto.Account{
				Name:   account.Name,
				UserId: account.UserId,
				Id:     account.Id,
			}, nil
		} else {
			return dto.Account{}, errors.New("check again your ID or Password")
		}
	}
}
