package main

import "fmt"

func main() {
	sum := 0
    for i, j := 0, 1; j < 4000000; i, j = i+j,i {
		if i % 2 == 0 {
			sum += i
		}
    }
	fmt.Println(sum)
}
