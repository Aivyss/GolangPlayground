package service

import (
	"com.playground/24_echo/repository"
	dto2 "com.playground/dto"
	"errors"
)

var AccountService = struct {
	Signup func(signup *dto2.Signup)
	Login  func(login *dto2.Login) (dto2.Account, error)
}{
	Signup: func(signup *dto2.Signup) {
		repository.AccountRepository.Save(signup)
	},
	Login: func(login *dto2.Login) (dto2.Account, error) {
		account, err := repository.AccountRepository.FindByUserId(login.Id)

		if err != nil {
			return dto2.Account{}, err
		} else {
			if account.Password == login.Password {
				return dto2.Account{
					Name:   account.Name,
					UserId: account.UserId,
					Id:     account.Id,
				}, nil
			} else {
				return dto2.Account{}, errors.New("check again your ID or Password")
			}
		}
	},
}
