package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/shiba-hiro/note-api-by-go/handler"
	"github.com/shiba-hiro/note-api-by-go/repository"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.SetLevel(log.INFO)
	e.Logger.SetOutput(os.Stdout)

	e.GET("/health", health)
	handler.Register(e.Group("/api/v1"))

	db, err := repository.OpenDbConnection()
	if err != nil {
		e.Logger.Fatalf("Cannot open Database: %v\n", db)
		return
	}
	defer db.Close()

	e.Logger.Fatal(e.Start(":1323"))
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "{\"app\": {\"message\": \"Application is running\", \"success\": true}}")
}
