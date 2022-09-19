package app

import (
	"time"
)

// Interface for Cars service
type CarsService interface {
	Car(plate string) (*Car, error)
	Cars() ([]*CarSummary, error)
	CreateOrUpdateCar(u *Car) (*Car, error)
	DeleteCar(plate string) error
}

// Main car object .
type Car struct {

	// Car plate
	Plate string `json:"plate"`

	Make string `json:"make"`

	Model string `json:"model"`

	Description string `json:"description,omitempty"`

	TypeOfUse string `json:"typeOfUse,omitempty"`

	NationalKey string `json:"nationalKey,omitempty"`

	SerialNumber string `json:"serialNumber,omitempty"`

	EngineSerialNumber string `json:"engineSerialNumber,omitempty"`

	Company string `json:"company,omitempty"`

	TimeCreated time.Time `json:"timeCreated,omitempty"`

	TimeUpdated time.Time `json:"timeUpdated,omitempty"`

	Images []Image `json:"images,omitempty"`
}

// Main CarSummary object .
type CarSummary struct {

	// Car plate
	Plate string `json:"plate"`

	Make string `json:"make"`

	Model string `json:"model"`

	Description string `json:"description,omitempty"`

	TypeOfUse string `json:"typeOfUse,omitempty"`
}

// Object that holds an image definition
type Image struct {
	Name string `json:"name"`

	Content []byte `json:"content"`
}
