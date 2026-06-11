package clinica

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDACEsp "tdas/cola_especialidad"
	TDADict "tdas/diccionario"
	"tp2/mensajes"
	"unicode"
)

const (
	PEDIRTURNO       = "PEDIR_TURNO"
	ATENDERSIGUIENTE = "ATENDER_SIGUIENTE"
	INFORME          = "INFORME"
	URGENTE          = "URGENTE"
	REGULAR          = "REGULAR"
)

type doctor struct {
	especialidad   string
	cant_pacientes int
}

func CargarPacientes(ruta_csv_pacientes string) (TDADict.Diccionario[string, int], error) {
	dictPacientes := TDADict.CrearHash[string, int]()
	err := procesarCSVPacientes(ruta_csv_pacientes, dictPacientes)
	return dictPacientes, err
}

func CargarDoctores(ruta_csv_doctores string) (TDADict.DiccionarioOrdenado[string, *doctor], TDADict.Diccionario[string, TDACEsp.ColaEspecialidad], error) {
	abbDoctores := TDADict.CrearABB[string, *doctor](strings.Compare)
	dictEspecialidades := TDADict.CrearHash[string, TDACEsp.ColaEspecialidad]()
	err := procesarCSVDoctores(ruta_csv_doctores, abbDoctores, dictEspecialidades)
	return abbDoctores, dictEspecialidades, err
}

func VerificarFormato(linea string, pacientes TDADict.Diccionario[string, int], doctores TDADict.DiccionarioOrdenado[string, *doctor], especialidades TDADict.Diccionario[string, TDACEsp.ColaEspecialidad]) (string, []string, error) {
	comandoPartes := strings.Split(linea, ":")
	if len(comandoPartes) != 2 {
		return "", nil, fmt.Errorf(mensajes.ENOENT_FORMATO, linea)
	}
	comando := comandoPartes[0]
	parametros := strings.Split(comandoPartes[1], ",")
	switch comando {
	case PEDIRTURNO:
		return PEDIRTURNO, parametros, verificarPedirTurno(parametros, pacientes, especialidades)
	case ATENDERSIGUIENTE:
		return ATENDERSIGUIENTE, parametros, verificarAtenderSiguiente(parametros, doctores)
	case INFORME:
		return INFORME, parametros, verificarInforme(parametros)
	default:
		return "", nil, fmt.Errorf(mensajes.ENOENT_CMD, comando)
	}
}
func PedirTurno(parametros []string, pacientes TDADict.Diccionario[string, int], especialidades TDADict.Diccionario[string, TDACEsp.ColaEspecialidad]) {
	nombrePaciente := parametros[0]
	nombreEspecialidad := parametros[1]
	tipoUrgencia := parametros[2]
	especialidad := especialidades.Obtener(nombreEspecialidad)
	especialidad.Encolar(nombrePaciente, pacientes.Obtener(nombrePaciente), tipoUrgencia)
	fmt.Printf(mensajes.PACIENTE_ENCOLADO, nombrePaciente)
	fmt.Printf(mensajes.CANT_PACIENTES_ENCOLADOS, especialidad.Cantidad(), nombreEspecialidad)
}

func AtenderSiguiente(parametros []string, doctores TDADict.DiccionarioOrdenado[string, *doctor], especialidades TDADict.Diccionario[string, TDACEsp.ColaEspecialidad]) {
	nombreDoctor := parametros[0]
	infoDoc := doctores.Obtener(nombreDoctor)
	especialidad := especialidades.Obtener(infoDoc.especialidad)
	if especialidad.EstaVacia() {
		fmt.Print(mensajes.SIN_PACIENTES)
		return
	}
	pacienteAtendido := especialidad.Desencolar()
	infoDoc.cant_pacientes++
	fmt.Printf(mensajes.PACIENTE_ATENDIDO, pacienteAtendido)
	fmt.Printf(mensajes.CANT_PACIENTES_ENCOLADOS, especialidad.Cantidad(), infoDoc.especialidad)
}

func Informe(parametros []string, doctores TDADict.DiccionarioOrdenado[string, *doctor]) {
	informeDoctores := make([]string, 0)
	contador := 0
	inicio := parametros[0]
	fin := parametros[1]
	if fin == "" {
		fin = string(unicode.MaxRune)
	}
	doctoresIter := doctores.IteradorRango(&inicio, &fin)
	for doctoresIter.HayAlgoMas() {
		nombreDoctor, infoDoc := doctoresIter.VerActual()
		contador++
		cadenaDoc := fmt.Sprintf(mensajes.INFORME_DOCTOR, contador, nombreDoctor, infoDoc.especialidad, infoDoc.cant_pacientes)
		informeDoctores = append(informeDoctores, cadenaDoc)
		doctoresIter.Avanzar()
	}
	fmt.Printf(mensajes.DOCTORES_SISTEMA, contador)
	fmt.Print(strings.Join(informeDoctores, ""))
}

// Funciones auxiliares

func procesarCSVPacientes(ruta_csv string, pacientes TDADict.Diccionario[string, int]) error {
	csv, err_apertura := os.Open(ruta_csv)
	if err_apertura != nil {
		return fmt.Errorf(mensajes.ENOENT_ARCHIVO, ruta_csv)
	}
	defer csv.Close()
	lector := bufio.NewScanner(csv)
	for lector.Scan() {
		cadena := lector.Text()
		datos := strings.Split(cadena, ",")
		anio, err_lectura := strconv.Atoi(datos[1])
		if err_lectura != nil {
			return fmt.Errorf(mensajes.ENOENT_ANIO, datos[1])
		}
		pacientes.Guardar(datos[0], anio)
	}
	return nil
}

func procesarCSVDoctores(ruta_csv string, doctores TDADict.DiccionarioOrdenado[string, *doctor], especialidades TDADict.Diccionario[string, TDACEsp.ColaEspecialidad]) error {
	csv, err_apertura := os.Open(ruta_csv)
	if err_apertura != nil {
		return fmt.Errorf(mensajes.ENOENT_ARCHIVO, ruta_csv)
	}
	defer csv.Close()
	lector := bufio.NewScanner(csv)
	for lector.Scan() {
		cadena := lector.Text()
		datos := strings.Split(cadena, ",")
		nombreDoctor := datos[0]
		especialidad := datos[1]
		doctores.Guardar(nombreDoctor, &doctor{especialidad: especialidad, cant_pacientes: 0})
		if !especialidades.Pertenece(especialidad) {
			especialidades.Guardar(especialidad, TDACEsp.CrearColaEspecialidad())
		}
	}
	return nil
}

func verificarPedirTurno(parametros []string, pacientes TDADict.Diccionario[string, int], especialidades TDADict.Diccionario[string, TDACEsp.ColaEspecialidad]) error {
	if len(parametros) != 3 {
		return fmt.Errorf(mensajes.ENOENT_PARAMS, PEDIRTURNO)
	}
	errores := make([]string, 0)
	nombrePaciente := parametros[0]
	nombreEspecialidad := parametros[1]
	tipoUrgencia := parametros[2]
	if !pacientes.Pertenece(nombrePaciente) {
		cadenaError := fmt.Sprintf(mensajes.ENOENT_PACIENTE, nombrePaciente)
		errores = append(errores, cadenaError)
	}
	if !especialidades.Pertenece(nombreEspecialidad) {
		cadenaError := fmt.Sprintf(mensajes.ENOENT_ESPECIALIDAD, nombreEspecialidad)
		errores = append(errores, cadenaError)
	}
	if tipoUrgencia != REGULAR && tipoUrgencia != URGENTE {
		cadenaError := fmt.Sprintf(mensajes.ENOENT_URGENCIA, tipoUrgencia)
		errores = append(errores, cadenaError)
	}
	if len(errores) == 0 {
		return nil
	} else {
		mensajeError := strings.Join(errores, "")
		return errors.New(mensajeError)
	}
}

func verificarAtenderSiguiente(parametros []string, doctores TDADict.DiccionarioOrdenado[string, *doctor]) error {
	if len(parametros) != 1 {
		return fmt.Errorf(mensajes.ENOENT_PARAMS, ATENDERSIGUIENTE)
	}
	doctor := parametros[0]
	if !doctores.Pertenece(doctor) {
		return fmt.Errorf(mensajes.ENOENT_DOCTOR, doctor)
	}
	return nil
}

func esString(cadena string) bool {
	for _, letra := range cadena {
		if !unicode.IsLetter(letra) && letra != ' ' {
			return false
		}
	}
	return true
}

func verificarInforme(parametros []string) error {
	if len(parametros) != 2 {
		return fmt.Errorf(mensajes.ENOENT_PARAMS, INFORME)
	}
	inicio := parametros[0]
	fin := parametros[1]
	if inicio != "" && fin != "" && inicio > fin {
		return fmt.Errorf(mensajes.ENOENT_RANGO, inicio, fin)
	}
	return nil
}
