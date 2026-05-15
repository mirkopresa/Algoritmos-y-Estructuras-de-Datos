// Implementar una primitiva para el ABB, que devuelva una lista con las claves del mismo,
// ordenadas tal que si insertáramos las claves en un ABB vacío, dicho ABB tendría la misma estructura que el árbol original.

package main

func (abb *abb[K, V]) Claves() Lista[K] {
	listaClaves := CrearListaEnlazada[K]()
	return nil
}

func (nodo *nodoAbb[K, V]) _obtenerClaves(lista Lista[K]) {
	if nodo == nil {
		return
	}
	lista.InsertarUltimo(nodo.clave)
	nodo.izq._obtenerClaves(lista)
	nodo.der._obtenerClaves(lista)
}
