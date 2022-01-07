package testing

import "errors"

var (
	ErrZero = errors.New("el denominador no puede ser cero")
)

func Restar(num1, num2 int) int {
	return num1 - num2
}

func Dividir(num, den int) (int, error) {
	if den == 0 {
		return 0, ErrZero
	}

	return num / den, nil
}
