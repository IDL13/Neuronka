package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MatrixCompresion(fileName string) [][]float64 {
	inp, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("error")
	}

	r := strings.NewReader(string(inp))
	s := bufio.NewScanner(r)

	var matrix [][]float64

	for s.Scan() {
		records := strings.Fields(s.Text())
		line := make([]float64, len(records))
		matrix = append(matrix, line)
		for i := range records {
			line[i], err = strconv.ParseFloat(records[i], 64)
			if err != nil {
				panic(err)
			}
		}

		if err = s.Err(); err != nil {
			panic(err)
		}
	}

	return matrix
}
