package main

import (
	"fmt"
	"strconv"
)

func main(){
	var total int = 20
	var items int = 7
	var average float64

	average = float64(total) / float64(items)
	fmt.Println(average)

	//string para int
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)
}