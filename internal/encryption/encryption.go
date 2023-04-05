package encryption

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/labstack/echo/v4"
)

func New() *Mas {
	return &Mas{}
}

type Mas struct {
	X1 float64 `json:"x1"`
	Y1 float64 `json:"y1"`
	X2 float64 `json:"x2"`
	Y2 float64 `json:"y2"`
	X3 float64 `json:"x3"`
	Y3 float64 `json:"y3"`
}

func (m *Mas) Encryption(c echo.Context) (*Mas, error) { // Получение и распарсинг json
	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Failed reading the request body: %s", err)
	}

	err = json.Unmarshal(b, &m)

	return m, nil
}
