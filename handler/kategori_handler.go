package handlers

import (
	"Crud/config"
	"Crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetKategori(c *gin.Context) {
	var kategori []models.Kategori
	config.DB.Preload("Produk").Find(&kategori)
	c.JSON(http.StatusOK, kategori)
}

func CreateKategori(c *gin.Context) {
	var kategori models.Kategori
	if err := c.ShouldBindJSON(&kategori); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&kategori)
	c.JSON(http.StatusOK, kategori)
}

func UpdateKategori(c *gin.Context) {
	id := c.Param("id")
	var kategori models.Kategori

	// Cek apakah kategori ada di database berdasarkan ID
	if err := config.DB.First(&kategori, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Kategori tidak ditemukan"})
		return
	}

	// Bind data JSON dari request body
	if err := c.ShouldBindJSON(&kategori); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak valid"})
		return
	}

	// Update kategori
	config.DB.Save(&kategori)

	// Kembalikan kategori yang telah diperbarui
	c.JSON(http.StatusOK, kategori)
}

func DeleteKategori(c *gin.Context) {
	id := c.Param("id")

	// 1. Hapus semua produk yang memiliki kategori_id yang sama
	if err := config.DB.Where("kategori_id = ?", id).Delete(&models.Produk{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus produk terkait: " + err.Error()})
		return
	}

	// 2. Hapus kategori
	if err := config.DB.Delete(&models.Kategori{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus kategori: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kategori dan produk terkait berhasil dihapus"})
}
