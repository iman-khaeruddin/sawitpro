package dto

type ProfileResponse struct {
	ID          int    `json:"userId"`
	FullName    string `json:"fullName"`
	PhoneNumber string `json:"phoneNumber"`
}

type UpdateProfileRequest struct {
	ID          int    `json:"userId"`
	FullName    string `json:"fullName" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
}
