package input

import (
	"regexp"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// LoginParam
// Summary: This is the structure which defines the login parameter.
type LoginParam struct {
	OperatorAccountID string `json:"operatorAccountId"`
	AccountPassword   string `json:"accountPassword"`
}

// Validate
// Summary: This is the function which validates the login parameter.
// output: (error) error object
func (i LoginParam) Validate() error {
	return i.validate()
}

// validate
// Summary: This is the function which validates the login parameter.
// output: (error) error object
func (i LoginParam) validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(
			&i.OperatorAccountID,
			validation.Required,
			validation.Length(1, 320), // between 1 and 320 characters
		),
		validation.Field(
			&i.AccountPassword,
			validation.Required,
			validation.Length(8, 20), // between 8 and 20 characters
			validation.Match(regexp.MustCompile(`[A-Z]`)).Error("must include at least one upper case letter"),         // at least one upper case letter
			validation.Match(regexp.MustCompile(`[a-z]`)).Error("must include at least one lower case letter"),         // at least one lower case letter
			validation.Match(regexp.MustCompile(`[0-9]`)).Error("must include at least one digit"),                     // at least one digit
			validation.Match(regexp.MustCompile(`[!@#\$%\^&\*]`)).Error("must include at least one special character"), // at least one special character
		),
	)
}

// Mask
// Summary: This is the function which masks the confidential information.
func (i *LoginParam) Mask() {
	i.AccountPassword = strings.Repeat("*", len(i.AccountPassword))
}

type ClientParam struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

// Validate
// Summary: This is the function which validates the login parameter.
// output: (error) error object
func (i ClientParam) Validate() error {
	return i.validate()
}

// validate
// Summary: This is the function which validates the login parameter.
// output: (error) error object
func (i ClientParam) validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(
			&i.ClientID,
			validation.Required,
		),
		validation.Field(
			&i.ClientSecret,
			validation.Required,
		),
	)
}

// RefreshParam
// Summary: This is the structure which defines the refresh parameter.
type RefreshParam struct {
	RefreshToken string `json:"refreshToken"`
}

// Validate
// Summary: This is the function which validates the refresh parameter.
// output: (error) error object
func (i RefreshParam) Validate() error {
	return i.validate()
}

// validate
// Summary: This is the function which validates the refresh parameter.
// output: (error) error object
func (i RefreshParam) validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(
			&i.RefreshToken,
			validation.Required,
		),
	)
}

// Mask
// Summary: This is the function which masks the confidential information.
func (i *RefreshParam) Mask() {
	i.RefreshToken = strings.Repeat("*", len(i.RefreshToken))
}

// ChangePasswordParam
// Summary: This is the structure which defines the change password parameter.
type ChangePasswordParam struct {
	UID         string
	NewPassword string
}

// Validate
// Summary: This is the function which validates the change password parameter.
// output: (error) error object
func (i ChangePasswordParam) Validate() error {
	if err := i.validate(); err != nil {
		return err
	}
	return nil
}

// validate
// Summary: This is the function which validates the change password parameter.
// output: (error) error object
func (i ChangePasswordParam) validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(
			&i.UID,
			validation.Required,
		),
		validation.Field(
			&i.NewPassword,
			validation.Required,
			validation.Length(8, 20), // between 8 and 20 characters
			validation.Match(regexp.MustCompile(`[A-Z]`)).Error("must include at least one upper case letter"),         // at least one upper case letter
			validation.Match(regexp.MustCompile(`[a-z]`)).Error("must include at least one lower case letter"),         // at least one lower case letter
			validation.Match(regexp.MustCompile(`[0-9]`)).Error("must include at least one digit"),                     // at least one digit
			validation.Match(regexp.MustCompile(`[!@#\$%\^&\*]`)).Error("must include at least one special character"), // at least one special character
		),
	)
}
