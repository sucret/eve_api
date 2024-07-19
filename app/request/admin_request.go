package request

type SaveAdmin struct {
	ID       uint   `form:"id" json:"id" binding:"exist=admin id"`
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Age      int    `form:"age" json:"age" binding:"required,min=1,max=100"`
	Gender   int    `form:"gender" json:"gender" binding:"required,oneof=1 2"`
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
}

func (SaveAdmin) GetMessages() ValidateErrorMessages {
	return ValidateErrorMessages{
		"ID.exist":          "记录不存在",
		"Name.required":     "用户名称不能为空",
		"Password.required": "密码不能为空",
		"Age.required":      "年龄不能为空",
		"Age.min":           "年龄最小为1",
		"Age.max":           "年龄最大为100",
		"Gender.required":   "性别不能为空",
		"Gender.oneof":      "性别只能是男或者女",
		"Mobile.required":   "手机号不能为空",
		"Mobile.mobile":     "手机号格式不正确",
	}
}
