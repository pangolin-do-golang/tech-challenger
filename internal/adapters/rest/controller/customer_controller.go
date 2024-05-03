package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/customer"
)

type CustomerController struct {
	service customer.IService
}

type CustomerPayload struct {
	Name  string `json:"name" validate:"required,min=5,max=20"`
	Cpf   string `json:"cpf" validate:"required,numeric,len=11"`
	Email string `json:"email" validate:"required,email"`
	Age   int    `json:"age" validate:"gte=18,lte=120"`
}

func NewCustomerController(service customer.IService) *CustomerController {
	return &CustomerController{
		service: service,
	}
}

// Create customer godoc
// @Summary Create customer
// @Description Create a new customer
// @Tags Customer
// @Param name body string true "Name example"
// @Param cpf body string true "03985594051"
// @Param email body string true "example@example.com"
// @Param age body int true "18"
// @Accept  json
// @Produce  json
// @Success 200 object customer.Customer "customer.Customer"
// @Failure 400 "Invalid cpf"
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

// Update customer godoc
// @Summary Update customer
// @Description Update a customer by id
// @Tags Customer
// @Param id path uint true "123456789"
// @Param name body string true "Name example"
// @Param cpf body string true "03985594051"
// @Param email body string true "example@example.com"
// @Param age body int true "18"
// @Accept  json
// @Produce  json
// @Success 200 object customer.Customer "customer.Customer"
// @Failure 400 "Invalid identifier informed"
// @Router /customer/:id [put]
func (ctrl CustomerController) Update(c *gin.Context) {
	id := c.Param("id")
	parsedId, err := strconv.ParseUint(id, 10, 64)

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

	updated, err := ctrl.service.Update(uint(parsedId), customer.Customer{
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

// Delete customer godoc
// @Summary Delete customer
// @Description Delete a customer by id
// @Tags Customer
// @Param id path uint true "123456789"
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 "Invalid identifier informed"
// @Router /customer/:id [delete]
func (ctrl CustomerController) Delete(c *gin.Context) {
	id := c.Param("id")
	parsedId, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid identifier informed.",
		})

		return
	}

	if err := ctrl.service.Delete(uint(parsedId)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.Status(http.StatusOK)
}

// Get customer list godoc
// @Summary Get customer list
// @Description Get customer list
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

// Get customer godoc
// @Summary Get customer by cpf
// @Description Get customer by cpf
// @Tags Customer
// @Param cpf path string true "customer cpf"
// @Accept json
// @Produce json
// @Success 200 {object} customer.Customer{}
// @Failure 404 "Customer not found"
// @Router /customer/:cpf [get]
func (ctrl CustomerController) GetByCpf(c *gin.Context) {
	cpf := c.Param("cpf")

	customer, err := ctrl.service.GetByCpf(cpf)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if customer.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"NotFound": "Customer not found.",
		})

		return
	}

	c.JSON(http.StatusOK, customer)
}
