package pets

import (
	"errors"
	"fmt"

	"github.com/shubham88fru/degign-patterns-go/models"
)

type AnimalInterface interface {
	Show() string
}

type DogFromFactory struct {
	Pet *models.Dog
}

func (dff *DogFromFactory) Show() string { //DogFromFactory implements AnimalInterface
	return fmt.Sprintf("This anial is a %s", dff.Pet.Breed.Breed)
}

type CatFromFactory struct {
	Pet *models.Cat
}

func (dff *CatFromFactory) Show() string { //CatFromFactory implements AnimalInterface
	return fmt.Sprintf("This anial is a %s", dff.Pet.Breed.Breed)
}

type PetFactoryInterface interface {
	newPet() AnimalInterface
}

type DogAbstractFactory struct{}

func (d *DogAbstractFactory) newPet() AnimalInterface {
	return &DogFromFactory{
		Pet: &models.Dog{
			DogName: "Rex",
			Breed: models.DogBreed{
				Breed: "German Shepherd",
			},
		},
	}
}

type CatAbstractFactory struct{}

func (c *CatAbstractFactory) newPet() AnimalInterface {
	return &CatFromFactory{
		Pet: &models.Cat{
			CatName: "Whiskers",
			Breed: models.CatBreed{
				Breed: "Siamese",
			},
		},
	}
}

func NewPetFromAbstractFactory(species string) (AnimalInterface, error) {
	switch species {
	case "dog":
		return (&DogAbstractFactory{}).newPet(), nil
	case "cat":
		return (&CatAbstractFactory{}).newPet(), nil
	}
	return nil, errors.New("Invalid pet type")
}

func NewPetWithBreedFromAbstractFactory(species, breed string) (AnimalInterface, error) {
	switch species {
	case "dog":
		return &DogFromFactory{}, nil
	case "cat":
		return &CatFromFactory{}, nil
	default:
		return nil, errors.New("Invalid species supplied")
	}
}
