package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/xtabs12/test_dans/internal/dto"
	"github.com/xtabs12/test_dans/internal/ucase"
	"net/http"
)

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func NewAuth(authLogic ucase.AUTH, group *echo.Group) {
	kk := authImpl{
		authLogic: authLogic,
	}
	RouterAuth(kk, group)
}

type authImpl struct {
	authLogic ucase.AUTH
}

func RouterAuth(c authImpl, group *echo.Group) {
	group.POST("/login", c.Authenticate)
}

func (i *authImpl) Authenticate(e echo.Context) error {
	var request LoginRequest
	if err := e.Bind(&request); err != nil {
		return e.JSON(http.StatusInternalServerError, buildResErr(err.Error()))
	}

	tokenString, authenticateErr := i.authLogic.Authenticate(e.Request().Context(),
		dto.AuthenticateParams{
			Username: request.UserName,
			Password: request.Password,
		})
	if authenticateErr != nil {
		return e.JSON(http.StatusInternalServerError, buildResErr("invalid username or password"))
	}
	e.SetCookie(&http.Cookie{})
	return e.JSON(http.StatusOK, buildResx(tokenString))
}
