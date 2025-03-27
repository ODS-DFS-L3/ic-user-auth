package authentication

// TokenResult
// Summary: This is structure which defines the tokenIntoroSpect result model.
type TokenResult struct {
	OperatorID   string `json:"operator_id"`
	OpenSystemId string `json:"open_system_id"`
	Active       bool   `json:"active"`
}
