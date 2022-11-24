package controller

import (
	"vinbigdata/internal/delivery/http/model"

	"github.com/gin-gonic/gin"
)

type MobileService interface {
	SaveUserCall(userName string, duration int64) error
	GetUserBilling(userName string) (*model.BillingData, error)
}

type TelecomController struct {
	Controller
	mobilService MobileService
}

func NewTelecomController(service MobileService) *TelecomController {
	return &TelecomController{
		mobilService: service,
	}
}

// SaveCall godoc
// @Summary track user call duration
// @Tags mobile
// @ID user-call
// @Accept json
// @Produce json
// @param user_name path string true "username"
// @Param json body model.CallRequest true "call information"
// @Success 200
// @Failure 400 {object} model.ErrorResponse
// @Router /mobile/{user_name}/call [post]
func (c *TelecomController) SaveCall(ctx *gin.Context) {
	var request model.CallRequest
	if c.WrapBindAndValidate(ctx, &request) {
		if err := c.mobilService.SaveUserCall(ctx.Param("user_name"), request.Duration); err != nil {
			c.Failure(ctx, err)
			return
		}
		c.Success(ctx, nil)
	}
}

// GetBill godoc
// @Summary get user mobile bill
// @Tags mobile
// @ID user-bill
// @Accept json
// @Produce json
// @param user_name path string true "username"
// @Success 200 {object} model.BillingData
// @Failure 400 {object} model.ErrorResponse
// @Router /mobile/{user_name}/billing [get]
func (c *TelecomController) GetBill(ctx *gin.Context) {
	data, err := c.mobilService.GetUserBilling(ctx.Param("user_name"))
	if err != nil {
		c.Failure(ctx, err)
		return
	}

	c.Success(ctx, data)
}
