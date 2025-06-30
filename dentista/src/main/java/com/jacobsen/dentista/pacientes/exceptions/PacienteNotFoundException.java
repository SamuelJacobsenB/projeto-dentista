package com.jacobsen.dentista.pacientes.exceptions;

public class PacienteNotFoundException extends RuntimeException {
    public PacienteNotFoundException() {}

    public PacienteNotFoundException(String message) {
        super(message);
    }
}
