package bootstrap

import (
	"fmt"
	"log"
	"time"

	"github.com/SamuelJacobsenB/projeto-dentista/backend/repositories"
	"github.com/SamuelJacobsenB/projeto-dentista/backend/services"
	cron "github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func InitScheduler(db *gorm.DB) {
	appointmentRepo := repositories.NewAppointmentRepository(db)
	appointmentService := services.NewAppointmentService(appointmentRepo)

	local, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		panic("Erro ao carregar localização")
	}

	scheduler := cron.New(cron.WithLocation(local))

	_, err = scheduler.AddFunc("0 7 * * *", func() {
		fmt.Println("Initialing scheduler")

		if err := appointmentService.SendReminderEmail(); err != nil {
			log.Println(err.Error())
		}

		if err := appointmentService.DeleteExpired(); err != nil {
			log.Println(err.Error())
		}
	})

	if err != nil {
		log.Println(err.Error())
	}

	scheduler.Start()
}
