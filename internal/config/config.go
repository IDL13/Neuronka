package config

import (
	"github.com/IDL13/Neuronka/pkg/utils"
)

func New() *DataSet {
	return &DataSet{}
}

type DataSet struct {
	Input  [][]float64
	Target [][]float64
}

func (d *DataSet) Read(fileNameInput, fileNameTarget string) *DataSet {
	d.Input = utils.MatrixCompresion(fileNameInput)
	d.Target = utils.MatrixCompresion(fileNameTarget)
	return d
}
