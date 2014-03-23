package main

import (
	"fmt"
	"math"
)

func isPrime(x int64) bool {
	if x <= 1 {
		return false
	}
	
	for i := squareRoot(x); i >= 1; i-- {
		if i == 1 {
			return true
		}
		if x % i == 0 {
			return false
		}
	}
	return true
}

func squareRoot(x int64) int64 {
	return int64(math.Sqrt(float64(x)))
}

func main() {
	var x int64 = 600851475143

	for i := squareRoot(x); i >= 1; i-- {
		if x % i == 0 && isPrime(i) {
			fmt.Println(i)
			break
		}
	}
}
