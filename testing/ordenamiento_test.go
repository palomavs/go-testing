package testing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenamiento(t *testing.T) {
	enteros := []int{5, 4, 88, 1, -8, 0}

	resultadoEsperado := []int{-8, 0, 1, 4, 5, 88}
	resultado := Ordenar(enteros)

	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")
}
