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

type WalletController interface {
	CreateNewWallet(ctx *gin.Context)
	GetWallet(ctx *gin.Context)
	GetWalletList(ctx *gin.Context)
	UpdateWallet(ctx *gin.Context)
	DeleteWallet(ctx *gin.Context)
}

type WalletControllerImpl struct {
	walletService services.WalletService
	response      response.HttpResponse
}

func (w *WalletControllerImpl) CreateNewWallet(ctx *gin.Context) {
	var newWallet models.Wallet
	err := ctx.ShouldBind(&newWallet)
	if err != nil {
		w.response.BadRequest(
			errors.Wrap(err, "could not bind request body into models.Wallet"),
			ctx,
		)
		return
	}

	if err := newWallet.VerifyWallet(); err != nil {
		w.response.BadRequest(err, ctx)
		return
	}

	err = w.walletService.CreateWallet(newWallet)
	if err != nil {
		w.response.BadRequest(
			errors.Wrap(err, fmt.Sprintf("could not create new wallet of %+v", newWallet)),
			ctx,
		)
		return
	}

	w.response.Success(nil, ctx)
}

func (w *WalletControllerImpl) GetWallet(ctx *gin.Context) {
	// TODO: verify if the wallet belongs to the current user
	id := ctx.Param("id")

	walletID, err := strconv.Atoi(id)
	if err != nil || walletID < 1 {
		w.response.BadRequest(errors.New("invalid wallet id"), ctx)
		return
	}

	walletDetail, err := w.walletService.GetWallet(walletID)
	if err != nil {
		w.response.InternalServerError(
			errors.Wrap(err, fmt.Sprintf("could not get detailed of walletID=%d", walletID)),
			ctx,
		)
		return
	}

	w.response.Success(walletDetail, ctx)
}

func (w *WalletControllerImpl) GetWalletList(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (w *WalletControllerImpl) UpdateWallet(ctx *gin.Context) {
	// TODO: verify if the wallet belongs to the current user

	var newWallet models.Wallet
	err := ctx.ShouldBind(&newWallet)
	if err != nil {
		w.response.BadRequest(
			errors.Wrap(err, "could not bind request body into models.Wallet"),
			ctx,
		)
		return
	}

	if err := newWallet.VerifyWallet(); err != nil {
		w.response.BadRequest(err, ctx)
		return
	}

	err = w.walletService.UpdateWallet(newWallet)
	if err != nil {
		w.response.BadRequest(
			errors.Wrap(err, fmt.Sprintf("could not update wallet of %+v", newWallet)),
			ctx,
		)
		return
	}

	w.response.Success(nil, ctx)
}

func (w *WalletControllerImpl) DeleteWallet(ctx *gin.Context) {
	// TODO: verify if the wallet belongs to the current user

	id := ctx.Param("id")

	walletID, err := strconv.Atoi(id)
	if err != nil || walletID < 1 {
		w.response.BadRequest(errors.New("invalid wallet id"), ctx)
		return
	}

	err = w.walletService.DeleteWallet(walletID)
	if err != nil {
		w.response.InternalServerError(
			errors.Wrap(err, fmt.Sprintf("could not delete walletID=%d", walletID)),
			ctx,
		)
		return
	}

	w.response.Success(nil, ctx)
}

func NewWalletController(db *bun.DB) (WalletController, error) {
	walletService, err := services.NewWalletService(db)
	if err != nil {
		return nil, errors.Wrap(err, "could not init new wallet service")
	}

	return &WalletControllerImpl{
		walletService: walletService,
	}, nil
}
