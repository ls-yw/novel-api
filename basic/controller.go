package basic

type Controller struct {
}

//func (bc Controller) Json(c *gin.Context, params interface{}) interface{} {
//	fmt.Printf("%+v =\n", params)
//	_ = c.ShouldBindJSON(&params)
//	fmt.Printf("%+v\n", params)
//	if err := request2.Validator(params); err != nil {
//		resp := errors.ErrorCustom
//		resp.Message = err.Error()
//		resp.ReturnJson(c)
//	}
//	return params
//}
