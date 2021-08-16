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

func (c *Controller) handleListOpenMarket(ctx *gin.Context) {
	response := struct {
		Message string `json:"message"`
	}{
		Message: "List",
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) handleCreateOpenMarket(ctx *gin.Context) {
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Create",
	}

	ctx.JSON(http.StatusCreated, response)
}

func (c *Controller) handleGetOpenMarket(ctx *gin.Context) {
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Get",
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) handleUpdateOpenMarket(ctx *gin.Context) {
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Update",
	}

	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) handleDeleteOpenMarket(ctx *gin.Context) {
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Delete",
	}

	ctx.JSON(http.StatusNoContent, response)
}
