package authUsecase

import (
	"github.com/Rayato159/neversuitup-e-commerce-test/config"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/auth"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/auth/authRepository"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/users"
	"github.com/Rayato159/neversuitup-e-commerce-test/pkg/lockkunchae"
)

type IAuthUsecase interface {
	GetPassport(cfg config.IConfig, req *users.UserForAll) (*users.UserPassport, error)
	FindAccessToken(userId, accessToken string) bool
	RefreshPassport(cfg config.IConfig, user *users.UserClaims, req *auth.UserRefreshCredential) (*users.UserPassport, error)
	FindUserId(refreshToken string) (string, error)
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

func (u *authUsecase) RefreshPassport(cfg config.IConfig, user *users.UserClaims, req *auth.UserRefreshCredential) (*users.UserPassport, error) {
	// Parse token
	claims, err := lockkunchae.ParseToken(cfg.Jwt(), req.RefreshToken)
	if err != nil {
		return nil, err
	}

	// Check oauth
	oauth, err := u.authRepository.FindOneOauth(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	newClaims := &users.UserClaims{
		Id:       user.Id,
		Username: user.Username,
	}

	accessToken, err := lockkunchae.Newlockkunchae(
		lockkunchae.Access,
		cfg.Jwt(),
		newClaims,
	)
	if err != nil {
		return nil, err
	}
	refreshToken := lockkunchae.RepeatToken(
		cfg.Jwt(),
		newClaims,
		claims.ExpiresAt.Unix(),
	)

	passport := &users.UserPassport{
		User: &users.User{
			Id:       user.Id,
			Username: user.Username,
		},
		Token: &users.UserToken{
			Id:           oauth.Id,
			AccessToken:  accessToken.SignToken(),
			RefreshToken: refreshToken,
		},
	}
	if err := u.authRepository.UpdateOauth(passport.Token); err != nil {
		return nil, err
	}
	return passport, nil
}

func (u *authUsecase) FindUserId(refreshToken string) (string, error) {
	userId, err := u.authRepository.FindUserId(refreshToken)
	if err != nil {
		return "", err
	}
	return userId, nil
}
