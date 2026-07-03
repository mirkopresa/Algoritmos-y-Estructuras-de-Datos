// Explicar en detalle cómo implementar una estructura de pila enlazada como implementación del TDA Pila tal que asegure que todas
// las primitivas del TDA se ejecuten en tiempo constante.
// explicado en mi cuaderno

package main

type pilaEnlazada[T any] struct {
	tope *nodoPila[T]
}

type nodoPila[T any] struct {
	dato     T
	anterior *nodoPila[T]
}

func CrearPilaEnlazada[T any]() Pila[T] {
	return &pilaEnlazada[T]{tope: nil}
}

func crearNodo[T any](dato T) *nodoPila[T] {
	return &nodoPila[T]{dato: dato, anterior: nil}
}

func (pila *pilaEnlazada[T]) EstaVacia() bool {
	return pila.tope == nil
}

func (pila *pilaEnlazada[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.tope.dato
}

func (pila *pilaEnlazada[T]) Apilar(dato T) {
	nuevoNodo := crearNodo(dato)
	nuevoNodo.anterior = pila.tope
	pila.tope = nuevoNodo
}

func (pila *pilaEnlazada[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	desapilado := pila.tope.dato
	pila.tope = pila.tope.anterior
	return desapilado
}
