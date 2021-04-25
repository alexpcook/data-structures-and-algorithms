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

func TestBSTValueNotExistsError(t *testing.T) {
	var err BSTValueNotExistsError
	want := "0 does not exist in binary search tree"
	if msg := err.Error(); msg != want {
		t.Fatalf("want error string == %s, got error string == %s", want, msg)
	}

	err = BSTValueNotExistsError(8)
	want = "8 does not exist in binary search tree"
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

func TestBinarySearchTreeLookup(t *testing.T) {
	bst := NewBinarySearchTree(0)

	data := []int{5, 9, 13, 3, 5, 1}
	for _, d := range data {
		bst.Insert(d)
	}

	for _, d := range append(data, 0) {
		err := bst.Lookup(d)
		if err != nil {
			t.Fatalf("want error == %v, got error == %v", nil, err)
		}
	}

	for _, d := range []int{7, 18, -4, 2} {
		err := bst.Lookup(d)
		if want := BSTValueNotExistsError(d); !errors.Is(err, want) {
			t.Fatalf("want error == %v, got error == %v", want, err)
		}
	}
}

func TestBinarySearchTreeDelete(t *testing.T) {
	bst := NewBinarySearchTree(7)

	rootNodeErr := BSTDeleteRootNodeError(7)
	if err := bst.Delete(7); !errors.Is(err, rootNodeErr) {
		t.Fatalf("want error == %v, got error == %v", rootNodeErr, err)
	}

	notExistsErr := BSTValueNotExistsError(8)
	if err := bst.Delete(8); !errors.Is(err, notExistsErr) {
		t.Fatalf("want error == %v, got error == %v", notExistsErr, err)
	}

	if err := bst.Insert(9); err != nil {
		t.Fatalf("want error == %v, got error == %v", nil, err)
	}

	switch err := bst.Delete(9); {
	case err != nil:
		t.Fatalf("want error == %v, got error == %v", nil, err)
	case bst.root.value != 7:
		t.Fatalf("want root value == %d, got root value == %d", 7, bst.root.value)
	case bst.root.left != nil:
		t.Fatalf("want bst root left == %v, got bst root left == %v", nil, bst.root.left)
	case bst.root.right != nil:
		t.Fatalf("want bst root right == %v, got bst root right == %v", nil, bst.root.right)
	}

	if err := bst.Insert(5); err != nil {
		t.Fatalf("want error == %v, got error == %v", nil, err)
	}

	switch err := bst.Delete(5); {
	case err != nil:
		t.Fatalf("want error == %v, got error == %v", nil, err)
	case bst.root.value != 7:
		t.Fatalf("want root value == %d, got root value == %d", 7, bst.root.value)
	case bst.root.left != nil:
		t.Fatalf("want bst root left == %v, got bst root left == %v", nil, bst.root.left)
	case bst.root.right != nil:
		t.Fatalf("want bst root right == %v, got bst root right == %v", nil, bst.root.right)
	}

	if err := bst.Insert(5); err != nil {
		t.Fatalf("want error == %v, got error == %v", nil, err)
	}

	switch err := bst.Delete(7); {
	case err != nil:
		t.Fatalf("want error == %v, got error == %v", nil, err)
	case bst.root.value != 5:
		t.Fatalf("want root value == %d, got root value == %d", 5, bst.root.value)
	case bst.root.left != nil:
		t.Fatalf("want bst root left == %v, got bst root left == %v", nil, bst.root.left)
	case bst.root.right != nil:
		t.Fatalf("want bst root right == %v, got bst root right == %v", nil, bst.root.right)
	}

	if err := bst.Insert(7); err != nil {
		t.Fatalf("want error == %v, got error == %v", nil, err)
	}

	switch err := bst.Delete(5); {
	case err != nil:
		t.Fatalf("want error == %v, got error == %v", nil, err)
	case bst.root.value != 7:
		t.Fatalf("want root value == %d, got root value == %d", 7, bst.root.value)
	case bst.root.left != nil:
		t.Fatalf("want bst root left == %v, got bst root left == %v", nil, bst.root.left)
	case bst.root.right != nil:
		t.Fatalf("want bst root right == %v, got bst root right == %v", nil, bst.root.right)
	}

	if err := bst.Insert(3); err != nil {
		t.Fatalf("want error == %v, got error == %v", nil, err)
	}
	if err := bst.Insert(5); err != nil {
		t.Fatalf("want error == %v, got error == %v", nil, err)
	}

	switch err := bst.Delete(3); {
	case err != nil:
		t.Fatalf("want error == %v, got error == %v", nil, err)
	case bst.root.value != 7:
		t.Fatalf("want root value == %d, got root value == %d", 7, bst.root.value)
	case *bst.root.left != node{5, nil, nil}:
		t.Fatalf("want bst root left == %v, got bst root left == %v", node{5, nil, nil}, *bst.root.left)
	case bst.root.right != nil:
		t.Fatalf("want bst root right == %v, got bst root right == %v", nil, bst.root.right)
	}

	bst = NewBinarySearchTree(7)
	for _, d := range []int{3, 1, 5, 4} {
		err := bst.Insert(d)
		if err != nil {
			t.Fatalf("want error == %v, got error == %v", nil, err)
		}
	}

	switch err := bst.Delete(3); {
	case err != nil:
		t.Fatalf("want error == %v, got error == %v", nil, err)
	case bst.root.left.value != 4:
		t.Fatalf("want bst root left value == %d, got bst root left value == %d", 4, bst.root.left.value)
	case bst.root.right != nil:
		t.Fatalf("want bst root right == %v, got bst root right == %v", nil, bst.root.right)
	case bst.root.left.left.value != 1:
		t.Fatalf("want bst root left left value == %d, got bst root left left value == %d", 1, bst.root.left.left.value)
	case bst.root.left.right.value != 5:
		t.Fatalf("want bst root left right value == %d, got bst root left right value == %d", 1, bst.root.left.right.value)
	case bst.root.left.right.left != nil:
		t.Fatalf("want bst root left right left == %v, got bst root left right left == %v", nil, bst.root.left.right.left)
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
