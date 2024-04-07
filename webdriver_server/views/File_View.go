package views

import (
	"github.com/gin-gonic/gin"
	"net/http"
	controllers "webdriver_server/controllers"
	"webdriver_server/db"
	"webdriver_server/service"
)

func FileList(ctx *gin.Context) {
	fileService := services.NewFileService(databases.DB)
	controller := controllers.NewFileController(ctx, fileService)

	result, pagination, err := controller.List()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":    "success",
		"data":       result,
		"pagination": pagination,
	})
}

func FileCreate(ctx *gin.Context) {
	fileService := services.NewFileService(databases.DB)
	controller := controllers.NewFileController(ctx, fileService)

	err := controller.Create()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func FileGet(ctx *gin.Context) {
	fileService := services.NewFileService(databases.DB)
	controller := controllers.NewFileController(ctx, fileService)

	result, err := controller.Get()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    result,
	})
}

func FileUpdate(ctx *gin.Context) {
	fileService := services.NewFileService(databases.DB)
	controller := controllers.NewFileController(ctx, fileService)

	err := controller.Update()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func FileDelete(ctx *gin.Context) {
	fileService := services.NewFileService(databases.DB)
	controller := controllers.NewFileController(ctx, fileService)

	count, err := controller.Delete()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"count":   count,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"count":   count,
	})
}
