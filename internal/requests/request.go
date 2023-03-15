package requests

import (
	"encoding/json"
	"fmt"
)

func New() *Request {
	return &Request{}
}

type Request struct {
	X1 float64 `json:"x1"`
	X2 float64 `json:"x2"`
	X3 float64 `json:"x3"`
	Y1 float64 `json:"y1"`
	Y2 float64 `json:"y2"`
	Y3 float64 `json:"y3"`
}

func (r *Request) Marshal(inf map[string]float64) []byte {
	date, err := json.Marshal(inf)
	if err != nil {
		fmt.Println(err)
	}
	return date
}

func (r *Request) Unmarshal(b []byte) {
	err := json.Unmarshal(b, &r)
	if err != nil {
		fmt.Println(err)
	}

}
