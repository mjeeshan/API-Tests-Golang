package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeeshan12/apimongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString string = "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"
const dbName string = "netflix"
const collectionName string = "watchlist"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected successfully")
	collection = client.Database(dbName).Collection(collectionName)
	fmt.Println("Collection Reference is ready")
}

func insertMovie(movie model.Netflix) {
	var mov model.Netflix
	filter := bson.M{"movie": movie.Movie}
	e := collection.FindOne(context.Background(), filter).Decode(&mov)
	if e == mongo.ErrNilDocument {
		fmt.Println("record does not exist")
		log.Fatal(e)
	}
	if mov.Movie == movie.Movie {
		fmt.Println("Movie is already existed. You can't add same movie again")
		log.Fatal("Movie is already existed. You can't add same movie again")
	}

	ins, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal((err))
	}
	fmt.Println("Inserted the movie in db", ins.InsertedID)
}

func updateMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal((err))
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count", result.ModifiedCount)
}

func deleteMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal((err))
	}
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted count", result.DeletedCount)
}

func deleteAllMovies() int32 {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted count for all movies", result.DeletedCount)
	return int32(result.DeletedCount)

}

func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
	var movies []primitive.M
	for cursor.Next(context.Background()) {
		var movie bson.M
		if err := cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)
	insertMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updateMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Mongo API for testing</h1>"))
}
