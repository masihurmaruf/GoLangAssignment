package main

import "fmt"

func main() {
	fmt.Println(sum())
}

func sum() int {
	sum := 0
	for i := 3; i <= 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}
