package authentication_test

import (
	"authenticator-backend/domain/model/authentication"
	f "authenticator-backend/test/fixtures"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 正常系
func TestProjectRepository_Firebase_VerifyIDToken(tt *testing.T) {

	tests := []struct {
		name   string
		token  string
		expect authentication.Claims
	}{
		{
			name:  "正常系",
			token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjc0MTI4NDUsIm9wZXJhdG9yX2lkIjoiZXhhbXBsZV9vcGVyYXRvcl9pZCIsInN1YiI6ImV4YW1wbGVfc3ViIn0.Ab2FKCgH7u2p82_zWgOvcatQMwW_8uHcLZi02iseNvY",
			expect: authentication.Claims{
				OperatorID: "example_operator_id",
				UID:        "example_sub",
			},
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				result, err := authentication.NewClaims(test.token)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expect.OperatorID, result.OperatorID, f.AssertMessage)
					assert.Equal(t, test.expect.UID, result.UID, f.AssertMessage)
				}
			},
		)
	}
}

// 異常系
func TestProjectRepository_Firebase_VerifyIDToken_Abnormal(tt *testing.T) {

	tests := []struct {
		name   string
		token  string
		expect error
	}{
		{
			name:   "異常系_トークンデコード失敗",
			token:  "token.error",
			expect: fmt.Errorf("failed to decode token payload: illegal base64 data at input byte 4"),
		},
		{
			name:   "異常系_ペイロード失敗",
			token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJvcGVyYXRvcl9pZCI6ICJleGFtcGxlX29wZXJhdG9yX2lkIiwgInN1YiI6ICJleGFtcGxlX3N1YiIs.ZHVtbXlfc2lnbmF0dXJl",
			expect: fmt.Errorf("failed to unmarshal token payload: unexpected end of JSON input"),
		},
		{
			name:   "異常系_operator_idの含まれないトークン",
			token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjc0MjExNTgsInN1YiI6ImV4YW1wbGVfc3ViIn0.oJHg7mJDGx3SF4meaomrjgMQpSVPo3csKMoFWdbwLJMr",
			expect: fmt.Errorf("token does not contain 'operator_id' in claims"),
		},
		{
			name:   "異常系_uidの含まれないトークン",
			token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjc0MjEyMjQsIm9wZXJhdG9yX2lkIjoiZXhhbXBsZV9vcGVyYXRvcl9pZCJ9.1yRiPDlqRRI_hEhcfb3q-ArxOwKvfjhc-de0IuCKAtw",
			expect: fmt.Errorf("token does not contain 'uid' in claims"),
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				_, err := authentication.NewClaims(test.token)
				if assert.Error(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}
