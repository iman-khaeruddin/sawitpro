package dto

type LoginRequest struct {
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Password    string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	UserID int    `json:"userID"`
}
