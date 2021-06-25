package model

import "time"

type VehicleEvent struct {
	Id               int16     `json:"id"`
	Vin              string    `json:"vin"`
	WorkDescription  string    `json:"workdescription"`
	Servicer         string    `json:"servicer"`
	Technician       string    `json:"technician"`
	SelectedFileName string    `json:"selectedfilename"`
	CreatedAt        time.Time `json:"created_at"`
}
