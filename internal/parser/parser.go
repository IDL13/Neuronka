package parser

import (
	"github.com/IDL13/Neuronka/pkg/utils"
)

func Parser(fileName string) [][]float64 {
	priceMatrix := utils.MatrixCompresion(fileName)
	return priceMatrix
}

// func main() {
// 	fileName := "parser.txt"
// 	priceMatrix := utils.MatrixCompresion(fileName)
// 	fmt.Println(priceMatrix)
// }
