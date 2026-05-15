// Se tiene un árbol binario de búsqueda con cadenas como claves y función de comparación strcmp.
// Implementar una primitiva func (abb *abb[K, V]) Mayores(clave K) Lista[K] que, dados un ABB y una clave,
// devuelva una lista ordenada con las claves del árbol estrictamente mayores a la recibida por parámetro
// (que no necesariamente está en el árbol).
// Implementar sin utilizar el iterador Interno del ABB.

package main

func (abb *abb[K, V]) Mayores(clave K) Lista[K] {
	resultado := CrearListaEnlazada[K]()
	abb.raiz._Mayores(resultado, clave)
	return resultado
}

func (nodo *nodoAbb[K, V]) _Mayores(lista listaEnlazada[K], clave K) {
	if nodo == nil {
		return
	}
	nodo.izq._Mayores(lista, clave)
	lista.InsertarUltimo(nodo.dato)
	nodo.der._Mayores(lista, clave)
}
