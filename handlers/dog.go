package handlers

import (
	"context"

	"github.com/morgansundqvist/service-dogtrain/database"
	"github.com/morgansundqvist/service-dogtrain/graph/model"
)

func GetDogs(ctx context.Context) ([]*model.Dog, error) {
	db := database.DB

	var dogs []database.DBDog
	db.Find(&dogs)

	var result []*model.Dog
	for _, dog := range dogs {
		result = append(result, DBDogToDog(dog))
	}

	return result, nil
}

func GetDog(ctx context.Context) (*model.Dog, error) {
	db := database.DB

	var dog database.DBDog

	db.Where("id = ?", 1).First(&dog)

	return DBDogToDog(dog), nil
}

func CreateDog(ctx context.Context, input model.DogInput) (*model.Dog, error) {
	db := database.DB

	dog := database.DBDog{
		Name: input.Name,
	}

	db.Create(&dog)

	x := DBDogToDog(dog)
	return x, nil
}

func DBDogToDog(dog database.DBDog) *model.Dog {
	x := &model.Dog{
		ID:   dog.ID,
		Name: dog.Name,
	}
	return x
}
