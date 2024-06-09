package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient() *mongo.Client {
	ctxRoot := context.Background()
	ctx, cancel := context.WithTimeout(ctxRoot, 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		panic(err)
	}
	/*
		defer func() {
			if err = client.Disconnect(ctx); err != nil {
				panic(err)
			}
		}()
	*/
	return client
}

/*
func GetSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	} else {
		return session
	}
}
*/
