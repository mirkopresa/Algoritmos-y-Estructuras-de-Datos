package cola

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

type colaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{primero: nil, ultimo: nil}
}

func crearNodo[T any](elemento T) *nodo[T] {
	return &nodo[T]{dato: elemento, siguiente: nil}
}

func (c colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil && c.ultimo == nil
}

func (c colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.primero.dato
}

func (c *colaEnlazada[T]) Encolar(elemento T) {
	nuevoNodo := crearNodo(elemento)
	if c.EstaVacia() {
		c.primero = nuevoNodo
	} else {
		c.ultimo.siguiente = nuevoNodo
	}
	c.ultimo = nuevoNodo
}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	desencolado := c.primero.dato
	// Correccion del profe que no voy a entregar
	c.primero = c.primero.siguiente
	if c.primero == nil {
		c.ultimo = nil
	}
	return desencolado
}
