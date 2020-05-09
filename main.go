package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/shiba-hiro/note-api-by-go/handler"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.SetLevel(log.INFO)
	e.Logger.SetOutput(os.Stdout)

	e.GET("/health", health)
	handler.Register(e.Group("/api/v1"))

	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func health(c echo.Context) error {
	return c.String(http.StatusOK, "{\"app\": {\"message\": \"Application is running\", \"success\": true}}")
}
