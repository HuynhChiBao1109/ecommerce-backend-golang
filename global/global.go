package global

import (
	"ecommerce-backend-golang/pkg/logger"
	"ecommerce-backend-golang/pkg/setting"

	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
