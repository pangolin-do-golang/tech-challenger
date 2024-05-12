package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/order"
)

type OrderHandler struct {
	service *order.Service
}

func RegisterOrderHandlers(router *gin.Engine, service *order.Service) {
	handler := &OrderHandler{
		service: service,
	}

	router.GET("/order/:id", handler.Get)
}

// Get Order godoc
// @Summary Get order by ID
// @Description Get order by ID
// @Tags order Routes
// @Param        id   path      string  true  "Order ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} order.Order{}
// @Router /order/:id [get]
func (handler *OrderHandler) Get(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid identifier informed.",
		})

		return
	}

	order, err := handler.service.Get(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went bad :(",
		})

		return
	}

	c.JSON(http.StatusOK, order)
}
