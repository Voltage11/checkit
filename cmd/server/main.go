package main

import (
	"checkit/internal/config"
	"checkit/pkg/logger"
	"log"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.New()
	if err != nil {
		log.Fatal("Ошибка загрузки кинфигурации: ", err)
	}

	// Логгер
	appLogger := logger.New(logger.AppLoggerLevelInfo)

	//Перед стартом сервера установим уровень логирования из конфигурации
	appLogger.SetLevel(getLogLevelFromStr(cfg.Logger.Level))

}

func getLogLevelFromStr(levelStr string) logger.AppLoggerLevel {
	switch levelStr {
	case "debug":
		return logger.AppLoggerLevelDebug
	case "info":
		return logger.AppLoggerLevelInfo
	case "warn":
		return logger.AppLoggerLevelWarn
	case "error":
		return logger.AppLoggerLevelError
	default:
		return logger.AppLoggerLevelInfo
	}
}
