package handler

import (
	"errors"
	"net/http"

	"authenticator-backend/domain/common"
	"authenticator-backend/extension/logger"
	"authenticator-backend/usecase/input"

	"github.com/labstack/echo/v4"
)

// TokenIntrospection
// Summary: This is the function which verifies the token.
// input: c(echo.Context): echo context
// output: (error) error object
func (h *authHandler) TokenIntrospection(c echo.Context) error {
	var param input.VerifyTokenParam
	method := c.Request().Method

	if err := c.Bind(&param); err != nil {
		logger.Set(c).Warnf(err.Error())
		errDetails := common.FormatBindErrMsg(err)

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceAuth, common.Err400InvalidRequest, "", "", method, errDetails))
	}
	if err := param.Validate(); err != nil {
		logger.Set(c).Warnf(err.Error())

		errDetails := err.Error()

		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusBadRequest, common.HTTPErrorSourceAuth, common.Err400Validation, "", "", method, errDetails))
	}

	output, err := h.VerifyUsecase.TokenIntrospection(param)
	if err != nil {
		var customErr *common.CustomError
		if errors.As(err, &customErr) {
			logger.Set(c).Errorf(err.Error())
			// Return a custom error in case of a 401 error
			if customErr.Code == http.StatusUnauthorized {
				return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusUnauthorized, common.HTTPErrorSourceAuth, common.Err401InvalidCredentials, "", "", method))
			}
		}
		// Return a 500 error for anything other than a custom error
		logger.Set(c).Errorf(err.Error())
		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusInternalServerError, common.HTTPErrorSourceAuth, common.Err500Unexpected, "", "", method))
	}

	if err := output.Validate(); err != nil {
		logger.Set(c).Errorf(err.Error())
		errDetails := err.Error()
		// Return a 500 error in case of a validation error
		return echo.NewHTTPError(common.HTTPErrorGenerate(http.StatusInternalServerError, common.HTTPErrorSourceAuth, common.Err500Unexpected, "", "", method, errDetails))
	}

	return c.JSON(http.StatusOK, output)
}
