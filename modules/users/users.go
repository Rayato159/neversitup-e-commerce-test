package users

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type usersErrCode string

const (
	RegisterErr usersErrCode = "users-001"
)

type User struct {
	Id       string `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
}

type UserPassport struct {
	User  *User      `json:"user"`
	Token *UserToken `json:"token"`
}

type UserCredential struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func (obj *UserCredential) BcryptHashing() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(obj.Password), 10)
	if err != nil {
		return fmt.Errorf("hashed password failed: %v", err)
	}
	obj.Password = string(hashedPassword)
	return nil
}

type UserClaims struct {
	Id       string `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
}

type UserToken struct {
	Id           string `db:"id" json:"id"`
	AccessToken  string `db:"access_token" json:"access_token"`
	RefreshToken string `db:"refresh_token" json:"refresh_token"`
}

type Oauth struct {
	Id     string `db:"id" json:"id"`
	UserId string `db:"user_id" json:"user_id"`
}

type UserRemoveCredential struct {
	OauthId string `json:"oauth_id" form:"oauth_id"`
}
