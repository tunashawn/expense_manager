package controller

import (
	"Expense_Manager/commons/response"
	"Expense_Manager/pkg/auth_service/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

type Auth interface {
	Login(ctx *gin.Context)
	VerifyJWTToken(ctx *gin.Context)
}

type AuthImpl struct {
	authService services.AuthService
	response    response.HttpResponse
}

func NewAuth() (Auth, error) {
	authService, err := services.NewAuthService()
	if err != nil {
		return nil, errors.Wrap(err, "could not create auth service")
	}
	return &AuthImpl{
		authService: authService,
	}, nil
}

func (a *AuthImpl) Login(ctx *gin.Context) {
	credential, err := a.authService.GetCredential(ctx)
	if err != nil {
		a.response.BadRequest(errors.Wrap(err, "could not get credential from basic auth token"), ctx)
		return
	}

	ok, err := a.authService.VerifyCredential(*credential)
	if err != nil {
		a.response.InternalServerError(errors.Wrap(err, "could not create jwt token"), ctx)
		return
	}
	if !ok {
		a.response.Unauthorized(errors.New("wrong username or password"), ctx)
		return
	}

	signedToken, err := a.authService.CreateJWTToken(*credential)
	if err != nil {
		a.response.InternalServerError(err, ctx)
	}

	a.response.Success(response.JWTToken{Token: signedToken}, ctx)
}

func (a *AuthImpl) VerifyJWTToken(ctx *gin.Context) {
	valid, err := a.authService.VerifyToken(ctx)

	switch {
	case valid:
		ctx.Next()
		return
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		a.response.Unauthorized(err, ctx)
		ctx.Abort()
		return
	default:
		a.response.InternalServerError(errors.Wrap(err, "could not verify token"), ctx)
		ctx.Abort()
		return
	}
}
