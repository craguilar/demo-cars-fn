package dynamo

import (
	"errors"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/craguilar/demo-cars-fn/internal/app"
)

type DBConfig struct {
	DbService *dynamodb.DynamoDB
	TableName string
}

func InitDb(db *dynamodb.DynamoDB, tableName string) *DBConfig {
	return &DBConfig{
		DbService: db,
		TableName: tableName,
	}
}

// CarsService represents a Dynamo DB implementation of internal.CarsService.
type CarsService struct {
	db *DBConfig
}

func NewCarService(db *DBConfig) *CarsService {
	return &CarsService{
		db: db,
	}
}

func (c *CarsService) Car(plate string) (*app.Car, error) {

	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(plate),
			},
		},
		TableName: &c.db.TableName,
	}

	result, err := c.db.DbService.GetItem(input)
	if err != nil {
		return nil, err

	}

	car := &app.Car{}
	err = dynamodbattribute.UnmarshalMap(result.Item, car)
	if err != nil {
		return nil, err
	}
	if car.Plate == "" {
		return nil, nil
	}
	log.Printf("Return car %v", car)
	return car, nil
}

func (c *CarsService) Cars() ([]*app.CarSummary, error) {
	input := &dynamodb.ScanInput{
		TableName: &c.db.TableName,
	}
	result, err := c.db.DbService.Scan(input)
	if err != nil {
		return nil, err
	}
	items := new([]app.Car)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, items)
	if err != nil {
		return nil, err
	}
	list := []*app.CarSummary{}
	for _, car := range *items {
		list = append(list, &app.CarSummary{Plate: car.Plate, Make: car.Make, Model: car.Model, Description: car.Description, TypeOfUse: car.TypeOfUse})
	}
	return list, nil
}

func (c *CarsService) CreateOrUpdateCar(u *app.Car) (*app.Car, error) {
	err := u.Validate()
	if err != nil {
		return nil, err
	}
	log.Printf("Creating car with plate %s", u.Plate)
	// Check if user exists
	currentCar, err := c.Car(u.Plate)
	if err != nil {
		return nil, err
	}
	if currentCar == nil {
		u.TimeCreated = time.Now()
	}
	// If it exists update the time stamp!
	u.TimeUpdated = time.Now()
	av, err := dynamodbattribute.MarshalMap(u)
	if err != nil {
		return nil, err
	}
	// Assign dynamo db key
	av["id"] = &dynamodb.AttributeValue{S: aws.String(u.Plate)}
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: &c.db.TableName,
	}

	output, err := c.db.DbService.PutItem(input)
	if err != nil {
		return nil, err
	}
	log.Infof("%s", output)
	return u, nil
}

func (c *CarsService) DeleteCar(plate string) error {
	return errors.New("Not implemented")
}
