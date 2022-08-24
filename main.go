package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/goml/gobrain"
	"github.com/goml/gobrain/persist"
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

	err = persist.Load("model.json", &ff)

	if err != nil {
		ff.Train(patterns, 100000, 0.6, 0.04, true)
		persist.Save("model.json", &ff)
	}

	result := ff.Update(X[0])

	var mf float64
	var mi int

	for i, v := range result {
		if mf < v {
			mf = v
			mi = i
		}
	}
	fmt.Println([]string{"Setosa", "Versicolor", "Virginica"}[mi])
}
