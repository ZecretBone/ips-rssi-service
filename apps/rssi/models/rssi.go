package models

import "time"

type RSSI struct {
	SSID        string    `bson:"ssid"`         //APs SSID
	Strength    float32   `bson:"strength"`     //RSSI signal strength in Dbm.
	PollingRate int       `bson:"polling_rate"` //Polling in ms
	CreatedAt   time.Time `bson:"created_at"`   //RSSI signal capture time
}

type Position struct {
	X float64 `bson:"x"` // X position in meter
	Y float64 `bson:"y"` // Y position in meter
	Z float64 `bson:"z"` // Z position in floor
}

type DeviceInfo struct {
	DeviceID string `bson:"device_id"` //Device ID
	Models   string `bson:"models"`    //Device models eg.Samsung
}
