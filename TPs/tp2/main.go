package main

import (
	"bufio"
	"fmt"
	"os"
	"tp2/clinica"
	"tp2/mensajes"
)

func imprimirError(err error) {
	fmt.Print(err.Error())
}

func main() {
	if len(os.Args) != 3 {
		fmt.Print(mensajes.ENOENT_CANT_PARAMS)
		return
	}
	rutaDoctores := os.Args[1]
	rutaPacientes := os.Args[2]
	pacientes, err := clinica.CargarPacientes(rutaPacientes)
	if err != nil {
		imprimirError(err)
		return
	}
	doctores, especialidades, err := clinica.CargarDoctores(rutaDoctores)
	if err != nil {
		imprimirError(err)
		return
	}
	// El programa se inicio correctamente
	lector := bufio.NewScanner(os.Stdin)
	for lector.Scan() {
		lineaComando := lector.Text()
		comando, parametros, err := clinica.VerificarFormato(lineaComando, pacientes, doctores, especialidades)
		if err != nil {
			imprimirError(err)
			continue
		}
		// formato correcto, chequear cual es y ejecutar la funcion
		if comando == clinica.PEDIRTURNO {
			clinica.PedirTurno(parametros, pacientes, especialidades)
		}
		if comando == clinica.ATENDERSIGUIENTE {
			clinica.AtenderSiguiente(parametros, doctores, especialidades)
		}
		if comando == clinica.INFORME {
			clinica.Informe(parametros, doctores)
		}
	}
}
