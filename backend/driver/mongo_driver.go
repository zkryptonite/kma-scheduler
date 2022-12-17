package driver

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"
	"log"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client *mongo.Client
}

var (
	Mongo  = &MongoDB{}
	once   sync.Once
	dbPass = os.Getenv("DB_PASS")
	dbUser = os.Getenv("DB_USER")
	dbUri  = fmt.Sprintf("mongodb+srv://%s:%s@cluster0.acyql.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
		dbUser, dbPass)
)

func ConnectMongoDB() *MongoDB {
	once.Do(func() {
		clientOptions := options.Client().ApplyURI(dbUri)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			log.Fatal(err)
		}

		Mongo.Client = client
	})
	return Mongo
}
