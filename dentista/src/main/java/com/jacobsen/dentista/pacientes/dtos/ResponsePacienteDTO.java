package com.jacobsen.dentista.pacientes.dtos;

import java.time.LocalDate;

public record ResponsePacienteDTO(Long id, String name, String cpf, Long cellphone, String email, LocalDate dateOfBirth) {}
