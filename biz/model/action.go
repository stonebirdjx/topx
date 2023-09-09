package model

import (
	"context"
	"fmt"
	"net/url"

	"github.com/stonebirdjx/topx/biz/dal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Action struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `json:"name" bson:"name"`
	ServiceName string             `json:"service_name" bson:"service_name"`
	Description string             `json:"description" bson:"description"`
	IsAuth      bool               `json:"is_auth" bson:"is_auth"`
	Path        string             `json:"path" bson:"path"`
	Proxy       string             `json:"proxy" bson:"proxy"`
	Timeout     int                `json:"timeout" bson:"timeout"`
	Version     string             `json:"version" bson:"version"`
}

func (a Action) Validate() error {
	if a.Name == "" {
		return fmt.Errorf("Action name can not be nil")
	}

	if a.Name != url.QueryEscape(a.Name) {
		return fmt.Errorf("Action name can not url special characters")
	}

	if a.ServiceName == "" {
		return fmt.Errorf("Action service name can not be nil")
	}

	if a.ServiceName != url.QueryEscape(a.ServiceName) {
		return fmt.Errorf("Action service name can not url special characters")
	}

	if a.Version == "" {
		return fmt.Errorf("Action version can not be nil")
	}

	if a.Version != url.QueryEscape(a.Version) {
		return fmt.Errorf("Action version can not url special characters")
	}

	return nil
}

// InsertOne .
func (a Action) InsertOne(ctx context.Context) error {
	a.ID = primitive.NewObjectID()
	_, err := dal.TopCol.InsertOne(ctx, a)
	return err
}
