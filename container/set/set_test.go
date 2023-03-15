package set

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	fmt.Println(s.Items())
}

func TestNewWith(t *testing.T) {
	s := NewWith[int](1, 2, 3)
	fmt.Println(s.Items())
}

func TestSet_Add(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	fmt.Println(s.Items())
}

func TestSet_Clear(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	fmt.Println(s.Items())
	s.Clear()
	fmt.Println(s.Items())
}

func TestSet_Contains(t *testing.T) {
	s := NewWith[int](1, 2, 3)
	if !s.Contains(1) || !s.Contains(3) {
		t.Errorf("contain but not")
	}

	if s.Contains(10) {
		t.Errorf("not contain but yes")
	}
}

func TestSet_Equals(t *testing.T) {
	s1 := NewWith[int](1, 2, 3)
	s2 := NewWith[int](1, 2, 3)

	if !s1.Equals(s2) {
		t.Errorf("equal but not")
	}
}

func TestSet_Items(t *testing.T) {
	s := NewWith[int](1, 2, 3)
	fmt.Println(s.Items())
}

func TestSet_Len(t *testing.T) {
	s := NewWith[int](1, 2, 3)
	fmt.Println(s.Len())
}

func TestSet_Remove(t *testing.T) {
	s := NewWith[int](1, 2, 3)
	fmt.Println(s.Items())
	s.Remove(1)
	fmt.Println(s.Items())
	s.Remove(10)
	fmt.Println(s.Items())
}
