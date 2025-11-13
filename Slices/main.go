package main

import "fmt"

func main() {

	numbers := make([]int, 5)

	fmt.Println(numbers)

	//subslicing

	numbers2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	subslice1 := numbers2[1:4]
	subslice2 := numbers2[:3]
	subslice3 := numbers2[4:]

	fmt.Println(subslice1)
	fmt.Println(subslice2)
	fmt.Println(subslice3)

	//append() copy()

	numbers3 := []int{}
	numbers3 = append(numbers3, 0, 1, 2, 3, 4)
	numbers4 := make([]int, len(numbers3), cap(numbers3)*2)

	copy(numbers4, numbers3)

	fmt.Println(numbers3)
	fmt.Println(numbers4)

	//nil slice
	var num []int //nil slice
	fmt.Println(num)

	if num == nil {
		fmt.Println("slice é nil")
	}
}
