package clinica

import (
	TDACola "tdas/cola"
	TDAHeap "tdas/cola_prioridad"
)

const (
	URGENTE = "URGENTE"
	REGULAR = "REGULAR"
)

type paciente struct {
	nombre      string
	inscripcion int
}

type ColaEspecialidad interface {

	// Cantidad devuelve la cantidad de pacientes en espera de ambas colas.
	Cantidad() int

	// Encolar agrega un nuevo paciente a la cola si recibe "URGENTE", o al heap si recibe "REGULAR" como tercer parametro.
	Encolar(nombre string, inscripcion int, urgencia string)

	// Desencolar saca el primer elemento de la cola/heap a un paciente dependiendo de si hay URGENTES encolados o no, y devuelve su nombre.
	Desencolar() string

	// EstaVacia devuelve verdadero si ambas colas no tienen elementos encolados, false en caso contrario.
	EstaVacia() bool
}

type colaEspecialidad struct {
	cola     TDACola.Cola[paciente]
	heap     TDAHeap.ColaPrioridad[paciente]
	cantidad int
}

func CrearColaEspecialidad() ColaEspecialidad {
	return &colaEspecialidad{cola: TDACola.CrearColaEnlazada[paciente](), heap: TDAHeap.CrearHeap(cmp), cantidad: 0}
}

func (c *colaEspecialidad) Cantidad() int {
	return c.cantidad
}

func (c *colaEspecialidad) Encolar(nombre string, inscripcion int, urgencia string) {
	paciente := paciente{nombre: nombre, inscripcion: inscripcion}
	if urgencia == URGENTE {
		c.cola.Encolar(paciente)
	}
	if urgencia == REGULAR {
		c.heap.Encolar(paciente)
	}
	c.cantidad++
}

func (c *colaEspecialidad) Desencolar() string {
	var atendido paciente
	if !c.cola.EstaVacia() {
		atendido = c.cola.Desencolar()
	} else {
		atendido = c.heap.Desencolar()
	}
	c.cantidad--
	return atendido.nombre
}

func (c *colaEspecialidad) EstaVacia() bool {
	return c.cantidad == 0
}

// Funciones auxiliares

func cmp(paciente1, paciente2 paciente) int {
	return paciente2.inscripcion - paciente1.inscripcion
}
