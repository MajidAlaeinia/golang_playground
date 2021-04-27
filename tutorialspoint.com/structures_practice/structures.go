package main

import "fmt"

type Book struct {
	title string
	author string
	subject string
	bookId int
}

func main() {
	var book1 Book
	var book2 Book

	setBook1Attributes(&book1)
	setBook2Attributes(book2)

	fmt.Printf( "Book 1 title: %s\n", book1.title)
	fmt.Printf( "Book 1 author: %s\n", book1.author)
	fmt.Printf( "Book 1 subject: %s\n", book1.subject)
	fmt.Printf( "Book 1 book_id: %d\n\n", book1.bookId)

	fmt.Printf( "Book 2 title: %s\n", book2.title)
	fmt.Printf( "Book 2 author: %s\n", book2.author)
	fmt.Printf( "Book 2 subject: %s\n", book2.subject)
	fmt.Printf( "Book 2 book_id: %d\n", book2.bookId)
}

func setBook1Attributes(book1 *Book) {
	book1.title = "Go Programming"
	book1.author = "Mahesh Kumar"
	book1.subject = "Go Programming Tutorial"
	book1.bookId = 6495407
}

func setBook2Attributes(book2 Book) {
	book2.title = "Telecom Billing"
	book2.author = "Zara Ali"
	book2.subject = "Telecom Billing Tutorial"
	book2.bookId = 6495700
}
