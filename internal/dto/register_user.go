package dto

type RegisterUserRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"  validate:"required"`
	Username  string `json:"username"   validate:"required"`
	Password  string `json:"password"   validate:"required"`
}

type RegisterUserResponse struct {
	ID string `json:"id"`
}
