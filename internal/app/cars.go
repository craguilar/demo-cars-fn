package app

import (
	"time"

	"github.com/go-playground/validator/v10"
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
	Plate string `json:"plate" validate:"required"`

	Make string `json:"make" validate:"required"`

	Model string `json:"model" validate:"required"`

	Description string `json:"description,omitempty"`

	TypeOfUse string `json:"typeOfUse" validate:"required"`

	NationalKey string `json:"nationalKey,omitempty"`

	SerialNumber string `json:"serialNumber,omitempty"`

	EngineSerialNumber string `json:"engineSerialNumber,omitempty"`

	Company string `json:"company,omitempty"`

	TimeCreated time.Time `json:"timeCreated,omitempty"`

	TimeUpdated time.Time `json:"timeUpdated,omitempty"`

	Images []Image `json:"images,omitempty"`

	v *validator.Validate
}

func (c *Car) Validate() error {
	if c.v == nil {
		c.v = validator.New()
	}
	return c.v.Struct(c)
}

// Main CarSummary object .
type CarSummary struct {

	// Car plate
	Plate string `json:"plate" validate:"required"`

	Make string `json:"make" validate:"required"`

	Model string `json:"model" validate:"required"`

	Description string `json:"description,omitempty"`

	TypeOfUse string `json:"typeOfUse" validate:"required"`
}

// Object that holds an image definition
type Image struct {
	Name string `json:"name"`

	Content []byte `json:"content"`
}
