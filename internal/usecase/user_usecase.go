package usecase

import (
	"context"
	"kevinhend97/mini-scan-asset/internal/entity"
	"kevinhend97/mini-scan-asset/internal/model"
	"kevinhend97/mini-scan-asset/internal/model/converter"
	"kevinhend97/mini-scan-asset/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUsecase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	UserRepository *repository.UserRepository
}

func NewUserUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, userRepository *repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		DB:             db,
		Log:            log,
		Validate:       validate,
		UserRepository: userRepository,
	}
}

func (c *UserUsecase) Login(ctx context.Context, request *model.LoginUserRequest) (*model.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, err
	}

	user := new(entity.User)

	if err := c.UserRepository.FindByEmail(tx, user, request.Email); err != nil {
		c.Log.Warnf("Failed find user by ID : %+v", err)
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		c.Log.Warnf("Failed to compare user password with bcrypt hash : %+v", err)
		return nil, err
	}

	user.Token = uuid.New().String()

	if err := c.UserRepository.Update(tx, user); err != nil {
		c.Log.Warnf("Failed save user: %+v", err)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, err
	}

	return converter.UserTokenResponse(user), nil
}
