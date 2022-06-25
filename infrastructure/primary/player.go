package infrastructure

import (
	"fmt"
	"telephonenumbers/usecases"
)

func main() {
	telephoneNumbers := parseInputs()

	minimumNumberCount := usecases.FindMinimumNumberCountUsecase(telephoneNumbers)

	fmt.Println(minimumNumberCount)
}

func parseInputs() []string {
	var telephoneNumberCount int
	fmt.Scan(&telephoneNumberCount)

	telephoneNumbers := make([]string, telephoneNumberCount)
	for i := 0; i < telephoneNumberCount; i++ {
		fmt.Scan(&telephoneNumbers[i])
	}

	return telephoneNumbers
}
