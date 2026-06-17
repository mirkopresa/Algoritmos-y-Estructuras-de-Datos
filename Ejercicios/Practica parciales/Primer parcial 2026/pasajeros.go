// La empresa de cruceros MB SuperStar ha construido el crucero más grande jamás imaginado, el Cinatit. El mismo es
// tan grande que puede albergar a una cantidad de pasajeros mas grande de lo que la gente de marketing puede contar a
// lo largo de varias categorías. Desde la misma, nos piden que ordenemos su lista de pasajeros para darles un orden para
// desembarcar. Debemos hacerlo lo más rápido posible ya que están por llegar al puerto. Para ordenarlos, nos dieron las
// siguientes estipulaciones:
// Primero se organizarán los pasajeros según su categoría de mayor a menor importancia (BLUE/YELLOW, BLACK, TURISTA).
// Para la categoría premium (BLUE/YELLOW), se ordenarán según la fecha en la que hayan empezado su suscripción, dada
// por el campo InicioSubscripion. La fecha tiene el formato AAAA-MM-DD.
// Para la categoría medium (BLACK) se ordenaran también según su fecha de suscripción pero los más nuevos primero.
// Para la categoría TURISTA, se ordenarán según que tan apurados estén por llegar a su casa (indicado por un coeficiente
// numérico, con el 0 indicando a los menos apurados). Mientras más apurados estén, más atrás en el orden estarán.

// Indicar y justificar la elección del algoritmo por sobre otras alternativas, así como su complejidad. Desarrollar completa-
// mente cada complejidad, no se admitirán resultados parciales.

package main

import (
	"strconv"
)

// Asumir que se tiene el siguiente struct:

type Categoria int

const (
	BLUEYELLOW Categoria = iota
	BLACK
	TURISTA
)

type Pasajero struct {
	Nombre             string
	InicioSubscripcion string
	CoefApurado        int8
	Categoria          Categoria
}

func extraerCoef(pasajero Pasajero) int {
	return int(pasajero.CoefApurado)
}

func extraerAnio(pasajero Pasajero) int {
	anio, _ := strconv.Atoi(pasajero.InicioSubscripcion[:4])
	return anio - 1
}

func extraerMes(pasajero Pasajero) int {
	mes, _ := strconv.Atoi(pasajero.InicioSubscripcion[5:7])
	return mes - 1
}

func extraerDia(pasajero Pasajero) int {
	dia, _ := strconv.Atoi(pasajero.InicioSubscripcion[8:])
	return dia - 1
}

func CountingSort(pasajeros []Pasajero, f func(a Pasajero) int, k int) []Pasajero {
	frecuencias := make([]int, k)
	for _, p := range pasajeros {
		frecuencias[f(p)]++
	}
	sumasAcumuladas := make([]int, k)
	for i := 1; i < k; i++ {
		sumasAcumuladas[i] = sumasAcumuladas[i-1] + frecuencias[i-1]
	}
	res := make([]Pasajero, len(pasajeros))
	for _, p := range pasajeros {
		res[sumasAcumuladas[f(p)]] = p
		sumasAcumuladas[f(p)]++
	}
	return res
}

func OrdenarPasajeros(pasajeros []Pasajero) []Pasajero {
	var turistas []Pasajero
	var blacks []Pasajero
	var blueYellows []Pasajero
	for _, pasajero := range pasajeros {
		if pasajero.Categoria == TURISTA {
			turistas = append(turistas, pasajero)
		} else if pasajero.Categoria == BLACK {
			blacks = append(blacks, pasajero)
		} else {
			blueYellows = append(blueYellows, pasajero)
		}
	}
	// Ordenamos turistas por coef ( de 0 a 127 )
	turistas = CountingSort(turistas, extraerCoef, 128)
	// Ordenamos los blacks por fecha, primero dia, luego mes, luego año
	blacks = CountingSort(blacks, extraerDia, 31)
	blacks = CountingSort(blacks, extraerMes, 12)
	blacks = CountingSort(blacks, extraerAnio, 2026)
	// Lo mismo para blue yellow
	blueYellows = CountingSort(blueYellows, extraerDia, 31)
	blueYellows = CountingSort(blueYellows, extraerMes, 12)
	blueYellows = CountingSort(blueYellows, extraerAnio, 2026)

	resultado := make([]Pasajero, 0, len(pasajeros))
	resultado = append(resultado, blueYellows...)
	// Ahora guardamos de menor a mayor, porque habiamos ordenado de menor a mayor (de mas viejo a mas nuevo)
	// y el enunciado nos pide lo contrario
	for i := len(blacks) - 1; i >= 0; i-- {
		resultado = append(resultado, blacks[i])
	}
	// Ahora guardamos de menor a mayor, porque habiamos ordenado de menor a mayor (de menos apurado a mas apurado)
	// y el enunciado nos pide lo contrario
	for i := len(turistas) - 1; i >= 0; i-- {
		resultado = append(resultado, turistas[i])
	}
	return resultado
}
