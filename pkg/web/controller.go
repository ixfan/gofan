package web

import (
	"net/http"
)

type Controller struct {
}

func (c *Controller) Response(context *Context, data interface{}, err error) {
	if err != nil {
		serviceError := err.(*ServiceError)
		context.Set("error", err)
		responseData := &Response{
			Code:    serviceError.Code,
			Message: serviceError.Error(),
		}
		context.Set("response", responseData)
		context.JSON(http.StatusOK, responseData)
	} else {
		responseData := &Response{
			Code:    http.StatusOK,
			Message: "操作成功",
			Data:    data,
		}
		context.Set("response", responseData)
		context.JSON(http.StatusOK, responseData)
	}
}
