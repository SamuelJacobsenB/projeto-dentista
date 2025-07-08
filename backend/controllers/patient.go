package controllers

import (
	"net/http"
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

func (c *PatientController) FindPagenedByName(ctx *gin.Context) {
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

	patients, err := c.service.FindPagenedByName(name, limit, offset)

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

func (c *PatientController) FindByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	patient, err := c.service.FindByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar paciente"})
		return
	}

	if patient == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Paciente não encontrado"})
		return
	}

	responsePatient := patient.ToResponseDTO()

	ctx.JSON(http.StatusOK, responsePatient)
}

func (c *PatientController) Create(ctx *gin.Context) {
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

	if err := c.service.Create(patient); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar paciente"})
		return
	}

	responsePatient := patient.ToResponseDTO()

	ctx.JSON(http.StatusCreated, responsePatient)
}

func (c *PatientController) Update(ctx *gin.Context) {
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

	if err := c.service.Update(patient); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar paciente"})
		return
	}

	responsePatient := patient.ToResponseDTO()

	ctx.JSON(http.StatusOK, responsePatient)
}

func (c *PatientController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.service.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar paciente"})
		return
	}

	ctx.Status(http.StatusNoContent)
}
