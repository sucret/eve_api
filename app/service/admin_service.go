package service

import "eve/app/model"

type adminService struct {
	baseService
}

func (a *adminService) Profile() (model model.Admin) {
	a.db.First(&model)
	return model
}
