package repository

import (
	"github.com/hencya/go-api-crud/models"
	"gorm.io/gorm"
)

//DestinationRepository is a  contract what DestinationRepository can do to db
type DestinationRepository interface {
	InsertDestination(d models.Destination) models.Destination
	UpdateDestination(d models.Destination) models.Destination
	DeleteDestination(d models.Destination)
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

func (db *destinationConnection) InsertDestination(d models.Destination) models.Destination {
	db.connection.Save(&d)
	db.connection.Preload("User").Find(&d)
	return d
}

func (db *destinationConnection) UpdateDestination(d models.Destination) models.Destination {
	db.connection.Save(&d)
	db.connection.Preload("User").Find(&d)
	return d
}

func (db *destinationConnection) DeleteDestination(d models.Destination) {
	db.connection.Delete(&d)
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
