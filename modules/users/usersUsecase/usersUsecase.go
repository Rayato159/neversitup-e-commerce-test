package usersUsecase

import (
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/users"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/users/usersRepository"
)

type IUsersUsecase interface {
	InsertUser(req *users.UserCredential) (*users.UserPassport, error)
}

type usersUsecase struct {
	usersRepository usersRepository.IUsersRepository
}

func NewUsersUsercase(usersRepository usersRepository.IUsersRepository) IUsersUsecase {
	return &usersUsecase{
		usersRepository: usersRepository,
	}
}

func (u *usersUsecase) InsertUser(req *users.UserCredential) (*users.UserPassport, error) {
	if err := req.BcryptHashing(); err != nil {
		return nil, err
	}

	user, err := u.usersRepository.InsertUser(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}
