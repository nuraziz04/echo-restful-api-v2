package web

type UsersCreateRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UsersUpdatePasswordRequest struct {
	Id              int    `json:"id" validate:"required"`
	Username        string `json:"username" validate:"required"`
	OldPassword     string `json:"oldPassword" validate:"required"`
	NewPassword     string `json:"newPassword" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
}
