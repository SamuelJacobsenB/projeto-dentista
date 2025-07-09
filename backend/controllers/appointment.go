package controllers

import (
	"net/http"
	"strconv"

	"github.com/SamuelJacobsenB/projeto-dentista/backend/dtos/request"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/services"
	"github.com/gin-gonic/gin"
)

type AppointmentController struct {
	service *services.AppointmentService
}

func NewAppointmentController(service *services.AppointmentService) *AppointmentController {
	return &AppointmentController{service}
}

func (controller *AppointmentController) FindAll(ctx *gin.Context) {
	appointments, err := controller.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao buscar agendamentos"})
		return
	}

	ctx.JSON(http.StatusOK, appointments)
}

func (controller *AppointmentController) FindOfToday(ctx *gin.Context) {
	appointments, err := controller.service.FindOfToday()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao buscar agendamentos de hoje"})
		return
	}

	ctx.JSON(http.StatusOK, appointments)
}

func (controller *AppointmentController) FindByID(ctx *gin.Context) {
	strId := ctx.Param("id")
	intId, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	appointment, err := controller.service.FindByID(uint(intId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "agendamento não encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, appointment)
}

func (controller *AppointmentController) Create(ctx *gin.Context) {
	var appointmentDTO request.AppointmentDTO

	if err := ctx.ShouldBindJSON(&appointmentDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	if err := appointmentDTO.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appointment := appointmentDTO.ToEntity()
	if err := controller.service.Create(appointment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao criar agendamento"})
		return
	}

	ctx.JSON(http.StatusCreated, appointment)
}

func (controller *AppointmentController) Update(ctx *gin.Context) {
	strId := ctx.Param("id")
	intId, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	var appointmentDTO request.AppointmentDTO
	if err := ctx.ShouldBindJSON(&appointmentDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "dados inválidos"})
		return
	}

	if err := appointmentDTO.ValidateUpdateDTO(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	appointment := appointmentDTO.ToEntity()
	if err := controller.service.Update(appointment, uint(intId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao atualizar agendamento"})
		return
	}

	ctx.JSON(http.StatusOK, appointment)
}

func (controller *AppointmentController) Delete(ctx *gin.Context) {
	strId := ctx.Param("id")
	intId, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}

	if err := controller.service.Delete(uint(intId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao excluir agendamento"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
