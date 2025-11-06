package main

import "fmt"

func main() {

	//declarar variaveis
	var mango string = "isso é uma manga gigante"
	var weight int = 54
	var height int = 23

	fmt.Println(mango)
	fmt.Println(weight)
	fmt.Println(height)

	//forma encurtada
	age := 32
	city := "garibaldi"
	fmt.Println(age)
	fmt.Println(city)

	//multiplas declarações
	var apples, oranges int = 23, 78
	fmt.Println("Eu tenho", apples, "maças e", oranges, "laranjas")

	var fruits = apples + oranges
	fmt.Println("minahs frutas totais:", fruits)

	var windows, mac, linux string = "Windows é normal\n", "Mac é meh\n", "Linux é top\n"

	print(windows, mac, linux)

	//declarações estaticas
	var x float64 = 20.5
	fmt.Println(x)
	fmt.Printf("x é do tipo: %T\n", x)

	//declarações dinamicas
	y := 89
	fmt.Println(y)
	fmt.Printf("y é do tipo: %T\n", y)

	//declarações mixed
	var a, b, c = 758.52, 8, "show"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a é do tipo: %T\n", a)
	fmt.Printf("b é do tipo: %T\n", b)
	fmt.Printf("c é do tipo: %T\n", c)
}
