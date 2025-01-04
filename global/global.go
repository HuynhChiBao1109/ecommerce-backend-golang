package global

import (
	"ecommerce-backend-golang/pkg/logger"
	"ecommerce-backend-golang/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
)
