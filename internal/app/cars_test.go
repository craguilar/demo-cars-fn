package app

import "testing"

func TestValidation(t *testing.T) {
	car := &Car{
		Description: "This is a failed test",
	}
	err := car.validate()
	if err == nil {
		t.Error("Car is expected to be validated %", err)
	}
}
