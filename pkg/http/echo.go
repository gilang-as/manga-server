package http

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)


type Server struct {
	Echo *echo.Echo
}


func NewEchoHTTPServer() Server {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Validator = &CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		if castedObject, ok := err.(validator.ValidationErrors); ok {
			for _, err := range castedObject {
				switch err.Tag() {
				case "required":
					report.Message = fmt.Sprintf("%s is required",
						err.Field())
				case "email":
					report.Message = fmt.Sprintf("%s is not valid email",
						err.Field())
				case "gte":
					report.Message = fmt.Sprintf("%s value must be greater than %s",
						err.Field(), err.Param())
				case "lte":
					report.Message = fmt.Sprintf( "%s value must be lower than %s",
						err.Field(), err.Param())
				}
				break
			}
		}
		c.Logger().Error(report)
		c.JSON(report.Code, report)
	}
	return Server{
		Echo: e,
	}
}

func (s *Server) SetupRoutes(fc func(s *Server)) {
	fc(s)
}




func (s *Server) Start(address string) error{
	return s.Echo.Start(address)
}