package com.jacobsen.dentista.pacientes.entities;

import java.time.LocalDate;

import com.jacobsen.dentista.pacientes.dtos.ResponsePacienteDTO;

import jakarta.persistence.*;

import jakarta.annotation.Nullable;
import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotNull;
import jakarta.validation.constraints.Pattern;

import org.hibernate.annotations.CreationTimestamp;
import org.hibernate.validator.constraints.br.CPF;

import lombok.Getter;
import lombok.Setter;

@Entity(name = "pacientes")
@Table(name = "pacientes")
@Getter
@Setter
public class Paciente {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @NotNull(message = "Nome é obrigatório")
    private String name;

    @CPF(message = "Cpf inválido")
    @NotNull(message = "Cpf é obrigatório")
    @Column(unique = true)
    private String cpf;

    @Pattern(regexp = "\\d{10,11}", message = "Telefone deve ter entre 10 e 11 dígitos")
    @NotNull(message = "Número de telefone é obrigatório")
    private Long cellphone;

    @Email(message = "Email inválido")
    private String email;

    @NotNull(message = "Data de nascimento inválida")
    private LocalDate dateOfBirth;

    @CreationTimestamp
    @Column(updatable = false)
    private LocalDate createdAt;

    public Paciente() {
    }

    public Paciente(String name, String cpf, Long cellphone, @Nullable String email, LocalDate dateOfBirth) {
        this.name = name;
        this.cpf = cpf;
        this.cellphone = cellphone;
        this.email = email;
        this.dateOfBirth = dateOfBirth;
    }

    public ResponsePacienteDTO toResponseDTO() {
        return new ResponsePacienteDTO(id, name, cpf, cellphone, email, dateOfBirth, createdAt);
    }
}
