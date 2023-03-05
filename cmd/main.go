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

	var x float64 = 0.1
	var y float64 = 0.1

	out := nn.Forward([]float64{x, y})
	fmt.Println(activation.Activation(out))
}
