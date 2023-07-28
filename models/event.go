package models

import (    
    "time"
)

type Event struct {    
	Id            int64     `gorm:"primaryKey" json:"event_id"`  
    Name          string    `gorm:"type:varchar(60)" json:"name"`
    Description   string    `json:"description"`
    Location      string    `json:"location"`
    EventDate   *time.Time  `json:"event_date"`    
    Capacity      int64     `json:"capacity"`
    Participant_Registered  int64 `json:"participant_registered"`
    Attedence     int64     `json:"attedence"`
    Absent        int64     `json:"absent"`
    Event_Status  string    `gorm:"type:varchar(10)" json:"event_status"`
    Participant   []Participant `gorm:"foreignKey:Event_id"`

    
}