package global

import (
	"eve/internal/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	Db     *gorm.DB
	Logger *zap.Logger
)
