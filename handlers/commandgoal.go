package handlers

import (
	"context"

	"github.com/morgansundqvist/service-dogtrain/database"
	"github.com/morgansundqvist/service-dogtrain/graph/model"
)

func GetCommandGoals(ctx context.Context, obj *model.Dog) ([]*model.CommandGoal, error) {
	db := database.DB

	var commandGoals []*database.DBCommandGoal

	if err := db.Where("dog_id = ?", obj.ID).Find(&commandGoals).Error; err != nil {
		return nil, err
	}

	var modelCommandGoals []*model.CommandGoal

	for _, commandGoal := range commandGoals {
		modelCommandGoals = append(modelCommandGoals, DBCommandGoalToCommandGoal(*commandGoal))
	}

	return modelCommandGoals, nil
}

func CreateCommandGoal(ctx context.Context, input model.CommandGoalInput) (*model.CommandGoal, error) {
	db := database.DB

	commandGoal := database.DBCommandGoal{
		DogId:            input.DogID,
		BaseCommand:      input.BaseCommand,
		Goal:             input.Goal,
		DefinitionOfDone: input.DefinitionOfDone,
		Priority:         input.Priority,
	}

	if err := db.Create(&commandGoal).Error; err != nil {
		return nil, err
	}

	x := DBCommandGoalToCommandGoal(commandGoal)
	return x, nil
}

func GetCommandGoalsByDogId(ctx context.Context, dogId string) ([]*model.CommandGoal, error) {
	db := database.DB

	var commandGoals []*database.DBCommandGoal

	if err := db.Where("dog_id = ?", dogId).Find(&commandGoals).Error; err != nil {
		return nil, err
	}

	var modelCommandGoals []*model.CommandGoal

	for _, commandGoal := range commandGoals {
		modelCommandGoals = append(modelCommandGoals, DBCommandGoalToCommandGoal(*commandGoal))
	}

	return modelCommandGoals, nil
}

func DBCommandGoalToCommandGoal(commandGoal database.DBCommandGoal) *model.CommandGoal {
	x := &model.CommandGoal{
		ID:               commandGoal.ID,
		DogID:            commandGoal.DogId,
		BaseCommand:      commandGoal.BaseCommand,
		Goal:             commandGoal.Goal,
		DefinitionOfDone: commandGoal.DefinitionOfDone,
		Priority:         commandGoal.Priority,
	}
	return x
}
