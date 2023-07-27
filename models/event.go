package models

import (    
    "time"
)

type Event struct {    
	Id            int64     `gorm:"primaryKey" json:"id"`  
    Name          string    `gorm:"type:varchar(100)" json:"name"`
    Description   string    `json:"description"`
    Location      string    `json:"location"`
    EventDate time.Time `json:"event_date"`
    Participant     []Participant `gorm:"foreignKey:EventID"` // Menyimpan relasi dengan participant menggunakan foreign key EventID
}