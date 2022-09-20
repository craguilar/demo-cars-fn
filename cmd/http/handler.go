package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/craguilar/demo-cars-fn/internal/app"
	"github.com/gorilla/mux"
)

type CarsServiceHandler struct {
	carService app.CarsService
}

func NewCarServiceHandler(service app.CarsService) *CarsServiceHandler {
	return &CarsServiceHandler{
		carService: service,
	}
}

func (c *CarsServiceHandler) AddCar(w http.ResponseWriter, r *http.Request) {
	var car app.Car

	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(SerializeError(http.StatusBadRequest, "Invalid Body parameter"))
		return
	}
	createdCar, err := c.carService.CreateOrUpdateCar(&car)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(SerializeData(createdCar))
}

func (c *CarsServiceHandler) GetCar(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	plate, ok := vars["carId"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(SerializeError(http.StatusBadRequest, "BadRequest"))
		return
	}
	car, err := c.carService.Car(plate)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	if car == nil {
		writeError(w, http.StatusNotFound, nil)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(SerializeData(car))
}

func (c *CarsServiceHandler) ListCars(w http.ResponseWriter, r *http.Request) {
	cars, err := c.carService.Cars()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(SerializeData(cars))
}

func (c *CarsServiceHandler) UpdateCar(w http.ResponseWriter, r *http.Request) {
	c.AddCar(w, r)
}

// Write HTTP Error out to http.ResponseWriter
func writeError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	var errorCode string
	if statusCode == 400 {
		errorCode = "InvalidParameter."
	} else if statusCode == 404 {
		errorCode = "NotFound or caller don't have access."
	} else if statusCode == 401 {
		errorCode = "Unauthorized"
	} else if statusCode == 409 {
		errorCode = "Conflict with resource"
	} else if statusCode == 500 {
		errorCode = "InternalServerError"
	}
	log.Printf("Received error from call %s", err)
	w.Write(SerializeError(statusCode, errorCode))

}
