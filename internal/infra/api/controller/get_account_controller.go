package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianojj/digital_wallet/internal/application/usecases"
)

type GetAccountController struct {
	getAccount *usecases.GetAccount
}

func NewGetAccountController(getAccount *usecases.GetAccount) *GetAccountController {
	return &GetAccountController{
		getAccount,
	}
}

func (g *GetAccountController) Handle(c *gin.Context) {
	id := c.Param("id")
	account, err := g.getAccount.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, account)
}
