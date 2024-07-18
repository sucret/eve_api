package request

type SaveAdmin struct {
	ID       uint   `form:"id" json:"id"`
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Age      int    `form:"age" json:"age" binding:"required,min=1,max=100"`
	Gender   int    `form:"gender" json:"gender" binding:"required,oneof=1 2"`
}
