package utils

import (
	"fmt"
	"math/rand/v2"
)

func RandomTicket(digits int) (int, error) {
	if digits > 9 || digits <= 0 {
		error := fmt.Errorf("Digits must be from 1 and 9")
		return 0, error
	}

	result := 0
	numbers := make([]int, 9)
	for i := 0; i < 9; i++ {
		numbers[i] = i + 1
	}

	for i := 0; i < digits; i++ {
		index := rand.IntN(9)
		for numbers[index%9] == 0 {
			index += 1
		}
		result = result*10 + numbers[index%9]
		numbers[index%9] = 0

	}

	return result, nil

}
