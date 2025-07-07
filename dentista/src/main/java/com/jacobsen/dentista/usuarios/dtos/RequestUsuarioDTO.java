package com.jacobsen.dentista.usuarios.dtos;

import com.jacobsen.dentista.usuarios.entities.Usuario;

public record RequestUsuarioDTO(String name, String email, String password) {
    public Usuario toEntity() {
        return new Usuario(name, email, password);
    }
}
