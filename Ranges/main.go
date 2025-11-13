package main
import "fmt"

func main(){
	//range com array e slices
	numbers := []int{1,2,3,4,5,6,7,8}

	for i := range numbers {
		fmt.Println(i, numbers[i])
	}
}