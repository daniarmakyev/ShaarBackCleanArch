package controller

import (
	"net/http"
	"shaar/bootstrap"
	"shaar/domain"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SigninController struct {
	SigninUsecase domain.SigninUsecase
	Env           *bootstrap.Env
}

func (lc *SigninController) Signin(c *gin.Context) {
	var request domain.SigninRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid request: " + err.Error()})
		return
	}

	user, err := lc.SigninUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found: " + err.Error()})
		return
	}

	if user == (domain.User{}) {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found: " + err.Error()})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	if lc.Env == nil || lc.Env.AccessTokenSecret == "" || lc.Env.RefreshTokenSecret == "" {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Environment configuration is missing"})
		return
	}

	accessToken, err := lc.SigninUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Error generating access token: " + err.Error()})
		return
	}

	refreshToken, err := lc.SigninUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Error generating refresh token: " + err.Error()})
		return
	}

	signinResponse := domain.SigninResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signinResponse)
}
