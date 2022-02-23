package mapuse

import (
	"testing"
)

func TestStringKey(t *testing.T) {
	aMap := make(map[string]int)
	keyA1 := "a"
	keyA2 := "a"
	t.Logf("&keyA1: %p, &keyA2: %p", &keyA1, &keyA2)
	aMap[keyA1] = 1
	t.Logf("%v: %v", keyA2, aMap[keyA2])
}

func TestStringPtrKey(t *testing.T) {
	aMap := make(map[*string]int)
	keyA1 := "a"
	keyA2 := "a"
	t.Logf("&keyA1: %p, &keyA2: %p", &keyA1, &keyA2)
	keyA1Ptr := &keyA1
	aMap[keyA1Ptr] = 1
	t.Logf("&keyA2[=%v]: %v", keyA2, aMap[&keyA2])
	keyA3 := &keyA1
	t.Logf("keyA3[=%v]: %v", *keyA3, aMap[keyA3])
	// Error: Cannot use '*keyA3' (type string) as the type *string
	//t.Logf("*keyA3[=%v]: %v", *keyA3, aMap[*keyA3])
	// Error: Cannot use '&keyA3' (type **string) as the type *string
	//t.Logf("&keyA3[=%v]: %v", *keyA3, aMap[&keyA3])
}

type BookTitle struct {
	title string
}

type BookTitlePtr struct {
	titlePtr *string
}

func TestBookTitleKey(t *testing.T) {
	aMap := make(map[BookTitle]string)
	title := BookTitle{
		title: "Go",
	}
	aMap[title] = "content"
	title2 := BookTitle{
		title: "Go",
	}
	t.Logf("&title: %p, &title2: %p", &title, &title2)
	t.Logf("title2: %v\n", aMap[title2])
}

func TestBookTitlePtrKey(t *testing.T) {
	aMap := make(map[BookTitlePtr]string)
	titleStr := "Go"
	titlePtr1 := &titleStr
	titlePtr2 := &titleStr
	title := BookTitlePtr{
		titlePtr: titlePtr1,
	}
	aMap[title] = "content"
	title2 := BookTitlePtr{
		titlePtr: titlePtr2,
	}
	t.Logf("&title: %p, &title2: %p", &title, &title2)
	// title.titlePtr 和 title2.titlePtr 这两个指针变量的值相同，所以以 title2 为键，也能取到值
	t.Logf("title2: %v\n", aMap[title2])
}
