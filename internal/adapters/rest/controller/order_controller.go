package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/order"
)

type OrderController struct {
	service order.IService
}

func NewOrderController(service order.IService) *OrderController {
	return &OrderController{
		service: service,
	}
}

// GetAll Get all order's list
// @Summary Get order list
// @Description Get all order's list
// @Tags Order
// @Accept json
// @Produce json
// @Success 200 {object} []order.Order{}
// @Router /order [get]
func (ctrl OrderController) GetAll(c *gin.Context) {
	orderSlice, err := ctrl.service.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, orderSlice)
}

// Get Get a order by ID
// @Summary Get order by ID
// @Description Get a order by ID
// @Tags Order
// @Param id path string true "ID of the order"
// @Accept json
// @Produce json
// @Success 200 {object} order.Order{}
// @Failure 400 {object} map[string]any "{\"error\": \Invalid identifier informed\"}"
// @Router /order/{id} [get]
func (ctrl OrderController) Get(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid identifier informed.",
		})

		return
	}

	order, err := ctrl.service.Get(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, order)
}
