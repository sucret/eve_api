package validator

import (
	"eve/internal/global"
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 自定义验证器注册
		_ = v.RegisterValidation("mobile", validateMobile)
		_ = v.RegisterValidation("exist", validateExist)
	}
}

// 校验手机号
func validateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	// 手机号为空则不验证，这里只验证有值的
	if mobile == "" {
		return true
	}

	ok, _ := regexp.MatchString(`^1[3-9]\d{9}$`, mobile)

	return ok
}

// 验证是否在数据库表中存在
func validateExist(fl validator.FieldLevel) bool {
	value := fmt.Sprint(fl.Field())
	// 如果值为空则不验证
	if value == "" {
		return true
	}

	param := strings.Split(fl.Param(), " ")

	sql := fmt.Sprintf(`SELECT 1 FROM %s WHERE %s = '%s' LIMIT 1`, param[0], param[1], value)

	var count int64
	if global.DB.Raw(sql).Count(&count); count > 0 {
		return true
	}

	return false
}
