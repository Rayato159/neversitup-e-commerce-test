package auth

type authErrCode string

const (
	LoginErr   authErrCode = "auth-001"
	RefreshErr authErrCode = "auth-002"
)

type UserRefreshCredential struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token"`
}
