package main
import (
	"fmt"
	"strings"
)

func main(){

	text := "isso ai"
 
	//valor hexadecimal da string
	fmt.Printf("HEX bytes: ")
	for i := 0; i < len(text); i++ {
		fmt.Printf("%x", text[i])
	}

	fmt.Printf("o tamanho da string é : ")
	fmt.Println(len(text))

	fruits := []string{"apples", "oranges", "and bananas"}
	fmt.Println(strings.Join(fruits, " "))
}
