package main

import "fmt"

func main() {
	type Book struct {
		Pages int
	}
	var book = Book{} // book is addressable
	p := &book.Pages
	*p = 123
	fmt.Println(book) // {123}

// The followings slice composite literals
// are equivalent to each other.
arr := [...]bool{3: false, 1: true, true}

// var zeroMap map[string]int
var zeroStruct Book

fmt.Println(arr)
fmt.Println(Book{})
fmt.Println(zeroStruct)

head1 := []byte{'P', 'N', 'G', ' '}

var heads = []*[]byte{
	{'P', 'N', 'G', ' '},
	&head1,
	{'J', 'P', 'E', 'G'},
}

idx := 1

fmt.Println(heads)
fmt.Printf("%p",&head1)
fmt.Printf("%p",heads[idx])
fmt.Println(heads[idx] == &head1)

	// The following two lines fail to compile, for
	// Book{} is unaddressable, so is Book{}.Pages.
}