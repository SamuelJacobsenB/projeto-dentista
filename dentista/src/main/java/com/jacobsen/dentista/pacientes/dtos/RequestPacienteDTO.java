package com.jacobsen.dentista.pacientes.dtos;

import com.jacobsen.dentista.pacientes.entities.Paciente;

import java.time.LocalDate;

public record RequestPacienteDTO(String name, String cpf, Long cellphone, String email, LocalDate dateOfBirth) {
    public Paciente toEntity() {
        return new Paciente(name, cpf, cellphone, email, dateOfBirth);
    }
}
