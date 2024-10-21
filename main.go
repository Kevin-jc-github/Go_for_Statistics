package main

import (
	"fmt"
	"time" // Importing the time package for measuring execution time

	"github.com/montanaflynn/stats"
)

// Anscombe's Quartet dataset, containing four sets of x and y values
// Each dataset demonstrates different statistical characteristics despite having the same linear regression values
var anscombeQuartet = map[string][][]float64{
	"x": {
		{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}, // x1 values
		{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}, // x2 values
		{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}, // x3 values
		{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},     // x4 values (different pattern from the other datasets)
	},
	"y": {
		{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}, // y1 values
		{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74},    // y2 values
		{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73}, // y3 values
		{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89},  // y4 values (different pattern from the other datasets)
	},
}

// calculateInterceptAndSlope calculates the linear regression for a given set of x and y values
// The function returns the slope and intercept of the regression line
func calculateInterceptAndSlope(x, y []float64) (float64, float64) {
	// Calculate the mean of the x and y values
	meanX, _ := stats.Mean(x)
	meanY, _ := stats.Mean(y)
	var numerator, denominator float64

	// Calculate the numerator and denominator for the slope formula
	for i := 0; i < len(x); i++ {
		numerator += (x[i] - meanX) * (y[i] - meanY)   // Numerator: sum of (x[i] - meanX) * (y[i] - meanY)
		denominator += (x[i] - meanX) * (x[i] - meanX) // Denominator: sum of (x[i] - meanX)^2
	}

	// Calculate the slope (m) and intercept (b) using the least-squares regression formula
	slope := numerator / denominator
	intercept := meanY - slope*meanX

	// Return the calculated slope and intercept
	return intercept, slope
}

func main() {
	// Run regression tests for each dataset in the Anscombe's Quartet
	fmt.Println("Linear regression results for each dataset:")

	// Loop through each dataset (x and y pairs)
	for i := 0; i < len(anscombeQuartet["x"]); i++ {
		// Get the x and y values for the current dataset
		x := anscombeQuartet["x"][i]
		y := anscombeQuartet["y"][i]

		// Start time measurement before calling the linear regression function
		start := time.Now()

		// Calculate the slope and intercept for the current dataset
		intercept, slope := calculateInterceptAndSlope(x, y)

		// Measure the duration it took to perform the regression calculation
		duration := time.Since(start)

		// Print out the calculated slope, intercept, and the time taken for the calculation
		fmt.Printf("Dataset %d: Slope = %.5f, Intercept = %.5f, Calculation Time = %s\n", i+1, slope, intercept, duration)
	}
}
