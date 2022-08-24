package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/goml/gobrain"
)

func loadData() ([][]float64, []string, error) {
	f, err := os.Open("./iris.csv")
	if err != nil {
		return nil, nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	// skip header
	scanner.Scan()

	var resultf [][]float64
	var results []string
	for scanner.Scan() {
		var f1, f2, f3, f4 float64
		var s string
		n, err := fmt.Sscanf(scanner.Text(), "%f,%f,%f,%f,%s", &f1, &f2, &f3, &f4, &s)
		if n != 5 || err != nil {
			return nil, nil, errors.New("cannot load data")
		}
		resultf = append(resultf, []float64{f1, f2, f3, f4})
		results = append(results, strings.Trim(s, `"`))

	}

	return resultf, results, nil
}

func main() {
	X, Y, err := loadData()
	if err != nil {
		log.Fatal(err)
	}
	_ = X
	_ = Y

	patterns := [][][]float64{}

	m := map[string][]float64{
		"Setosa":     {1, 0, 0},
		"Versicolor": {0, 1, 0},
		"Virginica":  {0, 0, 1},
	}

	for i, x := range X {
		patterns = append(patterns, [][]float64{
			x, m[Y[i]],
		})
	}

	ff := &gobrain.FeedForward{}
	ff.Init(4, 3, 3)
	ff.Train(patterns, 100000, 0.6, 0.4, true)

	result := ff.Update(X[0])

	fmt.Println(result)
}
