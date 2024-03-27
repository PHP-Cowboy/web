package global

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"web/config"
)

var (
	DB           *gorm.DB
	ServerConfig = &config.ServerConfig{}
	Logger       = make(map[string]*zap.SugaredLogger, 0)
)
