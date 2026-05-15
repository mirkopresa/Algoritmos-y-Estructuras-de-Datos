// Dado un árbol binario, escribir una primitiva recursiva que determine la altura del mismo.
// Indicar y justificar el orden de la primitiva.

package main

type ab struct {
	izq  *ab
	der  *ab
	dato int
}

func (arbol *ab) Altura() int {
	if arbol == nil {
		return 0
	}
	alturaIzquierda := arbol.izq.Altura()
	alturaDerecha := arbol.der.Altura()
	return max(alturaIzquierda, alturaDerecha) + 1
}
