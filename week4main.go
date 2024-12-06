package main
import (
    "fmt"
    "github.com/montanaflynn/stats"
)

func main() {
    // Dataset 1
    c1 := []stats.Coordinate{
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

    // Dataset 2
    c2 := []stats.Coordinate{
        {10, 9.14},
        {8, 8.14},
        {13, 8.74},
        {9, 8.77},
        {11, 9.26},
        {14, 8.10},
        {6, 6.13}, 
        {4, 3.10},
        {12, 9.13},
        {7, 7.26},
        {5, 4.74},
    }

    // Dataset 3
    c3 := []stats.Coordinate{
        {10, 7.46},
        {8, 6.77},
        {13, 12.74},
        {9, 7.11},
        {11, 7.81},
        {14, 8.84},
        {6, 6.08}, 
        {4, 5.39},
        {12, 8.15},
        {7, 6.42},
        {5, 5.73},
    }

    // Dataset 4
    c4 := []stats.Coordinate{
        {8, 6.58},
        {8, 5.76},
        {8, 7.71},
        {8, 8.84},
        {8, 8.47},
        {8, 7.04},
        {8, 5.25}, 
        {19, 12.50},
        {8, 5.56},
        {8, 7.91},
        {8, 6.89},
    }

    // Perform linear regression
    r1, err1 := stats.LinearRegression(c1)
    if err1 != nil {
        fmt.Println("Error in Dataset 1:", err1)
    } else {
        fmt.Printf("Linear Regression on Dataset 1: %v\n", r1)
    }

    r2, err2 := stats.LinearRegression(c2)
    if err2 != nil {
        fmt.Println("Error in Dataset 2:", err2)
    } else {
        fmt.Printf("Linear Regression on Dataset 2: %v\n", r2)
    }

    r3, err3 := stats.LinearRegression(c3)
    if err3 != nil {
        fmt.Println("Error in Dataset 3:", err3)
    } else {
        fmt.Printf("Linear Regression on Dataset 3: %v\n", r3)
    }

    r4, err4 := stats.LinearRegression(c4)
    if err4 != nil {
        fmt.Println("Error in Dataset 4:", err4)
    } else {
        fmt.Printf("Linear Regression on Dataset 4: %v\n", r4)
    }
}
