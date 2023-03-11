package auth

type UserAuth struct {
	Email       string `json:"email" validate:"email" binding:"required"`
	PasswordRaw string `json:"password" validate:"min=8" binding:"required"`
}
