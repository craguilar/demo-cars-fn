package internal

func GetCar(plate string) (Car, error) {
	return Car{
		Plate: plate,
		Make:  "Audi",
		Model: "A3",
	}, nil
}

func ListCars() []Car {
	return []Car{}
}

func AddOrUpdateCar(car *Car) (Car, error) {
	return Car{
		Plate: "GLD123",
		Make:  "Audi",
		Model: "A3",
	}, nil
}
