package main

import (
	"embed"
	"net/http"

	"github.com/labstack/echo"
)

//go:embed static
var local embed.FS

func main() {
	e := echo.New()
	e.GET("/", echo.WrapHandler(http.FileServer(http.FS(local))))
	e.Logger.Fatal(e.Start(":8989"))
}
