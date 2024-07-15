package initializers

// Подключение функций для ce mapper агрегатов
import (
	_ "SebStudy/domain/resume/ce_resume_mapper" // Если над можно убрать, туп закомментив
)

func InitializeCeMapperHandlers() {
	// log.Println("Handlers ce mapper has been connected")
}
