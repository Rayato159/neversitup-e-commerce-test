package lockkunchae

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/Rayato159/neversuitup-e-commerce-test/config"
	"github.com/Rayato159/neversuitup-e-commerce-test/modules/users"
	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

const (
	Access  TokenType = "access"
	Refresh TokenType = "refresh"
	ApiKey  TokenType = "apikey"
)

type lockkunchaeAuth struct {
	mapClaims *lockkunchaeMapClaims
	cfg       config.IJwtConfig
}

type lockkunchaeApiKey struct {
	*lockkunchaeAuth
}

type lockkunchaeMapClaims struct {
	Claims *users.UserClaims `json:"claims"`
	jwt.RegisteredClaims
}

type Ilockkunchae interface {
	SignToken() string
}

type IKawaiiApiKey interface {
	SignToken() string
}

func jwtTimeDurationCal(t int) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(time.Duration(int64(t) * int64(math.Pow10(9)))))
}

func jwtTimeRepeatAdapter(t int64) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Unix(t, 0))
}

func (a *lockkunchaeAuth) SignToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.mapClaims)
	ss, _ := token.SignedString(a.cfg.SecretKey())
	return ss
}

func (a *lockkunchaeApiKey) SignToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.mapClaims)
	ss, _ := token.SignedString(a.cfg.ApiKey())
	return ss
}

func ParseToken(cfg config.IJwtConfig, tokenString string) (*lockkunchaeMapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &lockkunchaeMapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method is invalid")
		}
		return cfg.SecretKey(), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, fmt.Errorf("token format is invalid")
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf("token had expired")
		} else {
			return nil, fmt.Errorf("parse token failed: %v", err)
		}
	}

	if claims, ok := token.Claims.(*lockkunchaeMapClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("claims type is invalid")
	}
}

func ParseApiKey(cfg config.IJwtConfig, tokenString string) (*lockkunchaeMapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &lockkunchaeMapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method is invalid")
		}
		return cfg.ApiKey(), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, fmt.Errorf("token format is invalid")
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf("token had expired")
		} else {
			return nil, fmt.Errorf("parse token failed: %v", err)
		}
	}

	if claims, ok := token.Claims.(*lockkunchaeMapClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("claims type is invalid")
	}
}

func RepeatToken(cfg config.IJwtConfig, claims *users.UserClaims, exp int64) string {
	obj := &lockkunchaeAuth{
		cfg: cfg,
		mapClaims: &lockkunchaeMapClaims{
			Claims: claims,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "nerversuitup-shop-api",
				Subject:   "refresh-token",
				Audience:  []string{"user"},
				ExpiresAt: jwtTimeRepeatAdapter(exp),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		},
	}
	return obj.SignToken()
}

func Newlockkunchae(tokenType TokenType, cfg config.IJwtConfig, claims *users.UserClaims) (Ilockkunchae, error) {
	switch tokenType {
	case Access:
		return newAccessToken(cfg, claims), nil
	case Refresh:
		return newRefreshToken(cfg, claims), nil
	case ApiKey:
		return newApiKey(cfg), nil
	default:
		return nil, fmt.Errorf("unknown token type")
	}
}

func newAccessToken(cfg config.IJwtConfig, claims *users.UserClaims) Ilockkunchae {
	return &lockkunchaeAuth{
		cfg: cfg,
		mapClaims: &lockkunchaeMapClaims{
			Claims: claims,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "nerversuitup-shop-api",
				Subject:   "access-token",
				Audience:  []string{"user"},
				ExpiresAt: jwtTimeDurationCal(cfg.AccessExpiresAt()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		},
	}
}

func newRefreshToken(cfg config.IJwtConfig, claims *users.UserClaims) Ilockkunchae {
	return &lockkunchaeAuth{
		cfg: cfg,
		mapClaims: &lockkunchaeMapClaims{
			Claims: claims,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "nerversuitup-shop-api",
				Subject:   "refresh-token",
				Audience:  []string{"user"},
				ExpiresAt: jwtTimeDurationCal(cfg.RefreshExpiresAt()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		},
	}
}

func newApiKey(cfg config.IJwtConfig) Ilockkunchae {
	return &lockkunchaeApiKey{
		lockkunchaeAuth: &lockkunchaeAuth{
			cfg: cfg,
			mapClaims: &lockkunchaeMapClaims{
				Claims: nil,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    "nerversuitup-shop-api",
					Subject:   "api-key",
					Audience:  []string{"user"},
					ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(2, 0, 0)),
					NotBefore: jwt.NewNumericDate(time.Now()),
					IssuedAt:  jwt.NewNumericDate(time.Now()),
				},
			},
		},
	}
}
