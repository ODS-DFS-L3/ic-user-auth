package usecase_test

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"testing"

	"authenticator-backend/domain/common"
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

// TestProjectUsecase_Login
// Summary: This is normal test class which confirm the operation of API Login.
// Target: auth_usecase_impl.go
// TestPattern:
// [x] 1-1. 200: 正常系
func TestProjectUsecase_Login(tt *testing.T) {

	var method = "GET"
	var endPoint = "/login"

	res := authentication.LoginResult{
		AccessToken:  f.Token,
		RefreshToken: f.Token,
	}
	expected := output.LoginResponse{
		AccessToken:  f.Token,
		RefreshToken: f.Token,
	}
	tests := []struct {
		name    string
		input   input.LoginParam
		receive authentication.LoginResult
		expect  output.LoginResponse
	}{
		{
			name:    "1-1. 200: 正常系",
			input:   f.NewInputLoginParam(),
			receive: res,
			expect:  expected,
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
				gocloakRepositoryMock.On("SignInWithPassword", mock.Anything, mock.Anything).Return(test.receive, nil)
				authusecase := usecase.NewAuthUsecase(gocloakRepositoryMock)

				actual, err := authusecase.Login(test.input)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expect.AccessToken, actual.AccessToken, f.AssertMessage)
					assert.Equal(t, test.expect.RefreshToken, actual.RefreshToken, f.AssertMessage)
				}
			},
		)
	}
}

// TestProjectUsecase_Login_Abnormal
// Summary: This is abnormal test class which confirm the operation of API Login.
// Target: auth_usecase_impl.go
// TestPattern:
// [x] 2-1. 500: 検証処理エラー
// [x] 2-2. 401: アクセストークン払い出し失敗
// [x] 2-3. 401: リフレッシュトークン払い出し失敗
func TestProjectUsecase_Login_Abnormal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/login"

	resNoAccessToken := authentication.LoginResult{
		AccessToken:  "",
		RefreshToken: "valid_token",
	}
	resNoRefreshToken := authentication.LoginResult{
		AccessToken:  "valid_token",
		RefreshToken: "",
	}
	tests := []struct {
		name         string
		input        input.LoginParam
		receive      authentication.LoginResult
		receiveError error
		expect       error
	}{
		{
			name:         "2-1. 200: 検証処理エラー",
			input:        f.NewInputLoginParam(),
			receiveError: fmt.Errorf("検証処理エラー"),
			expect:       fmt.Errorf("検証処理エラー"),
		},
		{
			name:    "2-2. 401: アクセストークン払い出し失敗",
			input:   f.NewInputLoginParam(),
			receive: resNoAccessToken,
			expect:  common.NewCustomError(common.CustomErrorCode401, common.Err401InvalidCredentials, nil, common.HTTPErrorSourceAuth),
		},
		{
			name:    "2-3. 401: リフレッシュトークン払い出し失敗",
			input:   f.NewInputLoginParam(),
			receive: resNoRefreshToken,
			expect:  common.NewCustomError(common.CustomErrorCode401, common.Err401InvalidCredentials, nil, common.HTTPErrorSourceAuth),
		},
		{
			// 実際に401エラーが返却された挙動のテスト
			name:         "2-4. 401: クライアント認証失敗",
			input:        f.NewInputLoginParam(),
			receiveError: fmt.Errorf("401 Unauthorized: invalid_grant: Invalid user credentials"),
			expect:       common.NewCustomError(common.CustomErrorCode401, common.Err401InvalidCredentials, nil, common.HTTPErrorSourceAuth),
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
				gocloakRepositoryMock.On("SignInWithPassword", mock.Anything, mock.Anything).Return(test.receive, test.receiveError)
				authusecase := usecase.NewAuthUsecase(gocloakRepositoryMock)

				_, err := authusecase.Login(test.input)
				if assert.Error(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// TestProjectUsecase_Client
// Summary: This is normal test class which confirm the operation of API Client.
// Target: auth_usecase_impl.go
// TestPattern:
// [x] 1-1. 200: 正常系
func TestProjectUsecase_Client(tt *testing.T) {

	var method = "GET"
	var endPoint = "/client"

	res := authentication.ClientResult{
		AccessToken: f.Token,
	}
	expected := output.ClientResponse{
		AccessToken: f.Token,
	}
	tests := []struct {
		name    string
		input   input.ClientParam
		receive authentication.ClientResult
		expect  output.ClientResponse
	}{
		{
			name:    "1-1. 200: 正常系",
			input:   f.NewInputClientParam(),
			receive: res,
			expect:  expected,
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
				gocloakRepositoryMock.On("SignInWithClient", mock.Anything, mock.Anything).Return(test.receive, nil)
				authusecase := usecase.NewAuthUsecase(gocloakRepositoryMock)

				actual, err := authusecase.Client(test.input)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expect.AccessToken, actual.AccessToken, f.AssertMessage)
				}
			},
		)
	}
}

// TestProjectUsecase_Client_Abnormal
// Summary: This is abnormal test class which confirm the operation of API Login.
// Target: auth_usecase_impl.go
// TestPattern:
// [x] 2-1. 500: 検証処理エラー
// [x] 2-2. 401: アクセストークン払い出し失敗
func TestProjectUsecase_Client_Abnormal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/client"

	resNoAccessToken := authentication.ClientResult{
		AccessToken: "",
	}
	tests := []struct {
		name         string
		input        input.ClientParam
		receive      authentication.ClientResult
		receiveError error
		expect       error
	}{
		{
			name:         "2-1. 200: 検証処理エラー",
			input:        f.NewInputClientParam(),
			receiveError: fmt.Errorf("検証処理エラー"),
			expect:       fmt.Errorf("検証処理エラー"),
		},
		{
			// 401エラーは発生しないがアクセストークンが取得できない場合401を返すようにする挙動のテスト
			name:    "2-2. 401: アクセストークン払い出し失敗",
			input:   f.NewInputClientParam(),
			receive: resNoAccessToken,
			expect:  common.NewCustomError(common.CustomErrorCode401, common.Err401InvalidCredentials, nil, common.HTTPErrorSourceAuth),
		},
		{
			// 実際に401エラーが返却された挙動のテスト
			name:         "2-3. 401: クライアント認証失敗",
			input:        f.NewInputClientParam(),
			receiveError: fmt.Errorf("401 Unauthorized: invalid_grant: Invalid user credentials"),
			expect:       common.NewCustomError(common.CustomErrorCode401, common.Err401InvalidCredentials, nil, common.HTTPErrorSourceAuth),
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
				gocloakRepositoryMock.On("SignInWithClient", mock.Anything, mock.Anything).Return(test.receive, test.receiveError)
				authusecase := usecase.NewAuthUsecase(gocloakRepositoryMock)

				_, err := authusecase.Client(test.input)
				if assert.Error(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// TestProjectUsecase_Refresh
// Summary: This is normal test class which confirm the operation of API Refresh.
// Target: auth_usecase_impl.go
// TestPattern:
// [x] 1-1. 200: 正常系
func TestProjectUsecase_Refresh(tt *testing.T) {

	var method = "GET"
	var endPoint = "/refresh"

	expected := output.RefreshResponse{
		AccessToken: "valid_token",
	}
	tests := []struct {
		name               string
		openOperatorSearch bool
		input              input.RefreshParam
		receive            string
		expect             output.RefreshResponse
	}{
		{
			name:               "1-1. 200: 正常系",
			openOperatorSearch: true,
			input:              f.NewInputRefreshParam(),
			receive:            "valid_token",
			expect:             expected,
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

				GoCloakRepositoryMock := new(mocks.GoCloakRepository)
				GoCloakRepositoryMock.On("RefreshToken", mock.Anything).Return(test.receive, nil)
				authusecase := usecase.NewAuthUsecase(GoCloakRepositoryMock)

				actual, err := authusecase.Refresh(test.input)
				if assert.NoError(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					// 順番が実行ごとに異なるため、順不同で中身を比較
					assert.Equal(t, test.expect.AccessToken, actual.AccessToken, f.AssertMessage)
				}
			},
		)
	}
}

// TestProjectUsecase_Refresh_Abnormal
// Summary: This is abnormal test class which confirm the operation of API Refresh.
// Target: auth_usecase_impl.go
// TestPattern:
// [x] 2-1. 500: 検証処理エラー
// [x] 2-2. 401: アクセストークン払い出し失敗
func TestProjectUsecase_Refresh_Abnormal(tt *testing.T) {

	var method = "GET"
	var endPoint = "/refresh"

	tests := []struct {
		name         string
		input        input.RefreshParam
		receive      string
		receiveError error
		expect       error
	}{
		{
			name:         "2-1. 200: 検証処理エラー",
			input:        f.NewInputRefreshParam(),
			receiveError: fmt.Errorf("検証処理エラー"),
			expect:       fmt.Errorf("検証処理エラー"),
		},
		{
			name:    "2-2. 401: アクセストークン払い出し失敗",
			input:   f.NewInputRefreshParam(),
			receive: "",
			expect:  common.NewCustomError(common.CustomErrorCode401, common.Err401InvalidCredentials, nil, common.HTTPErrorSourceAuth),
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

				GoCloakRepositoryMock := new(mocks.GoCloakRepository)
				GoCloakRepositoryMock.On("RefreshToken", mock.Anything).Return(test.receive, test.receiveError)
				authusecase := usecase.NewAuthUsecase(GoCloakRepositoryMock)

				_, err := authusecase.Refresh(test.input)
				if assert.Error(t, err) {
					// 実際のレスポンスと期待されるレスポンスを比較
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}

// TestProjectUsecase_ChangePassword
// Summary: This is normal test class which confirm the operation of API Change Password.
// Target: auth_usecase_impl.go
// TestPattern:
// [x] 1-1. 200: 正常系
func TestProjectUsecase_ChangePassword(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/auth/v1/authInfo"

	tests := []struct {
		name    string
		input   input.ChangePasswordParam
		receive error
		expect  error
	}{
		{
			name:  "1-1. 200: 正常系",
			input: f.NewInputChangePasswordParam(),
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
				gocloakRepositoryMock.On("ChangePassword", mock.Anything, mock.Anything).Return(test.receive)
				authusecase := usecase.NewAuthUsecase(gocloakRepositoryMock)

				err := authusecase.ChangePassword(test.input)
				assert.NoError(t, err)
			},
		)
	}
}

// TestProjectUsecase_ChangePassword_Abnormal
// Summary: This is abnormal test class which confirm the operation of API Change Password.
// Target: auth_usecase_impl.go
// TestPattern:
// [x] 2-1. 500: 変更処理エラー
func TestProjectUsecase_ChangePassword_Abnormal(tt *testing.T) {

	var method = "PUT"
	var endPoint = "/auth/v1/authInfo"

	tests := []struct {
		name    string
		input   input.ChangePasswordParam
		receive error
		expect  error
	}{
		{
			name:    "2-1. 500: 変更処理エラー",
			input:   f.NewInputChangePasswordParam(),
			receive: fmt.Errorf("変更処理エラー"),
			expect:  fmt.Errorf("変更処理エラー"),
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
				gocloakRepositoryMock.On("ChangePassword", mock.Anything, mock.Anything).Return(test.receive)
				authusecase := usecase.NewAuthUsecase(gocloakRepositoryMock)

				err := authusecase.ChangePassword(test.input)
				if assert.Error(t, err) {
					assert.Equal(t, test.expect.Error(), err.Error())
				}
			},
		)
	}
}
