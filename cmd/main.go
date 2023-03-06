package main

import (
	"fmt"

	"github.com/IDL13/Neuronka/internal/activation"
	"github.com/IDL13/Neuronka/internal/create"
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
