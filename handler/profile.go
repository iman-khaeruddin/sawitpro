package handler

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"os"
	"sawitpro/dto"
	"sawitpro/middleware"
	"sawitpro/repository"
	"sawitpro/service"
)

type ProfileHandler struct {
	db      *gorm.DB
	service service.ProfileInterface
}

func NewProfileHandler(db *gorm.DB) ProfileHandler {
	return ProfileHandler{
		db: db,
	}
}

func (h ProfileHandler) ProfileHandler(route *echo.Echo) {
	userRepo := repository.NewUser(h.db)
	useCase := service.ProfileService{UserRepo: userRepo}
	h.service = useCase

	e := route.Group("")

	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(middleware.AuthData)
		},
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
	}
	e.Use(echojwt.WithConfig(config))
	e.GET("/profile", h.getProfile)
	e.PUT("/profile", h.updateProfile)
}

func (h ProfileHandler) getProfile(c echo.Context) error {
	authData, _ := middleware.GetData(c)
	res, responseCode := h.service.GetProfile(c.Request().Context(), authData.ID)
	c.JSON(responseCode, res)
	return nil
}

func (h ProfileHandler) updateProfile(c echo.Context) error {
	authData, _ := middleware.GetData(c)
	request := new(dto.UpdateProfileRequest)
	if err := c.Bind(request); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	request.ID = authData.ID
	res, responseCode := h.service.UpdateProfile(c.Request().Context(), request)
	c.JSON(responseCode, res)
	return nil
}
