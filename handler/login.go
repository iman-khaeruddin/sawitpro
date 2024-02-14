package handler

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"regexp"
	"sawitpro/dto"
	"sawitpro/repository"
	"sawitpro/service"
)

type LoginHandler struct {
	db      *gorm.DB
	service service.LoginInterface
}

func NewLoginHandler(db *gorm.DB) LoginHandler {
	return LoginHandler{
		db: db,
	}
}

func (h LoginHandler) LoginHandler(route *echo.Echo) {
	userRepo := repository.NewUser(h.db)
	useCase := service.LoginService{UserRepo: userRepo}
	h.service = useCase

	route.POST("/login",
		h.login)
}

func (h LoginHandler) login(c echo.Context) error {
	request := new(dto.LoginRequest)
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

	// Validate the phone number using a regular expression
	re := regexp.MustCompile(`^(62)8[1-9][0-9]{6,10}$`)
	if !re.MatchString(request.PhoneNumber) {
		log.Println("Phone number is not valid:", request.PhoneNumber)
		return c.JSON(http.StatusBadRequest, dto.BaseResponse{
			Success:      false,
			MessageTitle: "error",
			Message:      "phone number not valid",
		})
	}

	res, responseCode := h.service.Login(c.Request().Context(), request)

	return c.JSON(responseCode, res)
}
