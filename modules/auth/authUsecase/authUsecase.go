package authUsecase

import (
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/auth/authRepository"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/users"
)

type IAuthUsecase interface {
	GetPassport(req *users.UserCredential) (*users.UserPassport, error)
}

type authUsecase struct {
	authRepository authRepository.IAuthRepository
}

func NewAuthUsecase(authRepository authRepository.IAuthRepository) IAuthUsecase {
	return &authUsecase{
		authRepository: authRepository,
	}
}

func (u *authUsecase) GetPassport(req *users.UserCredential) (*users.UserPassport, error) {

	return nil, nil
}
