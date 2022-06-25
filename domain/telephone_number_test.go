package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSize_ShouldReturnValidSize(t *testing.T) {
	telephoneNumber := NewTelephoneNumber("0621675898")

	actualSize := telephoneNumber.Size()

	expectedSize := 10
	assert.Equal(t, expectedSize, actualSize)
}

func TestGetNumberByIndex_ShouldReturnValidNumber(t *testing.T) {
	telephoneNumber := NewTelephoneNumber("0621675898")

	actualNumber := telephoneNumber.GetNumberByIndex(7)

	expectedNumber := 8
	assert.Equal(t, expectedNumber, actualNumber)
}

func TestGetNumberByIndex_ShouldPanicWhenIndexIsNegative(t *testing.T) {
	telephoneNumber := NewTelephoneNumber("0621675898")

	logicUnderTest := func() {
		telephoneNumber.GetNumberByIndex(-1)
	}

	assert.Panics(t, logicUnderTest, "The code did not panic")
}

func TestGetNumberByIndex_ShouldPanicWhenIndexIsEqualToSize(t *testing.T) {
	telephoneNumber := NewTelephoneNumber("0621675898")

	logicUnderTest := func() {
		telephoneNumber.GetNumberByIndex(10)
	}

	assert.Panics(t, logicUnderTest, "The code did not panic")
}
