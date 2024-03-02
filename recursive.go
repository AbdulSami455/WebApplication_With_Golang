package main

/*
import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/wcharczuk/go-chart/v2" // Update the import to use v2
)

func recursiveFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return recursiveFibonacci(n-1) + recursiveFibonacci(n-2)
}

func iterativeFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func measureTime(f func(int) int, n int) (int, time.Duration) {
	startTime := time.Now()
	result := f(n)
	elapsedTime := time.Since(startTime)
	return result, elapsedTime
}

func generateCSV(filename string, algorithm func(int) int, maxN, numTrials int) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Input Size", "Time (seconds)"})

	for n := 1; n <= maxN; n++ {
		totalTime := time.Duration(0)
		for i := 0; i < numTrials; i++ {
			_, elapsedTime := measureTime(algorithm, n)
			totalTime += elapsedTime
		}

		averageTime := totalTime.Seconds() / float64(numTrials)
		writer.Write([]string{fmt.Sprintf("%d", n), fmt.Sprintf("%f", averageTime)})
	}
}

func plotLineChart(csvFilename, title, xLabel, yLabel string) {
	file, err := os.Open(csvFilename)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	var xValues []float64
	var yValues []float64

	for _, record := range records[1:] {
		xValueStr := record[0]
		yValueStr := record[1]

		xValue, err := strconv.ParseFloat(xValueStr, 64)
		if err != nil {
			fmt.Println("Error converting X value to float:", err)
			return
		}

		yValue, err := strconv.ParseFloat(yValueStr, 64)
		if err != nil {
			fmt.Println("Error converting Y value to float:", err)
			return
		}

		xValues = append(xValues, xValue)
		yValues = append(yValues, yValue)
	}

	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: xValues,
				YValues: yValues,
			},
		},
	}

	yAxis := chart.YAxis{ValueFormatter: chart.FloatValueFormatter}

	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
		chart.Title{Title: title},
		chart.XAxis{ValueFormatter: chart.FloatValueFormatter},
		yAxis,
	}

	err = graph.Render(chart.PNG, os.Stdout)
	if err != nil {
		fmt.Println("Error rendering chart:", err)
	}
}
func main() {
	maxN := 250
	numTrials := 3

	recursiveCSVFilename := "recursive_times.csv"
	iterativeCSVFilename := "iterative_times.csv"

	generateCSV(recursiveCSVFilename, recursiveFibonacci, maxN, numTrials)
	generateCSV(iterativeCSVFilename, iterativeFibonacci, maxN, numTrials)

	plotLineChart(recursiveCSVFilename, "Recursive Fibonacci Time Analysis", "Input Size", "Time (seconds)")
	plotLineChart(iterativeCSVFilename, "Iterative Fibonacci Time Analysis", "Input Size", "Time (seconds)")
}
*/
