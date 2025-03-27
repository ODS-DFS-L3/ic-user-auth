package handler

import "github.com/labstack/echo/v4"

func (h *ouranosHandler) GetOperator(c echo.Context) error {
	return h.operatorHandler.GetOperator(c)
}

func (h *ouranosHandler) GetPlant(c echo.Context) error {
	return h.plantHandler.GetPlant(c)
}
