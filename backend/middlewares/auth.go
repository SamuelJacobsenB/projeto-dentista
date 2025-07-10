package middlewares

import (
	"net/http"

	"github.com/SamuelJacobsenB/projeto-dentista/backend/config"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/types"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const jwtByteSecret = config.JWT_SECRET

func AuthMiddleware(roles []types.Role) gin.HandlerFunc {
	if len(roles) == 0 {
		roles = []types.Role{types.RoleUser}
	}

	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "usuário deve estar logado"})
			return
		}

		strToken := authHeader[len("Bearer "):]
		if strToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}

		token, err := jwt.Parse(strToken, func(token *jwt.Token) (interface{}, error) {
			return jwtByteSecret, nil
		})
		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}

		rawRoles, ok := claims["roles"].([]interface{})
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}

		userRoles := make([]string, len(rawRoles))
		for i, role := range rawRoles {
			userRoles[i] = role.(string)
		}

		hasAllRoles := utils.HasAllRoles(roles, userRoles)

		if !hasAllRoles {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "usuário não tem permissão para acessar este recurso"})
			return
		}

		if userID, ok := claims["user_id"].(float64); ok {
			ctx.Set("user_id", uint(userID))
		}

		ctx.Set("user_roles", userRoles)

		ctx.Next()
	}
}
