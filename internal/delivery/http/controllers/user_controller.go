package controllers

import (
	"kevinhend97/mini-scan-asset/internal/usecase"

	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log     *logrus.Logger
	UseCase *usecase.UserUsecase
}

func NewUserController(log *logrus.Logger, useCase *usecase.UserUsecase) *UserController {
	return &UserController{
		Log:     log,
		UseCase: useCase,
	}
}
