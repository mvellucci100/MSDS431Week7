package main

import (
	"fmt"
	"github.com/montanaflynn/stats"
	"log"
)

// Function to perform linear regression and handle errors
func performLinearRegression(dataset []stats.Coordinate, datasetNumber int) (*stats.Regression, error) {
	r, err := stats.LinearRegression(dataset)
	if err != nil {
		// Returning a structured error with more detailed information
		return nil, fmt.Errorf("error performing linear regression on Dataset %d: %v", datasetNumber, err)
	}
	return r, nil
}

func main() {
	// Define datasets
	datasets := [][]stats.Coordinate{
		{
			{10, 8.04}, {8, 6.95}, {13, 7.58}, {9, 8.81}, {11, 8.33},
			{14, 9.96}, {6, 7.24}, {4, 4.26}, {12, 10.84}, {7, 4.82}, {5, 5.68},
		},
		{
			{10, 9.14}, {8, 8.14}, {13, 8.74}, {9, 8.77}, {11, 9.26},
			{14, 8.10}, {6, 6.13}, {4, 3.10}, {12, 9.13}, {7, 7.26}, {5, 4.74},
		},
		{
			{10, 7.46}, {8, 6.77}, {13, 12.74}, {9, 7.11}, {11, 7.81},
			{14, 8.84}, {6, 6.08}, {4, 5.39}, {12, 8.15}, {7, 6.42}, {5, 5.73},
		},
		{
			{8, 6.58}, {8, 5.76}, {8, 7.71}, {8, 8.84}, {8, 8.47},
			{8, 7.04}, {8, 5.25}, {19, 12.50}, {8, 5.56}, {8, 7.91}, {8, 6.89},
		},
	}

	// Loop through datasets and call the function for each
	for i, dataset := range datasets {
		r, err := performLinearRegression(dataset, i+1)
		if err != nil {
			// Log the error with dataset number for better context
			log.Println(err)
			continue // Continue with the next dataset even if there's an error
		}

		// Output the result of the regression
		fmt.Printf("Linear Regression on Dataset %d: Intercept: %v, Slope: %v\n", i+1, r.Intercept, r.Slope)
	}
}
