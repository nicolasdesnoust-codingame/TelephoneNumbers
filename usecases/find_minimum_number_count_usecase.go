package usecases

import (
	"telephonenumbers/domain"
)

func FindMinimumNumberCountUsecase(telephoneNumbers []string) int {

	tree := domain.NewTelephoneNumbersTree()
	for _, telephoneNumber := range telephoneNumbers {
		tree.RegisterTelephoneNumber(telephoneNumber)
	}

	return tree.Size()
}
