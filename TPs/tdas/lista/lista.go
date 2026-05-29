package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero inserta un nuevo elemento al inicio de la lista.
	InsertarPrimero(T)

	// InsertarUltimo inserta un nuevo elemento al final de la lista.
	InsertarUltimo(T)

	// BorrarPrimero elimina el primer elemento y devuelve el primer elemento de la lista.
	BorrarPrimero() T

	// VerPrimero devuelve el primer elemento de la lista.
	// Si la lista se encontraba vacía, entra en pánico con el mensaje "La lista esta vacia".
	VerPrimero() T

	// VerUltimo devuelve el último elemento de la lista.
	// Si la lista se encontraba vacía, entra en pánico con el mensaje "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos de la lista.
	Largo() int

	// Iterar recorre todos los elementos de la lista y le aplica a cada uno a la función visitar.
	Iterar(visitar func(T) bool)

	// Iterador devuelve un iterador de la lista.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual devuelve el elemento actual del iterador.
	// Si el iterador no tiene un elemento actual,
	// entra en pánico con el mensaje "El iterador termino de iterar".
	VerActual() T

	// HayAlgoMas devuelve verdadero si actualmente el iterador tiene un elemento, false en caso contrario.
	HayAlgoMas() bool

	// Avanzar mueve el iterador al siguiente elemento.
	// Si el iterador no tiene un elemento actual,
	// entra en pánico con el mensaje "El iterador termino de iterar".
	Avanzar()

	// Insertar inserta un nuevo elemento antes del actual.
	// El nuevo elemento se convierte en el actual.
	Insertar(T)

	// Borrar elimina el elemento actual y devuelve su valor.
	// El siguiente elemento se convierte en el actual.
	// Si el iterador no tiene un elemento actual,
	// entra en pánico con el mensaje "El iterador termino de iterar".
	Borrar() T
}
