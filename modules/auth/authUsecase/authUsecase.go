package authUsecase

import (
	"github.com/Rayato159/neversuitup-e-commerce-test/config"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/auth/authRepository"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/users"
	"github.com/Rayato159/neversuitup-e-commerce-test/pkg/lockkunchae"
)

type IAuthUsecase interface {
	GetPassport(cfg config.IConfig, req *users.UserForAll) (*users.UserPassport, error)
	FindAccessToken(userId, accessToken string) bool
}

type authUsecase struct {
	authRepository authRepository.IAuthRepository
}

func NewAuthUsecase(authRepository authRepository.IAuthRepository) IAuthUsecase {
	return &authUsecase{
		authRepository: authRepository,
	}
}

func (u *authUsecase) GetPassport(cfg config.IConfig, req *users.UserForAll) (*users.UserPassport, error) {
	// Sign token
	accessToken, err := lockkunchae.Newlockkunchae(lockkunchae.Access, cfg.Jwt(), &users.UserClaims{
		Id:       req.Id,
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}
	refreshToken, err := lockkunchae.Newlockkunchae(lockkunchae.Refresh, cfg.Jwt(), &users.UserClaims{
		Id:       req.Id,
		Username: req.Username,
	})
	if err != nil {
		return nil, err
	}

	// Set passport
	passport := &users.UserPassport{
		User: &users.User{
			Id:       req.Id,
			Username: req.Username,
		},
		Token: &users.UserToken{
			AccessToken:  accessToken.SignToken(),
			RefreshToken: refreshToken.SignToken(),
		},
	}

	if err := u.authRepository.InsertOauth(passport); err != nil {
		return nil, err
	}
	return passport, nil
}

func (u *authUsecase) FindAccessToken(userId, accessToken string) bool {
	return u.authRepository.FindAccessToken(userId, accessToken)
}
