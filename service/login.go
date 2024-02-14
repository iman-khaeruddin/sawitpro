package service

import (
	"context"
	"fmt"
	"net/http"
	"sawitpro/dto"
	"sawitpro/middleware"
	"sawitpro/repository"
	"sawitpro/util"
)

type LoginService struct {
	UserRepo repository.UserInterface
}

type LoginInterface interface {
	Login(ctx context.Context, request *dto.LoginRequest) (dto.BaseResponse, int)
}

func (ls LoginService) Login(ctx context.Context, request *dto.LoginRequest) (dto.BaseResponse, int) {
	var response dto.LoginResponse
	user, err := ls.UserRepo.FindByPhoneNumberPassword(ctx, request.PhoneNumber)

	if err != nil {
		return dto.FailedResponse(err.Error(), http.StatusInternalServerError)
	}

	if user.ID == 0 {
		return dto.FailedResponse("user not found", http.StatusBadRequest)
	}

	fmt.Println(user.Password)
	if user.Password != util.HashPassword(request.Password) {
		return dto.FailedResponse("wrong password", http.StatusBadRequest)
	}

	token, err := middleware.CreateToken(user.ID)

	if err != nil {
		return dto.FailedResponse(err.Error(), http.StatusInternalServerError)
	}

	response.UserID = user.ID
	response.Token = token

	ls.UserRepo.AddLoginAttempt(ctx, user.PhoneNumber)

	return dto.SuccessResponse(response, "login success", http.StatusOK)
}
