// Se tiene un árbol binario de búsqueda con cadenas como claves y función de comparación strcmp.
// Implementar una primitiva func (abb *abb[K, V]) Mayores(clave K) Lista[K] que, dados un ABB y una clave,
// devuelva una lista ordenada con las claves del árbol estrictamente mayores a la recibida por parámetro
// (que no necesariamente está en el árbol).
// Implementar sin utilizar el iterador Interno del ABB.

package main

func (abb *abb[K, V]) Mayores(clave K) Lista[K] {
	resultado := CrearListaEnlazada[K]()
	abb.raiz._Mayores(resultado, clave, abb.cmp)
	return resultado
}

func (nodo *nodoAbb[K, V]) _Mayores(lista Lista[K], clave K, cmp funcCmp[K]) {
	if nodo == nil {
		return
	}
	comparacion := cmp(clave, nodo.clave)
	if comparacion < 0 {
		nodo.izquierdo._Mayores(lista, clave)
		lista.InsertarUltimo(nodo.dato)
		nodo.derecho._Mayores(lista, clave)
	} else {
		nodo.derecho._Mayores(lista, clave)
	}
}
