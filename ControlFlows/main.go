package main

import "fmt"

func main() {

	//if else
	age := 17

	if age >= 18 {
		fmt.Println("você pode dirigir")
	} else {
		fmt.Println("você não pode dirigir")
	}

	score := 79

	if score >= 90 {
		fmt.Println("Grade:A")
	} else if score >= 80 {
		if score >= 85 {
			fmt.Println("Grade:B+")
		} else {
			fmt.Println("Grade:B")
		}
	} else {
		fmt.Println("Grade:C")
	}

	//loops
	//for loop
	//nested loop control
	//loop control breaks

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Println(i, "*", j, "=", i*j)
		}
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			goto end
		}
	}
end:
	fmt.Println("loop acabou")

	i := 0
start:
	if i < 10 {
		fmt.Println(i)
		i++
		goto start
	}

	//loop infinito
	for {
		fmt.Println("infinite loop")
	}

loop:
	fmt.Println("loop infinito diferente")
	goto loop
}
