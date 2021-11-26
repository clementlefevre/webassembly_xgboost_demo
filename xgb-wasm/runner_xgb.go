package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/dmitryikh/leaves"
	"github.com/dmitryikh/leaves/mat"
)

// RegressionRunner stores input data for the file
type XGBRunner struct {
	records     [][]string  // array of records
	y           []float64   // array of parsed data (y only)
	x           [][]float64 // array of parsed data (x only)
	header      []string    // first row of headers
	predictCol  int         // prediction column
	residuals   []float64   // final array of residuals
	predictions []float64
	dense       *mat.DenseMat    // Dense Matrix from X Data
	model       *leaves.Ensemble // the xgb model
}

// load XGB  model to save to runner
func (runner *XGBRunner) readModel() {

	resp, err := http.Get("model.bst")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error by reading file:\n%v\n", err)
	}
	model, err := leaves.XGEnsembleFromReader(bufio.NewReader(bytes.NewReader(b)), true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("XGB Model Loaded:\n%v\n", model)
	runner.model = model
}

func (runner *XGBRunner) readXData() {
	resp, err := http.Get("x_test_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error by reading file:\n%v\n", err)
	}

	runner.dense, err = mat.DenseMatFromCsv(bufio.NewReader(bytes.NewReader(b)), 0, false, ",", 0.0)
	if err != nil {
		fmt.Printf("Error by reading file:\n%v\n", err)
	}

	fmt.Printf("Dense data X loaded:  rows :\n%d\n", runner.dense.Rows)
	fmt.Print("start computing predictions...\n")
	nThreads := 1

	var nRows int

	nRows = runner.dense.Rows

	predictions := make([]float64, nRows)
	start := time.Now()
	runner.model.PredictDense(runner.dense.Values, runner.dense.Rows, runner.dense.Cols, predictions, 0, nThreads)
	elapsed := time.Since(start)
	fmt.Printf("Finished computing predictions :\n%d\n", runner.dense.Rows)
	fmt.Printf(" computing predictions took %s ", elapsed)
	fmt.Printf("for %d churn\n", nRows)
	var totalOut, total int
	totalOut = 0
	total = 0
	var retentionRate float64
	retentionRate = 0

	for index, value := range predictions {
		proba := math.Round(value*100*100) / 100
		if value < 0.5 {
			totalOut++
		}
		total++
		retentionRate = float64(totalOut) / float64(total) * 100.0
		fmt.Printf("Prediction : %f  %%  chances  customers #%d will churn before end of fiscal Year.\n", proba, index)

	}
	fmt.Printf("\n *************************************\nYour retention rate :  : %f %%\n", retentionRate)

}
