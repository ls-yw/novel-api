package request

type Pages struct {
	Page int `form:"page" json:"page"`
	Size int `form:"size" json:"size"`
}

type Id struct {
	Id uint `form:"id" verify:"required"`
}
