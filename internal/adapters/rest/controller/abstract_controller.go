package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/pangolin-do-golang/tech-challenge/internal/domainerrors"
)

type AbstractController struct{}

type HTTPError struct {
	Message string `json:"error"`
}

func (ctrl *AbstractController) Error(c *gin.Context, err error) {
	var e *domainerrors.Error
	if domainerrors.As(err, &e) {
		switch e.Type {
		case "BUSINESS":
			c.JSON(422, &HTTPError{Message: e.Message})
		case "INPUT":
			c.JSON(400, &HTTPError{Message: e.Message})
		default:
			c.JSON(500, &HTTPError{Message: "Internal Server Error"})
		}
	} else {
		c.JSON(500, &HTTPError{Message: "Internal Server Error"})
	}
}
