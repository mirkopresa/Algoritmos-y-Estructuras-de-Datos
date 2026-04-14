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

func EsNumero(caracter rune) bool {
	return unicode.IsDigit(caracter)
}

func EsEspacio(caracter rune) bool {
	return unicode.IsSpace(caracter)
}

func EsOperadorOParentesis(caracter rune) TipoCaracter {
	switch caracter {
	case '(':
		return ParentesisIzq
	case ')':
		return ParentesisDer
	default:
		return Operador
	}
}

// Funcion que devuelve la precedencia de un operador
func Precedencia(operador string) int {
	switch operador {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 3 // Para el operador ^
	}
}

func CargarNumero(numero *strings.Builder, vectorInfix []Caracter) []Caracter {
	if numero.Len() > 0 {
		numeroAGuardar := Caracter{tipo: Numero, valor: numero.String()}
		vectorInfix = append(vectorInfix, numeroAGuardar)
		numero.Reset()
	}
	return vectorInfix
}

// Funcion que convierte una linea de texto en un array de strings,
// donde cada string es un numero o un operador
func TextoAVector(linea string) []Caracter {
	vectorInfix := make([]Caracter, 0)
	numero := strings.Builder{}
	for _, caracter := range linea {
		if EsEspacio(caracter) {
			continue
		} else if EsNumero(caracter) {
			numero.WriteRune(caracter)
		} else { // Es un operador o un parentesis
			vectorInfix = CargarNumero(&numero, vectorInfix)
			caracterTipo := EsOperadorOParentesis(caracter)
			caracterAGuardar := Caracter{tipo: caracterTipo, valor: string(caracter)}
			vectorInfix = append(vectorInfix, caracterAGuardar)
		}
	}
	vectorInfix = CargarNumero(&numero, vectorInfix)
	return vectorInfix
}

// Funcion que devuelve la asociatividad de un operador, true para izquierda y false para derecha
func EsAsociativoPorIzquierda(operador string) bool {
	if operador == "^" {
		return false
	}
	return true
}

func ReacomodarOperadores(vectorPostfix []Caracter, operador Caracter, pila TDAPila.Pila[Caracter]) []Caracter {
	for !pila.EstaVacia() && pila.VerTope().tipo == Operador {
		operadorTopePrecedencia := Precedencia(pila.VerTope().valor)
		operadorNuevoPrecedencia := Precedencia(operador.valor)
		if operadorTopePrecedencia > operadorNuevoPrecedencia || (operadorTopePrecedencia == operadorNuevoPrecedencia && EsAsociativoPorIzquierda(operador.valor)) {
			vectorPostfix = append(vectorPostfix, pila.Desapilar())
		} else {
			break
		}
	}
	pila.Apilar(operador)
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
		default: // Es un operador
			vectorPostfix = ReacomodarOperadores(vectorPostfix, caracter, pila)
		}
	}
	for !pila.EstaVacia() {
		vectorPostfix = append(vectorPostfix, pila.Desapilar())
	}
	return vectorPostfix
}

// Funcion que imprime el vectorPostfix por Stdout
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
		vectorInfix := TextoAVector(linea)
		vectorPostfix := ConvertirInfixToPostfix(vectorInfix)
		ImprimirResultado(vectorPostfix)
	}
}
