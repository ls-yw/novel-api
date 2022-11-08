package errors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Data map[string]interface{}

func (e *Error) Error() *Error {
	return e
}

func (e *Error) ReturnJson(c *gin.Context, data ...map[string]interface{}) {
	returnJson := make(Data)
	if len(data) > 0 {
		returnJson = data[0]
	}
	returnJson["Code"] = e.Code
	returnJson["Message"] = e.Message

	c.JSON(http.StatusOK, returnJson)
	panic(nil)
}
