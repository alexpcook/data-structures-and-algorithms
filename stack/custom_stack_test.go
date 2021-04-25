package stack

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestStackIsEmptyError(t *testing.T) {
	err := StackIsEmptyError{}
	val := fmt.Sprint(err)
	if val != "stack is empty" {
		t.Fatalf("want %q, got %q", "stack is empty", val)
	}
}

func TestStack1Peek(t *testing.T) {
	s := new(Stack1)
	val, err := s.Peek()
	if !errors.Is(err, StackIsEmptyError{}) {
		t.Fatalf("want StackIsEmptyError, got %v", err)
	}
	if val != "" {
		t.Fatalf("want %q, got %q", "", val)
	}

	s.Push("call1")
	s.Push("call2")

	val, err = s.Peek()
	if err != nil {
		t.Fatalf("want nil error, got %v", err)
	}
	if val != "call2" {
		t.Fatalf("want %q, got %q", "call2", val)
	}
	if s.top.value != "call2" {
		t.Fatalf("want stack top value %q, got %q", "call2", s.top.value)
	}
	if s.bottom.value != "call1" {
		t.Fatalf("want stack bottom value %q, got %q", "call1", s.bottom.value)
	}
}

func TestStack1Push(t *testing.T) {
	s := new(Stack1)
	s.Push("call1")
	if sDepth := s.length; sDepth != 1 {
		t.Fatalf("want stack depth 1, got %d", sDepth)
	}
	if s.top == nil || s.bottom == nil || s.top != s.bottom {
		t.Fatalf("want non-nil and equal stack top and bottom, got top=%#v and bottom=%#v", s.top, s.bottom)
	}

	s.Push("call2")
	if sDepth := s.length; sDepth != 2 {
		t.Fatalf("want stack depth 2, got %d", sDepth)
	}
	if s.top == s.bottom {
		t.Fatalf("want different stack top and bottom, got top=%#v and bottom=%#v", s.top, s.bottom)
	}
}

func TestStack1Pop(t *testing.T) {
	s := new(Stack1)
	val, err := s.Pop()
	if !errors.Is(err, StackIsEmptyError{}) {
		t.Fatalf("want StackIsEmptyError, got %v", err)
	}
	if val != "" {
		t.Fatalf("want %q, got %q", "", val)
	}

	s.Push("call1")
	s.Push("call2")

	val, err = s.Pop()
	if err != nil {
		t.Fatalf("want nil error, got %v", err)
	}
	if val != "call2" {
		t.Fatalf("want %q, got %q", "call2", val)
	}
	if s.top.value != "call1" {
		t.Fatalf("want stack top value %q, got %q", "call1", s.top.value)
	}
	if s.bottom.value != "call1" {
		t.Fatalf("want stack bottom value %q, got %q", "call1", s.bottom.value)
	}

	val, err = s.Pop()
	if err != nil {
		t.Fatalf("want nil error, got %v", err)
	}
	if val != "call1" {
		t.Fatalf("want %q, got %q", "call1", val)
	}
	if s.top != nil || s.bottom != nil {
		t.Fatalf("want nil stack top and bottom, got top=%#v and bottom=%#v", s.top, s.bottom)
	}
}

func TestStack1String(t *testing.T) {
	s := new(Stack1)
	s.Push("call1")
	fmt.Print(s)
}

func TestStack2Peek(t *testing.T) {
	s := new(Stack2)
	val, err := s.Peek()
	if !errors.Is(err, StackIsEmptyError{}) {
		t.Fatalf("want StackIsEmptyError, got %v", err)
	}
	if val != "" {
		t.Fatalf("want %q, got %q", "", val)
	}

	s.Push("call1")
	s.Push("call2")

	val, err = s.Peek()
	if err != nil {
		t.Fatalf("want nil error, got %v", err)
	}
	if val != "call2" {
		t.Fatalf("want %q, got %q", "call2", val)
	}
	if want := []string{"call1", "call2"}; !reflect.DeepEqual(want, s.data) {
		t.Fatalf("want %v, got %v", want, s.data)
	}
}

func TestStack2Push(t *testing.T) {
	s := new(Stack2)
	s.Push("call1")
	if sDepth := s.length; sDepth != 1 {
		t.Fatalf("want stack depth 1, got %d", sDepth)
	}
	if want := []string{"call1"}; !reflect.DeepEqual(want, s.data) {
		t.Fatalf("want %v, got %v", want, s.data)
	}

	s.Push("call2")
	if sDepth := s.length; sDepth != 2 {
		t.Fatalf("want stack depth 2, got %d", sDepth)
	}
	if want := []string{"call1", "call2"}; !reflect.DeepEqual(want, s.data) {
		t.Fatalf("want %v, got %v", want, s.data)
	}
}

func TestStack2Pop(t *testing.T) {
	s := new(Stack2)
	val, err := s.Pop()
	if !errors.Is(err, StackIsEmptyError{}) {
		t.Fatalf("want StackIsEmptyError, got %v", err)
	}
	if val != "" {
		t.Fatalf("want %q, got %q", "", val)
	}

	s.Push("call1")
	s.Push("call2")

	val, err = s.Pop()
	if err != nil {
		t.Fatalf("want nil error, got %v", err)
	}
	if val != "call2" {
		t.Fatalf("want %q, got %q", "call2", val)
	}
	if want := []string{"call1"}; !reflect.DeepEqual(want, s.data) {
		t.Fatalf("want %v, got %v", want, s.data)
	}

	val, err = s.Pop()
	if err != nil {
		t.Fatalf("want nil error, got %v", err)
	}
	if val != "call1" {
		t.Fatalf("want %q, got %q", "call1", val)
	}
	var want []string = nil
	if !reflect.DeepEqual(want, s.data) {
		t.Fatalf("want %v, got %v", want, s.data)
	}
}

func TestStack2String(t *testing.T) {
	s := new(Stack2)
	s.Push("call1")
	fmt.Print(s)
}
