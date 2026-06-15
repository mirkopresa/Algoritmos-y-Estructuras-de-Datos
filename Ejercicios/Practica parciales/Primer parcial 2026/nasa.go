// La NASA está preparando el lanzamiento de Artemis III y necesita de nuestra ayuda para que el sistema no sufra
// desperfectos. El satélite utilizará una implementación de ListaEnlazada con esta estructura interna:
// type ListaEnlazada[T any] struct {
// 		primero *nodoLista[T]
// }
// type nodoLista[T any] struct {
// 		dato T
// 		siguiente *nodoLista[T]
// }

// Debido a la ingerencia de rayos cósmicos, la integridad de esta estructura enlazada puede verse afectada. Estos rayos
// pueden generar que el siguiente de cada NodoLista se vea afectado y en vez de apuntar al nodo siguiente que le
// corresponde, apuntar a otro anterior, generando ciclos.
// Se pide implementar una primitiva que detecte si la lista tiene un ciclo. La función debe tener una complejidad temporal
// de O(n) (siendo n la cantidad de elementos de la lista) y espacial de O(1).
// Pista: pensar en el ejercicio (y resolución) para obtener el K-último de una lista sin tener largo, en una única iteración.

package main

type ListaEnlazada[T any] struct {
	primero *nodoLista[T]
}

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

func (lista *ListaEnlazada[T]) HayCiclo() bool {
	lento := lista.primero
	rapido := lista.primero
	for lento != nil {
		lento := lento.siguiente
		rapido := rapido.siguiente.siguiente
		if lento == rapido {
			return true
		}
	}
	return false
}
