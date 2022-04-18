package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/PankajKumar9/lenslocked.com/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://pnkjdh:pankaj2mongo@cluster0.aroxb.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

const dbName = "Lenslocked"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	//client options

	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongodb

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")
	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")

}

//this is where mongodb hel  pers ß
//insert 1 record

func insertOneMovie(movie models.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie in db with id", inserted.InsertedID)

}

//update 1 record

func updateOneMovie(movieID string) {
	id, _ := primitive.ObjectIDFromHex("movieID")

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Println("Modified count :", result.ModifiedCount)

}

// delete one record

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	fmt.Println("movie got deleted with delete count", deleteCount)
}

//delete all records from mongoDB

func deleteAllMovie() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("number of movies delete:", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// get all movies from database

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {

			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies

}

//Actual controllers

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Netflix
	fmt.Println("prinintg request")
	fmt.Println(r)
	json.NewDecoder(r.Body).Decode(&movie)
	fmt.Println("Movies is :", movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)

	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)

	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}
func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)

}


//lenslocked stuff

func CreateOrder(user models.Users,amount int,desc string){}