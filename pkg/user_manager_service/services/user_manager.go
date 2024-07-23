package services

import (
	"Expense_Manager/pkg/user_manager_service/models"
	"Expense_Manager/pkg/user_manager_service/repositories"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

const (
	usernameRegex = `^[^\W_][\w.]{0,29}$`
	passwordRegex = `^[A-Za-z0-9_]{8,16}$`
	emailRegex    = `^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`
)

type UserManagerService interface {
	GetNewUserInformation(ctx *gin.Context) (*models.User, error)
	CreateNewUser(user models.User) error
	ValidateUserInformation(user models.User) (bool, error)
	GetNewPassword(ctx *gin.Context) ([]byte, error)
	UpdateUserPassword(username string, password []byte) error
}

type UserManagerServiceImpl struct {
	mysql repositories.UserManagerMySQLRepo
}

func NewUserManagerService() (UserManagerService, error) {
	mongo, err := repositories.NewUserManagerMySQLRepo()
	if err != nil {
		return nil, err
	}

	return &UserManagerServiceImpl{mysql: mongo}, nil
}

func (u *UserManagerServiceImpl) UpdateUserPassword(username string, password []byte) error {
	err := u.mysql.UpdateUserPassword(username, password)
	if err != nil {
		return errors.Wrap(err, "could not perform update user's password")
	}

	return nil
}

func (u *UserManagerServiceImpl) GetNewPassword(ctx *gin.Context) ([]byte, error) {
	var user models.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		return nil, err
	}

	newPassword := user.Password

	if err := u.validatePassword(newPassword); err != nil {
		return nil, err
	}
	// hash and salt the password
	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.MinCost)
	if err != nil {
		return nil, errors.Wrap(err, "could not hash and salt new password")
	}
	return hashedNewPassword, nil
}

func (u *UserManagerServiceImpl) GetNewUserInformation(ctx *gin.Context) (*models.User, error) {
	var user models.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		return nil, errors.Wrap(err, "could not get new user information")
	}

	if err := user.FormatNewUserInformation(); err != nil {
		return nil, errors.Wrap(err, "could not format new user information")
	}

	return &user, nil
}

// if ok, err == nil => 200
// if ok, err != nil => 500
// if !ok, err != nil => 400
func (u *UserManagerServiceImpl) ValidateUserInformation(user models.User) (bool, error) {

	if ok, err := u.validateUsername(user.Username); err != nil {
		return ok, err
	}

	if err := u.validatePassword(user.Password); err != nil {
		return false, err
	}

	if ok, err := u.validateEmail(user.Email); err != nil {
		return ok, err
	}

	return false, nil
}

// if username is ok and err == nil => username is accepted
// if username is ok and err != nil => code 500
// if username is !ok and err != nil => code 400
func (u *UserManagerServiceImpl) validateUsername(username string) (bool, error) {
	// check format
	if matched, err := regexp.MatchString(usernameRegex, username); !matched {
		if err != nil {
			return true, err
		}
		return false, errors.New("wrong username format")
	}

	// check if username already existed
	exist, err := u.mysql.IsUsernameExist(username)
	if err != nil {
		return true, errors.Wrap(err, "could not check username extinction")
	}
	if exist {
		return false, errors.New(fmt.Sprintf("username %s is already used", username))
	}

	return true, nil
}

func (u *UserManagerServiceImpl) validatePassword(password string) error {
	if matched, err := regexp.MatchString(passwordRegex, password); !matched {
		if err != nil {
			return err
		}
		return errors.New("wrong password format")
	}

	return nil
}

// if email is ok and err == nil => email is accepted
// if email is ok and err != nil => code 500
// if email is !ok and err != nil => code 400
func (u *UserManagerServiceImpl) validateEmail(email string) (bool, error) {
	if matched, err := regexp.MatchString(emailRegex, email); !matched {
		if err != nil {
			return true, err
		}
		return false, errors.New("wrong email format")
	}

	// check if username already existed
	exist, err := u.mysql.IsEmailExist(email)
	if err != nil {
		return true, errors.Wrap(err, "could not check username extinction")
	}
	if exist {
		return false, errors.New(fmt.Sprintf("email %s is already used", email))
	}

	return true, nil
}

func (u *UserManagerServiceImpl) CreateNewUser(user models.User) error {
	return u.mysql.CreateNewUser(user)
}
