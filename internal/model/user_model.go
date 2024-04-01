package model

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,max=20"`
	Password string `json:"password" validate:"required,max=100"`
}

type UserResponse struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Token string `json:"token,omitempty"`
}

type VerifyUserRequest struct {
	Token string `validate:"required,max=255"`
}

type LogoutRequest struct {
	ID int `json:"id" validate:"required,max=11"`
}

type GetUserRequest struct {
	ID int `json:"id" validate:"required,max=11"`
}
