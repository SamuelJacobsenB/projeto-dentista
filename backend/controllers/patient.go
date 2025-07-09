package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/SamuelJacobsenB/projeto-dentista/backend/dtos/request"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/dtos/response"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/services"
	"github.com/gin-gonic/gin"
)

type PatientController struct {
	service *services.PatientService
}

func NewPatientController(service *services.PatientService) *PatientController {
	return &PatientController{service}
}

func (controller *PatientController) FindPagenedByName(ctx *gin.Context) {
	name := ctx.Query("name")
	strLimit := ctx.DefaultQuery("limit", "20")
	strOffset := ctx.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(strLimit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Limite inválido"})
		return
	}

	offset, err := strconv.Atoi(strOffset)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Página inválida"})
		return
	}

	patients, err := controller.service.FindPagenedByName(name, limit, offset)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar pacientes"})
		return
	}

	patientsResponse := make([]response.PatientDTO, len(patients))

	for i, patient := range patients {
		patientsResponse[i] = *patient.ToResponseDTO()
	}

	ctx.JSON(http.StatusOK, patientsResponse)
}

func (controller *PatientController) FindByID(ctx *gin.Context) {
	strId := ctx.Param("id")
	intId, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	patient, err := controller.service.FindByID(uint(intId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao buscar paciente"})
		return
	}

	responsePatient := patient.ToResponseDTO()

	ctx.JSON(http.StatusOK, responsePatient)
}

func (controller *PatientController) Create(ctx *gin.Context) {
	var patientDTO request.PatientDTO

	if err := ctx.ShouldBindJSON(&patientDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	if err := patientDTO.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patient := patientDTO.ToEntity()

	if err := controller.service.Create(patient); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar paciente"})
		return
	}

	responsePatient := patient.ToResponseDTO()

	ctx.JSON(http.StatusCreated, responsePatient)
}

func (controller *PatientController) Update(ctx *gin.Context) {
	strId := ctx.Param("id")
	intId, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var patientDTO request.PatientDTO

	if err := ctx.ShouldBindJSON(&patientDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	if err := patientDTO.ValidateUpdateDTO(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patient := patientDTO.ToEntity()

	if err := controller.service.Update(patient, uint(intId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar paciente"})
		return
	}

	responsePatient := patient.ToResponseDTO()

	ctx.JSON(http.StatusOK, responsePatient)
}

func (controller *PatientController) UploadImage(ctx *gin.Context) {
	strId := ctx.Param("id")
	id, err := strconv.Atoi(strId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	profileImage, err := ctx.FormFile("profile_image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao obter arquivo de imagem"})
		return
	}

	if err := controller.service.UploadImage(profileImage, uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao fazer upload da imagem"})
		return
	}

	dsn := fmt.Sprintf("uploads/patients/%s%s", strId, filepath.Ext(profileImage.Filename))
	if err := ctx.SaveUploadedFile(profileImage, dsn); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar arquivo de imagem"})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (controller *PatientController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := controller.service.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar paciente"})
		return
	}

	ctx.Status(http.StatusNoContent)
}
