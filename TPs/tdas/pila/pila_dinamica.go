package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

const (
	_CAPACIDAD_MINIMA = 10
	_REDIMENSION      = 2
	_REDUCCION        = 4
)

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{datos: make([]T, _CAPACIDAD_MINIMA), cantidad: 0}
}

func (p pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(dato T) {
	if cap(p.datos) == p.cantidad {
		p.redimensionar(cap(p.datos) * _REDIMENSION)
	}
	p.datos[p.cantidad] = dato
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	desapilado := p.datos[p.cantidad-1]
	p.cantidad--
	if p.cantidad <= cap(p.datos)/_REDUCCION && cap(p.datos) > _CAPACIDAD_MINIMA {
		p.redimensionar(cap(p.datos) / _REDIMENSION)
	}
	return desapilado
}

// Funcion auxiliar

func (p *pilaDinamica[T]) redimensionar(redimension int) {
	pilaNueva := make([]T, redimension)
	copy(pilaNueva, p.datos)
	p.datos = pilaNueva
}
