package models

import (    
    "time"    
)

type Participant struct {     
    Id          int64     `gorm:"primaryKey" json:"participant_id"`  
    Name        string    `gorm:"type:varchar(60)" json:"name"`
    Email       string    `gorm:"type:varchar(60)" json:"email"`
    Number      string    `gorm:"type:varchar(13)" json:"number"`
    Status      string    `gorm:"type:varchar(20)" json:"status"` // Status pendaftaran (misalnya: "Terdaftar", "Dikonfirmasi", "Dibatalkan")
    RegisterDate time.Time`json:"register_date" time_format:"2006-01-02 00:00:00" time_utf:"8"`    
    Unique_Code string    `gorm:"type:varchar(60)" json:"unique_code"`
    EventID     uint       
    CreatedAt time.Time `json:"created_at" time_format:"2006-01-02 00:00:00" time_utf:"8"`
    UpdatedAt time.Time `json:"updated_at" time_format:"2006-01-02 00:00:00" time_utf:"8"`
    

}

