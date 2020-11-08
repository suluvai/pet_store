package repositories

import (
	"errors"
	"pet_store_rest_api/models"
	"sync"

	"fmt"
)

type PetRepository struct {
	petsData concurrentPetDataContainer
}

// A concurrent Data container, that can be safely accessed.
// Using in memory storage instead of a database for the assignment purporse.
// Do not access this directly, use the get and update methods.
type concurrentPetDataContainer struct {
	mutex sync.RWMutex
	pets  map[int64]models.Pet // Keyed on Pet Id.
}

var PetRepoData = &PetRepository{
	petsData: concurrentPetDataContainer{
		pets: make(map[int64]models.Pet)}}

// AddNewPet add new pet information to pet data container.
// Retuns nil on success, error if the pet data with corresponding already exists.
func (repository *PetRepository) AddNewPet(newPet models.Pet) error {
	repository.petsData.mutex.Lock()
	defer repository.petsData.mutex.Unlock()

	if _, found := repository.petsData.pets[newPet.Id]; !found {
		repository.petsData.pets[newPet.Id] = newPet
		return nil
	}
	return errors.New("A Pet with given Id already exists")
}

// GetPet reads the pet info from pet data container based on the requested pet id.
// Returns pet info if data for the requested pet id is found, otherwise error.
func (repository *PetRepository) GetPet(petId int64) (interface{}, error) {
	repository.petsData.mutex.Lock()
	defer repository.petsData.mutex.Unlock()

	if pet, found := repository.petsData.pets[petId]; found {
		return pet, nil
	}
	return nil, fmt.Errorf("Pet not found with id: %v", petId)
}

// GetPetByStatus reads the pet info from pet data container based on the requested status list.
// Returns pet info if data for the requested pet id is found, otherwise error.
func (repository *PetRepository) GetPetByStatus(status []string) []models.Pet {
	repository.petsData.mutex.Lock()
	defer repository.petsData.mutex.Unlock()

	// Using a simple O(n) iteration here, typically a db would be used.
	var pets []models.Pet
	for _, val := range repository.petsData.pets {
		for _, v := range status {
			if v == val.Status {
				pets = append(pets, val)
			}
		}
	}
	return pets
}

// UpdatePet updates the existing pet info in the pet data container.
// Returns error if the pet id of the requested pet is not found.
func (repository *PetRepository) UpdatePet(petInfo models.Pet) error {
	repository.petsData.mutex.Lock()
	defer repository.petsData.mutex.Unlock()

	if _, found := repository.petsData.pets[petInfo.Id]; found {
		repository.petsData.pets[petInfo.Id] = petInfo
		return nil
	}
	return fmt.Errorf("Pet not found with id: %v", petInfo.Id)
}

// DeletePet deletes the pet data for the requested petId.
// Returns nil on success, error if pet id not found.
func (repository *PetRepository) DeletePet(petId int64) error {
	repository.petsData.mutex.Lock()
	defer repository.petsData.mutex.Unlock()

	if _, found := repository.petsData.pets[petId]; found {
		delete(repository.petsData.pets, petId)
		return nil
	}
	return fmt.Errorf("Pet not found with id: %v", petId)
}
