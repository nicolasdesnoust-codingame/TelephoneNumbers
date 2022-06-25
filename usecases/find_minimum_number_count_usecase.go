package usecases

import (
	"telephonenumbers/domain"
)

func FindMinimumNumberCountUsecase(rawTelephoneNumbers []string) int {
	tree := domain.NewTelephoneNumbersTree()

	for _, rawTelephoneNumber := range rawTelephoneNumbers {
		telephoneNumber := domain.NewTelephoneNumber(rawTelephoneNumber)
		tree.RegisterTelephoneNumber(telephoneNumber)
	}

	return tree.Size()
}
