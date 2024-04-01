package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Global DB connection
var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("dogtrain.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migration to create tables for schema model
	DB.AutoMigrate(&DBDog{})
	DB.AutoMigrate(&DBCommandGoal{})
}

// Base contains common columns for all tables.
type Base struct {
	ID        string     `gorm:"primary_key"; json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(tx *gorm.DB) error {
	newUuid := uuid.NewString()
	base.ID = newUuid
	return nil
}

type DBDog struct {
	Base

	Name string `json:"name"`
}

type DBCommandGoal struct {
	Base

	DogId            string `json:"dog_id"`
	BaseCommand      string `json:"base_command"`
	Goal             string `json:"goal"`
	DefinitionOfDone string `json:"definition_of_done"`
	Priority         int    `json:"priority"`
}

type DBTrainingSession struct {
	Base

	DogId         string    `json:"dog_id"`
	CommandGoalId string    `json:"command_goal_id"`
	Date          time.Time `json:"date"`
	Note          string    `json:"note"`
	Duration      int       `json:"duration"`
	SuccessScale  int       `json:"success_scale"`
}
