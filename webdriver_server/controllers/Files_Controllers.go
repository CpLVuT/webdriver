package controllers

import (
	"net/http"
	"strconv"
	"webdriver_server/models"
	"webdriver_server/service"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	ctx         *gin.Context
	fileService *services.FileService
}

func NewFileController(ctx *gin.Context, fileService *services.FileService) *FileController {
	return &FileController{
		ctx:         ctx,
		fileService: fileService,
	}
}

func (c *FileController) Create() {
	var file models.Files
	if err := c.ctx.ShouldBindJSON(&file); err != nil {
		c.ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.fileService.CreateFile(&file); err != nil {
		c.ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (c *FileController) GetByID() {
	ownerIdStr, _ := c.ctx.Params.Get("ownerId")
	ownerId, err := strconv.Atoi(ownerIdStr)
	if err != nil {
		c.ctx.JSON(http.StatusBadRequest, gin.H{"error": "ownerId must be a number."})
		return
	}

	file, err := c.fileService.GetFileByID(uint(ownerId))
	if err != nil {
		c.ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.ctx.JSON(http.StatusOK, file)
}

func (c *FileController) Update() {
	var file models.Files
	if err := c.ctx.ShouldBindJSON(&file); err != nil {
		c.ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.fileService.UpdateFile(&file); err != nil {
		c.ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (c *FileController) Delete() {
	ownerIdStr, _ := c.ctx.Params.Get("ownerId")
	ownerId, err := strconv.Atoi(ownerIdStr)
	if err != nil {
		c.ctx.JSON(http.StatusBadRequest, gin.H{"error": "ownerId must be a number."})
		return
	}

	if err := c.fileService.DeleteFile(uint(ownerId)); err != nil {
		c.ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (c *FileController) List() {
	files, err := c.fileService.GetAllFiles()
	if err != nil {
		c.ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.ctx.JSON(http.StatusOK, files)
}
