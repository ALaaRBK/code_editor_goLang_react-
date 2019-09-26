package schemas

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

//Submission
type Submission struct {
	Lang      string `json:"lang"`
	CreatedAt time.Time
	FilePath  string
	CodeText  string `json:"codeText"`
	Stdin     string `json:"stdin"`
}

// Save ..
func Save(db *mongo.Database, submission Submission) {
	collection := db.Collection("submission")
	_, err := collection.InsertOne(context.TODO(), submission)

	if err != nil {
		fmt.Println("DB Insert Error", err)
	}
}
