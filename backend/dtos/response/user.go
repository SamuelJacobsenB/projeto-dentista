package response

import (
	"time"

	"github.com/SamuelJacobsenB/projeto-dentista/backend/types"
)

type UserDTO struct {
	ID        uint         `json:"id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	Roles     []types.Role `json:"roles"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
