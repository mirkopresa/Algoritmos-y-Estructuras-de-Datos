package clinica

type Doctor interface {
	// Especialidad devuelve la especialidad del doctor.
	Especialidad() string

	// CantidadAtendidos devuelve el total de pacientes atendidos.
	CantidadAtendidos() int

	// AtenderPaciente aumenta en uno la cantidad de pacientes atendidos.
	AtenderPaciente()
}

type doctor struct {
	especialidad   string
	cant_pacientes int
}

func crearDoctor(especialidad string) Doctor {
	return &doctor{especialidad: especialidad, cant_pacientes: 0}
}

func (d *doctor) Especialidad() string {
	return d.especialidad
}

func (d *doctor) CantidadAtendidos() int {
	return d.cant_pacientes
}

func (d *doctor) AtenderPaciente() {
	d.cant_pacientes++
}
