package mensajes

const (
	PACIENTE_ENCOLADO        = "Paciente %s encolado\n"                            // listo
	CANT_PACIENTES_ENCOLADOS = "%d paciente(s) en espera para %s\n"                // listo
	ENOENT_PACIENTE          = "ERROR: no existe el paciente '%s'\n"               // listo
	ENOENT_ESPECIALIDAD      = "ERROR: no existe la especialidad '%s'\n"           // listo
	ENOENT_URGENCIA          = "ERROR: grado de urgencia no identificado ('%s')\n" // listo
	ENOENT_RANGO             = "ERROR: rango alfabetico invalido ('%s,%s')\n"      // listo

	PACIENTE_ATENDIDO = "Se atiende a %s\n"                 // listo
	SIN_PACIENTES     = "No hay pacientes en espera\n"      // listo
	ENOENT_DOCTOR     = "ERROR: no existe el doctor '%s'\n" // listo

	DOCTORES_SISTEMA = "%d doctor(es) en el sistema\n"                         // listo
	INFORME_DOCTOR   = "%d: %s, especialidad %s, %d paciente(s) atendido(s)\n" // listo

	ENOENT_CANT_PARAMS = "No se recibieron los 2 (dos) parametros: <archivo doctores> y <archivo pacientes>\n" // listo
	ENOENT_ARCHIVO     = "No se pudo leer archivo %s\n"                                                        // listo
	ENOENT_ANIO        = "Valor no numerico en campo de anio: %s\n"                                            // listo                                            // listo
	ENOENT_FORMATO     = "ERROR: formato de comando incorrecto ('%s')\n"                                       // listo
	ENOENT_CMD         = "ERROR: no existe el comando '%s'\n"                                                  // listo
	ENOENT_PARAM_RANGO = "ERROR: Valor no alfabetico en campo de rango alfabetico '%s'\n"                      // listo

	ENOENT_PARAMS = "ERROR: cantidad de parametros invalidos para comando '%s'\n" // listo
)
