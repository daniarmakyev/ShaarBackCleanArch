package controller

import (
	"net/http"
	"shaar/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PlacesController struct {
	PlacesUsecase domain.PlacesUsecase
}

func NewPlacesController(pu domain.PlacesUsecase) *PlacesController {
	return &PlacesController{
		PlacesUsecase: pu,
	}
}

func (pc *PlacesController) GetAllPlaces(c *gin.Context) {
	places, err := pc.PlacesUsecase.GetAllPlaces(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Error fetching places: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, places)
}

func (pc *PlacesController) GetPlacesByCategory(c *gin.Context) {
	category := c.Param("categories")
	places, err := pc.PlacesUsecase.GetPlacesByCategory(c.Request.Context(), category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Error fetching places by category: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, places)
}

func (pc *PlacesController) GetPlacesByPrice(c *gin.Context) {
	priceStr := c.Param("price")
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid price parameter"})
		return
	}
	places, err := pc.PlacesUsecase.GetPlacesByPrice(c.Request.Context(), price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Error fetching places by price: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, places)
}

func (pc *PlacesController) GetAllCategories(c *gin.Context) {
	categories, err := pc.PlacesUsecase.GetAllCategories(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Error fetching categories: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}
