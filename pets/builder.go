package pets

import "errors"

type PetInterface interface {
	SetSpecies(s string) *Pet
	SetBreed(s string) *Pet
	SetMinWeight(s int) *Pet
	SetMaxWeight(s int) *Pet
	SetWeight(s int) *Pet
	SetDescription(s string) *Pet
	SetLifeSpan(s int) *Pet
	SetGeographicOrigin(s string) *Pet
	SetColor(s string) *Pet
	SetAge(s int) *Pet
	SetAgeEstimated(s bool) *Pet
	Build() (*Pet, error)
}

func NewPetBuilder() PetInterface {
	return &Pet{}
}

func (p *Pet) SetSpecies(s string) *Pet {
	p.Species = s
	return p
}

func (p *Pet) SetBreed(b string) *Pet {
	p.Breed = b
	return p
}

func (p *Pet) SetMinWeight(w int) *Pet {
	p.MinWeight = w
	return p
}

func (p *Pet) SetMaxWeight(w int) *Pet {
	p.MaxWeight = w
	return p
}

func (p *Pet) SetWeight(w int) *Pet {
	p.Weight = w
	return p
}

func (p *Pet) SetDescription(d string) *Pet {
	p.Description = d
	return p
}

func (p *Pet) SetLifeSpan(l int) *Pet {
	p.LifeSpan = l
	return p
}

func (p *Pet) SetGeographicOrigin(g string) *Pet {
	p.GeographicOrigin = g
	return p
}

func (p *Pet) SetColor(c string) *Pet {
	p.Color = c
	return p
}

func (p *Pet) SetAge(a int) *Pet {
	p.Age = a
	return p
}

func (p *Pet) SetAgeEstimated(a bool) *Pet {
	p.AgeEstimated = a
	return p
}

func (p *Pet) Build() (*Pet, error) {
	if p.MinWeight < p.MaxWeight {
		return nil, errors.New("MinWeight must be less than MaxWeight")
	}

	p.AverageWeight = (p.MinWeight + p.MaxWeight) / 2
	return p, nil
}
