package middlewaresUsecase

import "github.com/Rayato159/neversuitup-e-commerce-test/modules/middlewares/middlewaresRepository"

type IMiddlewaresUsecase interface {
	FindAccessToken(userId, accessToken string) bool
}

type middlewaresUsecase struct {
	middlewaresRepository middlewaresRepository.IMiddlewaresRepository
}

func MiddlewaresUsecase(middlewaresRepository middlewaresRepository.IMiddlewaresRepository) IMiddlewaresUsecase {
	return &middlewaresUsecase{
		middlewaresRepository: middlewaresRepository,
	}
}

func (u *middlewaresUsecase) FindAccessToken(userId, accessToken string) bool {
	return u.middlewaresRepository.FindAccessToken(userId, accessToken)
}
