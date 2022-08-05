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
		context.JSON(http.StatusOK, &Response{
			Code:    serviceError.Code,
			Message: serviceError.Error(),
		})
	} else {
		context.JSON(http.StatusOK, &Response{
			Code:    http.StatusOK,
			Message: "success.",
			Data:    data,
		})
	}
}
