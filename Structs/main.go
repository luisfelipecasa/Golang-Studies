package main

import "fmt"

func main() {
	type Book struct {
		title   string
		author  string
		subject string
		bookId  int
		read    bool
	}

	book1 := Book{
		title:   "Livro top",
		author:  "Luisinho",
		subject: "ação",
		bookId:  1,
		read:    true,
	}

	fmt.Println(book1.title)
	fmt.Println(book1.author)
	fmt.Println(book1.subject)
	fmt.Println(book1.bookId)
	fmt.Println(book1.read)
}
