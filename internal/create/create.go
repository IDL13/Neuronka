package create

import (
	"github.com/fxsjy/gonn/gonn"
)

func CreateNN() {
	nn := gonn.DefaultNetwork(2, 16, 2, false)

	input := [][]float64{}

	target := [][]float64{}

	nn.Train(input, target, 10000)

	gonn.DumpNN("gonn", nn)
}
