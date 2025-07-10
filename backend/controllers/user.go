package controllers

import (
	"net/http"
	"strconv"

	"github.com/SamuelJacobsenB/projeto-dentista/backend/dtos/request"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service}
}

func (controller *UserController) FindByID(ctx *gin.Context) {
	strId := ctx.Param("id")
	intId, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id do usuário inválido"})
		return
	}

	requestedID := uint(intId)

	userIDValue, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "usuário não autenticado"})
		return
	}
	userID := userIDValue.(uint)

	rolesValue, exists := ctx.Get("user_roles")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "cargos não encontrados"})
		return
	}
	userRoles := rolesValue.([]string)

	isAdmin := false
	for _, role := range userRoles {
		if role == "admin" {
			isAdmin = true
			break
		}
	}

	if !isAdmin && userID != requestedID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "acesso negado"})
		return
	}

	user, err := controller.service.FindByID(requestedID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	responseUser := user.ToResponseDTO()

	ctx.JSON(http.StatusOK, responseUser)
}

func (controller *UserController) Create(ctx *gin.Context) {
	var userDTO request.UserDTO

	if err := ctx.ShouldBindJSON(&userDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	if err := userDTO.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := userDTO.ToEntity()

	err := controller.service.Create(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"})
		return
	}

	responseUser := user.ToResponseDTO()

	ctx.JSON(http.StatusCreated, responseUser)
}

func (controller *UserController) Promote(ctx *gin.Context) {
	strId := ctx.Param("id")
	intId, err := strconv.Atoi(strId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id do usuário inválido"})
		return
	}

	id := uint(intId)

	err = controller.service.Promote(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao promover usuário"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Usuário promovido com sucesso"})
}

func (controller *UserController) Delete(ctx *gin.Context) {
	strId := ctx.Param("id")
	intId, err := strconv.Atoi(strId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Id do usuário inválido"})
		return
	}

	id := uint(intId)

	err = controller.service.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar usuário"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso"})
}
