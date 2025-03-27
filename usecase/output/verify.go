package output

import (
	"authenticator-backend/domain/common"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// VerifyTokenResponse
// Summary: This is the structure which defines the verify token response.
type VerifyTokenResponse struct {
	OperatorID   string `json:"operatorId,omitempty"`
	OpenSystemId string `json:"openSystemId,omitempty"`
	Active       bool   `json:"active"`
}

// Validate
// Summary: This is the function which validates the verify token parameter.
// output: (error) error object
func (p VerifyTokenResponse) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(
			&p.OperatorID,
			validation.By(common.StringEmptyOrUUIDValid),
		),
		validation.Field(
			&p.OpenSystemId,
			validation.Length(21, 36), // between 21 and 36 characters
		),
	)
}

// VerifyApiKeyResponse
// Summary: This is the structure which defines the verify API key response.
type VerifyApiKeyResponse struct {
	IsAPIKeyValid    bool `json:"isApiKeyValid"`
	IsIPAddressValid bool `json:"isIpAddressValid"`
}
