/*
	Большие числа и операции

	Разработать программу, которая перемножает, делит, складывает,
	вычитает две числовых переменных a, b, значения которых > 2^20 (больше 1 миллион).

	Комментарий: в Go тип int справится с такими числами,
	но обратите внимание на возможное переполнение для ещё больших значений.
	Для очень больших чисел можно использовать math/big.
*/

package main

import (
	"fmt"
	"math"
	"math/big"
)

func addition(a, b int64) any {
	if (b > 0 && a > math.MaxInt64-b) || (b < 0 && a < math.MinInt64-b) {
		result := new(big.Int).Add(big.NewInt(a), big.NewInt(b))
		return result
	}
	return a + b
}

func subtraction(a, b int64) any {
	if (b > 0 && a < math.MinInt64+b) || (b < 0 && a > math.MaxInt64+b) {
		result := new(big.Int).Sub(big.NewInt(a), big.NewInt(b))
		return result
	}
	return a - b
}

func multiplication(a, b int64) any {
	if a == 0 || b == 0 {
		return int64(0)
	}
	if a == -1 && b == math.MinInt64 || b == -1 && a == math.MinInt64 {
		result := new(big.Int).Mul(big.NewInt(a), big.NewInt(b))
		return result
	}
	if a > math.MaxInt64/b || a < math.MinInt64/b {
		result := new(big.Int).Mul(big.NewInt(a), big.NewInt(b))
		return result
	}
	return a * b
}

func division(a, b int64) (any, error) {
	if b == 0 {
		return nil, fmt.Errorf("деление на ноль")
	}
	if a == math.MinInt64 && b == -1 {
		result := new(big.Int).Div(big.NewInt(a), big.NewInt(b))
		return result, nil
	}
	return a / b, nil
}

func main() {
	a := int64(math.MaxInt64 - 5)
	b := int64(10)

	sum := addition(a, b)
	sub := subtraction(a, b)
	mul := multiplication(a, b)
	div, _ := division(a, b)

	fmt.Println("Сумма:", sum)
	fmt.Println("Разность:", sub)
	fmt.Println("Произведение:", mul)
	fmt.Println("Частное:", div)
}
