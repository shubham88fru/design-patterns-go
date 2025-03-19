package models

import "database/sql"

type Repository interface {
	AllDogBreeds() ([]*DogBreed, error)
	GetBreedByName(name string) (*DogBreed, error)
	GetDogOfMonthByID(id int) (*DogOfMonth, error)
}

type mysqlRepository struct {
	DB *sql.DB
}

func newMySqlRepository(db *sql.DB) Repository {
	return &mysqlRepository{
		DB: db,
	}
}

type testRepository struct {
	DB *sql.DB
}

func newTestRepository(db *sql.DB) Repository {
	return &testRepository{
		DB: nil,
	}
}
