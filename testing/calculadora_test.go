package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResta(t *testing.T) {
	num1 := 10
	num2 := 5
	resultadoEsperado := 5

	resultado := Restar(num1, num2)

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}

func TestDividirCero(t *testing.T) {
	num := 10
	den := 0
	resultadoEsperado := 0

	resultado, err := Dividir(num, den)

	assert.Equal(t, ErrZero, err, "debería dar error y ser el mismo")
	assert.Equal(t, resultadoEsperado, resultado, "debería ser el mismo")
}

func TestDividir(t *testing.T) {
	num := 10
	den := 2
	resultadoEsperado := 5

	resultado, err := Dividir(num, den)

	assert.Nil(t, err, "no debería dar error")
	assert.Equal(t, resultadoEsperado, resultado, "debería ser el mismo")
}
