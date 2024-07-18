package services

import (
	"Expense_Manager/commons/config"
	"Expense_Manager/commons/response"
	"Expense_Manager/pkg/auth_service/models"
	"Expense_Manager/pkg/auth_service/repository"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type AuthService interface {
	GetCredential(ctx *gin.Context) (*models.Credential, error)
	CreateJWTToken(credential models.Credential) (string, error)
	VerifyToken(ctx *gin.Context) (bool, error)
	VerifyCredential(credential models.Credential) (bool, error)
	GetCredentialFromToken(ctx *gin.Context) (*models.Credential, error)
}

type AuthServiceImpl struct {
	secretKey       []byte
	tokenExpireTime int
	response        response.HttpResponse
	mongo           repository.AuthServiceMongoRepo
}

func NewAuthService() (*AuthServiceImpl, error) {
	authConfig := new(config.AuthConfig)
	err := config.GetConfig(authConfig)
	if err != nil {
		return nil, errors.Wrap(err, "could not get auth config")
	}

	mongo, err := repository.NewAuthServiceMongoRepository()
	if err != nil {
		return nil, errors.Wrap(err, "could not create auth service mongo repository")
	}

	return &AuthServiceImpl{
		secretKey:       []byte(authConfig.JWTSecretKey),
		tokenExpireTime: authConfig.JWTTokenExpireTime,
		mongo:           mongo,
	}, nil
}

func (a *AuthServiceImpl) GetCredential(ctx *gin.Context) (*models.Credential, error) {
	basicToken, err := a.getAuthToken("Basic", ctx)
	if err != nil {
		return nil, errors.Wrap(err, "could not get basic auth token")
	}

	decodedToken, err := base64.StdEncoding.DecodeString(basicToken)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode basic auth token")
	}

	// extract username and password
	credentials := strings.Split(string(decodedToken), ":")
	if len(credentials) != 2 {
		return nil, errors.New("invalid basic auth token")
	}

	// hash and salt the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(credentials[1]), bcrypt.MinCost)
	if err != nil {
		return nil, errors.Wrap(err, "could not hash and salt password")
	}

	return &models.Credential{
		Username:       credentials[0],
		Password:       credentials[1],
		HashedPassword: hashedPassword,
	}, nil
}

func (a *AuthServiceImpl) GetCredentialFromToken(ctx *gin.Context) (*models.Credential, error) {
	bearerToken, err := a.getAuthToken("Bearer", ctx)
	if err != nil {
		return nil, err
	}
	token, _, err := new(jwt.Parser).ParseUnverified(bearerToken, jwt.MapClaims{})
	if err != nil {
		return nil, errors.Wrap(err, "could not parse token")
	}

	// Extract the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.Wrap(err, "invalid token claims")
	}

	// Print the payload
	var res models.Credential
	err = mapstructure.Decode(claims["credential"], &res)
	if err != nil {
		return nil, errors.Wrap(err, "could not decode credential claim to struct Credential")
	}

	return &res, nil
}

func (a *AuthServiceImpl) getAuthToken(authType string, ctx *gin.Context) (string, error) {
	authHeader := ctx.Request.Header.Get("Authorization")

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != authType {
		return "", errors.New("invalid auth token")
	}

	return parts[1], nil
}

func (a *AuthServiceImpl) CreateJWTToken(credential models.Credential) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"credential": credential.ToJWTPayload(),
			"iat":        time.Now().Unix(),
			"exp":        time.Now().Add(time.Minute * time.Duration(a.tokenExpireTime)).Unix(),
		})

	signedToken, err := token.SignedString(a.secretKey)
	if err != nil {
		return "", errors.Wrap(err, "could not sign jwt token with the secret key")
	}

	return signedToken, nil
}

func (a *AuthServiceImpl) VerifyToken(ctx *gin.Context) (bool, error) {
	bearerToken, err := a.getAuthToken("Bearer", ctx)
	if err != nil {
		return false, errors.Wrap(err, "could not get bearer token")
	}

	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) { return a.secretKey, nil })

	if err != nil {
		return false, errors.Wrap(err, "could not validate jwt token")
	}

	return token.Valid, nil
}

func (a *AuthServiceImpl) VerifyCredential(credential models.Credential) (bool, error) {
	hashedPassword, err := a.mongo.GetHashedPasswordOfUser(credential.Username)
	if err != nil {
		return false, errors.Wrap(err, "could not get hashed password of user")
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(credential.Password))
	if err != nil {
		return false, nil
	}

	return true, nil
}
