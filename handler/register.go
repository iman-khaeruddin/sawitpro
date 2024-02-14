package handler

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"regexp"
	"sawitpro/dto"
	"sawitpro/repository"
	"sawitpro/service"
	"sawitpro/util"
)

type RegisterHandler struct {
	db      *gorm.DB
	service service.RegisterInterface
}

func NewRegisterHandler(db *gorm.DB) RegisterHandler {
	return RegisterHandler{
		db: db,
	}
}

func (h RegisterHandler) RegisterHandler(route *echo.Echo) {
	userRepo := repository.NewUser(h.db)
	useCase := service.RegisterService{UserRepo: userRepo}
	h.service = useCase

	route.POST("/register",
		h.register)
}

func (h RegisterHandler) register(c echo.Context) error {
	request := new(dto.RegisterRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BaseResponse{
			Success:      false,
			MessageTitle: "error",
			Message:      err.Error(),
		})
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.BaseResponse{
			Success:      false,
			MessageTitle: "error",
			Message:      err.Error(),
		})
	}

	// validate the phone number using a regular expression
	re := regexp.MustCompile(`^(62)8[1-9][0-9]{6,10}$`)
	if !re.MatchString(request.PhoneNumber) {
		return c.JSON(http.StatusBadRequest, dto.BaseResponse{
			Success:      false,
			MessageTitle: "error",
			Message:      "phone number not valid, please used country code",
		})
	}

	six, num, upper, special := util.VerifyPassword(request.Password)
	if six == false {
		return c.JSON(http.StatusBadRequest, dto.BaseResponse{
			Success:      false,
			MessageTitle: "error",
			Message:      "password min 6 character",
		})
	}
	if num == false {
		return c.JSON(http.StatusBadRequest, dto.BaseResponse{
			Success:      false,
			MessageTitle: "error",
			Message:      "password at least contain 1 number",
		})
	}
	if upper == false {
		return c.JSON(http.StatusBadRequest, dto.BaseResponse{
			Success:      false,
			MessageTitle: "error",
			Message:      "password at least contain 1 capital character",
		})
	}
	if special == false {
		return c.JSON(http.StatusBadRequest, dto.BaseResponse{
			Success:      false,
			MessageTitle: "error",
			Message:      "password at least contain 1 symbol",
		})
	}

	res, responseCode := h.service.Register(c.Request().Context(), request)

	c.JSON(responseCode, res)
	return nil
}
