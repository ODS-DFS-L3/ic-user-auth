package usecase

import (
	"authenticator-backend/domain/model/authentication"
	"authenticator-backend/domain/repository"
	"authenticator-backend/extension/logger"
	"authenticator-backend/usecase/input"
	"authenticator-backend/usecase/output"
)

// verifyUsecase
// Summary: This is the structure which defines the verify usecase.
type verifyUsecase struct {
	gocloakRepository repository.GoCloakRepository
	authRepository    repository.AuthRepository
}

// NewVerifyUsecase
// Summary: This is the function which creates the verify usecase.
// input: f(repository.gocloakRepository) gocloak repository
// input: a(repository.AuthRepository) auth repository
// output: (IVerifyUsecase) verify usecase
func NewVerifyUsecase(f repository.GoCloakRepository, a repository.AuthRepository) IVerifyUsecase {
	return &verifyUsecase{f, a}
}

// TokenIntrospection
// Summary: This is the function which verifies the token.
// input: input(input.VerifyTokenParam) input parameters
// output: (output.VerifyTokenResponse) output response
// output: (error) error object
func (u verifyUsecase) TokenIntrospection(input input.VerifyTokenParam) (output.VerifyTokenResponse, error) {

	res, err := u.gocloakRepository.TokenIntrospection(input.IDToken)

	if err != nil {
		logger.Set(nil).Errorf(err.Error())
		return output.VerifyTokenResponse{}, err
	}

	return output.VerifyTokenResponse{
		OperatorID:   res.OperatorID,
		OpenSystemId: res.OpenSystemId,
		Active:       res.Active,
	}, nil
}

// IDToken
// Summary: This is the function which verifies the ID token.
// input: input(input.VerifyIDTokenParam) input parameters
// output: (authentication.Claims) claims
// output: (error) error object
func (u verifyUsecase) IDToken(input input.VerifyIDTokenParam) (authentication.Claims, error) {
	claims, err := u.gocloakRepository.VerifyIDToken(input.IDToken)
	if err != nil {
		logger.Set(nil).Warnf(err.Error())

		return authentication.Claims{}, err
	}
	return claims, nil
}
