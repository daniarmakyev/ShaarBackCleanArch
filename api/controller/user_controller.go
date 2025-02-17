package controller

import (
	"fmt"
	"log"
	"net/http"
	"shaar/bootstrap"
	"shaar/domain"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.UserUseCase
	Env         *bootstrap.Env
}

func extractToken(r *gin.Context) (string, error) {
	authHeader := r.GetHeader("Authorization")
	if authHeader == "" {
		return "", gin.Error{
			Err:  fmt.Errorf("Authorization header not found"),
			Type: gin.ErrorTypePublic,
		}
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", gin.Error{
			Err:  fmt.Errorf("Invalid authorization header format"),
			Type: gin.ErrorTypePublic,
		}
	}

	return parts[1], nil
}

func (uc *UserController) GetUser(c *gin.Context) {
	token, err := extractToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Authorization token missing or malformed"})
		return
	}

	id, err := uc.UserUsecase.ExtractIDFromToken(token, uc.Env.AccessTokenSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid token"})
		return
	}

	user, err := uc.UserUsecase.GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User not found: " + err.Error()})
		return
	}

	userResponse := domain.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
		Payment:  user.Payment,
		Name:     user.Name,
		Surname:  user.Surname,
	}

	c.JSON(http.StatusOK, userResponse)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	var request domain.UserUpdateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: fmt.Sprintf("Invalid request data: %v", err)})
		log.Print("Binding error:", err)
		return
	}
	log.Printf("Request data: %+v", request)

	token, err := extractToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Authorization token missing or malformed"})
		return
	}

	id, err := uc.UserUsecase.ExtractIDFromToken(token, uc.Env.AccessTokenSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid token"})
		return
	}

	_, err = uc.UserUsecase.GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User not found"})
		return
	}

	updatedUser, err := uc.UserUsecase.UpdateUser(c, id, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: fmt.Sprintf("Error updating user: %v", err)})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}
