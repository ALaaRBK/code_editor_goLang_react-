package schemas

import (
	"context"
	"db1"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

//Submission
type Submission struct {
	Lang      string `json:"lang"`
	CreatedAt int64  `json:"createdAt"`
	FilePath  string `json:"filePath"`
}

//Save ..
// func Save(db mongo.Database, submission Submission) {

// }
func JustTest() {
	collection := db1.DB.Collection("SSSSSSSSSSSSS")
	res, err := collection.InsertOne(context.TODO(), bson.M{"name": "pi", "value": 3.14159})
	fmt.Println(res, err)
}
