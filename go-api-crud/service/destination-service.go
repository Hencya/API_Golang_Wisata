package service

import (
	"fmt"
	"log"

	"github.com/hencya/go-api-crud/dto"
	"github.com/hencya/go-api-crud/models"
	"github.com/hencya/go-api-crud/repository"
	"github.com/mashingan/smapping"
)

//DestinationService is a ....
type DestinationService interface {
	Insert(b dto.DestinationCreateDTO) models.Destination
	Update(b dto.DestinationUpdateDTO) models.Destination
	Delete(b models.Destination)
	All() []models.Destination
	FindByID(destinationID uint64) models.Destination
	IsAllowedToEdit(userID string, destinationID uint64) bool
}

type destinationService struct {
	destinationRepository repository.DestinationRepository
}

//NewDestinationService .....
func NewDestinationService(destinationRepo repository.DestinationRepository) DestinationService {
	return &destinationService{
		destinationRepository: destinationRepo,
	}
}

func (service *destinationService) Insert(b dto.DestinationCreateDTO) models.Destination {
	destination := models.Destination{}
	err := smapping.FillStruct(&destination, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.destinationRepository.InsertDestination(destination)
	return res
}

func (service *destinationService) Update(b dto.DestinationUpdateDTO) models.Destination {
	destination := models.Destination{}
	err := smapping.FillStruct(&destination, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.destinationRepository.UpdateDestination(destination)
	return res
}

func (service *destinationService) Delete(b models.Destination) {
	service.destinationRepository.DeleteDestination(b)
}

func (service *destinationService) All() []models.Destination {
	return service.destinationRepository.AllDestination()
}

func (service *destinationService) FindByID(destinationID uint64) models.Destination {
	return service.destinationRepository.FindDestinationByID(destinationID)
}

func (service *destinationService) IsAllowedToEdit(userID string, destinationID uint64) bool {
	b := service.destinationRepository.FindDestinationByID(destinationID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
