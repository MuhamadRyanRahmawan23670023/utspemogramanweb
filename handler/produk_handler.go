package handlers

import (
	"Crud/config"
	"Crud/models"

	"github.com/gin-gonic/gin"

	"net/http"
)

func GetProduk(c *gin.Context) {
	var produk []models.Produk
	config.DB.Find(&produk)
	c.JSON(http.StatusOK, produk)
}

func CreateProduk(c *gin.Context) {
	var produk models.Produk
	if err := c.ShouldBindJSON(&produk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&produk)
	c.JSON(http.StatusOK, produk)
}

func UpdateProduk(c *gin.Context) {
	id := c.Param("id")
	var produk models.Produk
	config.DB.First(&produk, id)
	c.BindJSON(&produk)
	config.DB.Save(&produk)
	c.JSON(http.StatusOK, produk)
}

func DeleteProduk(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.Produk{}, id)
	c.JSON(http.StatusOK, gin.H{"deleted": id})
}
