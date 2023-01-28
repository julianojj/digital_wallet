package unit

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/julianojj/digital_wallet/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNotShouldCreateANewTransactionIfTransactionIDIsRequired(t *testing.T) {
	currentTime := time.Now()
	transaction, err := entity.NewTransaction("", uuid.NewString(), "Credit", 1000.00, currentTime, currentTime)
	assert.EqualError(t, err, "transaction id is required")
	assert.Nil(t, transaction)
}

func TestNotShouldCreateANewTransactionIfAccountIDIsRequired(t *testing.T) {
	currentTime := time.Now()
	transaction, err := entity.NewTransaction(uuid.NewString(), "", "Credit", 1000.00, currentTime, currentTime)
	assert.EqualError(t, err, "account id is required")
	assert.Nil(t, transaction)
}

func TestNotShouldCreateANewTransactionIfTransactionTypeIsRequired(t *testing.T) {
	currentTime := time.Now()
	transaction, err := entity.NewTransaction(uuid.NewString(), uuid.NewString(), "", 1000.00, currentTime, currentTime)
	assert.EqualError(t, err, "transaction type is required")
	assert.Nil(t, transaction)
}

func TestNotShouldCreateANewTransactionIfAmountIsRequired(t *testing.T) {
	currentTime := time.Now()
	transaction, err := entity.NewTransaction(uuid.NewString(), uuid.NewString(), "Credit", 0.00, currentTime, currentTime)
	assert.EqualError(t, err, "amount must be greater than zero")
	assert.Nil(t, transaction)
}

func TestShouldCreateANewTransaction(t *testing.T) {
	currentTime := time.Now()
	transaction, err := entity.NewTransaction(uuid.NewString(), uuid.NewString(), "Credit", 1000.00, currentTime, currentTime)
	assert.NoError(t, err)
	assert.Equal(t, "Credit", transaction.Type)
	assert.Equal(t, 1000.00, transaction.Amount)
}
