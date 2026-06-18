// Implementar un algoritmo de ordenamiento, que sea lineal,
// que permita definir el orden en una fila de personas para comprar
// una Cajita CampeónFeliz en un establecimiento de comida rápida.
// Los datos (structs) a ordenar cuentan con edad (número), nombre (string) y nacionalidad
// (enumerativo, de 32 valores posibles). Primero deben ir los niños
// (todos con edad menor o igual a 12), y estos deben ordenarse por edad
// (de menor a mayor), independientemente de la nacionalidad. Luego deben ir los
// "no niños", que primero deben estar ordenados por nacionalidad (segundo
// Francia, por ejemplo) y en caso de igualdad de nacionalidad, por edad,
// también de menor a mayor.

// Justificar la complejidad de la función implementada. El desarrollo de la
// complejidad debe estar completo para el problema en cuestión, no se aceptarán
// resultados parciales genéricos.

package main

import "fmt"

type Nacionalidad int

const (
	ARGENTINA Nacionalidad = iota
	FRANCIA
	CROACIA
	MARRUECOS
	PAISES_BAJOS
	INGLATERRA
	BRASIL
	PORTUGAL
	JAPON
	SENEGAL
	AUSTRALIA
	SUIZA
	ESPANA
	ESTADOS_UNIDOS
	POLONIA
	COREA_DEL_SUR
	ALEMANIA
	ECUADOR
	CAMERUN
	URUGUAY
	TUNEZ
	MEXICO
	BELGICA
	GHANA
	ARABIA_SAUDITA
	IRAN
	COSTA_RICA
	DINAMARCA
	SERBIA
	GALES
	CANADA
	CATAR
)

type Persona struct {
	edad         int
	nombre       string
	nacionalidad Nacionalidad
}

func extraerEdad(p Persona) int {
	return p.edad
}

func extraerNacionalidad(p Persona) int {
	return int(p.nacionalidad)
}

// Primero ordenamos por no niños, y luego por nacionalidad, por ultimo por niños
// Complejidad O(d * (n)) siendo d la cantidad de ordenamientos, n la cantidad de personas
// Ordenamos 3 veces, entonces la complejidad es O(n)
func CountingSort(arr []Persona, rango int, f func(Persona) int) []Persona {
	frecuencias := make([]int, rango)
	for _, num := range arr {
		frecuencias[f(num)]++
	}

	sumasAcumuladas := make([]int, rango)
	for j := 1; j < len(frecuencias); j++ {
		sumasAcumuladas[j] = sumasAcumuladas[j-1] + frecuencias[j-1]
	}

	arrResultado := make([]Persona, len(arr))
	for _, num := range arr {
		posicion := sumasAcumuladas[f(num)]
		arrResultado[posicion] = num
		sumasAcumuladas[f(num)]++
	}
	return arrResultado
}

func OrdenarFila(personas []Persona) []Persona {
	ninos := make([]Persona, 0)
	adultos := make([]Persona, 0)
	for _, p := range personas {
		if p.edad <= 12 {
			ninos = append(ninos, p)
		} else {
			adultos = append(adultos, p)
		}
	}
	adultos = CountingSort(adultos, 87, extraerEdad)
	adultos = CountingSort(adultos, 32, extraerNacionalidad)
	ninos = CountingSort(ninos, 13, extraerEdad)
	res := make([]Persona, 0)
	res = append(res, ninos...)
	res = append(res, adultos...)
	return res
}

func main() {
	lamine := Persona{nombre: "Lamine Yamal", edad: 10, nacionalidad: ESPANA}
	messi := Persona{nombre: "Lionel Messi", edad: 13, nacionalidad: ARGENTINA}
	mbappe := Persona{nombre: "Kylian Mbappé", edad: 20, nacionalidad: FRANCIA}
	endrick := Persona{nombre: "Endrick", edad: 8, nacionalidad: BRASIL}
	ronaldo := Persona{nombre: "Cristiano Ronaldo", edad: 8, nacionalidad: PORTUGAL}
	kane := Persona{nombre: "Harry Kane", edad: 12, nacionalidad: INGLATERRA}
	neymar := Persona{nombre: "Neymar Jr", edad: 12, nacionalidad: BRASIL}
	vinicius := Persona{nombre: "Vinicius Jr", edad: 10, nacionalidad: BRASIL}
	julian := Persona{nombre: "Julián Álvarez", edad: 8, nacionalidad: ARGENTINA}
	foden := Persona{nombre: "Phil Foden", edad: 13, nacionalidad: INGLATERRA}
	bellingham := Persona{nombre: "Jude Bellingham", edad: 8, nacionalidad: INGLATERRA}
	musiala := Persona{nombre: "Jamal Musiala", edad: 10, nacionalidad: ALEMANIA}
	wirtz := Persona{nombre: "Florian Wirtz", edad: 12, nacionalidad: ALEMANIA}
	gavi := Persona{nombre: "Gavi", edad: 13, nacionalidad: ESPANA}
	pedri := Persona{nombre: "Pedri", edad: 20, nacionalidad: ESPANA}
	dibumartinez := Persona{nombre: "Emiliano Martínez", edad: 20, nacionalidad: ARGENTINA}
	personas := []Persona{lamine, messi, mbappe, endrick, ronaldo, kane, neymar, vinicius, julian, foden, bellingham, musiala, wirtz, gavi, pedri, dibumartinez}
	personas = OrdenarFila(personas)
	fmt.Println(personas)
}
