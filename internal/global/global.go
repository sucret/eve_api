package global

import (
	"eve/internal/config"
	"eve/internal/event"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Config          *config.Config
	DB              *gorm.DB
	Logger          *zap.Logger
	EventDispatcher *event.Dispatcher
)
