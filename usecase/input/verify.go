package input

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// VerifyTokenParam
// Summary: This is the structure which defines the verify token parameter.
type VerifyTokenParam struct {
	IDToken string `json:"idToken"`
}

// Validate
// Summary: This is the function which validates the verify token parameter.
// output: (error) error object
func (p VerifyTokenParam) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(
			&p.IDToken,
			validation.Required,
		),
	)
}

// VerifyIDTokenParam
// Summary: This is the structure which defines the verify ID token parameter.
type VerifyIDTokenParam struct {
	IDToken string `json:"idToken"`
}
