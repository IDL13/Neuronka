package main

import (
	"fmt"

	"github.com/IDL13/Neuronka/internal/activation"
	"github.com/IDL13/Neuronka/internal/create"
	"github.com/IDL13/Neuronka/internal/requests"
	"github.com/fxsjy/gonn/gonn"
)

func main() {
	create.CreateNN()
	nn := gonn.LoadNN("gonn")

	var x1 float64 = 1.0
	var y1 float64 = 5.0
	var x2 float64 = 2.0
	var y2 float64 = 4.0
	var x3 float64 = 5.0
	var y3 float64 = 5.0

	out := nn.Forward([]float64{x1, y1, x2, y2, x3, y3})
	fmt.Println(activation.Activation(out))
}

func New(arr_x, arr_y []float64) {
	create.CreateNN()
	nn := gonn.LoadNN("gonn")
	r := requests.New()

	m := map[string]float64{"x1": arr_x[0], "x2": arr_x[1], "x3": arr_x[2], "y1": arr_y[0], "y2": arr_y[1], "y3": arr_y[2]}

	date := r.Marshal(m)
	r.Unmarshal(date)

	out := nn.Forward([]float64{r.X1, r.Y1, r.X2, r.Y2, r.X3, r.Y3})
	fmt.Println(activation.Activation(out))
}
