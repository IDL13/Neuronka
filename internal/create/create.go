package create

import (
	"github.com/fxsjy/gonn/gonn"
)

func CreateNN() {
	nn := gonn.DefaultNetwork(6, 18, 2, false)

	input := [][]float64{
		[]float64{0.0, 0.0, 1.0, 1.0, 2.0, 2.0}, []float64{1.0, 1.0, 2.0, 2.0, 3.0, 1.0}, []float64{1.0, 1.0, 2.0, 2.0, 3.0, 3.0},
		[]float64{2.0, 4.0, 3.0, 3.0, 4.0, 2.0}, []float64{1.0, 5.0, 2.0, 4.0, 3.0, 5.0}}

	target := [][]float64{
		[]float64{1, 0}, []float64{0, 1}, []float64{1, 0},
		[]float64{0, 1}, []float64{1, 0}}

	nn.Train(input, target, 10000)

	gonn.DumpNN("gonn", nn)
}
