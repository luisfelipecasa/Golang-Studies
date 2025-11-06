package main

import "fmt"

func main() {
	var a float64 = 10
	var b float64 = 20
	c := 11
	d := 3

	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)

	//modulos
	fmt.Println("c % d =", c%d)

	a++
	fmt.Println("a++ =", a)
	a--
	fmt.Println("a-- =", a)

	//relational operators
	fmt.Println("a == b:", a == b)
	fmt.Println("a != b:", a != b)
	fmt.Println("a > b:", a > b)
	fmt.Println("a < b:", a < b)
	fmt.Println("a >= b:", a >= b)
	fmt.Println("a <= b:", a <= b)

	//logic operators
	x := true
	y := false

	//and
	fmt.Println("x && y =", x && y)
	//or
	fmt.Println("x || y =", x || y)
	//not
	fmt.Println("!x =", !x)
	fmt.Println("!y =", !y)

	//assignment operators
	a += b
	fmt.Println("a += b:", a)
	a -= b
	fmt.Println("a -= b:", a)

	//bitwise operators
	//funcionam bit por bit
	z := 30
	
	//and
	fmt.Println("z & 2 =", z&2)
	//or
	fmt.Println("z | 2 =", z|2)
	//xor
	fmt.Println("z ^ 2 =", z^2)
	//not
	fmt.Println("^z =", ^z)
	//left shift
	fmt.Println("z << 1 =", z<<1)
	//right shift
	fmt.Println("z >> 1 =", z>>1)

	//Miscellaneous operators
	e := 10

	//address-of
	ptr := &e
	fmt.Println(ptr)

	//pointer reference operator
	fmt.Println(*ptr)
	
}
