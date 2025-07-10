package utils

import "github.com/SamuelJacobsenB/projeto-dentista/backend/types"

func HasAllRoles(requiredRoles []types.Role, userRoles []string) bool {
	roleMap := make(map[string]bool)

	for _, role := range userRoles {
		roleMap[role] = true
	}

	for _, requiredRole := range requiredRoles {
		if !roleMap[string(requiredRole)] {
			return false
		}
	}

	return true
}
