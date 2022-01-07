package testing

import (
	"sort"
)

func Ordenar(enteros []int) []int {
	sort.Slice(enteros, func(i, j int) bool {
		return enteros[i] < enteros[j]
	})

	return enteros
}
