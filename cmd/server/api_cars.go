package main

import (
	"encoding/json"
	"net/http"

	cars "github.com/craguilar/demo-cars-fn/internal"
	"github.com/gorilla/mux"
)

func AddCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	//TODO
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	plate, ok := vars["carId"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(serializeError(http.StatusBadRequest, "BadRequest"))
		return
	}
	car, error := cars.GetCar(plate)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(serializeError(http.StatusInternalServerError, "Internal server error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(serializeData(car))
}

func ListCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func serializeData(data interface{}) []byte {
	result, error := json.Marshal(data)
	if error != nil {
		return nil
	}
	return result
}

func serializeError(code int, message string) []byte {
	result, error := json.Marshal(cars.ModelError{Code: string(code), Message: message})
	if error != nil {
		return []byte{}
	}
	return result
}
