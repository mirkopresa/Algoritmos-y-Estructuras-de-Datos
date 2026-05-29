package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	largo   int
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
}

type iteradorListaEnlazada[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{largo: 0, primero: nil, ultimo: nil}
}

func crearNodo[T any](elemento T, siguiente *nodoLista[T]) *nodoLista[T] {
	return &nodoLista[T]{dato: elemento, siguiente: siguiente}
}

func (l listaEnlazada[T]) EstaVacia() bool {
	return l.largo == 0
}

func (l *listaEnlazada[T]) InsertarPrimero(nuevoDato T) {
	nuevoNodo := crearNodo(nuevoDato, l.primero)
	if l.EstaVacia() {
		l.ultimo = nuevoNodo
	}
	l.primero = nuevoNodo
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(nuevoDato T) {
	nuevoNodo := crearNodo(nuevoDato, nil)
	if l.EstaVacia() {
		l.primero = nuevoNodo
	} else {
		l.ultimo.siguiente = nuevoNodo
	}
	l.ultimo = nuevoNodo
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	datoBorrado := l.primero.dato
	l.primero = l.primero.siguiente
	l.largo--
	if l.EstaVacia() {
		l.ultimo = nil
	}
	return datoBorrado
}

func (l listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.primero.dato
}

func (l listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.ultimo.dato
}

func (l listaEnlazada[T]) Largo() int {
	return l.largo
}

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := l.primero
	for actual != nil {
		if !visitar(actual.dato) {
			break
		}
		actual = actual.siguiente
	}
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorListaEnlazada[T]{lista: l, actual: l.primero, anterior: nil}
}

func (iter *iteradorListaEnlazada[T]) VerActual() T {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.dato
}

func (iter *iteradorListaEnlazada[T]) HayAlgoMas() bool {
	return iter.actual != nil
}

func (iter *iteradorListaEnlazada[T]) Avanzar() {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iteradorListaEnlazada[T]) Insertar(nuevoDato T) {
	nuevoNodo := crearNodo(nuevoDato, iter.actual)
	if iter.lista.primero == iter.actual {
		iter.lista.primero = nuevoNodo
	} else {
		iter.anterior.siguiente = nuevoNodo
	}
	if iter.actual == nil {
		iter.lista.ultimo = nuevoNodo
	}
	iter.actual = nuevoNodo
	iter.lista.largo++
}

func (iter *iteradorListaEnlazada[T]) Borrar() T {
	if !iter.HayAlgoMas() {
		panic("El iterador termino de iterar")
	}
	datoBorrado := iter.actual.dato
	if iter.lista.primero == iter.actual {
		iter.lista.primero = iter.actual.siguiente
	} else {
		iter.anterior.siguiente = iter.actual.siguiente
	}
	if iter.actual.siguiente == nil {
		iter.lista.ultimo = iter.anterior
	}
	iter.actual = iter.actual.siguiente
	iter.lista.largo--
	return datoBorrado
}
