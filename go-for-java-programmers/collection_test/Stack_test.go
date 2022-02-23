package collection_test

import (
	"fmt"
	"go-for-java-programmers/collection"
	"testing"
)

func TestStack(t *testing.T) {
	var s collection.Stack
	s.Push("World!")
	s.Push("Hello ")
	for s.Size() > 0 {
		fmt.Print(s.Pop())
	}
	fmt.Println()
}
