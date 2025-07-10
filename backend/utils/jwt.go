package utils

import (
	"time"

	"github.com/SamuelJacobsenB/projeto-dentista/backend/config"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/types"
	"github.com/golang-jwt/jwt"
)

const jwtByteSecret = config.JWT_SECRET

func GenerateJwt(id uint, roles []types.Role) (string, error) {
	claims := jwt.MapClaims{
		"user_id": id,
		"roles":   roles,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtByteSecret)
}
