package controller

import (
	"github.com/pangolin-do-golang/tech-challenge/internal/application/cart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CartController struct {
	service cart.IService
}

func NewCartController(cart cart.IService) *CartController {
	return &CartController{
		service: cart,
	}
}

type AddProductPayload struct {
	ClientID  uuid.UUID `json:"client_id" binding:"required"`
	ProductID uuid.UUID `json:"product_id" binding:"required"`
	// TODO validação de quantidade >= 0 / estoque
	Quantity int    `json:"quantity" binding:"required"`
	Comments string `json:"comments"`
}

// AddProduct adds a Product to Customer's Cart
// @Description Adds a Product to Customer's Cart
// @Tags Cart
// @Param payload body controller.AddProductPayload true "AddProductPayload"
// @Accept json
// @Produce json
// @Success 200
// @Failure 500 {object} map[string]any "{\"error\": \"Internal Server Error\"}"
// @Router /cart/add-product [post]
func (ctrl CartController) AddProduct(c *gin.Context) {
	payload := &AddProductPayload{}
	err := c.BindJSON(payload)
	if err != nil {
		// TODO melhorar retorno de validação
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.service.AddProduct(c.Request.Context(), &cart.Product{
		ClientID:  payload.ClientID,
		ProductID: payload.ProductID,
		Quantity:  payload.Quantity,
		Comments:  payload.Comments,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.Status(http.StatusOK)
}

type EditProductPayload struct {
	ClientID  uuid.UUID `json:"client_id" binding:"required"`
	ProductID uuid.UUID `json:"product_id" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required"`
	Comments  string    `json:"comments"`
}

// EditProduct edits a Product from Cart by ID
// @Description Edits a Product from Customer's Cart
// @Tags Cart
// @Param payload body controller.EditProductPayload true "EditProductPayload"
// @Accept json
// @Produce json
// @Success 200
// @Failure 500 {object} map[string]any "{\"error\": \"Internal Server Error\"}"
// @Router /cart/edit-product [post]
func (ctrl CartController) EditProduct(c *gin.Context) {
	payload := &EditProductPayload{}
	err := c.BindJSON(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.service.EditProduct(c.Request.Context(), &cart.Product{
		ClientID:  payload.ClientID,
		ProductID: payload.ProductID,
		Quantity:  payload.Quantity,
		Comments:  payload.Comments,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.Status(http.StatusOK)
}

type RemoveProductPayload struct {
	ClientID  uuid.UUID `json:"client_id" binding:"required"`
	ProductID uuid.UUID `json:"product_id" binding:"required"`
}

// RemoveProduct removes a Product from Customer's Cart
// @Description Removes a Product from Customer's Cart
// @Tags Cart
// @Param payload body controller.RemoveProductPayload true "RemoveProductPayload"
// @Accept json
// @Produce json
// @Success 200
// @Failure 500 "Internal Server Error"
// @Router /cart/remove-product [post]
func (ctrl CartController) RemoveProduct(c *gin.Context) {
	payload := &RemoveProductPayload{}
	err := c.BindJSON(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.service.RemoveProduct(c.Request.Context(), payload.ClientID, payload.ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.Status(http.StatusOK)
}
