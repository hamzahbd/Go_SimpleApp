package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		// Mencetak informasi tentang bahasa dan versi Go sebelum menampilkan pesan "Hello, Docker!"
		goVersion := runtime.Version()
		fmt.Println("Bahasa Pemrograman: Go")
		fmt.Println("Versi Go:", goVersion)

		return c.HTML(http.StatusOK, fmt.Sprintf("Hello, Docker! <3<br> Bahasa Pemrograman: Go<br> Versi Go: %s", goVersion))
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
