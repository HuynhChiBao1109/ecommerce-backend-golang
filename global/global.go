package global

import (
	"ecommerce-backend-golang/pkg/logger"
	"ecommerce-backend-golang/pkg/setting"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
	Rdb    *redis.Client
)
