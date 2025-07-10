package controllers

import (
	"net/http"

	"github.com/SamuelJacobsenB/projeto-dentista/backend/dtos/request"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/services"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController(service *services.AuthService) *AuthController {
	return &AuthController{service}
}

func (controller *AuthController) Login(ctx *gin.Context) {
	var loginDTO request.LoginDTO

	if err := ctx.ShouldBindJSON(&loginDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	if err := loginDTO.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := controller.service.Login(loginDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"access_token": token})
}

func (controller *AuthController) VerifyUser(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (controller *AuthController) VerifyAdmin(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}
