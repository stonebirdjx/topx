package dal

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	topxColName = "topx"
)

type MongoOption struct {
	URI string
	DB  string
}

var (
	Mongoclient *mongo.Client
	TopCol      *mongo.Collection
)

func MongoInit(ctx context.Context, opt MongoOption) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(opt.URI))
	if err != nil {
		return err
	}

	Mongoclient = client
	TopCol = Mongoclient.Database(opt.DB).Collection(topxColName)
	return nil
}
