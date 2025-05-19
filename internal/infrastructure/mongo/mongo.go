package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/event"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

type Config struct {
	Dsn string
}

type MongoClient interface {
	GetClient() *mongo.Client
	Disconnect() error
}

type mongoClient struct {
	client *mongo.Client
}

func NewClient(cfg *Config) (MongoClient, error) {
	monitor := &event.CommandMonitor{
		Started: func(_ context.Context, evt *event.CommandStartedEvent) {
			log.Print(evt.Command)
		},
	}
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		ApplyURI(cfg.Dsn).
		SetServerAPIOptions(serverAPI).SetMonitor(monitor)

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return &mongoClient{
		client: client,
	}, nil
}

func (m *mongoClient) GetClient() *mongo.Client {
	return m.client
}

func (m *mongoClient) Disconnect() error {
	return m.client.Disconnect(context.Background())
}
