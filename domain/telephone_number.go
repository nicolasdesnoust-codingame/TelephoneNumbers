package domain

import "strconv"

type TelephoneNumber struct {
	numbers []int
}

func NewTelephoneNumber(rawTelephoneNumber string) *TelephoneNumber {
	numbers := parseRawTelephoneNumber(rawTelephoneNumber)

	return &TelephoneNumber{numbers: numbers}
}

func parseRawTelephoneNumber(rawTelephoneNumber string) []int {
	numbers := make([]int, len(rawTelephoneNumber))

	for i, rawNumber := range rawTelephoneNumber {
		numbers[i], _ = strconv.Atoi(string(rawNumber))
	}

	return numbers
}

func (telephoneNumber *TelephoneNumber) GetNumberByIndex(index int) int {
	return telephoneNumber.numbers[index]
}

func (telephoneNumber *TelephoneNumber) Size() int {
	return len(telephoneNumber.numbers)
}
