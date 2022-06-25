package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinimumNumberCount_ShouldReturnNoNumbersForEmptyInput(t *testing.T) {

	phoneNumbers := []string{}

	actualMinimumNumberCount := FindMinimumNumberCountUsecase(phoneNumbers)

	assert.Zero(t, actualMinimumNumberCount)
}

func TestFindMinimumNumberCount_ShouldReturnValidResultForASinglePhoneNumber(t *testing.T) {

	phoneNumbers := []string{"0687653423"}

	actualMinimumNumberCount := FindMinimumNumberCountUsecase(phoneNumbers)

	assert.Equal(t, 10, actualMinimumNumberCount)
}

func TestFindMinimumNumberCount_ShouldReturnValidResultWhenPhoneNumbersDoesNotShareCommonPrefix(t *testing.T) {

	phoneNumbers := []string{"0687653423", "3687653423"}

	actualMinimumNumberCount := FindMinimumNumberCountUsecase(phoneNumbers)

	assert.Equal(t, 20, actualMinimumNumberCount)
}

func TestFindMinimumNumberCount_ShouldReturnValidResultWhenPhoneNumbersShareCommonPrefix(t *testing.T) {

	commonPrefix := "068765"
	phoneNumbers := []string{commonPrefix + "3423", commonPrefix + "5190"}

	actualMinimumNumberCount := FindMinimumNumberCountUsecase(phoneNumbers)

	assert.Equal(t, 14, actualMinimumNumberCount)
}
