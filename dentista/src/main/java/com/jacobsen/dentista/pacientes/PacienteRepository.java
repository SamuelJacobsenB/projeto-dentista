package com.jacobsen.dentista.pacientes;

import com.jacobsen.dentista.pacientes.entities.Paciente;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface PacienteRepository extends JpaRepository<Paciente, Long> {
}
