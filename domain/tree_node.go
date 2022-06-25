package domain

import (
	"fmt"
)

type TreeNode struct {
	value    int
	children map[int]*TreeNode
}

func NewTreeNode(value int) *TreeNode {
	return &TreeNode{value, map[int]*TreeNode{}}
}

func (node *TreeNode) hasChildWithValue(value int) bool {
	child := node.children[value]

	return child != nil
}

func (node *TreeNode) addChildWithValue(value int) *TreeNode {
	child := NewTreeNode(value)

	if node.hasChildWithValue(value) {
		panic(fmt.Sprintf("Node with value %d already has a child with value %d", node.value, value))
	}
	node.children[value] = child

	return child
}

func (node *TreeNode) getChildByValue(value int) *TreeNode {
	return node.children[value]
}
