package infrastructure

import (
	"fmt"
	"telephonenumbers/usecases"
)

func main() {
	var N int
	fmt.Scan(&N)

	telephoneNumbers := make([]string, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&telephoneNumbers[i])
	}

	minimumNumberCount := usecases.FindMinimumNumberCountUsecase(telephoneNumbers)

	fmt.Println(minimumNumberCount)
}
