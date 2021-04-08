package config

import (
	"manga-server/handler/http_handler"
	"manga-server/pkg/http"
)

func Routes(s *http.Server, handler http_handler.HttpHandler) {
	s.Echo.GET("/test", handler.GetTest)

}