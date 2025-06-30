package com.jacobsen.dentista.pacientes;

import com.jacobsen.dentista.pacientes.dtos.RequestPacienteDTO;
import com.jacobsen.dentista.pacientes.dtos.ResponsePacienteDTO;
import com.jacobsen.dentista.pacientes.entities.Paciente;
import com.jacobsen.dentista.pacientes.exceptions.PacienteNotFoundException;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class PacienteService {
    private final PacienteRepository repository;

    public PacienteService(PacienteRepository repository) {
        this.repository = repository;
    }

    public List<ResponsePacienteDTO> getAll() {
        return repository.findAll().stream().map(Paciente::toResponseDTO).toList();
    }

    public ResponsePacienteDTO getOne(Long id) {
        Paciente paciente = repository.findById(id)
                .orElseThrow(() -> new PacienteNotFoundException("Paciente não foi encontrado"));

        return paciente.toResponseDTO();
    }

    public void save(RequestPacienteDTO dto) {
        Paciente paciente = dto.toEntity();
        repository.save(paciente);
    }

    public void deleteOne(Long id) {
        repository.deleteById(id);
    }
}
