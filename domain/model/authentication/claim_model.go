package authentication

import (
	base64 "encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// Claims
// Summary: This is structure which defines the claims model.
type Claims struct {
	OperatorID string `json:"operator_id"`
	UID        string `json:"sub"`
}

// NewClaims
// Summary: This is the function which creates the Claims model from the token.
// input: token(*auth.Token): token
// output: (Claims) Claims model
// output: (error) error object
func NewClaims(token string) (Claims, error) {
	part := strings.Split(token, ".")
	payload, err := base64.RawURLEncoding.DecodeString(part[1])
	if err != nil {
		return Claims{}, fmt.Errorf("failed to decode token payload: %v", err)
	}
	var claimsMap map[string]interface{}
	if err := json.Unmarshal(payload, &claimsMap); err != nil {
		return Claims{}, fmt.Errorf("failed to unmarshal token payload: %v", err)
	}

	// Extract the operator_id claim
	operatorID, ok := claimsMap["operator_id"].(string)
	log.Println("operatorID入手", operatorID)
	if !ok {
		return Claims{}, fmt.Errorf("token does not contain 'operator_id' in claims")
	}
	uid, ok := claimsMap["sub"].(string)
	if !ok {
		return Claims{}, fmt.Errorf("token does not contain 'uid' in claims")
	}
	return Claims{
		OperatorID: operatorID,
		UID:        uid,
	}, nil
}
