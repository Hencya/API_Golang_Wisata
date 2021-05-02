package service

import (
	"fmt"
	"log"

	"github.com/hencya/go-api-crud/dto"
	"github.com/hencya/go-api-crud/models"
	"github.com/hencya/go-api-crud/repository"
	"github.com/mashingan/smapping"
)

//DestinationService is a of what DestinationService can do
type DestinationService interface {
	Insert(d dto.DestinationCreateDTO) models.Destination
	Update(d dto.DestinationUpdateDTO) models.Destination
	Delete(d models.Destination)
	All() []models.Destination
	FindByID(destinationID uint64) models.Destination
	IsAllowedToEdit(userID string, destinationID uint64) bool
}

type destinationService struct {
	destinationRepository repository.DestinationRepository
}

//NewDestinationService method is creates a new instance of DestinationService
func NewDestinationService(destinationRepo repository.DestinationRepository) DestinationService {
	return &destinationService{
		destinationRepository: destinationRepo,
	}
}

func (service *destinationService) Insert(d dto.DestinationCreateDTO) models.Destination {
	destination := models.Destination{}
	err := smapping.FillStruct(&destination, smapping.MapFields(&d))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.destinationRepository.InsertDestination(destination)
	return res
}

func (service *destinationService) Update(d dto.DestinationUpdateDTO) models.Destination {
	destination := models.Destination{}
	err := smapping.FillStruct(&destination, smapping.MapFields(&d))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.destinationRepository.UpdateDestination(destination)
	return res
}

func (service *destinationService) Delete(d models.Destination) {
	service.destinationRepository.DeleteDestination(d)
}

func (service *destinationService) All() []models.Destination {
	return service.destinationRepository.AllDestination()
}

func (service *destinationService) FindByID(destinationID uint64) models.Destination {
	return service.destinationRepository.FindDestinationByID(destinationID)
}

func (service *destinationService) IsAllowedToEdit(userID string, destinationID uint64) bool {
	d := service.destinationRepository.FindDestinationByID(destinationID)
	id := fmt.Sprintf("%v", d.UserID)
	return userID == id
}
