package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"manga-server/config"
	"manga-server/handler/graphql_handler"
	"manga-server/handler/http_handler"
	"manga-server/pkg/http"
	"manga-server/pkg/mysql"
	"manga-server/repository/mysql_repository"
	"manga-server/usecase"
)

func main() {
	environment := config.GetEnvironment()

	db := mysql.Connect(environment.DBHost, environment.DBUsername, environment.DBPassword, environment.DBName)
	defer db.Close()
	//config.DropAll(db)
	config.DBMigrate(db)

	anihup_repository := mysql_repository.NewMysqlRepository(db)
	anihup_usecase := usecase.NewUseCaseImpl(anihup_repository)
	anihup_restapi := http_handler.NewHttpHandler()
	anihup_graphql := graphql_handler.NewHandler(anihup_usecase)

	e := http.NewEchoHTTPServer()

	e.SetupRoutes(func(s *http.Server) {
		config.Routes(s, anihup_restapi)
	})

	e.Echo.POST("/graphql", echo.WrapHandler(anihup_graphql))
	e.Echo.GET("/graphql", echo.WrapHandler(anihup_graphql))

	fmt.Println(e.Start("127.0.0.1:" + environment.HostPort))
}

func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}