// Package fame is a media catalog package.
package fame

import (
	"net/http"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

// Config sets the Fame package up
type Config struct {
	Port uint
}

func hello(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

// Run is the main function
func Run(c Config) {
	// Handler

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// Routes
	e.Get("/", hello)

	// Start server
	e.Run(":8097")
}
