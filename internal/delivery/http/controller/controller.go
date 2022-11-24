package controller

import (
	"encoding/json"
	"net/http"
	"vinbigdata/internal/delivery/http/model"
	error2 "vinbigdata/package/error"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
}

func (h *Controller) Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
	ctx.Abort()
}

func (h *Controller) Failure(ctx *gin.Context, err error) {
	code := http.StatusBadRequest
	if e, isInstance := err.(error2.CError); isInstance {
		code = e.Code
	}

	response := model.ErrorResponse{
		Code:    int64(code),
		Message: err.Error(),
	}
	ctx.JSON(http.StatusBadRequest, response)
	ctx.Abort()
}

func (h *Controller) WrapBindAndValidate(ctx *gin.Context, request interface{}) bool {
	body, _ := ctx.GetRawData()

	err := json.Unmarshal(body, &request)
	if err != nil {
		h.Failure(ctx, err)
		return false
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		h.Failure(ctx, err)
		return false
	}

	return true
}
