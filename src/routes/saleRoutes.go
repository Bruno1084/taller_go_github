package routes

import (
	"github.com/gin-gonic/gin"
	"go_test/src/domain"
	"go_test/src/services"
	"net/http"
)

func SaleSetup(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/sales", GetSales)
		api.GET("/sales/:id", GetSaleById)
		api.POST("/sales", CreateSale)
		api.PATCH("/sales/:id", PatchSale)
	}
}

func GetSales(c *gin.Context) {
	userID := c.Query("user_id")
	status := c.Query("status")

	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id es requerido"})
		return
	}

	if status != "" && status != "pending" && status != "approved" && status != "rejected" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status inválido"})
		return
	}

	sales, metadata, err := services.SearchSales(userID, status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sales":    sales,
		"metadata": metadata,
	})
}

func GetSaleById(c *gin.Context) {
	id := c.Param("id")
	sale, err := services.GetSaleById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, sale)
}

func CreateSale(c *gin.Context) {
	var sale domain.Sale
	if err := c.ShouldBindJSON(&sale); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newSale, err := services.CreateSale(&sale)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newSale)
}

func PatchSale(c *gin.Context) {
	id := c.Param("id")
	var sale domain.SaleUpdateFields
	if err := c.ShouldBindJSON(&sale); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedSale, err := services.PatchSale(id, &sale)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedSale)
}
