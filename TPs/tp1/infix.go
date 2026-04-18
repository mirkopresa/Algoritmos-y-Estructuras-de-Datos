package main

import (
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

const (
	_SUMA           = "+"
	_RESTA          = "-"
	_MULTIPLICACION = "*"
	_DIVISION       = "/"
	_POTENCIA       = "^"
	_PARENTESIS_IZQ = '('
	_PARENTESIS_DER = ')'
)

type Caracter struct {
	tipo  TipoCaracter
	valor string
}

// Funciones auxiliares para identificar el tipo de caracter
func esNumero(caracter rune) bool {
	return unicode.IsDigit(caracter)
}

func esEspacio(caracter rune) bool {
	return unicode.IsSpace(caracter)
}

func esOperador(caracter rune) bool {
	return precedencia(string(caracter)) > 0
}

func esParentesisIzq(caracter rune) bool {
	return caracter == _PARENTESIS_IZQ
}

func esParentesisDer(caracter rune) bool {
	return caracter == _PARENTESIS_DER
}

// Funcion que devuelve la precedencia de un operador
func precedencia(operador string) int {
	switch operador {
	case _SUMA, _RESTA:
		return 1
	case _MULTIPLICACION, _DIVISION:
		return 2
	case _POTENCIA:
		return 3
	default:
		return 0
	}
}

// Funcion que obtiene un numero completo en caso de que haya mas de un digito
func obtenerNumero(inicio int, linea string) (string, int) {
	numero := ""
	fin := inicio
	for fin+1 < len(linea) && esNumero(rune(linea[fin+1])) {
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
		case esEspacio(caracter):
			continue
		case esNumero(caracter):
			var numero string
			numero, i = obtenerNumero(i, linea)
			caracterAGuardar = Caracter{tipo: Numero, valor: numero}
		case esOperador(caracter):
			caracterAGuardar = Caracter{tipo: Operador, valor: string(linea[i])}
		case esParentesisIzq(caracter):
			caracterAGuardar = Caracter{tipo: ParentesisIzq, valor: string(linea[i])}
		case esParentesisDer(caracter):
			caracterAGuardar = Caracter{tipo: ParentesisDer, valor: string(linea[i])}
		default:
			panic("Caracter no reconocido")
		}
		vectorInfix = append(vectorInfix, caracterAGuardar)
	}
	return vectorInfix
}

// Funcion que devuelve la asociatividad de un operador, true para derecha y false para izquierda
func esAsociativoPorDerecha(operador string) bool {
	if operador == _POTENCIA {
		return true
	}
	return false
}

func desapilarYGuardarOperadoresMayorOIgual(vectorPostfix []Caracter, operador Caracter, pila TDAPila.Pila[Caracter]) []Caracter {
	for !pila.EstaVacia() && pila.VerTope().tipo == Operador {
		operadorTopePrecedencia := precedencia(pila.VerTope().valor)
		operadorNuevoPrecedencia := precedencia(operador.valor)
		if operadorTopePrecedencia < operadorNuevoPrecedencia {
			break
		}
		if operadorTopePrecedencia == operadorNuevoPrecedencia && esAsociativoPorDerecha(operador.valor) {
			break
		}
		vectorPostfix = append(vectorPostfix, pila.Desapilar())
	}
	return vectorPostfix
}

func desapilarYGuardarHastaParentesisIzq(vectorPostfix []Caracter, pila TDAPila.Pila[Caracter]) []Caracter {
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
			vectorPostfix = desapilarYGuardarHastaParentesisIzq(vectorPostfix, pila)
		case Operador:
			vectorPostfix = desapilarYGuardarOperadoresMayorOIgual(vectorPostfix, caracter, pila)
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
