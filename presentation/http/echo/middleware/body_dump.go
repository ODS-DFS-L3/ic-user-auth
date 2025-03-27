package middleware

import (
	"authenticator-backend/domain/common"
	"authenticator-backend/extension/logger"

	"github.com/labstack/echo/v4"
)

// dumpHandler
// Summary: This is the function which dump the request and response body.
// input: c(echo.Context): echo context
// input: reqBody([]byte): request body
// input: resBody([]byte): response body
func dumpHandler(c echo.Context, reqBody, resBody []byte) {
	var operatorID string
	i := c.Get("operatorID")
	if i != nil {
		operatorID = i.(string)
	}
	logger.Set(c).Infof("DataSpaceAPI Access Path: %v, OperatorId: %v, Header: %v, Request Body: %v, Response Body: %v", c.Request().URL.String(), operatorID, c.Request().Header, string(reqBody), string(resBody))
}

// dumpSkipper
// Summary: This is the function which skips the dump.
// output: (bool) true if skip, false otherwise
func dumpSkipper(c echo.Context) bool {
	return !common.IsOutputDump()
}
