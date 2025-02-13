package controller

import (
	"net/http"
	"shaar/domain"

	"github.com/gin-gonic/gin"
)

type AirController struct {
	AirUsecase domain.AirUsecase
}

func NewAirController(airUsecase domain.AirUsecase) *AirController {
	return &AirController{
		AirUsecase: airUsecase,
	}
}

func (wc *AirController) GetAir(c *gin.Context) {
	err, airResponse := wc.AirUsecase.GetAir(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, airResponse)
}
