package clinica

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDADict "tdas/diccionario"
	"tp2/mensajes"
	"unicode"
)

type Clinica interface {
	PedirTurno(parametros []string) (string, int, string, error)
	AtenderSiguiente(parametros []string) (string, int, string, error)
	Informe(parametros []string) (int, string, error)
}

type clinica struct {
	pacientes      TDADict.Diccionario[string, int]
	doctores       TDADict.DiccionarioOrdenado[string, Doctor]
	especialidades TDADict.Diccionario[string, ColaEspecialidad]
}

func CrearClinica(ruta_csv_pacientes, ruta_csv_doctores string) (Clinica, error) {
	dictPacientes := TDADict.CrearHash[string, int]()
	abbDoctores := TDADict.CrearABB[string, Doctor](strings.Compare)
	dictEspecialidades := TDADict.CrearHash[string, ColaEspecialidad]()
	clinica := &clinica{
		pacientes:      dictPacientes,
		doctores:       abbDoctores,
		especialidades: dictEspecialidades,
	}
	errPacientes := clinica.procesarCSVPacientes(ruta_csv_pacientes)
	if errPacientes != nil {
		return nil, errPacientes
	}
	errDoctores := clinica.procesarCSVDoctores(ruta_csv_doctores)
	if errDoctores != nil {
		return nil, errDoctores
	}
	return clinica, nil
}

func (clinica *clinica) PedirTurno(parametros []string) (string, int, string, error) {
	err := clinica.verificarPedirTurno(parametros)
	if err != nil {
		return "", 0, "", err
	}
	nombrePaciente := parametros[0]
	nombreEspecialidad := parametros[1]
	tipoUrgencia := parametros[2]
	especialidad := clinica.especialidades.Obtener(nombreEspecialidad)
	especialidad.Encolar(nombrePaciente, clinica.pacientes.Obtener(nombrePaciente), tipoUrgencia)
	return nombrePaciente, especialidad.Cantidad(), nombreEspecialidad, nil
}

func (clinica *clinica) AtenderSiguiente(parametros []string) (string, int, string, error) {
	err := clinica.verificarAtenderSiguiente(parametros)
	if err != nil {
		return "", 0, "", err
	}
	nombreDoctor := parametros[0]
	infoDoc := clinica.doctores.Obtener(nombreDoctor)
	especialidad := clinica.especialidades.Obtener(infoDoc.Especialidad())
	if especialidad.EstaVacia() {
		return "", 0, "", fmt.Errorf(mensajes.SIN_PACIENTES)
	}
	pacienteAtendido := especialidad.Desencolar()
	infoDoc.AtenderPaciente()
	return pacienteAtendido, especialidad.Cantidad(), infoDoc.Especialidad(), nil
}

func (clinica *clinica) Informe(parametros []string) (int, string, error) {
	err := clinica.verificarInforme(parametros)
	if err != nil {
		return 0, "", err
	}
	informeDoctores := make([]string, 0)
	contador := 0
	inicio := parametros[0]
	fin := parametros[1]
	if fin == "" {
		fin = string(unicode.MaxRune)
	}
	doctoresIter := clinica.doctores.IteradorRango(&inicio, &fin)
	for doctoresIter.HayAlgoMas() {
		nombreDoctor, infoDoc := doctoresIter.VerActual()
		contador++
		cadenaDoc := fmt.Sprintf(mensajes.INFORME_DOCTOR, contador, nombreDoctor, infoDoc.Especialidad(), infoDoc.CantidadAtendidos())
		informeDoctores = append(informeDoctores, cadenaDoc)
		doctoresIter.Avanzar()
	}
	informeFinal := strings.Join(informeDoctores, "")
	return contador, informeFinal, nil
}

// Funciones auxiliares

func (clinica *clinica) procesarCSVPacientes(ruta_csv string) error {
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
		clinica.pacientes.Guardar(datos[0], anio)
	}
	return nil
}

func (clinica *clinica) procesarCSVDoctores(ruta_csv string) error {
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
		clinica.doctores.Guardar(nombreDoctor, &doctor{especialidad: especialidad, cant_pacientes: 0})
		if !clinica.especialidades.Pertenece(especialidad) {
			clinica.especialidades.Guardar(especialidad, CrearColaEspecialidad())
		}
	}
	return nil
}

func (clinica *clinica) verificarPedirTurno(parametros []string) error {
	errores := make([]string, 0)
	nombrePaciente := parametros[0]
	nombreEspecialidad := parametros[1]
	tipoUrgencia := parametros[2]
	if !clinica.pacientes.Pertenece(nombrePaciente) {
		cadenaError := fmt.Sprintf(mensajes.ENOENT_PACIENTE, nombrePaciente)
		errores = append(errores, cadenaError)
	}
	if !clinica.especialidades.Pertenece(nombreEspecialidad) {
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

func (clinica *clinica) verificarAtenderSiguiente(parametros []string) error {
	doctor := parametros[0]
	if !clinica.doctores.Pertenece(doctor) {
		return fmt.Errorf(mensajes.ENOENT_DOCTOR, doctor)
	}
	return nil
}

func (clinica *clinica) verificarInforme(parametros []string) error {
	inicio := parametros[0]
	fin := parametros[1]
	if inicio != "" && fin != "" && inicio > fin {
		return fmt.Errorf(mensajes.ENOENT_RANGO, inicio, fin)
	}
	return nil
}
