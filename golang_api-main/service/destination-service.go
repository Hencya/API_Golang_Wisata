package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/ydhnwb/golang_api/dto"
	"github.com/ydhnwb/golang_api/entity"
	"github.com/ydhnwb/golang_api/repository"
)

//DestinationService is a ....
type DestinationService interface {
	Insert(b dto.DestinationCreateDTO) entity.Destination
	Update(b dto.DestinationUpdateDTO) entity.Destination
	Delete(b entity.Destination)
	All() []entity.Destination
	FindByID(destinationID uint64) entity.Destination
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

func (service *destinationService) Insert(b dto.DestinationCreateDTO) entity.Destination {
	destination := entity.Destination{}
	err := smapping.FillStruct(&destination, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.destinationRepository.InsertDestination(destination)
	return res
}

func (service *destinationService) Update(b dto.DestinationUpdateDTO) entity.Destination {
	destination := entity.Destination{}
	err := smapping.FillStruct(&destination, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.destinationRepository.UpdateDestination(destination)
	return res
}

func (service *destinationService) Delete(b entity.Destination) {
	service.destinationRepository.DeleteDestination(b)
}

func (service *destinationService) All() []entity.Destination {
	return service.destinationRepository.AllDestination()
}

func (service *destinationService) FindByID(destinationID uint64) entity.Destination {
	return service.destinationRepository.FindDestinationByID(destinationID)
}

func (service *destinationService) IsAllowedToEdit(userID string, destinationID uint64) bool {
	b := service.destinationRepository.FindDestinationByID(destinationID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
