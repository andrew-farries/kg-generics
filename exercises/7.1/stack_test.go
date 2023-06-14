package stack_test

import (
	"fmt"
	"stack"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPushOneValueToEmptyStackLeavesTheStackWithLength1(t *testing.T) {
	t.Parallel()
	s := stack.Stack[int]{}
	s.Push(0)
	if s.Len() != 1 {
		t.Fatal("Push didn't add value to stack")
	}
}

func TestPushTwiceThenPopTwiceOnEmptyStackLeavesTheStackEmpty(t *testing.T) {
	t.Parallel()
	s := stack.Stack[string]{}
	s.Push("a", "b")
	if s.Len() != 2 {
		t.Fatal("Push didn't add all values to stack")
	}
	s.Pop()
	want := "a"
	got, ok := s.Pop()
	if !ok {
		t.Fatal("Pop returned not ok on non-empty stack")
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestPopReturnsNotOKOnEmptyStack(t *testing.T) {
	t.Parallel()
	s := stack.Stack[int]{}
	_, ok := s.Pop()
	if ok {
		t.Fatal("Pop returned ok on empty stack")
	}
}

func TestSet(t *testing.T) {
	s := stack.NewSet[int](3)

	s.Add(1)
	s.Add(2)

	fmt.Println(s.Contains(1))
	fmt.Println(s.Contains(3))
	fmt.Println(s.Members())
	fmt.Println(s)
	fmt.Println(s.Union(stack.NewSet([]int{4, 5, 6}...)))
	fmt.Println(s.Intersection(stack.NewSet([]int{1, 5, 6}...)))
}
