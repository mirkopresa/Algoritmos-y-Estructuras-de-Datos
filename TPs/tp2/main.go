package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tp2/clinica"
	"tp2/mensajes"
)

const (
	PEDIRTURNO       = "PEDIR_TURNO"
	ATENDERSIGUIENTE = "ATENDER_SIGUIENTE"
	INFORME          = "INFORME"
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
	clinica, err := clinica.CrearClinica(rutaPacientes, rutaDoctores)
	if err != nil {
		imprimirError(err)
		return
	}
	// El programa se inicio correctamente
	lector := bufio.NewScanner(os.Stdin)
	for lector.Scan() {
		lineaComando := strings.Split(lector.Text(), ":")
		if len(lineaComando) != 2 {
			fmt.Printf(mensajes.ENOENT_FORMATO, lector.Text())
			continue
		}
		comando := lineaComando[0]
		parametros := strings.Split(lineaComando[1], ",")
		switch comando {
		case PEDIRTURNO:
			if len(parametros) != 3 {
				imprimirError(fmt.Errorf(mensajes.ENOENT_PARAMS, PEDIRTURNO))
				continue
			}
			nombrePaciente, cantidad, nombreEspecialidad, err := clinica.PedirTurno(parametros)
			if err != nil {
				imprimirError(err)
				continue
			}
			fmt.Printf(mensajes.PACIENTE_ENCOLADO, nombrePaciente)
			fmt.Printf(mensajes.CANT_PACIENTES_ENCOLADOS, cantidad, nombreEspecialidad)
		case ATENDERSIGUIENTE:
			if len(parametros) != 1 {
				imprimirError(fmt.Errorf(mensajes.ENOENT_PARAMS, ATENDERSIGUIENTE))
				continue
			}
			pacienteAtendido, cantidad, especialidadDoctor, err := clinica.AtenderSiguiente(parametros)
			if err != nil {
				imprimirError(err)
				continue
			}
			fmt.Printf(mensajes.PACIENTE_ATENDIDO, pacienteAtendido)
			fmt.Printf(mensajes.CANT_PACIENTES_ENCOLADOS, cantidad, especialidadDoctor)
		case INFORME:
			if len(parametros) != 2 {
				imprimirError(fmt.Errorf(mensajes.ENOENT_PARAMS, INFORME))
				continue
			}
			contador, informe, err := clinica.Informe(parametros)
			if err != nil {
				imprimirError(err)
				continue
			}
			fmt.Printf(mensajes.DOCTORES_SISTEMA, contador)
			fmt.Print(informe)
		default:
			fmt.Printf(mensajes.ENOENT_CMD, comando)
		}
	}
}
