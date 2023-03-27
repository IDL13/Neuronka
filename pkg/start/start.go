package start

import (
	"github.com/IDL13/Neuronka/internal/activation"
	"github.com/IDL13/Neuronka/internal/create"
	"github.com/fxsjy/gonn/gonn"
)

func Start(x1, y1, x2, y2, x3, y3 float64) string {
	create.CreateNN()
	nn := gonn.LoadNN("gonn")

	out := nn.Forward([]float64{x1, y1, x2, y2, x3, y3})
	return activation.Activation(out)
}
