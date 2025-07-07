package com.jacobsen.dentista.usuarios.entities;

import java.util.UUID;

import org.hibernate.validator.constraints.Length;

import com.jacobsen.dentista.usuarios.dtos.ResponseUsuarioDTO;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Table;
import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotNull;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.Id;

import lombok.Getter;
import lombok.Setter;

@Entity(name = "usuarios")
@Table(name = "usuarios")
@Getter
@Setter
public class Usuario {
    @Id
    @GeneratedValue
    private UUID id;

    @NotNull(message = "Nome é obrigatório")
    @Length(max = 80, message = "Nome deve ter no máximo 80 caracteres")
    private String name;

    @NotNull(message = "Email obrigatório")
    @Email(message = "Email inválido")
    @Column(unique = true)
    private String email;

    @NotNull(message = "Senha é obrigatória")
    @Length(min = 8, max = 15, message = "Senha deve ter entre 8 e 15 caracteres")
    private String password;

    public Usuario() {
    }

    public Usuario(String name, String email, String password) {
        this.name = name;
        this.email = email;
        this.password = password;
    }

    public ResponseUsuarioDTO toResponseDTO() {
        return new ResponseUsuarioDTO(id, name, email);
    }
}
