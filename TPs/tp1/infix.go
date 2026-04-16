package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	TDAPila "tdas/pila"
	"unicode"
)

type TipoCaracter int

const (
	Numero TipoCaracter = iota
	Operador
	ParentesisIzq
	ParentesisDer
)

type Caracter struct {
	tipo  TipoCaracter
	valor string
}

// Funciones auxiliares para identificar el tipo de caracter
func EsNumero(caracter rune) bool {
	return unicode.IsDigit(caracter)
}

func EsEspacio(caracter rune) bool {
	return unicode.IsSpace(caracter)
}

func EsOperador(caracter rune) bool {
	return Precedencia(string(caracter)) > 0
}

func EsParentesisIzq(caracter rune) bool {
	return caracter == '('
}

func EsParentesisDer(caracter rune) bool {
	return caracter == ')'
}

// Funcion que devuelve la precedencia de un operador
func Precedencia(operador string) int {
	switch operador {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	case "^":
		return 3
	default:
		return 0
	}
}

// Funcion que obtiene un numero completo en caso de que haya mas de un digito
func ObtenerNumero(inicio int, linea string) (string, int) {
	numero := ""
	fin := inicio
	for fin+1 < len(linea) && EsNumero(rune(linea[fin+1])) {
		fin++
	}
	numero = linea[inicio : fin+1]
	return numero, fin
}

// Funcion que convierte un string en un vector de Caracteres
func LineaAVector(linea string) []Caracter {
	vectorInfix := make([]Caracter, 0)
	for i := 0; i < len(linea); i++ {
		var caracterAGuardar Caracter
		caracter := rune(linea[i])
		switch {
		case EsEspacio(caracter):
			continue
		case EsNumero(caracter):
			var numero string
			numero, i = ObtenerNumero(i, linea)
			caracterAGuardar = Caracter{tipo: Numero, valor: numero}
		case EsOperador(caracter):
			caracterAGuardar = Caracter{tipo: Operador, valor: string(linea[i])}
		case EsParentesisIzq(caracter):
			caracterAGuardar = Caracter{tipo: ParentesisIzq, valor: string(linea[i])}
		case EsParentesisDer(caracter):
			caracterAGuardar = Caracter{tipo: ParentesisDer, valor: string(linea[i])}
		default:
			panic("Caracter no reconocido")
		}
		vectorInfix = append(vectorInfix, caracterAGuardar)
	}
	return vectorInfix
}

// Funcion que devuelve la asociatividad de un operador, true para derecha y false para izquierda
func EsAsociativoPorDerecha(operador string) bool {
	if operador == "^" {
		return true
	}
	return false
}

func DesapilarYGuardarOperadoresMayorOIgual(vectorPostfix []Caracter, operador Caracter, pila TDAPila.Pila[Caracter]) []Caracter {
	for !pila.EstaVacia() && pila.VerTope().tipo == Operador {
		operadorTopePrecedencia := Precedencia(pila.VerTope().valor)
		operadorNuevoPrecedencia := Precedencia(operador.valor)
		if operadorTopePrecedencia < operadorNuevoPrecedencia {
			break
		}
		if operadorTopePrecedencia == operadorNuevoPrecedencia && EsAsociativoPorDerecha(operador.valor) {
			break
		}
		vectorPostfix = append(vectorPostfix, pila.Desapilar())
	}
	return vectorPostfix
}

func DesapilarYGuardarHastaParentesisIzq(vectorPostfix []Caracter, pila TDAPila.Pila[Caracter]) []Caracter {
	for !pila.EstaVacia() && pila.VerTope().tipo != ParentesisIzq {
		vectorPostfix = append(vectorPostfix, pila.Desapilar())
	}
	pila.Desapilar()
	return vectorPostfix
}

func ConvertirInfixToPostfix(vectorInfix []Caracter) []Caracter {
	pila := TDAPila.CrearPilaDinamica[Caracter]()
	vectorPostfix := make([]Caracter, 0)
	for _, caracter := range vectorInfix {
		switch caracter.tipo {
		case Numero:
			vectorPostfix = append(vectorPostfix, caracter)
		case ParentesisIzq:
			pila.Apilar(caracter)
		case ParentesisDer:
			vectorPostfix = DesapilarYGuardarHastaParentesisIzq(vectorPostfix, pila)
		case Operador:
			vectorPostfix = DesapilarYGuardarOperadoresMayorOIgual(vectorPostfix, caracter, pila)
			pila.Apilar(caracter)
		default:
			panic("Tipo de caracter no reconocido")
		}
	}
	for !pila.EstaVacia() {
		vectorPostfix = append(vectorPostfix, pila.Desapilar())
	}
	return vectorPostfix
}

func ImprimirResultado(vectorPostfix []Caracter) {
	valores := make([]string, 0)
	for _, caracter := range vectorPostfix {
		valores = append(valores, caracter.valor)
	}
	lineaPostfix := strings.Join(valores, " ")
	fmt.Fprintf(os.Stdout, "%s\n", lineaPostfix)
}

func main() {
	lector := bufio.NewScanner(os.Stdin)
	for lector.Scan() {
		linea := lector.Text()
		vectorInfix := LineaAVector(linea)
		vectorPostfix := ConvertirInfixToPostfix(vectorInfix)
		ImprimirResultado(vectorPostfix)
	}
}
