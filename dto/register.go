package dto

type RegisterRequest struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	FullName    string `json:"fullName" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type RegisterResponse struct {
	UserID   int    `json:"userId"`
	FullName string `json:"fullName"`
}
