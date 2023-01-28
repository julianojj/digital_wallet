package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/julianojj/digital_wallet/internal/application/usecases"
)

type TransferAccountsController struct {
	transferAccounts *usecases.TransferAccounts
}

func NewTransferAccountsController(
	transferAccounts *usecases.TransferAccounts,
) *TransferAccountsController {
	return &TransferAccountsController{
		transferAccounts,
	}
}

func (t *TransferAccountsController) Handle(c *gin.Context) {
	var input usecases.TransferAccountsInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = t.transferAccounts.Execute(input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(204)
}
