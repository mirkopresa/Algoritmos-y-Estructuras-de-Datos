// Escribir una primitiva que reciba una lista y devuelva
// el elemento que esté a k posiciones del final (el ante-k-último),
// recorriendo la lista una sola vez y sin usar estructuras auxiliares.
// Considerar que k es siempre menor al largo de la lista.
// Por ejemplo, si se recibe la lista [ 1, 5, 10, 3, 6, 8 ], y k = 4, debe devolver 10.
// Indicar el orden de complejidad de la primitiva.

package main

func (lista *ListaEnlazada[T]) AnteKUltimo(k int) T {
	listaPrimera := lista.prim
	listaSegunda := lista.prim
	for i := 0; i < k; i++ {
		listaPrimera = listaPrimera.prox
	}
	for listaPrimera != nil {
		listaPrimera = listaPrimera.prox
		listaSegunda = listaSegunda.prox
	}
	return listaSegunda.dato
}
