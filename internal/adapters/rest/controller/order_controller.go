package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/order"
)

type OrderController struct {
	service order.IOrderService
}

func NewOrderController(service order.IOrderService) *OrderController {
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
		c.JSON(http.StatusInternalServerError, gin.H{
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

	o, err := ctrl.service.Get(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, o)
}

// Create Order godoc
// @Summary Create order from Cart
// @Description Create order from Cart
// @Tags order Routes
// @Accept  json
// @Produce  json
// @Success 200 {object} order.Order{}
// @Router /order [post]
func (ctrl OrderController) Create(c *gin.Context) {
	type Payload struct {
		ClientID uuid.UUID `json:"client_id" binding:"required"`
	}

	payload := &Payload{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	o, err := ctrl.service.Create(payload.ClientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, o)
}
