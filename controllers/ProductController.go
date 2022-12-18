package controllers

import (
	"example/todo_go/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
	Store(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type productcontroller struct{}

func NewProductController() ProductController {
	return &productcontroller{}
}

func Index(ctx *gin.Context) {
	var product []models.Product
	models.DB.Find(&product)
	ctx.JSON(http.StatusOK, gin.H{"data": product})
}

func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	if err := models.DB.First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"data":    nil,
				"message": "Data not fount",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})

}

func Store(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Terjadi sesuatu yang salah",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    product,
		"message": "Data Product " + product.Name + " Berhasil diperbaharui",
	})

}

func Delete(c *gin.Context) {
	var product models.Product
	requestId := c.Param("id")
	id, _ := strconv.ParseInt(requestId, 10, 64)
	if models.DB.Delete(product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Id tidak ditemukan !",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Data berhasil dihapus " + product.Name,
	})
}
