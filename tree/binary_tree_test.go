package tree

import "testing"

func TestNewBinarySearchTree(t *testing.T) {
	root := 7
	bst := NewBinarySearchTree(root)

	if bst.height != 1 {
		t.Fatalf("want bst height == 1, got bst height == %d", bst.height)
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
