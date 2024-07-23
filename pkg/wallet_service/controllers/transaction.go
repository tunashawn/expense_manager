package controllers

import (
	"Expense_Manager/commons/response"
	"Expense_Manager/pkg/wallet_service/models"
	"Expense_Manager/pkg/wallet_service/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"strconv"
)

type TransactionController interface {
	CreateNewTransaction(ctx *gin.Context)
	GetTransaction(ctx *gin.Context)
	GetTransactionList(ctx *gin.Context)
	UpdateTransaction(ctx *gin.Context)
	DeleteTransaction(ctx *gin.Context)
}

type TransactionControllerImpl struct {
	transactionService services.TransactionService
	response           response.HttpResponse
}

func (t *TransactionControllerImpl) CreateNewTransaction(ctx *gin.Context) {
	var newTransaction models.Transaction
	err := ctx.ShouldBind(&newTransaction)
	if err != nil {
		t.response.BadRequest(
			errors.Wrap(err, "could not bind request body into models.Transaction"),
			ctx,
		)
		return
	}

	if err := newTransaction.VerifyTransaction(); err != nil {
		t.response.BadRequest(
			errors.Wrap(err, "invalid transaction"),
			ctx,
		)
		return
	}

	err = t.transactionService.CreateNewTransaction(newTransaction)
	if err != nil {
		t.response.BadRequest(
			errors.Wrap(err, fmt.Sprintf("could not create new transaction of %+v", newTransaction)),
			ctx,
		)
		return
	}

	t.response.Success(nil, ctx)
}

func (t *TransactionControllerImpl) GetTransaction(ctx *gin.Context) {
	// TODO: verify if the transaction belongs to the current user
	id := ctx.Param("id")

	transactionID, err := strconv.Atoi(id)
	if err != nil || transactionID < 1 {
		t.response.BadRequest(errors.New("invalid transaction id"), ctx)
		return
	}

	transactionDetail, err := t.transactionService.GetTransaction(transactionID)
	if err != nil {
		t.response.InternalServerError(
			errors.Wrap(err, fmt.Sprintf("could not get detailed of transactionID=%d", transactionID)),
			ctx,
		)
		return
	}

	t.response.Success(transactionDetail, ctx)
}

func (t *TransactionControllerImpl) GetTransactionList(ctx *gin.Context) {
	// TODO: verify if the transaction belongs to the current user

	//TODO implement me
	panic("implement me")
}

func (t *TransactionControllerImpl) UpdateTransaction(ctx *gin.Context) {
	// TODO: verify if the transaction belongs to the current user

	var transaction models.Transaction
	err := ctx.ShouldBind(&transaction)
	if err != nil {
		t.response.BadRequest(errors.Wrap(err, "invalid request body"), ctx)
		return
	}

	if err := transaction.VerifyTransaction(); err != nil {
		t.response.BadRequest(
			errors.Wrap(err, "invalid transaction"),
			ctx,
		)
		return
	}
}

func (t *TransactionControllerImpl) DeleteTransaction(ctx *gin.Context) {
	id := ctx.Param("id")

	transactionID, err := strconv.Atoi(id)
	if err != nil || transactionID < 1 {
		t.response.BadRequest(errors.New("invalid transaction id"), ctx)
		return
	}

	err = t.transactionService.DeleteTransaction(transactionID)
	if err != nil {
		t.response.InternalServerError(
			errors.Wrap(err, fmt.Sprintf("could not delete transactionID=%d", transactionID)),
			ctx,
		)
		return
	}

	t.response.Success(nil, ctx)
}

func NewTransactionController(db *bun.DB) (TransactionController, error) {
	transactionService, err := services.NewTransactionService(db)
	if err != nil {
		return nil, errors.Wrap(err, "could not init new wallet service")
	}

	return &TransactionControllerImpl{
		transactionService: transactionService,
	}, nil
}
