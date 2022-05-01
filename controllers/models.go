package controllers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrorNotFount      = errors.New("models: resource not found")
	ErrInvalidID       = errors.New("models: ID provided was invalid")
	ErrInvalidEmail    = errors.New("models : invalid email address provided")
	ErrInvalidPassword = errors.New("models: incorrect password provided")
)

type UserService struct{}

type Netflix struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson: "_id,omitempty"`
	Movie   string             `json:"movie,omitempty" bson:"movie"`
	Watched bool               `json: "watched,omitempty" bson: "watched"`
}

type Users struct {
	_id          primitive.ObjectID `json:"_id,omitempty" bson: "_id,omitempty"`
	Name         string             `json:"name,omitempty" bson:"name"`
	Email        string             `json:"email,omitempty" bson:"email"`
	Color        string             `json:"color,omitempty" bson:"color"`
	Orders       []Order            `json:"orders,omitempty" bson:"orders"`
	Password     string             `json:"password" bson:"password"`
	PasswordHash string             `json:"passwordhash" bson:"passwordhash"`
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

	var User Users
	err := Collection.FindOne(ctx, filter).Decode(&User)
	if err != nil {
		panic(err)
	}

	return &User, nil
}

func (us *UserService) Create(user *Users) (*Users, error) {
	pwBytes := []byte(user.Password + userPwPepper)
	tempx := 10
	hashedBytes, err := bcrypt.GenerateFromPassword(pwBytes, tempx)
	//hashedBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(hashedBytes)
	user._id = primitive.NewObjectID()
	for i := range user.Orders {
		user.Orders[i].UserID = user._id.Hex()
	}
	//user.Password = ""
	z := true
	inserted, err := Collection.InsertOne(context.Background(), user, &options.InsertOneOptions{BypassDocumentValidation: &z})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie in db with id", inserted.InsertedID)
	return user, nil

}

func (us *UserService) Update(user *Users) error {
	id := user._id

	filter := bson.M{"_id": id}
	ud := bson.M{}
	ud["name"] = user.Name
	ud["email"] = user.Email
	ud["color"] = user.Color
	ud["orders"] = user.Orders

	update := bson.M{"$set": ud}

	result, err := Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Println("Modified count :", result.ModifiedCount)

	return nil

}

func (us *UserService) ByEmail(email string) (*Users, error) {

	filter := bson.M{"email": email}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	var User Users
	err := Collection.FindOne(ctx, filter).Decode(&User)
	if err != nil {
		panic(err)
	}

	return &User, nil

}
func (us *UserService) Authenticate(email, password string) (*Users, error) {

	foundUser, err := us.ByEmail(email)
	if err != nil {
		return nil, err
	}
	fmt.Println("&&&&&&&&")
	fmt.Println(fmt.Sprintf("%v", foundUser))
	fmt.Println(foundUser.Password)
	fmt.Println("&&&&&&&")
	fmt.Println(password)
	fmt.Println("compare these :")
	fmt.Println(fmt.Sprintf("this is stored %v", []byte(foundUser.PasswordHash)))
	fmt.Println(fmt.Sprintf("this is input %v", []byte(password+userPwPepper)))
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password+userPwPepper))
	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return nil, ErrInvalidPassword

		default:
			return nil, err
		}
	}
	fmt.Println("yaha tak aaya na @ hased")
	return foundUser, nil
}

func (us *UserService) Delete(id primitive.ObjectID) error {

	filter := bson.M{"_id": id}
	deleteCount, err := Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	fmt.Println("movie got deleted with delete count", deleteCount)
	return nil
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
