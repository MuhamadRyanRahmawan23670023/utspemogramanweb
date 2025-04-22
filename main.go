package main

import (
	"Crud/config"
	handlers "Crud/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config.ConnectDatabase()

	// Routes untuk kategori
	r.GET("/kategori", handlers.GetKategori)
	r.POST("/kategori", handlers.CreateKategori)
	r.PUT("/kategori/:id", handlers.UpdateKategori)
	r.DELETE("/kategori/:id", handlers.DeleteKategori) // hanya satu route untuk DELETE kategori

	// Routes untuk produk
	r.GET("/produk", handlers.GetProduk)
	r.POST("/produk", handlers.CreateProduk)
	r.PUT("/produk/:id", handlers.UpdateProduk)
	r.DELETE("/produk/:id", handlers.DeleteProduk)

	// Serve static files
	r.Static("/public", "./public")
	r.LoadHTMLFiles("public/index.html")

	// Route untuk halaman utama
	r.GET("/", func(c *gin.Context) {
		c.File("./public/index.html")
	})

	// Menjalankan server
	r.Run(":8080")
}
