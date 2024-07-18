package service

import "gorm.io/gorm"

// service基类，其他service都继承它
type baseService struct {
	db *gorm.DB
}
