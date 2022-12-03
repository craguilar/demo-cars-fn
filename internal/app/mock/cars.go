package mock

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/craguilar/demo-cars-fn/internal/app"
)

// CarsService represents a mock implementation of internal.CarsService.
type CarsService struct {
	db   map[string]*app.Car
	lock sync.RWMutex
}

func NewCarService() *CarsService {
	return &CarsService{
		db: make(map[string]*app.Car),
	}
}

func (c *CarsService) Car(plate string) (*app.Car, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	// Get the car !
	car, exists := c.db[plate]
	if !exists {
		return nil, nil
	}
	return car, nil
}

func (c *CarsService) Cars() ([]*app.CarSummary, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	list := []*app.CarSummary{}
	for _, car := range c.db {
		list = append(list, &app.CarSummary{Plate: car.Plate, Make: car.Make, Model: car.Model, Description: car.Description, TypeOfUse: car.TypeOfUse})
	}
	return list, nil
}

func (c *CarsService) CreateOrUpdateCar(u *app.Car) (*app.Car, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	err := u.Validate()
	if err != nil {
		return nil, err
	}

	_, exists := c.db[u.Plate]
	if !exists {
		u.TimeCreated = time.Now()
		u.TimeUpdated = time.Now()
		c.db[u.Plate] = u
		return u, nil
	}
	// If it exists update the time stamp and return, we should be more strict about validations but dah!
	u.TimeUpdated = time.Now()
	c.db[u.Plate] = u
	log.Printf("Created car %s", u.Plate)
	return u, nil
}

func (c *CarsService) DeleteCar(plate string) error {
	c.lock.RLock()
	defer c.lock.RUnlock()

	_, exists := c.db[plate]
	if !exists {
		return errors.New("Car doesn't exist!")
	}
	delete(c.db, plate)
	return nil
}
