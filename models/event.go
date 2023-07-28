package models

import (    
    "time"
    "gorm.io/gorm"
)

type Event struct {    
    gorm.Model
	// Id            int64     `gorm:"primaryKey" json:"event_id"`  
    Name          string    `gorm:"type:varchar(60)" json:"name"`
    Description   string    `json:"description"`
    Location      string    `json:"location"`
    EventDate     time.Time  `json:"event_date" time_format:"2006-01-02 00:00:00" time_utf:"8"`    
    Capacity      int64     `json:"capacity"`
    Participant_Registered  int64 `json:"participant_registered"`
    Attedence     int64     `json:"attedence"`
    Absent        int64     `json:"absent"`
    Event_Status  string    `gorm:"type:varchar(10)" json:"event_status"`
    Participants   []Participant    //`gorm:"foreignKey:EventRefer"`
}