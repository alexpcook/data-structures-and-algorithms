package stack

import (
	"errors"
	"fmt"
	"testing"
)

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
}

func TestStack1String(t *testing.T) {
	s := new(Stack1)
	s.Push("call1")
	fmt.Print(s)
}
