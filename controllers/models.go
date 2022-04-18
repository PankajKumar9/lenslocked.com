package controllers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrorNotFount = errors.New("models: resource not found")
)

type UserService struct{}

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson: "_id,omitempty"`
	Movie   string             `json:"movie,omitempty" bson:"movie"`
	Watched bool               `json: "watched,omitempty" bson: "watched"`
}

type Users struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson: "_id,omitempty"`
	Name   string             `json:"name,omitempty" bson:"name"`
	Email  string             `json:"email,omitempty" bson:"email"`
	Color  string             `json:"color,omitempty" bson:"color"`
	Orders []Order            `json:"orders,omitempty" bson:"orders"`
}
type Order struct {
	UserID      string `json:"userid,omitempty" bson:"userid"`
	Amount      int    `json:"amount,omitempty" bson:"amount"`
	Description string `json:"description,omitempty" bson:"description"`
}

func (us *UserService) ByID(id primitive.ObjectID) (*Users, error) {
	filter := bson.M{"_id": id}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := Collection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

	var User Users
	result.Decode(&User)

	return &User, nil
}

func (us *UserService) Create(user *Users) (*Users, error) {

	user.ID = primitive.NewObjectID()
	for i := range user.Orders {
		user.Orders[i].UserID = user.ID.Hex()
	}

	inserted, err := Collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie in db with id", inserted.InsertedID)
	return user,nil

}



/*


func toDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}

	err = bson.Unmarshal([]byte(data), &doc)
	fmt.Println(doc, "here check")
	return
}


	res, err := db.Collection(BATCH_COLLECTION_NAME).InsertOne(ctx, doc)


*/
