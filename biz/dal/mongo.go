package dal

import "context"

type MongoOption struct {
	URI string
}

func MongoConnect(ctx context.Context, opt MongoOption) error {
	return nil
}
