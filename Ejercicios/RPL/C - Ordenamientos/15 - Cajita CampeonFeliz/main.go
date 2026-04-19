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

const (
	_EDAD_MAXIMA_NIÑO        int = 12
	_EDAD_MAXIMA_NO_NIÑO     int = 100
	_CANTIDAD_NACIONALIDADES int = 32
)

type Persona struct {
	edad         int
	nombre       string
	nacionalidad Nacionalidad
}

// Primero ordenamos por no niños, y luego por nacionalidad, por ultimo por niños
// Complejidad O(d * (n)) siendo d la cantidad de ordenamientos, n la cantidad de personas
// Ordenamos 3 veces, entonces la complejidad es O(n)
func CountingSort[T any](arr []T, rango int, f func(T) int) []T {
	frecuencias := make([]int, rango)
	for _, num := range arr {
		frecuencias[f(num)]++
	}

	sumasAcumuladas := make([]int, rango)
	for j := 1; j < len(frecuencias); j++ {
		sumasAcumuladas[j] = sumasAcumuladas[j-1] + frecuencias[j-1]
	}

	arrResultado := make([]T, len(arr))
	for _, num := range arr {
		posicion := sumasAcumuladas[f(num)]
		arrResultado[posicion] = num
		sumasAcumuladas[f(num)]++
	}
	return arrResultado
}

func OrdenarFila(personas []Persona) []Persona {
	// Le enviamos _EDAD_MAXIMA_NO_NIÑO-_EDAD_MAXIMA_NIÑO+1 porque el rango de edades para los no niños es de 14 a _EDAD_MAXIMA_NIÑO,
	// y el bucket reservado para los niños es el 0, ya que no nos importa su edad
	personas = CountingSort(personas, _EDAD_MAXIMA_NO_NIÑO-_EDAD_MAXIMA_NIÑO+1, func(p Persona) int {
		if p.edad <= 12 { // Si la edad es menor o igual a 12, guarda a todos en el bucket 0
			return 0
		}
		return p.edad - 12 // Para los no niños, guarda en el bucket correspondiente a su edad menos 12
	})
	// Le enviamos _CANTIDAD_NACIONALIDADES+1 porque el bucket reservado para los niños es el 0,
	// y luego tenemos un bucket para cada nacionalidad (32 buckets)
	personas = CountingSort(personas, _CANTIDAD_NACIONALIDADES+1, func(p Persona) int {
		if p.edad <= 12 { // Si es un niño, lo guardamos en el bucket 0, ya que no nos importa su nacionalidad
			return 0
		}
		return int(p.nacionalidad) + 1 // Para los no niños, guardamos en el bucket correspondiente a su nacionalidad (el bucket 0 reservado para los niños)
	})
	// Le enviamos _EDAD_MAXIMA_NIÑO+2 porque el rango de edades para los niños es de 0 a _EDAD_MAXIMA_NIÑO, (13 buckets)
	// y el bucket reservado para los no niños es el _EDAD_MAXIMA_NIÑO + 1,
	personas = CountingSort(personas, _EDAD_MAXIMA_NIÑO+2, func(p Persona) int {
		if p.edad <= 12 { // Si es un niño, lo guardamos en el bucket correspondiente a su edad
			return p.edad
		}
		return _EDAD_MAXIMA_NIÑO + 1 // Para los no niños, lo guardamos en el bucket reservado para ellos, ya que no nos importa su edad
	})
	return personas
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
