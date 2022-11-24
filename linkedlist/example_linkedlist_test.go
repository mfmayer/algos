package linkedlist_test

import (
	"fmt"

	"github.com/mfmayer/algos/linkedlist"
)

func Example() {
	// Create linked list and return head element
	l1 := linkedlist.NewElement(1, 2, 3, 4)
	l5 := linkedlist.NewElement(5)
	l6 := linkedlist.NewElement(6, 7, 8, 9)
	l5.InsertBefore(l1)
	l5.InsertAfter(l6)

	next := l1
	for i := 1; next != nil; i++ {
		fmt.Printf("%v,", next.Data)
		next = next.Next()
	}
	// Output: 1,2,3,4,5,6,7,8,9,
}
