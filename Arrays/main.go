package main

import "fmt"

//declaração
// var balance = []float32{12.2, 23.34, 45.54, 67.66, 99.55}

func main() {
	var n [10]int

	for i := 0; i < 10; i++ {
		n[i] = i + 100
	}

	for i := 0; i < 10; i++ {
		fmt.Println(n[i])
	}
}
