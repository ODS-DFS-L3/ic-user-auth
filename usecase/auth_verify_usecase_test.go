package usecase_test

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"testing"

	"authenticator-backend/domain/model/authentication"
	f "authenticator-backend/test/fixtures"
	mocks "authenticator-backend/test/mock"
	"authenticator-backend/usecase"
	"authenticator-backend/usecase/input"
	"authenticator-backend/usecase/output"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestProjectUsecase_TokenIntrospection
// Summary: This is normal test class which confirm the operation of API TokenIntrospection.
// Target: auth_verify_usecase_impl.go
// TestPattern:
// [x] 1-1. 200: 正常系（リソースオーナ用）
// [x] 1-2. 200: 正常系（クライアントクレデンシャルフロー用）
func TestProjectUsecase_TokenIntrospection(tt *testing.T) {

	var method = "GET"
	var endPoint = "/tokenIntrospection"

	res := authentication.TokenResult{
		OperatorID:   "e03cc699-7234-31ed-86be-cc18c92208e5",
		OpenSystemId: "",
		Active:       true,
	}
	res2 := authentication.TokenResult{
		OperatorID:   "",
		OpenSystemId: "cooperationSystemA101",
		Active:       true,
	}
	expected := output.VerifyTokenResponse{
		OperatorID: "e03cc699-7234-31ed-86be-cc18c92208e5",
		Active:     true,
	}
	expected2 := output.VerifyTokenResponse{
		OpenSystemId: "cooperationSystemA101",
		Active:       true,
	}
	tests := []struct {
		name    string
		input   input.VerifyTokenParam
		receive authentication.TokenResult
		expect  output.VerifyTokenResponse
	}{
		{
			name:    "1-1. 200: 正常系（リソースオーナ用）",
			input:   f.NewInputVerifyTokenParam(),
			receive: res,
			expect:  expected,
		},
		{
			name:    "1-2. 200: 正常系（クライアントクレデンシャルフロー用）",
			input:   f.NewInputVerifyTokenParam(),
			receive: res2,
			expect:  expected2,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				q := make(url.Values)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)

				gocloakRepositoryMock := new(mocks.GoCloakRepository)
				authRepositoryMock := new(mocks.AuthRepository)
				gocloakRepositoryMock.On("TokenIntrospection", mock.Anything).Return(test.receive, nil)
				verifyUsecase := usecase.NewVerifyUsecase(gocloakRepositoryMock, authRepositoryMock)

				actual, err := verifyUsecase.TokenIntrospection(test.input)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expect.OperatorID, actual.OperatorID, f.AssertMessage)
					assert.Equal(t, test.expect.OpenSystemId, actual.OpenSystemId, f.AssertMessage)
					assert.Equal(t, test.expect.Active, actual.Active, f.AssertMessage)
				}
			},
		)
	}
}

// TestProjectUsecase_TokenIntrospection_Abnormal
// Summary: This is abnormal test class which confirm the operation of API TokenIntrospection.
// Target: auth_verify_usecase_impl.go
// TestPattern:
// [x] 2-1. 500: 検証処理エラー
func TestProjectUsecase_TokenIntrospection_Abnormal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/tokenIntrospection"

	tests := []struct {
		name         string
		input        input.VerifyTokenParam
		receive      authentication.TokenResult
		receiveError error
		expect       error
	}{
		{
			name:         "2-1. 500: 検証処理エラー",
			input:        f.NewInputVerifyTokenParam(),
			receive:      authentication.TokenResult{},
			receiveError: fmt.Errorf("検証処理エラー"),
			expect:       fmt.Errorf("検証処理エラー"),
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				q := make(url.Values)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)

				gocloakRepositoryMock := new(mocks.GoCloakRepository)
				authRepositoryMock := new(mocks.AuthRepository)
				gocloakRepositoryMock.On("TokenIntrospection", mock.Anything).Return(test.receive, test.receiveError)
				verifyUsecase := usecase.NewVerifyUsecase(gocloakRepositoryMock, authRepositoryMock)

				_, err := verifyUsecase.TokenIntrospection(test.input)
				if assert.Error(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// TestProjectUsecase_IDToken
// Summary: This is normal test class which confirm the operation of API IDToken.
// Target: auth_verify_usecase_impl.go
// TestPattern:
// [x] 1-1. 200: 正常系
func TestProjectUsecase_IDToken(tt *testing.T) {

	var method = "GET"
	var endPoint = "/tokenIntrospection"

	res := authentication.Claims{
		OperatorID: "e03cc699-7234-31ed-86be-cc18c92208e5",
	}
	tests := []struct {
		name    string
		input   input.VerifyIDTokenParam
		receive authentication.Claims
		expect  authentication.Claims
	}{
		{
			name:    "1-1. 200: 正常系",
			input:   f.NewInputVerifyIDTokenParam(),
			receive: res,
			expect:  res,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				q := make(url.Values)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)

				gocloakRepositoryMock := new(mocks.GoCloakRepository)
				authRepositoryMock := new(mocks.AuthRepository)
				gocloakRepositoryMock.On("VerifyIDToken", mock.Anything).Return(test.receive, nil)
				verifyUsecase := usecase.NewVerifyUsecase(gocloakRepositoryMock, authRepositoryMock)

				actual, err := verifyUsecase.IDToken(test.input)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expect.OperatorID, actual.OperatorID, f.AssertMessage)
				}
			},
		)
	}
}

// TestProjectUsecase_IDToken_Abnormal
// Summary: This is abnormal test class which confirm the operation of API IDToken.
// Target: auth_verify_usecase_impl.go
// TestPattern:
// [x] 2-1. 500: 検証処理エラー
func TestProjectUsecase_IDToken_Abnormal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/token"

	tests := []struct {
		name         string
		input        input.VerifyIDTokenParam
		receive      authentication.Claims
		receiveError error
		expect       error
	}{
		{
			name:         "2-1. 500: 検証処理エラー",
			input:        f.NewInputVerifyIDTokenParam(),
			receive:      authentication.Claims{},
			receiveError: fmt.Errorf("検証処理エラー"),
			expect:       fmt.Errorf("検証処理エラー"),
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()

				q := make(url.Values)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)

				gocloakRepositoryMock := new(mocks.GoCloakRepository)
				authRepositoryMock := new(mocks.AuthRepository)
				gocloakRepositoryMock.On("VerifyIDToken", mock.Anything).Return(test.receive, test.receiveError)
				verifyUsecase := usecase.NewVerifyUsecase(gocloakRepositoryMock, authRepositoryMock)

				_, err := verifyUsecase.IDToken(test.input)
				if assert.Error(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}
