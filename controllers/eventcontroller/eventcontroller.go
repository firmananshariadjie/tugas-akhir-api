package eventcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/firmananshariadjie/tugas-akhir-api/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// Index mengambil daftar semua Event dari database dan mengembalikannya dalam format JSON.
func Index(c *gin.Context) {
	var events []models.Event
	models.DB.Find(&events)
	c.JSON(http.StatusOK, gin.H{"events": events})
}

// Show mengambil satu data Event berdasarkan ID yang diberikan dan mengembalikannya dalam format JSON.
func Show(c *gin.Context) {
	var event models.Event
	id := c.Param("id")

	if err := models.DB.First(&event, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"event": event})
}

// Create membuat data Event baru berdasarkan data JSON yang diterima dan menyimpannya ke database.
func Create(c *gin.Context) {
	var event models.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&event)
	c.JSON(http.StatusOK, gin.H{"event": event})
}

// Update mengupdate data Event berdasarkan ID yang diberikan dengan data JSON yang diterima.
func Update(c *gin.Context) {
	var event models.Event
	id := c.Param("id")

	if err := c.ShouldBindJSON(&event); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&event).Where("id = ?", id).Updates(&event).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate Event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

// Delete menghapus data Event berdasarkan ID yang diberikan.
func Delete(c *gin.Context) {
	var event models.Event

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&event, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus Event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
