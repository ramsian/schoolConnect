package healthHelper

import (
	"context"
	"fmt"
	"main/database"
	healthModel "main/health/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HealthHelper struct{}

func (h *HealthHelper) GetServerPing() string {
	return "I'm alive!"
}

func (h *HealthHelper) CreateHealthData() error {
	fmt.Println("Test")
	database := database.GetDatabase()
	collection := database.Collection(healthModel.Name)
	dummyHealth := healthModel.HealthModel{
		ID:          primitive.NewObjectID(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       "Test Title",
		Description: "Test Description",
	}
	_, err := collection.InsertOne(context.TODO(), dummyHealth)
	if err != nil {
		return err
	}
	return nil
}
