package main

import "fmt"

func capital(s string) {
	fmt.Printf("%s, é a capital.\n", s)
}

func country(s string) {
	fmt.Printf("%s, é o país.\n", s)
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	} else if n2 > n1 {
		return n2
	} else {
		return 0
	}
}

func swap(x, y string)(string, string){
	return y, x
}

func increment(number int){
	number++
	fmt.Println("Dentro do increment:", number)
}

func modify(slice []int){
	slice[0] = 999
	fmt.Println("Dentro do modify:", slice)
}

func main() {
	fmt.Println("Funções em GO")

	country("Brasil")
	capital("Brasília")
	a, b := 100, 200
	result := max(a, b)
	fmt.Printf("O valor máximo é: %d\n", result)
	firstName, lastName := swap("Luís Felipe", "Casagrande")
	fmt.Printf("Nome invertido: %s %s\n", firstName, lastName)

	x:=10
	increment(x)
	fmt.Println("Depois do increment:", x)

	mySlice := []int{1, 2, 3}
	modify(mySlice)
	fmt.Println("Depois do modify:", mySlice)
}

// func nomeDaFuncao(parametros) tipoDoRetorno {
// 	//corpo da função
// }
