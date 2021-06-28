package model

type EventType struct {
	Category string          `json:"eventcategory"`
	Type     [][]interface{} `json:"eventtype"`
}

/*
Event Categories/Types
Ownership Change
	Initial Purchase
	Transfer from Sale
	Repossesion
Compliance
	Emissions
Milage Milestone
	50K
	100K
	150K
	200K
	250K
Maintenance/Servicing
	Replace air filter
	Scheduled maintenance
	Electrical work
	New tires
	Battery replacement
	Brake work
	Fluid added/replaced
	Wheels aligned/balanced
	Other
Damage
	Accident
	Vandalism
	Weather
	Other
Usage Summary
	Self Driving Miles
	Manual Driving Miles
	Average Speed
	Max Speed
	Min Speed
	Speed Violations
	Lbs Towed
Vehicle Alerts
	Air Bags Deployed
	Check Engine Alert
	Battery Alert
	Brake Alert
	Other
*/
