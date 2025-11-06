package main

import "fmt"

func main() {
	//declarar constantes string e numeric
	const NAME string = "Maria Eduarda"
	const PRICE int = 100
	fmt.Println("O nome é:", NAME)
	fmt.Println("O preço é:", PRICE)

	//declarar integers literals
	//decimal, octal, hexadecimal
	const (
		DECIMAL     = 255  //decimais sem prefixo
		OCTAL       = 0377 //octal com o 0
		HEXADECIMAL = 0xff //hexadecimal com o 0x
	)

	fmt.Println("Decimal:", DECIMAL, "Octal:", OCTAL, "Hexadecimal:", HEXADECIMAL)

	//floating-point literals
	//um floatin-point pode ter uma parte integer, um ponto decimal, uma parte fracionaria e uma parte exponencial
	const PI float64 = 3.141
	fmt.Println("O pi é:", PI)
	const AVOGRADO = 6.022e23
	fmt.Println("O avogrado é:", AVOGRADO)

	//escape sequences em string literals
	const GREETING = "Hello, Earth\n"
	const QUOTE = "\"Go é simples\" - Eu que disse"
	fmt.Println(GREETING)
	fmt.Println(QUOTE)

	const BELL = "Bell is \a"
	fmt.Println(BELL)

	const LB = "I\nAM\nBATMAN\n!"
	fmt.Println(LB)

	//Multi-line e string literals concatenadas
	const MULTILINE = "Isso é um comentário " + "que é muito legal demais"

	const CONCATENATED = "string " + "Concatenada"
	fmt.Println(MULTILINE)
	fmt.Println(CONCATENATED)

	//Boolean constants
	const ACTIVE = true
	const READY = false
	fmt.Println("ACTIVE:", ACTIVE, "READY:", READY)

	//constantes para calculos
	const LENGTH = 50
	const WIDTH = 5
	const AREA = LENGTH * WIDTH
	fmt.Println("AREA:", AREA)
}
