package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
)

func main() {
	// Example dataset of x and y coordinates
	data := []stats.Coordinate{
		{10, 8.04},
		{8, 6.95},
		{13, 7.58},
		{9, 8.81},
		{11, 8.33},
		{14, 9.96},
		{6, 7.24},
		{4, 4.26},
		{12, 10.84},
		{7, 4.82},
		{5, 5.68},
	}

	// Perform linear regression on the dataset
	r, err := stats.LinearRegression(data)
	if err != nil {
		fmt.Println("Error performing linear regression:", err)
		return
	}

	// Output the result of the regression
	fmt.Printf("Linear Regression Result: Intercept: %v, Slope: %v\n", r.Intercept, r.Slope)
}
