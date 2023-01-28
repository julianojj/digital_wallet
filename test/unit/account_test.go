package unit

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/julianojj/digital_wallet/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNotShouldCreateANewAccountIfAccountIdIsRequired(t *testing.T) {
	currentTime := time.Now()
	account, errAccount := entity.NewAccount("", "237", "0001", "0123456-7", currentTime, currentTime)
	assert.EqualError(t, errAccount, "account id is required")
	assert.Nil(t, account)
}

func TestNotShouldCreateANewAccountIfBankCodeIsRequired(t *testing.T) {
	currentTime := time.Now()
	account, errAccount := entity.NewAccount(uuid.NewString(), "", "0001", "0123456-7", currentTime, currentTime)
	assert.EqualError(t, errAccount, "bank code is required")
	assert.Nil(t, account)
}

func TestNotShouldCreateANewAccountIfBranchIsRequired(t *testing.T) {
	currentTime := time.Now()
	account, errAccount := entity.NewAccount(uuid.NewString(), "237", "", "0123456-7", currentTime, currentTime)
	assert.EqualError(t, errAccount, "branch is required")
	assert.Nil(t, account)
}

func TestNotShouldCreateANewAccountIfAccountNumberIsRequired(t *testing.T) {
	currentTime := time.Now()
	account, errAccount := entity.NewAccount(uuid.NewString(), "237", "0001", "", currentTime, currentTime)
	assert.EqualError(t, errAccount, "account number is required")
	assert.Nil(t, account)
}

func TestNotShouldCreateANewAccountAndCreditIfAmountMustBeGreaterThanZero(t *testing.T) {
	currentTime := time.Now()
	account, errAccount := entity.NewAccount(uuid.NewString(), "237", "0001", "0123456-7", currentTime, currentTime)
	assert.NoError(t, errAccount)
	errCredit := account.Credit(0)
	assert.EqualError(t, errCredit, "amount must be greater than zero")
}

func TestNotShouldCreateANewAccountAndDebitIfAmountMustBeGreaterThanZero(t *testing.T) {
	currentTime := time.Now()
	account, errAccount := entity.NewAccount(uuid.NewString(), "237", "0001", "0123456-7", currentTime, currentTime)
	assert.NoError(t, errAccount)
	errDebit := account.Debit(0)
	assert.EqualError(t, errDebit, "amount must be greater than zero")
}

func TestShouldCreateANewAccountAndCredit(t *testing.T) {
	currentTime := time.Now()
	account, errAccount := entity.NewAccount(uuid.NewString(), "237", "0001", "0123456-7", currentTime, currentTime)
	account.Credit(1000.00)
	assert.Equal(t, 1000.00, account.GetBalance())
	assert.NoError(t, errAccount)
}

func TestShouldCreateANewAccountAndDebit(t *testing.T) {
	currentTime := time.Now()
	account, errAccount := entity.NewAccount(uuid.NewString(), "237", "0001", "0123456-7", currentTime, currentTime)
	account.Credit(1000.00)
	account.Debit(100.00)
	assert.Equal(t, 900.00, account.GetBalance())
	assert.NoError(t, errAccount)
}
