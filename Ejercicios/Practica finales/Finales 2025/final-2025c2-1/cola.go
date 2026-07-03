package main

type colaEnlazada[T any] struct {
	primero       *nodoCola[T]
	ultimo        *nodoCola[T]
	anteriorMedio *nodoCola[T]
	cant          int
}

type nodoCola[T any] struct {
	dato      T
	siguiente *nodoCola[T]
	anterior  *nodoCola[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada{primero: nil, ultimo: nil, anteriorMedio: nil, cant: 0}
}

func crearNodo[T any](dato T) *nodoCola[T] {
	return &nodoCola[T]{dato: dato, siguiente: nil, anterior: nil}
}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.cant == 0
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.primero.dato
}

func (c *colaEnlazada[T]) Encolar(dato T) {
	nuevoNodo := crearNodo(dato)
	if c.EstaVacia() {
		c.primero = nuevoNodo
		c.ultimo = nuevoNodo
	} else if c.cant%2 != 0 {
		if c.cant == 1 {
			c.anteriorMedio = c.primero
		} else {
			c.anteriorMedio = c.anteriorMedio.siguiente
		}
	}
	if !c.EstaVacia() {
		nuevoNodo.anterior = c.ultimo
		c.ultimo.siguiente = nuevoNodo
	}
	c.ultimo = nuevoNodo
	c.cant++
}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	desencolado := c.primero.dato
	if c.cant%2 != 0 {
		c.anteriorMedio = c.anteriorMedio.siguiente
	}
	c.cant--
	if c.cant <= 1 {
		c.anteriorMedio = nil
	}
	c.primero = c.primero.siguiente
	if c.primero == nil {
		c.ultimo = nil
	} else {
		c.primero.anterior = nil
	}
	return desencolado
}

func (c *colaEnlazada[T]) EliminarMedio() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	medio := c.anteriorMedio.siguiente.dato
	if c.cant == 1 {
		return c.Desencolar()
	}
	// falta mas logica aca
	c.anteriorMedio.siguiente = c.anteriorMedio.siguiente.siguiente
	if c.cant%2 != 0 {
		c.anteriorMedio = c.anteriorMedio.anterior
	}

	c.cant--
	return medio
}
