// Ejercicio 2: ¿Árbol paga ABL?
// Implementar una primitiva para el Árbol Binario que determine si el mismo cumple propiedad de AVL.
// Indicar y justificar la complejidad de algoritmo. Ojo, cuidado con la complejidad.

package ejparcialclase

type ab[T any] struct {
	izq  *ab[T]
	der  *ab[T]
	dato T
	cmp  func(T, T) int
}

func (ab *ab[T]) EsAVL() bool {
	_, avl := ab.esAVL(nil, nil) // pseudo codigo
	return avl
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// A = 2, B = 2, C = 0, log en base 2 de 2 = 1, 1 > C, entonces O(n ^ (log 2 de 2)) -> O(n)
func (ab *ab[T]) esAVL(minimo, maximo *T) (int, bool) {
	if ab == nil {
		return 0, true
	}
	if minimo != nil && ab.cmp(ab.dato, *minimo) < 0 || maximo != nil && ab.cmp(ab.dato, *maximo) > 0 {
		return 0, false
	}
	alturaIzq, avlIzq := ab.izq.esAVL(minimo, &ab.dato)
	alturaDer, avlDer := ab.der.esAVL(&ab.dato, maximo)
	alturaActual := max(alturaIzq, alturaDer) + 1
	if abs(alturaIzq-alturaDer) > 1 {
		return alturaActual, false
	}
	return alturaActual, avlIzq && avlDer
}

// ejercicio feo
