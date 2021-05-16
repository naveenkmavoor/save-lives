package models

type AlertMssg struct {
	AccidentOccured bool `json:"accident_occured"`
	People          struct {
		Numbers int   `json:"numbers"`
		Age     []int `json:"age"`
	} `json:"people"`
	VehicleID string    `json:"vehicle_id"`
	Location  []float64 `json:"location"`
	Sensor    struct {
		Accelerometer []float64 `json:"accelerometer"`
		Gyroscope     []float64 `json:"gyroscope"`
		Sound         int       `json:"sound"`
	} `json:"sensor"`
}
