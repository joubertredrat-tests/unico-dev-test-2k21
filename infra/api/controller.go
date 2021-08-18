package api

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	domainService "github.com/joubertredrat-tests/unico-dev-test-2k21/domain/fair/service"
	"github.com/joubertredrat-tests/unico-dev-test-2k21/infra/domain/fair/repository"
	"github.com/joubertredrat-tests/unico-dev-test-2k21/infra/mysql"
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
	var request OpenMarketListSearchCriteria
	ctx.ShouldBindQuery(&request)

	db, err := mysql.NewMysqlConnection(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DBNAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
		return
	}

	searchCriteria := createOpenMarkeSearchCriteriatFromListRequest(request)

	openMarketRepositoryMysql := repository.NewOpenMarketRepositoryMysql(db)
	openMarketService := domainService.NewOpenMarketService(openMarketRepositoryMysql)
	openMarketList, _ := openMarketService.GetListByCriteria(searchCriteria)

	response := createResponseFromOpenMarketList(openMarketList)
	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) handleCreateOpenMarket(ctx *gin.Context) {
	var request OpenMarketCreateRequest
	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint(err)})
		return
	}

	if err := validator.New().Struct(request); err != nil {
		errors := []string{}
		for _, fieldErr := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprint(fieldErr))

		}
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	openMarket := createOpenMarketFromCreateRequest(request)
	db, err := mysql.NewMysqlConnection(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DBNAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
		return
	}

	openMarketRepositoryMysql := repository.NewOpenMarketRepositoryMysql(db)
	openMarketService := domainService.NewOpenMarketService(openMarketRepositoryMysql)
	openMarketCreated, err := openMarketService.Create(openMarket)
	if err != nil {
		if errors.Is(err, domainService.OpenMarketServiceAlreadyExistError) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": fmt.Sprint(err)})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
		return
	}

	response := createResponseFromOpenMarket(*openMarketCreated)
	ctx.JSON(http.StatusCreated, response)
}

func (c *Controller) handleGetOpenMarket(ctx *gin.Context) {
	registryID := ctx.Param("id")
	if registryID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "registry id required"})
		return
	}

	db, err := mysql.NewMysqlConnection(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DBNAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
		return
	}

	openMarketRepositoryMysql := repository.NewOpenMarketRepositoryMysql(db)
	openMarketService := domainService.NewOpenMarketService(openMarketRepositoryMysql)
	openMarketFound, err := openMarketService.GetByRegistryID(registryID)
	if err != nil {
		if errors.Is(err, domainService.OpenMarketServiceNotFoundError) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprint(err)})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
		return
	}

	response := createResponseFromOpenMarket(*openMarketFound)
	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) handleUpdateOpenMarket(ctx *gin.Context) {
	registryID := ctx.Param("id")
	if registryID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "registry id required"})
		return
	}

	var request OpenMarketUpdateRequest
	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprint(err)})
		return
	}

	openMarket := createOpenMarketFromUpdateRequest(request)
	openMarket.RegistryID = registryID
	db, err := mysql.NewMysqlConnection(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DBNAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
		return
	}

	openMarketRepositoryMysql := repository.NewOpenMarketRepositoryMysql(db)
	openMarketService := domainService.NewOpenMarketService(openMarketRepositoryMysql)
	openMarketUpdated, err := openMarketService.Update(openMarket)
	if err != nil {
		if errors.Is(err, domainService.OpenMarketServiceNotFoundError) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprint(err)})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
		return
	}

	response := createResponseFromOpenMarket(*openMarketUpdated)
	ctx.JSON(http.StatusOK, response)
}

func (c *Controller) handleDeleteOpenMarket(ctx *gin.Context) {
	registryID := ctx.Param("id")
	if registryID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "registry id required"})
		return
	}

	db, err := mysql.NewMysqlConnection(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DBNAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
		return
	}

	openMarketRepositoryMysql := repository.NewOpenMarketRepositoryMysql(db)
	openMarketService := domainService.NewOpenMarketService(openMarketRepositoryMysql)

	if err := openMarketService.Delete(registryID); err != nil {
		if errors.Is(err, domainService.OpenMarketServiceNotFoundError) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprint(err)})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprint(err)})
		return
	}

	ctx.JSON(http.StatusNoContent, struct{}{})
}
