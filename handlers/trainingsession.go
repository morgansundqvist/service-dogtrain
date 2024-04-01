package handlers

import (
	"context"
	"time"

	"github.com/morgansundqvist/service-dogtrain/database"
	"github.com/morgansundqvist/service-dogtrain/graph/model"
)

func ParseDateFromString(date string) time.Time {
	returnDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Now()
	}
	return returnDate
}

func CreateTrainingSession(ctx context.Context, input model.TrainingSessionInput) (*model.TrainingSession, error) {
	db := database.DB

	trainingSession := &database.DBTrainingSession{
		DogId:         input.DogID,
		CommandGoalId: input.CommandGoalID,
		Date:          ParseDateFromString(input.Date),
		Note:          input.Note,
		Duration:      input.Duration,
		SuccessScale:  input.SuccessScale,
	}

	db.Create(&trainingSession)

	returnTrainingSession := DBTrainingSessionToTrainingSession(trainingSession)
	return returnTrainingSession, nil

}

func UpdateTrainingSession(ctx context.Context, id string, input model.TrainingSessionInput) (*model.TrainingSession, error) {
	db := database.DB

	trainingSession := &database.DBTrainingSession{}

	db.Where("id = ?", id).First(&trainingSession)

	trainingSession.DogId = input.DogID
	trainingSession.CommandGoalId = input.CommandGoalID
	trainingSession.Date = ParseDateFromString(input.Date)
	trainingSession.Note = input.Note
	trainingSession.Duration = input.Duration
	trainingSession.SuccessScale = input.SuccessScale

	db.Save(&trainingSession)

	returnTrainingSession := DBTrainingSessionToTrainingSession(trainingSession)
	return returnTrainingSession, nil
}

func GetTrainingSessionsByDogId(ctx context.Context, dogID string) ([]*model.TrainingSession, error) {
	db := database.DB

	trainingSessions := []*database.DBTrainingSession{}
	db.Where("dog_id = ?", dogID).Find(&trainingSessions)

	returnTrainingSessions := []*model.TrainingSession{}
	for _, trainingSession := range trainingSessions {
		returnTrainingSession := DBTrainingSessionToTrainingSession(trainingSession)
		returnTrainingSessions = append(returnTrainingSessions, returnTrainingSession)
	}

	return returnTrainingSessions, nil
}

func GetTrainingSessionsByCommandGoalId(ctx context.Context, commandGoalID string) ([]*model.TrainingSession, error) {
	db := database.DB

	trainingSessions := []*database.DBTrainingSession{}
	db.Where("command_goal_id = ?", commandGoalID).Find(&trainingSessions)

	returnTrainingSessions := []*model.TrainingSession{}
	for _, trainingSession := range trainingSessions {
		returnTrainingSession := DBTrainingSessionToTrainingSession(trainingSession)
		returnTrainingSessions = append(returnTrainingSessions, returnTrainingSession)
	}

	return returnTrainingSessions, nil
}

func DBTrainingSessionToTrainingSession(trainingSession *database.DBTrainingSession) *model.TrainingSession {
	returnTrainingSession := &model.TrainingSession{
		ID:            trainingSession.ID,
		DogID:         trainingSession.DogId,
		CommandGoalID: trainingSession.CommandGoalId,
		Date:          trainingSession.Date.Format("2006-01-02"),
		Note:          trainingSession.Note,
		Duration:      trainingSession.Duration,
		SuccessScale:  trainingSession.SuccessScale,
	}
	return returnTrainingSession
}
