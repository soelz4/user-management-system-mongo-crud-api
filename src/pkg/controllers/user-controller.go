package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"go-mongo/src/pkg/models"
)

type UserController struct {
	client *mongo.Client
}

func NewUserController(s *mongo.Client) *UserController {
	return &UserController{s}
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	// Decode Request Body and Put in user Variable and Generate ObjectId (_id) for this Variable
	json.NewDecoder(r.Body).Decode(&user)
	user.ID = primitive.NewObjectID()
	// Choose Your Database and Collection
	collection := uc.client.Database("bank").Collection("user")
	// Insert Document into the Collection
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		panic(err)
	}
	// Encode Go Server Data-Object into the JSON (for API)
	user_json, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(user_json)
}

func (uc UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	// Slice of User Struct
	users := []models.User{}
	// Choose Your Database and Collection
	collection := uc.client.Database("bank").Collection("user")
	// Find All Documents - Filter ~> bson.D{} (Nothing)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	// Put All Documents into the users Slice
	if err = cursor.All(context.TODO(), &users); err != nil {
		panic(err)
	}
	// Encode Go Server Data-Object into the JSON (for API)
	users_json, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(users_json)
}

func (uc UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	// Choose Your Database and Collection
	collection := uc.client.Database("bank").Collection("user")
	// Save given ID in /user/{userID} in params
	params := mux.Vars(r)
	userID := params["userID"]
	// ObjectIDFromHex Creates a New ObjectID from a Hex String. It Returns an Error if the Hex String is not a Valid ObjectID.
	objectUserID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	filter := bson.M{"_id": objectUserID}
	// Find Document by ID - Filter by ID
	err = collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(404)
		} else {
			panic(err)
		}
	}
	// Encode Go Server Data-Object into the JSON (for API)
	user_json, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(user_json)
}

func (uc UserController) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	updateUser := models.User{}
	// Decode Request Body and Put in user Variable and Generate ObjectId (_id) for this Variable
	json.NewDecoder(r.Body).Decode(&updateUser)
	// Choose Your Database and Collection
	collection := uc.client.Database("bank").Collection("user")
	// Save given ID in /user/{userID} in params
	params := mux.Vars(r)
	userID := params["userID"]
	// userID = updateUser.ID.Hex()
	objectUserID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		panic(err)
	} else {
		updateUser.ID = objectUserID
		update := bson.D{{"$set", updateUser}}
		filter := bson.M{"_id": objectUserID}
		// Update Document by ID - Filter with ID
		_, err = collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			panic(err)
		}
	}
}

func (uc UserController) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	// Choose Your Database and Collection
	collection := uc.client.Database("bank").Collection("user")
	// Save given ID in /user/{userID} in params
	params := mux.Vars(r)
	userID := params["userID"]
	objectUserID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	filter := bson.M{"_id": objectUserID}
	// Delete Document by ID - Filter with ID
	_, err = collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "User ~> ", objectUserID, " Deleted", "\n")
}
