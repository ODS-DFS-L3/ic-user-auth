package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"authenticator-backend/domain/common"
	f "authenticator-backend/test/fixtures"
	mocks "authenticator-backend/test/mock"
	"authenticator-backend/usecase/input"
	"authenticator-backend/usecase/output"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// /////////////////////////////////////////////////////////////////////////////////
// POST /api/v2/systemAuth/token のテストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 正常系
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_SystemAuthToken_Normal(tt *testing.T) {
	var method = "POST"
	var endPoint = "/api/v2/systemAuth/token"

	tests := []struct {
		name         string
		input        input.VerifyTokenParam
		expectStatus int
	}{
		{
			name: "1-1. 200: 正常系",
			input: input.VerifyTokenParam{
				IDToken: f.Token,
			},
			expectStatus: http.StatusOK,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()
				inputJSON, _ := json.Marshal(test.input)
				q := make(url.Values)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), strings.NewReader(string(inputJSON)))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Set("operatorID", f.OperatorId)

				authUsecase := new(mocks.IAuthUsecase)
				verifyUsecase := new(mocks.IVerifyUsecase)
				res := output.VerifyTokenResponse{
					OperatorID:   f.OperatorId,
					OpenSystemId: f.OpenSystemId,
					Active:       f.Active,
				}
				authHandler := NewAuthHandler(
					authUsecase,
					verifyUsecase,
				)
				verifyUsecase.On("TokenIntrospection", test.input).Return(res, nil)

				err := authHandler.TokenIntrospection(c)
				if assert.NoError(t, err) {
					assert.Equal(t, test.expectStatus, rec.Code)
				}
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// POST /api/v2/systemAuth/token のテストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 400: バリデーションエラー: idTokenが含まれていない場合
// [x] 1-2. 400: バリデーションエラー: idTokenがstring形式でない場合
// [x] 1-3. 401: 認証情報無効エラー：資格情報が無効
// [x] 1-4. 500: 内部システムエラー：返却値バリデーションエラー（operatorIdが36桁未満）
// [x] 1-5. 500: 内部システムエラー：返却値バリデーションエラー（operatorIdが36桁超過）
// [x] 1-6. 500: 内部システムエラー：返却値バリデーションエラー（operatorIdが非UUID形式）
// [x] 1-7. 500: 内部システムエラー：返却値バリデーションエラー（openSystemIdが21桁未満）
// [x] 1-8. 500: 内部システムエラー：返却値バリデーションエラー（openSystemIdが36桁超過）
// [x] 1-9. 500: その他：想定外のエラー（Keycloak通信中に通信が切断）
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_SystemAuthToken(tt *testing.T) {
	var method = "POST"
	var endPoint = "/api/v2/systemAuth/token"

	tests := []struct {
		name         string
		input        input.VerifyTokenParam
		output       output.VerifyTokenResponse
		invalidInput any
		receive      error
		expectError  string
		expectStatus int
	}{
		{
			name: "1-1. 400: バリデーションエラー：idTokenが含まれていない場合",
			input: input.VerifyTokenParam{
				IDToken: "",
			},
			output:       output.VerifyTokenResponse{},
			expectError:  "code=400, message={[auth] BadRequest Validation failed, idToken: cannot be blank.",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-2. 400: バリデーションエラー：idTokenがstring形式でない場合",
			invalidInput: struct {
				IdToken int
			}{
				1,
			},
			output:       output.VerifyTokenResponse{},
			expectError:  "code=400, message={[auth] BadRequest Invalid request parameters",
			expectStatus: http.StatusBadRequest,
		},
		{
			name: "1-3. 401: 認証情報無効エラー：資格情報が無効",
			input: input.VerifyTokenParam{
				IDToken: f.Token,
			},
			output:       output.VerifyTokenResponse{},
			receive:      common.NewCustomError(common.CustomErrorCode401, common.Err401InvalidCredentials, nil, common.HTTPErrorSourceAuth),
			expectError:  "code=401, message={[auth] Unauthorized",
			expectStatus: http.StatusUnauthorized,
		},
		{
			name: "1-4. 500: 内部システムエラー：返却値バリデーションエラー（operatorIdが36桁未満）",
			input: input.VerifyTokenParam{
				IDToken: f.Token,
			},
			output: output.VerifyTokenResponse{
				OperatorID:   "123456789012345678901234567890-abcd",
				OpenSystemId: f.OpenSystemId,
				Active:       f.Active,
			},
			receive:      nil,
			expectError:  "code=500, message={[auth] InternalServerError",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-5. 500: 内部システムエラー：返却値バリデーションエラー（operatorIdが36桁超過）",
			input: input.VerifyTokenParam{
				IDToken: f.Token,
			},
			output: output.VerifyTokenResponse{
				OperatorID:   "1234567890123456789012345678901234567",
				OpenSystemId: f.OpenSystemId,
				Active:       f.Active,
			},
			receive:      nil,
			expectError:  "code=500, message={[auth] InternalServerError",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-6. 500: 内部システムエラー：返却値バリデーションエラー（operatorIdが非UUID形式）",
			input: input.VerifyTokenParam{
				IDToken: f.Token,
			},
			output: output.VerifyTokenResponse{
				OperatorID:   "e03cc699-7234-31ed-86be0cc18c92208e5",
				OpenSystemId: f.OpenSystemId,
				Active:       f.Active,
			},
			receive:      nil,
			expectError:  "code=500, message={[auth] InternalServerError",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-7. 500: 内部システムエラー：返却値バリデーションエラー（openSystemIdが21桁未満）",
			input: input.VerifyTokenParam{
				IDToken: f.Token,
			},
			output: output.VerifyTokenResponse{
				OperatorID:   f.OperatorId,
				OpenSystemId: "1234567890123456789o",
				Active:       f.Active,
			},
			receive:      nil,
			expectError:  "code=500, message={[auth] InternalServerError",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-8. 500: 内部システムエラー：返却値バリデーションエラー（openSystemIdが36桁超過）",
			input: input.VerifyTokenParam{
				IDToken: f.Token,
			},
			output: output.VerifyTokenResponse{
				OperatorID:   f.OperatorId,
				OpenSystemId: "1234567890123456789012345678901234567",
				Active:       f.Active,
			},
			receive:      nil,
			expectError:  "code=500, message={[auth] InternalServerError",
			expectStatus: http.StatusInternalServerError,
		},
		{
			name: "1-9. 500: その他：想定外のエラー（Keycloak通信中に通信が切断）",
			input: input.VerifyTokenParam{
				IDToken: f.Token,
			},
			output:       output.VerifyTokenResponse{},
			receive:      echo.ErrNotFound,
			expectError:  "code=500, message={[auth] InternalServerError",
			expectStatus: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		test := test
		tt.Run(
			test.name,
			func(t *testing.T) {
				t.Parallel()
				var inputJSON []byte
				if test.invalidInput != nil {
					inputJSON, _ = json.Marshal(test.invalidInput)
				} else {
					inputJSON, _ = json.Marshal(test.input)
				}
				q := make(url.Values)

				e := echo.New()
				rec := httptest.NewRecorder()
				req := httptest.NewRequest(method, endPoint+"?"+q.Encode(), strings.NewReader(string(inputJSON)))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Set("operatorID", f.OperatorId)

				authUsecase := new(mocks.IAuthUsecase)
				verifyUsecase := new(mocks.IVerifyUsecase)

				authHandler := NewAuthHandler(
					authUsecase,
					verifyUsecase,
				)
				verifyUsecase.On("TokenIntrospection", test.input).Return(test.output, test.receive)

				err := authHandler.TokenIntrospection(c)
				e.HTTPErrorHandler(err, c)
				if err != nil {
					assert.Equal(t, test.expectStatus, rec.Code)
					assert.ErrorContains(t, err, test.expectError)
				} else {
					assert.Equal(t, test.expectStatus, rec.Code)
				}
			},
		)
	}
}
