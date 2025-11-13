package main
import "fmt"

func main(){
	//declaring pointers
	var ip *int
	// var fp *float32

	//operações basicas com pointers
	var a int = 20
	
	ip = &a
	fmt.Println(&a)
	fmt.Println(ip)
	fmt.Println(*ip)

	//nil value
	var ptr *int
	fmt.Println(ptr)
}