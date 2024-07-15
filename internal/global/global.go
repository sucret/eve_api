package global

import (
	"eve/internal/config"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	Db     *gorm.DB
)
