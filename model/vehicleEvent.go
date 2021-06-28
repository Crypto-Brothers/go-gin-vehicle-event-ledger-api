package model

import "time"

type VehicleEvent struct {
	Id             int16     `json:"id"`
	Vin            string    `json:"vin"`
	EventCategory  string    `json:"eventcategory"`
	EventType      string    `json:"eventtype"`
	Description    string    `json:"description"`
	RelateFileName string    `json:"relatedfilename"`
	CreatedAt      time.Time `json:"created_at"`
}
