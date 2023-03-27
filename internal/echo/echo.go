package echo

import (
	"fmt"
	"net/http"

	"github.com/IDL13/Neuronka/internal/encryption"
	"github.com/IDL13/Neuronka/pkg/start"
	"github.com/labstack/echo/v4"
)

func StartServer() {
	e := echo.New()

	e.GET("/", Start)
	e.POST("/request", Request)

	e.Logger.Fatal(e.Start(":8080"))
}

func Start(c echo.Context) error {
	return c.String(http.StatusOK, "[SERVER_STARTED]")
}

func Request(c echo.Context) error {
	en := encryption.New()
	m, err := en.Encryption(c)
	if err != nil {
		fmt.Println("err")
	}
	s := start.Start(m.X1, m.Y1, m.X2, m.Y2, m.X3, m.Y3)
	err = c.String(http.StatusOK, s)
	if err != nil {
		fmt.Println("error from Request")
	}
	return nil
}
