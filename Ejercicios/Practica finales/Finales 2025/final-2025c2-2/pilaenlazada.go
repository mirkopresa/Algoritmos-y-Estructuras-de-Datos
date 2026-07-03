package main

type Pila[T any] interface {
	EstaVacia() bool
	VerTope() T
	Desapilar() T
	Apilar(T)
}
