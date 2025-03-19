package models

import (
	"context"
	"log"
	"time"
)

func (d *mysqlRepository) AllDogBreeds() ([]*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `select id, breed, weight_low_lbs, weight_high_lbs, 
		cast(((weight_low_lbs + weight_high_lbs) / 2) as unsigned) as average_weight, 
		lifespan, coalesce(details, ''), coalesce(alternate_names, ''), geographic_origin from dog_breeds`

	var breeds []*DogBreed
	rows, err := d.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var b DogBreed
		err := rows.Scan(
			&b.ID,
			&b.Breed,
			&b.WeigthLowLbs,
			&b.WeigthHighLbs,
			&b.AverageWeight,
			&b.LifeSpan,
			&b.Details,
			&b.AlternateNames,
			&b.GeographicOrigin,
		)

		if err != nil {
			return nil, err
		}

		breeds = append(breeds, &b)
	}

	return breeds, nil
}

func (m *mysqlRepository) GetBreedByName(b string) (*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `select id, breed, weight_low_lbs, weight_high_lbs, 
		cast(((weight_low_lbs + weight_high_lbs) / 2) as unsigned) as average_weight, 
		lifespan, coalesce(details, ''), coalesce(alternate_names, ''), geographic_origin from dog_breeds where breed = ?`

	row := m.DB.QueryRowContext(ctx, query, b)
	var dogBreed DogBreed
	err := row.Scan(
		&dogBreed.ID,
		&dogBreed.Breed,
		&dogBreed.WeigthLowLbs,
		&dogBreed.WeigthHighLbs,
		&dogBreed.AverageWeight,
		&dogBreed.LifeSpan,
		&dogBreed.Details,
		&dogBreed.AlternateNames,
		&dogBreed.GeographicOrigin,
	)

	if err != nil {
		log.Println("Error getting breed by name: ", err)
		return nil, err
	}

	return &dogBreed, nil
}

func (m *mysqlRepository) GetDogOfMonthByID(id int) (*DogOfMonth, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `select id, video, image from dog_of_month where id = ?`

	row := m.DB.QueryRowContext(ctx, query, id)
	var dog DogOfMonth
	err := row.Scan(
		&dog.ID,
		&dog.Video,
		&dog.Image,
	)

	if err != nil {
		log.Println("Error getting dom by id:", err)
		return nil, err
	}

	return &dog, nil
}
