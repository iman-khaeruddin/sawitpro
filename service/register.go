package service

import (
	"context"
	"errors"
	"net/http"
	"sawitpro/dto"
	"sawitpro/entity"
	"sawitpro/repository"
	"sawitpro/util"
	"strings"
)

type RegisterService struct {
	UserRepo repository.UserInterface
}

type RegisterInterface interface {
	Register(ctx context.Context, request *dto.RegisterRequest) (dto.BaseResponse, int)
}

func (rs RegisterService) Register(ctx context.Context, request *dto.RegisterRequest) (dto.BaseResponse, int) {
	var response dto.RegisterResponse
	user := entity.User{
		FullName:     request.FullName,
		PhoneNumber:  request.PhoneNumber,
		Password:     util.HashPassword(request.Password),
		LoginAttempt: 0,
	}

	result, err := rs.UserRepo.Create(ctx, &user)

	if err != nil {
		if strings.Contains(err.Error(), errors.New("duplicate key value violates unique constraint").Error()) {
			return dto.FailedResponse("phone number has registered", http.StatusBadRequest)
		}
		return dto.FailedResponse(err.Error(), http.StatusInternalServerError)
	}

	response.UserID = result.ID
	response.FullName = result.FullName

	return dto.SuccessResponse(response, "register success", http.StatusOK)
}
