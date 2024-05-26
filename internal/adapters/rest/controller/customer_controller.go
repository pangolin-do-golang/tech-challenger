package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/customer"
)

type CustomerController struct {
	service customer.IService
}

type CustomerPayload struct {
	Name  string `json:"name" binding:"required,min=5,max=20"`
	Cpf   string `json:"cpf" binding:"required,numeric,len=11"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"gte=18,lte=120"`
}

func NewCustomerController(service customer.IService) *CustomerController {
	return &CustomerController{
		service: service,
	}
}

// Create a new customer
// @Summary Create customer
// @Description Create a new customer
// @Tags Customer
// @Param payload body controller.CustomerPayload true "CustomerPayload"
// @Accept json
// @Produce json
// @Success 200 {object} customer.Customer "customer.Customer"
// @Failure 400 {object} map[string]any "{\"error\": \"Invalid CPF\"}"
// @Router /customer [post]
func (ctrl CustomerController) Create(c *gin.Context) {
	var payload CustomerPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	created, err := ctrl.service.Create(customer.Customer{
		Name:  payload.Name,
		Cpf:   payload.Cpf,
		Email: payload.Email,
		Age:   payload.Age,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, created)
}

// Update a customer by id
// @Summary Update customer
// @Description Update a customer by id
// @Tags Customer
// @Param id path string true "ID of the customer"
// @Param payload body controller.CustomerPayload true "CustomerPayload"
// @Accept json
// @Produce json
// @Success 200 {object} customer.Customer "customer.Customer"
// @Failure 400 {object} map[string]any "{\"error\": \"Invalid CPF\"}"
// @Router /customer/{id} [put]
func (ctrl CustomerController) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid identifier informed.",
		})

		return
	}

	var payload CustomerPayload

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	updated, err := ctrl.service.Update(id, customer.Customer{
		Name:  payload.Name,
		Cpf:   payload.Cpf,
		Email: payload.Email,
		Age:   payload.Age,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, updated)
}

// Delete a customer by id
// @Summary Delete customer
// @Description Delete a customer by id
// @Tags Customer
// @Param id path string true "123456789"
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 "Invalid identifier informed"
// @Router /customer/{id} [delete]
func (ctrl CustomerController) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid identifier informed.",
		})

		return
	}

	if err := ctrl.service.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.Status(http.StatusOK)
}

// GetAll Get all customer's list
// @Summary Get customer list
// @Description Get all customer's list
// @Tags Customer
// @Accept json
// @Produce json
// @Success 200 {object} []customer.Customer{}
// @Router /customer [get]
func (ctrl CustomerController) GetAll(c *gin.Context) {
	customerSlice, err := ctrl.service.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, customerSlice)
}

// GetByCpf Get a customer by cpf
// @Summary Get customer by cpf
// @Description Get a customer by cpf
// @Tags Customer
// @Param cpf path string true "customer cpf"
// @Accept json
// @Produce json
// @Success 200 {object} customer.Customer{}
// @Failure 404 "Customer not found"
// @Router /customer/{cpf} [get]
func (ctrl CustomerController) GetByCpf(c *gin.Context) {
	cpf := c.Param("cpf")

	customerRecord, err := ctrl.service.GetByCpf(cpf)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if customerRecord.Id == uuid.Nil {
		c.JSON(http.StatusNotFound, gin.H{
			"NotFound": "Customer not found.",
		})

		return
	}

	c.JSON(http.StatusOK, customerRecord)
}
