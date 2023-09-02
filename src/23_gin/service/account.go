package service

import (
	"com.playground/23_gin/repository"
	"com.playground/dto"
	"errors"
	"sync"
)

var once sync.Once
var a *AccountService

type AccountService struct {
	accountRepository *repository.AccountRepository
}

func AccountServe() *AccountService {
	once.Do(func() {
		a = &AccountService{
			accountRepository: repository.AccountRepo(),
		}
	})

	return a
}

func (a AccountService) Signup(signup *dto.Signup) {
	a.accountRepository.Save(signup)
}

func (a AccountService) Login(login *dto.Login) (dto.Account, error) {
	account, err := a.accountRepository.FindByUserId(login.Id)

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
