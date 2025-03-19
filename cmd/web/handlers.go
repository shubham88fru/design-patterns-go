package main

import (
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/shubham88fru/degign-patterns-go/models"
	"github.com/shubham88fru/degign-patterns-go/pets"
	"github.com/tsawler/toolbox"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	app.Render(w, "home.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
	page := chi.URLParam(r, "page")
	app.Render(w, page+".page.gohtml", nil)
}

func (app *application) CreateDogFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("dog"))
}

func (app *application) CreateCatFromFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	_ = t.WriteJSON(w, http.StatusOK, pets.NewPet("cat"))
}

func (app *application) TestPatterns(w http.ResponseWriter, r *http.Request) {
	app.Render(w, "test.page.gohtml", nil)
}

func (app *application) CreateDogFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dog, err := pets.NewPetFromAbstractFactory("dog")
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, dog)

}

func (app *application) CreateCatFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	cat, err := pets.NewPetFromAbstractFactory("cat")
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, cat)

}

func (app *application) GetAllDogBreedsJSON(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dogBreeds, err := app.App.Models.DogBreed.All()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_ = t.WriteJSON(w, http.StatusOK, dogBreeds)
}

func (app *application) CreateDogWithBuilder(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	dogBuilder := pets.NewPetBuilder()
	dog, err := dogBuilder.SetAge(5).
		SetBreed("Golden Retriever").
		SetColor("Golden").
		SetDescription("A friendly dog").
		SetLifeSpan(12).
		SetMaxWeight(75).
		SetMinWeight(55).
		SetSpecies("Dog").
		SetWeight(65).
		Build()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	_ = t.WriteJSON(w, http.StatusOK, dog)
}

func (app *application) GetAllCatBreeds(w http.ResponseWriter, r *http.Request) {
	var t toolbox.Tools
	catBreeds, err := app.App.CatService.GetAllBreeds()
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
	}

	_ = t.WriteJSON(w, http.StatusOK, catBreeds)

}

func (app *application) AnimalFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	//Steup toolbox.
	var t toolbox.Tools

	//Get species from URL itself.
	species := chi.URLParam(r, "species")

	//Get breed from the URL.
	b := chi.URLParam(r, "breed")
	breed, _ := url.QueryUnescape(b)

	//Create a pet from abstract factory.
	pet, err := pets.NewPetWithBreedFromAbstractFactory(species, breed)
	if err != nil {
		_ = t.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	//Write the result as JSON.
	_ = t.WriteJSON(w, http.StatusOK, pet)
}

func (app *application) DogOfMonth(w http.ResponseWriter, r *http.Request) {
	// Get the breed
	breed, _ := app.App.Models.DogBreed.GetBreedByName("German Shepherd Dog")

	//Get the dog of the month
	dom, _ := app.App.Models.Dog.GetDogOfMonthByID(1)

	layout := "2006-01-02"
	dob, _ := time.Parse(layout, "2023-11-01")

	//Create the do and decorate
	dog := models.DogOfMonth{
		Dog: &models.Dog{
			ID:                1,
			DogName:           "Sam",
			BreedID:           breed.ID,
			Color:             "Black and Tan",
			DateOfBirth:       dob,
			SprayedOrNeutered: 0,
			Description:       "A loyal and friendly dog",
			Weight:            20,
			Breed:             *breed,
		},
		Video: dom.Video,
		Image: dom.Image,
	}

	//Serve the webpage
	data := make(map[string]any)
	data["dog"] = dog

	app.Render(w, "dog-of-month.page.gohtml", &templateData{Data: data})
}
