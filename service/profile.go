package service

import (
	"context"
	"errors"
	"net/http"
	"sawitpro/dto"
	"sawitpro/entity"
	"sawitpro/repository"
	"strings"
)

type ProfileService struct {
	UserRepo repository.UserInterface
}

type ProfileInterface interface {
	GetProfile(ctx context.Context, ID int) (dto.BaseResponse, int)
	UpdateProfile(ctx context.Context, request *dto.UpdateProfileRequest) (dto.BaseResponse, int)
}

func (ps ProfileService) GetProfile(ctx context.Context, ID int) (dto.BaseResponse, int) {
	user, err := ps.UserRepo.FindByID(ctx, ID)
	if err != nil {
		return dto.FailedResponse(err.Error(), http.StatusInternalServerError)
	}
	return dto.SuccessResponse(user, "get profile success", http.StatusOK)
}

func (ps ProfileService) UpdateProfile(ctx context.Context, request *dto.UpdateProfileRequest) (dto.BaseResponse, int) {
	user, err := ps.UserRepo.FindByID(ctx, request.ID)
	if err != nil {
		return dto.FailedResponse(err.Error(), http.StatusInternalServerError)
	}

	usr := &entity.User{
		ID:          user.ID,
		FullName:    request.FullName,
		PhoneNumber: request.PhoneNumber,
	}

	_, err = ps.UserRepo.UpdateSelectedFields(ctx, usr, "FullName", "PhoneNumber")
	if err != nil {
		if strings.Contains(err.Error(), errors.New("duplicate key value violates unique constraint").Error()) {
			return dto.FailedResponse("phone number has registered", http.StatusConflict)
		}
		return dto.FailedResponse(err.Error(), http.StatusInternalServerError)
	}

	return dto.SuccessResponse(nil, "update profile success", http.StatusOK)
}
