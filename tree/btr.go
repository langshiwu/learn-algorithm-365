// tree provides implement of binary search trees
package tree

import (
	"fmt"
	"strings"
)

// 树中每一个节点包含的属性
type Node struct {
	// 可以存储任意类型的数据
	element interface{}
	// 双亲
	parent *Node
	// 左孩子
	left *Node
	// 右孩子
	right *Node
}

//比较器,判断是走左子树还是右子树
type comparer func(interface{}, interface{}) bool

type Bst struct {
	compare comparer
	root    *Node
}

// 创建一个空的二叉树
func (b *Bst) New(compare comparer) *Bst {
	return &Bst{compare: compare}
}

// 中序遍历
// 子树根的关键字位于左子树关键字和右子树关键字的中间
func inorder_tree_walk(tree *Node) {
	if tree == nil {
		return
	}

	inorder_tree_walk(tree.left)
	fmt.Println(tree.element)
	inorder_tree_walk(tree.right)
}

// 查找一个具有给定关键字的节点,运行时间最坏情况为这颗树的高h, O(h)
func (b *Bst) tree_search(tree *Node, element interface{}) *Node {
	if tree == nil {
		return nil
	}

	if tree.element == element {
		return tree
	}

	// 为真则走左，为假则走右
	if b.compare(element, tree.element) {
		return b.tree_search(tree.right, element)
	} else {
		return b.tree_search(tree.left, element)
	}
}

// 迭代法，避免递归的性能消耗， 你懂的
func (b *Bst) iterative_tree_search(tree *Node, element interface{}) *Node {
	if tree == nil {
		return nil
	}

	for element != tree.element {
		if b.compare(element, tree.element) {
			tree = tree.right
		} else {
			tree = tree.left
		}
	}
	return tree
}
