package service

import (
	"eve/internal/global"
)

var (
	BaseService = &baseService{db: global.DB}

	AdminService = &adminService{baseService: *BaseService}
)
