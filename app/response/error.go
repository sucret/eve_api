package response

const (
	SystemError = iota + 400000
	BusinessError
	HaveNoPermission
	ValidateError
)

var ErrorMap = map[int]string{
	SystemError:      "系统错误",
	BusinessError:    "业务错误",
	HaveNoPermission: "没有操作权限",
	ValidateError:    "数据校验失败",
}
