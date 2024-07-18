package controllers

import (
	"Expense_Manager/commons/response"
	services2 "Expense_Manager/pkg/auth_service/services"
	"Expense_Manager/pkg/user_manager_service/services"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type UserManager interface {
	CreateNewUser(ctx *gin.Context)
	ChangePassword(ctx *gin.Context)
}

type UserManagerImpl struct {
	response           response.HttpResponse
	userManagerService services.UserManagerService
	auth               services2.AuthService
}

func NewUserManager() (UserManager, error) {
	userManagerService, err := services.NewUserManagerService()
	if err != nil {
		return nil, err
	}

	return &UserManagerImpl{
		userManagerService: userManagerService,
		auth:               &services2.AuthServiceImpl{},
	}, nil
}

func (u *UserManagerImpl) ChangePassword(ctx *gin.Context) {
	credential, err := u.auth.GetCredentialFromToken(ctx)
	if err != nil {
		u.response.InternalServerError(errors.Wrap(err, "could not get credential from auth token"), ctx)
		return
	}

	newHashedPassword, err := u.userManagerService.GetNewPassword(ctx)
	if err != nil {
		u.response.BadRequest(err, ctx)
		return
	}

	err = u.userManagerService.UpdateUserPassword(credential.Username, newHashedPassword)
	if err != nil {
		u.response.InternalServerError(errors.Wrap(err, "could not update password"), ctx)
		return
	}

	u.response.Success(nil, ctx)
}

func (u *UserManagerImpl) CreateNewUser(ctx *gin.Context) {
	user, err := u.userManagerService.GetNewUserInformation(ctx)
	if err != nil {
		u.response.BadRequest(err, ctx)
		return
	}

	ok, err := u.userManagerService.ValidateUserInformation(*user)
	if err != nil {
		if !ok {
			u.response.BadRequest(err, ctx)
			return
		}
		u.response.InternalServerError(errors.Wrap(err, "could not check for user extinction"), ctx)
		return
	}

	err = u.userManagerService.CreateNewUser(*user)
	if err != nil {
		u.response.InternalServerError(errors.Wrap(err, "could not create new user"), ctx)
		return
	}

	u.response.Success(nil, ctx)
}
