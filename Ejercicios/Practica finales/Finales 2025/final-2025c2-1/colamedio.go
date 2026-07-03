package main

type Cola[T any] interface {
	Encolar(T)
	Desencolar() T
	EstaVacia() bool
	EliminarMedio() T
	VerPrimero() T
}
