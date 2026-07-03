// Implementar una primitiva para el árbol binario esABB(comparacion func(T, T) int) bool que determine si el mismo cumple
// propiedad de ABB dada la función de comparación pasada por parámetro. Indicar y justificar la complejidad del algoritmo
// implementado.

package main

type Arbol[T any] struct {
	izq  *Arbol[T]
	der  *Arbol[T]
	dato T
}

func (arbol *Arbol[T]) esABB(comparacion func(T, T) int) bool {
	return arbol._esABB(comparacion, nil, nil)
}

func (arbol *Arbol[T]) _esABB(cmp func(T, T) int, max, min *T) bool {
	if arbol == nil {
		return true
	}
	if max != nil {
		if cmp(arbol.dato, *max) > 0 {
			return false
		}
	}
	if min != nil {
		if cmp(arbol.dato, *min) < 0 {
			return false
		}
	}
	abbIzq := arbol.izq._esABB(cmp, &arbol.dato, min)
	abbDer := arbol.der._esABB(cmp, max, &arbol.dato)
	return abbIzq && abbDer
}
