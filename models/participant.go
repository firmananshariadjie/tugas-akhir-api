package models

import (    
    "time"
)

type Participant struct {  
    Id          int64     `gorm:"primaryKey" json:"id"`  
    Name        string    `gorm:"type:varchar(100)" json:"name"`
    Email       string    `gorm:"type:varchar(100)" json:"email"`
    Number      string    `gorm:"type:int(13)" json:"number"`
    Status      string    `gorm:"type:varchar(100)" json:"status"` // Status pendaftaran (misalnya: "Terdaftar", "Dikonfirmasi", "Dibatalkan")
    RegisterDate time.Time`json:"register_date"`
    EventID     uint      `json:"event_id"` // Foreign key untuk relasi dengan Acara
}

