package handler_test

import (
	"authenticator-backend/presentation/http/echo/handler"
	f "authenticator-backend/test/fixtures"
	mocks "authenticator-backend/test/mock"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// /////////////////////////////////////////////////////////////////////////////////
// Put /api/v2/ テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-1. 200: 正常系：plantの場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_PutOparator(tt *testing.T) {
	var method = "PUT"
	var endPoint = "/api/v2/authinfo/plant"

	tests := []struct {
		name string
	}{
		{
			name: "1-1. 201: 正常系：operatorの場合",
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
				c.Set("operatorID", f.OperatorId)

				operatorHandler := new(mocks.IOperatorHandler)
				operatorHandler.On("PutOperator", mock.Anything).Return(nil)
				plantHandler := new(mocks.IPlantHandler)
				plantHandler.On("PutPlant", mock.Anything).Return(nil)
				resetHandler := new(mocks.IResetHandler)
				h := handler.NewOuranosHandler(operatorHandler, plantHandler, resetHandler)

				err := h.PutOperator(c)
				assert.NoError(t, err)
			},
		)
	}
}

// /////////////////////////////////////////////////////////////////////////////////
// Put /api/v2/ テストケース
// /////////////////////////////////////////////////////////////////////////////////
// [x] 1-2. 200: 正常系：plantの場合
// /////////////////////////////////////////////////////////////////////////////////
func TestProjectHandler_PutPlant(tt *testing.T) {
	var method = "PUT"
	var endPoint = "/api/v2/authInfo/plant"

	tests := []struct {
		name string
	}{
		{
			name: "1-1. 201: 正常系：plantの場合",
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
				req := httptest.NewRequest(method, endPoint+q.Encode(), nil)
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				c := e.NewContext(req, rec)
				c.SetPath(endPoint)
				c.Set("operatorID", f.OperatorId)

				operatorHandler := new(mocks.IOperatorHandler)
				operatorHandler.On("PutOperator", mock.Anything).Return(nil)
				plantHandler := new(mocks.IPlantHandler)
				plantHandler.On("PutPlant", mock.Anything).Return(nil)
				resetHandler := new(mocks.IResetHandler)
				h := handler.NewOuranosHandler(operatorHandler, plantHandler, resetHandler)

				err := h.PutPlant(c)
				assert.NoError(t, err)
			},
		)
	}
}
