package com.jacobsen.dentista.pacientes;

import com.jacobsen.dentista.pacientes.dtos.RequestPacienteDTO;
import com.jacobsen.dentista.pacientes.dtos.ResponsePacienteDTO;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("pacientes")
public class PacienteController {
    private final PacienteService service;

    public PacienteController(PacienteService service) {
        this.service = service;
    }

    @GetMapping
    public ResponseEntity<List<ResponsePacienteDTO>> getAll() {
        List<ResponsePacienteDTO> pacientes = service.getAll();
        return ResponseEntity.ok(pacientes);
    }

    @GetMapping("/{id}")
    public ResponseEntity<ResponsePacienteDTO> getOne(@PathVariable Long id) {
        ResponsePacienteDTO paciente = service.getOne(id);
        return ResponseEntity.ok(paciente);
    }

    @PostMapping
    public ResponseEntity<Void> save(@RequestBody RequestPacienteDTO dto) {
        service.save(dto);
        return ResponseEntity.status(HttpStatus.CREATED).build();
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteOne(@PathVariable Long id) {
        service.deleteOne(id);
        return ResponseEntity.status(HttpStatus.NO_CONTENT).build();
    }
}