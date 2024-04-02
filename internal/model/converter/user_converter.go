package converter

import (
	"kevinhend97/mini-scan-asset/internal/entity"
	"kevinhend97/mini-scan-asset/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		ID:    int(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}
}

func UserTokenResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		Token: user.Token,
	}
}
