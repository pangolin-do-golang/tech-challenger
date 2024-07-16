package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge/internal/core/product"
	"net/http"
)

type ProductController struct {
	service *product.Service
}

func NewProductController(service *product.Service) *ProductController {
	return &ProductController{
		service: service,
	}
}

// Search for products
// @Summary Search products with given criteria
// @Tags Product
// @Param search path string false "Name of Product"
// @Param category path string false "Category of Product"
// @Accept json
// @Produce json
// @Success 200 {object} product.Product "product.Product"
// @Failure 500 {object} map[string]any "{\"error\": \"something went bad :(\"}"
// @Router /product [get]
func (ctrl *ProductController) Search(c *gin.Context) {
	search := c.Query("search")
	category := c.Query("category")

	products, err := ctrl.service.Search(search, category)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went bad :(",
		})

		return
	}

	c.JSON(http.StatusOK, products)
}

// Delete a Product
// @Summary Delete a Product with given ID
// @Tags Product
// @Param id path string true "ID of Product"
// @Accept json
// @Produce json
// @Success 204
// @Failure 500 {object} map[string]any "{\"error\": \"something went bad :(\"}"
// @Router /product/{id} [delete]
func (ctrl *ProductController) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid identifier informed.",
		})

		return
	}

	err = ctrl.service.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "something went bad :(",
		})

		return
	}

	c.JSON(http.StatusNoContent, gin.H{})

}
