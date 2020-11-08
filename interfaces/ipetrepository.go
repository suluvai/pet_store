package interfaces

import (
	"pet_store_rest_api/models"
)

type IPetRepository interface {
	AddNewPet(newPet models.Pet) error
	GetPet(petId int64) (interface{}, error)
	GetPetByStatus(status []string) []models.Pet
	UpdatePet(petInfo models.Pet) error
	DeletePet(petId int64) error
}
