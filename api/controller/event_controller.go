package controller

import (
	"net/http"
	"shaar/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	EventUsecase domain.EventUsecase
}

func NewEventController(eu domain.EventUsecase) *EventController {
	return &EventController{
		EventUsecase: eu,
	}
}

func (ec *EventController) Create(c *gin.Context) {
	var request domain.EventRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid input data"})
		return
	}

	event := domain.EventRequest{
		Name:        request.Name,
		Description: request.Description,
		Date:        request.Date,
		Address:     request.Address,
	}

	ec.EventUsecase.Create(c, &event)
	c.JSON(http.StatusOK, gin.H{"message": "Event created successfully!"})
}

func (ec *EventController) GetAllEvents(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "6")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid page parameter"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid limit parameter"})
		return
	}

	events, err := ec.EventUsecase.GetAllEvents(c.Request.Context(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Error fetching events: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}
