package create

import (
	"github.com/IDL13/Neuronka/internal/config"
	"github.com/fxsjy/gonn/gonn"
)

func CreateNN() {
	nn := gonn.DefaultNetwork(6, 18, 2, false)

	config := config.New()
	config.Read("input.txt", "target.txt")

	input := config.Input
	target := config.Target

	nn.Train(input, target, 10000)

	gonn.DumpNN("gonn", nn)
}
