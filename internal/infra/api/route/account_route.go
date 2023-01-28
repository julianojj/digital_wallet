package route

import (
	"github.com/gin-gonic/gin"
	"github.com/julianojj/digital_wallet/internal/infra/api/controller"
)

type AccountRoute struct {
	r                          *gin.Engine
	getAccountController       *controller.GetAccountController
	transferAccountsController *controller.TransferAccountsController
}

func NewAccountRoute(
	r *gin.Engine,
	getAccountController *controller.GetAccountController,
	transferAccountsController *controller.TransferAccountsController,
) *AccountRoute {
	return &AccountRoute{
		r,
		getAccountController,
		transferAccountsController,
	}
}

func (a *AccountRoute) Register() {
	a.r.GET("/accounts/:id", a.getAccountController.Handle)
	a.r.POST("/accounts/transfer", a.transferAccountsController.Handle)
}
