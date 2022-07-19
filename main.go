package main

import (
	"constraints"
	"fmt"
	. "generics/book"
	. "generics/number"
)

type Comparable[T any] interface {
	Less(c T) bool
}

type SliceOfT[T any] []T

func Sort[Elem Comparable[Elem]](list []Elem) {
	for i := 0; i < len(list); i++ {
		lesser := i
		for j := i + 1; j < len(list); j++ {
			if list[j].Less(list[lesser]) {
				lesser = j
			}
		}
		if lesser != i {
			Swap(list, i, lesser)
		}
	}
}

func Swap[Elem Comparable[Elem]](list []Elem, i, j int) {
	list[i], list[j] = list[j], list[i]
}

func main() {
	a := add(3, 4)
	fmt.Println(a)
	var bookshelf = []Book{{300}, {100}, {200}}
	var anotherBookshelf = []Book{{700}, {800}, {60}, {10}, {5}, {1000}}
	var intshelf = []MyInt{500, 60, 70}

	/*
		Step #1: Instantiation:
			Sort[book]
				   ↓									-- pass type argument
			Sort[Elem interface{ Less(y Elem) bool }]
				   ↓									-- substitute book for Elem
			Sort[book interface{ Less(y book) bool }]
				   ↓									-- verify that book satisfies type parameter constraint
			#Sort[book] -- Instantiated function (normal function)

		Step #2: Invocation (as usual):
			#Sort[book](bookshelf)

	*/

	// First way: steps 1 and 2 in one line
	Sort[Book](bookshelf) // Step 1 + Step 2 in one line

	// Second way: step 1 then step 2
	booksort := Sort[Book]
	booksort(anotherBookshelf)

	// Third way: omitting instantiation (type inference Elem == MyInt)
	Sort(intshelf)

	fmt.Println("Generics:")
	fmt.Printf("bookshelf: %v\n", bookshelf)
	fmt.Printf("anotherBookshelf: %v\n", anotherBookshelf)
	fmt.Printf("intshelf: %v\n", anotherBookshelf)

	fmt.Println("\n\nWithout Generics:")
	var bookshelfWithSortableBooks = []SortableBook{{300}, {100}, {200}}
	SortWithoutGeneric(SortableBooks(bookshelfWithSortableBooks))
	fmt.Printf("bookshelf: %v\n", bookshelfWithSortableBooks)

	fmt.Println("a greater than sb?", greaterThan("a", "sb"))
	fmt.Println("sb greater than a?", greaterThan("sb", "a"))
	fmt.Println("1 greater than 2?", greaterThan(1, 2))
	fmt.Println("1.2 greater than 1.22?", greaterThan(1.2, 1.22))
}

type Sortable interface {
	Less(i, j int) bool
	Swap(i, j int)
	Len() int
}

func SortWithoutGeneric(list Sortable) {

	for i := 0; i < list.Len(); i++ {
		lesser := i
		for j := i + 1; j < list.Len(); j++ {
			if list.Less(j, lesser) {
				lesser = j
			}
		}
		if lesser != i {
			list.Swap(i, lesser)
		}
	}
}

func greaterThan[C constraints.Ordered](a, b C) bool {
	return a > b
}

func add[T constraints.Integer](x, y T) T {
	return x + y
}
