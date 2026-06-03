// En nuestro juego de rol táctico, Final Fantasy Algorithms, el orden de ataque se decide por el atributo iniciativa: los de mayor
// iniciativa atacan primero. En esta batalla, un hechizo global está a punto de activarse, por lo que solo quedan T turnos en los que se
// pueda realizar un ataque. Se desea saber, de los N personajes, cuáles lograrán atacar antes de que se active el hechizo Final Fantástico
// Algorítmico. Se tiene una lista con los personajes que participarán en el combate como una estructura de formato
// (nombre string, iniciativa int):
// [ ('Ma-Go Lang', 95), ('Bárbara', 75), ('Cléri-Go Lang', 60), ('Arquera de bugs', 90) ]
// Se pide realizar una función determinarOrdenDeAtaque que reciba la lista de combatientes, y la cantidad T turnos de turnos restantes.
// La función debe devolver una lista con los nombres de los personajes que logran actuar en esa ventana de tiempo, ordenados por turno
// en el que actúan. Indicar y justificar la complejidad del algoritmo implementado, expresada con las variables N y T del problema.

package main

type personaje struct {
	nombre     string
	iniciativa int
}

func determinarOrdenDeAtaque(lista Lista[personaje], t int) Lista[personaje] {
	arr := make([]personaje, 0)
	for iter := lista.Iterador(); iter.HayAlgoMas(); iter.Avanzar() { // O(N) siendo N la cantidad de personajes
		personaje := lista.VerActual()
		arr = append(arr, personaje)
	}
	heap := CrearHeapArr(arr, func(a, b personaje) int { return a.iniciativa - b.iniciativa }) // O(N) por el heapify
	resultado := CrearListaEnlazada[personaje]()
	for t > 0 && !heap.EstaVacia() { // O(T*log n) si t < n, O(n log n) si t == n
		resultado.InsertarUltimo(heap.Desencolar())
		t--
	}
	return resultado
}
