package _struct

import (
	"fmt"
	"testing"
)

type book struct {
	title   string
	price   int
	authors [1]author
}

type author struct {
	name string
}

func TestAssignValObjToVar(t *testing.T) {
	a := author{
		name: "x-man",
	}
	b := a
	fmt.Printf("a: %v, b: %v\n", a, b)
	fmt.Printf("&a: %p, &b: %p\n", &a, &b)
	b.name = "batman"
	fmt.Printf("a: %v, b: %v\n", a, b)
}

func TestNewBook(t *testing.T) {
	theAuthor := author{
		name: "peter",
	}
	authors := [1]author{0: theAuthor}
	var price = 10
	var title = "go programming"
	fmt.Printf("&theAuthor: %p, &authors: %p, &(authors[0]): %p, &price: %p, &title: %p\n",
		&theAuthor, &authors, &(authors[0]), &price, &title)
	book := book{}
	book.title = title
	book.authors = authors
	book.price = price
	fmt.Printf("&book.authors: %p, &book.authors[0]: %p, &book.price: %p, &book.title: %p\n",
		&(book.authors), &(book.authors[0]), &(book.price), &(book.title))
}

func TestNewBook2(t *testing.T) {
	theAuthor := author{
		name: "peter",
	}
	authors := [1]author{0: theAuthor}
	var price = 10
	var title = "go programming"
	fmt.Printf("&theAuthor: %p, &authors: %p, &(authors[0]): %p, &price: %p, &title: %p\n",
		&theAuthor, &authors, &(authors[0]), &price, &title)
	book := book{
		title:   title,
		authors: authors,
		price:   price,
	}
	fmt.Printf("&book.authors: %p, &book.authors[0]: %p, &book.price: %p, &book.title: %p\n",
		&(book.authors), &(book.authors[0]), &(book.price), &(book.title))
	// Deep copy, not shallow copy
	fmt.Printf("&theAuthor.name: %p, &(book.authors[0].name): %p\n", &theAuthor.name, &(book.authors[0].name))
}
