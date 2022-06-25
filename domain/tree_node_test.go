package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasChildWithValue_ShouldTellValueIsNotAmongItsChildren(t *testing.T) {
	node := NewTreeNode(10)

	hasChildWithValue := node.hasChildWithValue(10)

	assert.False(t, hasChildWithValue)
}

func TestHasChildWithValue_ShouldTellValueIsAmongItsChildren(t *testing.T) {
	node := NewTreeNode(10)
	node.addChildWithValue(12)

	hasChildWithValue := node.hasChildWithValue(10)

	assert.False(t, hasChildWithValue)
}

func TestAddChildWithValue_ShouldNotAddSameChildValueTwice(t *testing.T) {
	node := NewTreeNode(10)
	node.addChildWithValue(12)

	logicUnderTest := func() {
		node.addChildWithValue(12)
	}

	assert.Panics(t, logicUnderTest, "The code did not panic")
}

func TestGetChildByValue_ShouldReturnNilForMissingValue(t *testing.T) {
	node := NewTreeNode(10)

	child := node.getChildByValue(10)

	assert.Nil(t, child)
}

func TestGetChildByValue_ShouldReturnChild(t *testing.T) {
	node := NewTreeNode(10)
	actualChild := node.addChildWithValue(12)

	expectedChild := node.getChildByValue(12)

	assert.Equal(t, expectedChild, actualChild)
}
