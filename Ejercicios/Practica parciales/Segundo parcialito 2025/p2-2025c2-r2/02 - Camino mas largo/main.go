// Implementar una primitiva para el árbol binario que nos devuelva una lista con los datos de los nodos que forman del camino más
// largo que exista desde la raíz del árbol hasta alguna hoja.
// Justificar el orden de complejidad de la primitiva implementada. Indicar en qué orden queda descrito dicho camino en la lista que
// devuelve la primitiva.

package main

type ab[T any] struct {
	izq  *ab[T]
	der  *ab[T]
	dato T
}

func (ab *ab[T]) CaminoLargo() Lista[T] {
	return caminoLargo(ab)
}

func caminoLargo[T any](nodo *ab[T]) Lista[T] {
	if nodo == nil {
		return CrearListaEnlazada[T]()
	}
	listaIzq := caminoLargo(nodo.izq)
	listaDer := caminoLargo(nodo.der)
	if listaIzq.Cantidad() >= listaDer.Cantidad() {
		listaIzq.InsertarPrimero(nodo.dato)
		return listaIzq
	} else {
		listaDer.InsertarPrimero(nodo.dato)
		return listaDer
	}
}
