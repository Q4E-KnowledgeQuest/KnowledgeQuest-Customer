package server

import (
	"fmt"
	"main/components/courses"
	"net/http"

	"github.com/labstack/echo/v4"
	middleware "github.com/labstack/echo/v4/middleware"
)

var e *echo.Echo

func Start(port int) {
	e = echo.New()
	e.HideBanner = true

	e.Use(middleware.Recover())

	DefaultCORSConfig := middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}

	e.Use(middleware.CORSWithConfig(DefaultCORSConfig))

	initializeRoutes()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

func initializeRoutes() {
	e.Static("/", "public")
	e.GET("/register/:key", func (c echo.Context) error {
		key := c.Param("key")
		output, err := courses.RegisterLicense(key)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		courses.DownloadCourses()
		return c.String(http.StatusOK, output)
	})
}
