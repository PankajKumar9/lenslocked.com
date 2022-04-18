package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson: "_id,omitempty"`
	Movie   string             `json:"movie,omitempty" bson:"movie"`
	Watched bool               `json: "watched,omitempty" bson: "watched"`
}

type Users struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson: "_id,omitempty"`
	Name  string             `json:"name,omitempty" bson:"name"`
	Email string             `json:"email,omitempty" bson:"email"`
	Color string             `json:"color,omitempty" bson:"color"`
	Orders	[]Order	`json:"orders,omitempty" bson:"orders"`
}
type Order struct {
	
	UserID      uint `json:"userid,omitempty" bson:"userid"`
	Amount      string             `json:"amount,omitempty" bson:"amount"`
	Description string             `json:"description,omitempty" bson:"description"`
}
