package list

import (
	"fmt"
	"testing"
)

func TestList_HeadAdd(t *testing.T) {
	l := New[int]()
	l.HeadAdd(1, 2, 3, 4)
	fmt.Println(l.List())
}

func TestList_Append(t *testing.T) {
	l := New[int]()
	l.Append(1, 2, 3, 4)
	fmt.Println(l.List())
}

func TestList_ForEach(t *testing.T) {
	l := New[string]()
	l.Append("1", "2", "3", "4")
	l.ForEach(func(v string) {
		fmt.Println(v)
	})
}

func TestList_Len(t *testing.T) {
	l := New[int]()
	l.Append(1, 2, 3, 4)
	fmt.Println(l.Len())
}

func Test_DeleteIndex(t *testing.T) {
	l := New[int]()
	l.Append(1, 2, 3, 4)

	// l.DeleteIndex(2)
	// fmt.Println(l.List())

	// l.DeleteIndex(1)
	// fmt.Println(l.List())

	l.DeleteIndex(10)
	fmt.Println(l.List())
}
