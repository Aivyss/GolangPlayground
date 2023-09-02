package service

import (
	"com.playground/24_echo/repository"
	"com.playground/dto"
	"errors"
)

var AccountService = struct {
	Signup func(signup *dto.Signup)
	Login  func(login *dto.Login) (dto.Account, error)
}{
	Signup: func(signup *dto.Signup) {
		repository.AccountRepository.Save(signup)
	},
	Login: func(login *dto.Login) (dto.Account, error) {
		account, err := repository.AccountRepository.FindByUserId(login.Id)

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
	},
}
