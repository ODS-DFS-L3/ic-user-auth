package handler

import "github.com/labstack/echo/v4"

func (h *ouranosHandler) PutOperator(c echo.Context) error {
	return h.operatorHandler.PutOperator(c)
}

func (h *ouranosHandler) PutPlant(c echo.Context) error {
	return h.plantHandler.PutPlant(c)
}
