package main

import (
	"errors"
	"fmt"
	"math"
)

type error interface {
	Error() string
}

func Sqrt(value float64)(float64, error){
	if value < 0 {
		return 0, errors.New("número negativo para raiz quadrada")
	}

	return math.Sqrt(value), nil
}

func main() {

	result, err := Sqrt(9)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
