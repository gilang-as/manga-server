package http_handler

import (
	"github.com/labstack/echo/v4"
	http2 "manga-server/pkg/http"
	"net/http"
)

func (h *HttpHandler) GetTest(c echo.Context) error {
	return c.JSON(http.StatusOK, http2.BaseResponse{
		Message: "Hello Word",
		Status:  true,
		Data:    nil,
	})
}