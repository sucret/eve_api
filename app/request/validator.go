package request

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

// ValidateErrorMessages 错误消息类型
type ValidateErrorMessages map[string]string

// Validator 验证规则
type Validator interface {
	GetMessages() ValidateErrorMessages
}

// GetErrorMsg 获取错误信息的方法
func GetErrorMsg(request interface{}, err error) string {
	// 定义 ValidationErrors 变量，可能包含错误数据
	var validationErrors validator.ValidationErrors

	// 如果request是Validator类型，并且err是validator.ValidationErrors类型
	if _, isValidator := request.(Validator); isValidator && errors.As(err, &validationErrors) && len(validationErrors) > 0 {
		// 获取第一个错误信息
		errMsg := validationErrors[0]
		if message, exist := request.(Validator).GetMessages()[errMsg.Field()+"."+errMsg.Tag()]; exist {
			return message
		} else {
			// 如果没有自定义消息，返回错误成员本身的错误信息
			return errMsg.Error()
		}
	}

	// 其他类型的错误直接返回错误信息
	return err.Error()
}
