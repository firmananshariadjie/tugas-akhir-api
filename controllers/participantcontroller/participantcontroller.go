package participantcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/firmananshariadjie/tugas-akhir-api/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Index mengambil daftar semua Participant acara dari database dan mengembalikannya dalam format JSON.
func Index(c *gin.Context) {
	var participants []models.Participant
	models.DB.Find(&participants)
	c.JSON(http.StatusOK, gin.H{"participants": participants})
}

// Show mengambil satu data Participant acara berdasarkan ID yang diberikan dan mengembalikannya dalam format JSON.
func Show(c *gin.Context) {
	var participant models.Participant
	id := c.Param("id")

	if err := models.DB.First(&participant, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"participant": participant})
}

// Create membuat data Participant acara baru berdasarkan data JSON yang diterima dan menyimpannya ke database.
func Create(c *gin.Context) {
	var participant models.Participant

	if err := c.ShouldBindJSON(&participant); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	participant.Status = "Terdaftar"
	participant.Unique_Code = uuid.New().String()
	participant.CreatedAt = time.Now()
	participant.RegisterDate = time.Now()
	
	models.DB.Create(&participant)

	// Mengupdate kolom "participant_registered" di tabel "event"
	// dengan menambahkan 1 pada nilai saat ini untuk event ID tertentu
	eventID := participant.EventID
	models.DB.Model(&models.Event{}).
		Where("id = ?", eventID).
		Update("participant_registered", models.DB.Raw("participant_registered + ?", 1))

	c.JSON(http.StatusOK, gin.H{"participant": participant})
}

// Update mengupdate data Participant acara berdasarkan ID yang diberikan dengan data JSON yang diterima.
func Update(c *gin.Context) {
	var participant models.Participant
	id := c.Param("id")

	if err := c.ShouldBindJSON(&participant); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&participant).Where("id = ?", id).Updates(&participant).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate Participant acara"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})
}

// Delete menghapus data Participant acara berdasarkan ID yang diberikan.
func Delete(c *gin.Context) {
	var participant models.Participant
	eventID := participant.EventID
	models.DB.Model(&models.Event{}).
		Where("id = ?", eventID).
		Update("participant_registered", models.DB.Raw("participant_registered - ?", 1))

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&participant, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus Participant acara"})
		return
	}
	

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
