package main
import "fmt"

//variaveis globais
var g int
var x int = 25

func main(){

	var a, b int

	//variaveis locais
	a = 10
	b = 20
	g = a + b
	x := 100

	//a varivael local cobre a variavel global
	fmt.Println(a, b, g, x)

}