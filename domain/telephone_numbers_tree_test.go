package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTelephoneNumberTree_ShouldReturnSizeOfZeroWhenInitialized(t *testing.T) {
	tree := NewTelephoneNumbersTree()

	actualSize := tree.Size()

	assert.Zero(t, actualSize)
}

func TestTelephoneNumberTree_ShouldReturnSizeOfOneRegisteredTelephoneNumber(t *testing.T) {
	tree := NewTelephoneNumbersTree()
	tree.RegisterTelephoneNumber(NewTelephoneNumber("0657839210"))

	actualSize := tree.Size()

	expectedSize := 10
	assert.Equal(t, expectedSize, actualSize)
}

func TestTelephoneNumberTree_ShouldReturnSizeOfMultipleRegisteredTelephoneNumberThatAreIndependent(t *testing.T) {
	tree := NewTelephoneNumbersTree()
	tree.RegisterTelephoneNumber(NewTelephoneNumber("0657839210"))
	tree.RegisterTelephoneNumber(NewTelephoneNumber("3657839210"))

	actualSize := tree.Size()

	expectedSize := 20
	assert.Equal(t, expectedSize, actualSize)
}

func TestTelephoneNumberTree_ShouldReturnSizeOfMultipleRegisteredTelephoneNumberThatShareACommonPrefix(t *testing.T) {
	tree := NewTelephoneNumbersTree()
	commonPrefix := "065783"
	tree.RegisterTelephoneNumber(NewTelephoneNumber(commonPrefix + "9210"))
	tree.RegisterTelephoneNumber(NewTelephoneNumber(commonPrefix + "5431"))

	actualSize := tree.Size()

	expectedSize := 14
	assert.Equal(t, expectedSize, actualSize)
}
