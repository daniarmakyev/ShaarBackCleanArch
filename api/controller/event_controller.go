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

func (ec *EventController) GetEvents(c *gin.Context) {
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

	events, total, err := ec.EventUsecase.GetAllEvents(c.Request.Context(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Error fetching events: " + err.Error()})
		return
	}

	totalPages := (total + limit - 1) / limit

	response := gin.H{
		"events":      events,
		"total":       total,
		"page":        page,
		"limit":       limit,
		"totalPages":  totalPages,
		"hasNextPage": page < totalPages,
		"hasPrevPage": page > 1,
	}

	c.JSON(http.StatusOK, response)
}
