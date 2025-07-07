package com.jacobsen.dentista.usuarios.dtos;

import java.util.UUID;

public record ResponseUsuarioDTO(UUID id, String name, String email) {
}
