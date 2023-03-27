package start

import (
	"github.com/IDL13/Neuronka/internal/activation"
	"github.com/IDL13/Neuronka/internal/create"
	"github.com/IDL13/Neuronka/internal/parser"
	"github.com/fxsjy/gonn/gonn"
)

func Start() string {
	create.CreateNN()
	nn := gonn.LoadNN("gonn")
	priceMatrix := parser.Parser("parser.txt")

	var x1 float64 = priceMatrix[1][0]
	var y1 float64 = priceMatrix[0][0]
	var x2 float64 = priceMatrix[1][1]
	var y2 float64 = priceMatrix[0][1]
	var x3 float64 = priceMatrix[1][2]
	var y3 float64 = priceMatrix[0][2]

	out := nn.Forward([]float64{x1, y1, x2, y2, x3, y3})
	return activation.Activation(out)
}
