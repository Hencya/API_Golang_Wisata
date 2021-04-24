package repository

import (
	"github.com/ydhnwb/golang_api/entity"
	"gorm.io/gorm"
)

//DestinationRepository is a ....
type DestinationRepository interface {
	InsertDestination(b entity.Destination) entity.Destination
	UpdateDestination(b entity.Destination) entity.Destination
	DeleteDestination(b entity.Destination)
	AllDestination() []entity.Destination
	FindDestinationByID(destinationID uint64) entity.Destination
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

func (db *destinationConnection) InsertDestination(b entity.Destination) entity.Destination {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *destinationConnection) UpdateDestination(b entity.Destination) entity.Destination {
	db.connection.Save(&b)
	db.connection.Preload("User").Find(&b)
	return b
}

func (db *destinationConnection) DeleteDestination(b entity.Destination) {
	db.connection.Delete(&b)
}

func (db *destinationConnection) FindDestinationByID(destinationID uint64) entity.Destination {
	var destination entity.Destination
	db.connection.Preload("User").Find(&destination, destinationID)
	return destination
}

func (db *destinationConnection) AllDestination() []entity.Destination {
	var destinations []entity.Destination
	db.connection.Preload("User").Find(&destinations)
	return destinations
}
