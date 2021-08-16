package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

func NewController() Controller {
	return Controller{}
}

func (c *Controller) handleHealth(ctx *gin.Context) {
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Hi, you are you?",
	}

	ctx.JSON(http.StatusOK, response)
}
