package slice

import (
	"fmt"
	"testing"
)

type book struct {
	title string
}

func revise(books []book) []book {
	fmt.Printf("[revise func] addr of books: %p, books as reference: %p\n", &books, books)
	books[0].title = "Go programming"
	return books
}

func TestMakeSlice(t *testing.T) {
	// make(sliceType, len):
	// 底层数组大小为 len。slice 容量和长度都为 len，start 为 0
	slice1 := make([]int, 2)
	fmt.Printf("%v\n", slice1)
	// make(sliceType, len, cap):
	// 底层数组大小为 cap。slice 容量为 cap，长度为 len，start 为 0
	slice2 := make([]int, 2, 2)
	fmt.Printf("%v\n", slice2)
	slice3 := make([]int, 0, 2)
	fmt.Printf("%v\n", slice3)
}

func TestPassSlice(t *testing.T) {
	books := []book{
		book{},
	}
	fmt.Printf("[before revise] addr of books: %p, books as reference: %p\n", &books, books)
	fmt.Printf("[before revise] addr of books[0]: %p, books[0]: %v\n", &(books[0]), books[0])
	returnedBooks := revise(books)
	fmt.Printf("[after revise] addr of returnedBooks: %p, retunedBooks as reference: %p\n", &returnedBooks, returnedBooks)
	fmt.Printf("[after revise] addr of books[0]: %p, books[0]: %v\n", &(books[0]), books[0])
}