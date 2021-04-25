package tree

import (
	"errors"
	"testing"
)

func TestNewBinarySearchTree(t *testing.T) {
	root := 7
	bst := NewBinarySearchTree(root)

	if height := bst.Height(); height != 1 {
		t.Fatalf("want bst height == %d, got bst height == %d", 1, height)
	}
	if bst.root.value != root {
		t.Fatalf("want bst root value == %d, got bst root value == %d", root, bst.root.value)
	}
	if bst.root.left != nil {
		t.Fatalf("want bst root left == nil, got bst root left == %v", bst.root.left)
	}
	if bst.root.right != nil {
		t.Fatalf("want bst root right == nil, got bst root right == %v", bst.root.right)
	}
}

func TestBSTValueExistsError(t *testing.T) {
	var err BSTValueExistsError
	want := "0 already exists in binary search tree"
	if msg := err.Error(); msg != want {
		t.Fatalf("want error string == %s, got error string == %s", want, msg)
	}

	err = BSTValueExistsError(8)
	want = "8 already exists in binary search tree"
	if msg := err.Error(); msg != want {
		t.Fatalf("want error string == %s, got error string == %s", want, msg)
	}
}

func TestBinarySearchTreeInsert(t *testing.T) {
	root := 7
	bst := NewBinarySearchTree(root)

	left := 3
	err := bst.Insert(left)
	if err != nil {
		t.Fatalf("want error == %v, got error == %v", nil, err)
	}
	if wantLeft := (&node{left, nil, nil}); *bst.root.left != *wantLeft {
		t.Fatalf("want bst root left == %v, got bst root left == %v", wantLeft, bst.root.left)
	}
	if bst.root.right != nil {
		t.Fatalf("want bst root right == %v, got bst root right == %v", nil, bst.root.right)
	}
	if height := bst.Height(); height != 2 {
		t.Fatalf("want bst height == %d, got bst height == %d", 2, height)
	}

	right := 9
	err = bst.Insert(right)
	if err != nil {
		t.Fatalf("want error == %v, got error == %v", nil, err)
	}
	if wantLeft := (&node{left, nil, nil}); *bst.root.left != *wantLeft {
		t.Fatalf("want bst root left == %v, got bst root left == %v", wantLeft, bst.root.left)
	}
	if wantRight := (&node{right, nil, nil}); *bst.root.right != *wantRight {
		t.Fatalf("want bst root right == %v, got bst root right == %v", wantRight, bst.root.right)
	}
	if height := bst.Height(); height != 2 {
		t.Fatalf("want bst height == %d, got bst height == %d", 2, height)
	}

	duplicate := 9
	err = bst.Insert(duplicate)
	if !errors.Is(err, BSTValueExistsError(duplicate)) {
		t.Fatalf("want error == %v, got error == %v", BSTValueExistsError(duplicate), err)
	}
	if wantLeft := (&node{left, nil, nil}); *bst.root.left != *wantLeft {
		t.Fatalf("want bst root left == %v, got bst root left == %v", wantLeft, bst.root.left)
	}
	if wantRight := (&node{right, nil, nil}); *bst.root.right != *wantRight {
		t.Fatalf("want bst root right == %v, got bst root right == %v", wantRight, bst.root.right)
	}
	if height := bst.Height(); height != 2 {
		t.Fatalf("want bst height == %d, got bst height == %d", 2, height)
	}

	err = bst.Insert(root)
	if !errors.Is(err, BSTValueExistsError(root)) {
		t.Fatalf("want error == %v, got error == %v", BSTValueExistsError(root), err)
	}
	if wantLeft := (&node{left, nil, nil}); *bst.root.left != *wantLeft {
		t.Fatalf("want bst root left == %v, got bst root left == %v", wantLeft, bst.root.left)
	}
	if wantRight := (&node{right, nil, nil}); *bst.root.right != *wantRight {
		t.Fatalf("want bst root right == %v, got bst root right == %v", wantRight, bst.root.right)
	}
	if height := bst.Height(); height != 2 {
		t.Fatalf("want bst height == %d, got bst height == %d", 2, height)
	}

	another := 5
	err = bst.Insert(another)
	if err != nil {
		t.Fatalf("want error == %v, got error == %v", nil, err)
	}
	if want := (&node{another, nil, nil}); *bst.root.left.right != *want {
		t.Fatalf("want bst root left == %v, got bst root left == %v", want, *bst.root.left.right)
	}
	if wantRight := (&node{right, nil, nil}); *bst.root.right != *wantRight {
		t.Fatalf("want bst root right == %v, got bst root right == %v", wantRight, bst.root.right)
	}
	if height := bst.Height(); height != 3 {
		t.Fatalf("want bst height == %d, got bst height == %d", 3, height)
	}
}

func TestBinarySearchTreeHeight(t *testing.T) {
	data := []int{1, 2, 3}
	bst := NewBinarySearchTree(0)
	if got := bst.Height(); got != 1 {
		t.Fatalf("want bst height == %d, got bst height == %d", 1, got)
	}

	for i, d := range data {
		bst.Insert(d)
		if got := bst.Height(); got != i+2 {
			t.Fatalf("want bst height == %d, got bst height == %d", i+2, got)
		}
	}

	// TODO: test height after deletion
}
