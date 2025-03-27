package repository

import (
	"authenticator-backend/domain/model/authentication"
)

// GoCloakRepository
// Summary: This is interface which defines GoCloakRepository　functions.
//
//go:generate mockery --name GoCloakRepository --output ../../test/mock --case underscore
type GoCloakRepository interface {
	SignInWithPassword(operatorAccountID string, password string) (authentication.LoginResult, error)
	SignInWithClient(clientID, clientSecret string) (authentication.ClientResult, error)
	VerifyIDToken(idToken string) (authentication.Claims, error)
	TokenIntrospection(idToken string) (authentication.TokenResult, error)
	RefreshToken(refreshToken string) (string, error)
	ChangePassword(uid, newPassword string) error
}
