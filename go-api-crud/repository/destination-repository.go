package repository

import (
	"github.com/hencya/go-api-crud/models"
	"gorm.io/gorm"
)

//DestinationRepository is a ....
type DestinationRepository interface {
	InsertDestination(b models.Destination) models.Destination
	UpdateDestination(b models.Destination) models.Destination
	DeleteDestination(b models.Destination)
	AllDestination() []models.Destination
	FindDestinationByID(destinationID uint64) models.Destination
}

type destinationConnection struct {
	connection *gorm.DB
}

//NewDestinationRepository creates an instance DestinationRepository
func NewDestinationRepository(dbConn *gorm.DB) DestinationRepository {
	return &destinationConnection{
		connection: dbConn,
	}
}

func (db *destinationConnection) InsertDestination(b models.Destination) models.Destination {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *destinationConnection) UpdateDestination(b models.Destination) models.Destination {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *destinationConnection) DeleteDestination(b models.Destination) {
	db.connection.Delete(&b)
}

func (db *destinationConnection) FindDestinationByID(destinationID uint64) models.Destination {
	var destination models.Destination
	db.connection.Preload("User").Find(&destination, destinationID)
	return destination
}

func (db *destinationConnection) AllDestination() []models.Destination {
	var destinations []models.Destination
	db.connection.Preload("User").Find(&destinations)
	return destinations
}
