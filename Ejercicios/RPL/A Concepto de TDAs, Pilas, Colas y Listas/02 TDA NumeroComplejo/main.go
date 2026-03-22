package main

type Complejo struct {
	real       float64
	imaginario float64
}

// CrearComplejo devuelve un puntero a la estructura que representa el numero complejo
// dado por la parte real e imaginaria pasadas por parámetro.
func CrearComplejo(real float64, img float64) *Complejo {
	return &Complejo{real: real, imaginario: img}
}

// Multiplicar dados dos números complejos, modifica el primero (c1) dejando el resultado de hacer
// la multiplicacion entre ambos números. Recordar que:
// c1 * c2 = (a + i b) * (c + i d) = (a * c - b * d) + i (b * c + a * d)
func (comp *Complejo) Multiplicar(otro Complejo) {
	parteReal := (comp.real * otro.real) - (comp.imaginario * otro.imaginario)
	parteImaginaria := (comp.imaginario * otro.real) + (comp.real * otro.imaginario)
	comp.real, comp.imaginario = parteReal, parteImaginaria
}

// Dados dos números complejos, modifica el primero (c1) dejando el resultado de hacer
// la suma entre ambos números. Recordar que:
// c1 + c2 = (a + i b) + (c + i d) = (a + c) + i (b + d)
func (comp *Complejo) Sumar(otro Complejo) {
	parteReal := comp.real + otro.real
	parteImaginaria := comp.imaginario + otro.imaginario
	comp.real, comp.imaginario = parteReal, parteImaginaria
}

// ParteReal Obtiene la parte real de un numero complejo
func (comp Complejo) ParteReal() float64 {
	// ?
}

// ParteImaginaria Obtiene la parte imaginaria de un numero complejo
func (comp Complejo) ParteImaginaria() float64 {
	// ?
}

// Modulo Obtiene el modulo de un numero complejo
func (comp Complejo) Modulo() float64 {
	// ?
}

// Angulo Obtiene el angulo de un numero complejo
func (comp Complejo) Angulo() float64 {
	// ?
}
