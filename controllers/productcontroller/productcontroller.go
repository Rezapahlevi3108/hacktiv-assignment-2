package productcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/Rezapahlevi3108/hacktiv-assignment-2/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var items []models.Item

	models.DB.Find(&items)
	c.JSON(http.StatusOK, gin.H{"items": items})

}

func Show(c *gin.Context) {
	var item models.Item
	item_id := c.Param("item_id")

	if err := models.DB.First(&item, item_id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"item": item})
}

func Create(c *gin.Context) {

	var item models.Item

	if err := c.ShouldBindJSON(&item); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&item)
	c.JSON(http.StatusOK, gin.H{"item": item})
}

func Update(c *gin.Context) {
	var item models.Item
	item_id := c.Param("item_id")

	if err := c.ShouldBindJSON(&item); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&item).Where("item_id = ?", item_id).Updates(&item).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

}

func Delete(c *gin.Context) {

	var item models.Item

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	item_id, _ := input.Id.Int64()
	if models.DB.Delete(&item, item_id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
