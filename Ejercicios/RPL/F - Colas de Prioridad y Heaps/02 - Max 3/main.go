// Implementar una primitiva para el heap (de máximos) que obtenga los 3 elementos
// más grandes del heap en tiempo constante (O(1)).

package main

type fcmpHeap[T any] func(T, T) int

type heap[T any] struct {
	datos    []T
	cantidad int
	cmp      fcmpHeap[T]
}

func (h heap[T]) Max3() []T {
	resultado := make([]T, 0)
	// Primer maximo
	resultado = append(resultado, h.datos[0])

	// Segundo maximo
	comparacion := h.cmp(h.datos[1], h.datos[2])
	var menor T
	var hizqMayor int
	var hderMayor int
	if comparacion < 0 {
		resultado = append(resultado, h.datos[2])
		menor = h.datos[1]
		hizqMayor = 5
		hderMayor = 6
	} else {
		resultado = append(resultado, h.datos[1])
		menor = h.datos[2]
		hizqMayor = 3
		hderMayor = 4
	}

	// Tercer maximo
	if h.cmp(menor, h.datos[hizqMayor]) > 0 && h.cmp(menor, h.datos[hderMayor]) > 0 {
		resultado = append(resultado, menor)
	} else {
		comparacion := h.cmp(h.datos[hizqMayor], h.datos[hderMayor])
		if comparacion < 0 {
			resultado = append(resultado, h.datos[hderMayor])
		} else {
			resultado = append(resultado, h.datos[hizqMayor])
		}
	}
	return resultado
}
