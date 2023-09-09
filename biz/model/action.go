package model

import (
	"context"
	"fmt"
	"net/url"

	"github.com/stonebirdjx/topx/biz/dal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// ListActions
type ListOption struct {
	PapeSize int64
	PageNum  int64
}

func ListActions(ctx context.Context, opt ListOption) (*[]Action, error) {
	if opt.PapeSize > 0 {
		return listActionLimit(ctx, opt)
	}

	return listAction(ctx)
}

func listAction(ctx context.Context) (*[]Action, error) {
	actions := &[]Action{}
	filter := bson.M{}
	cursor, err := dal.TopCol.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, actions); err != nil {
		return nil, err
	}

	return actions, nil
}

func listActionLimit(ctx context.Context, opt ListOption) (*[]Action, error) {
	filter := bson.M{}
	skip := (opt.PageNum - 1) * opt.PapeSize
	findOpt := &options.FindOptions{
		Limit: &opt.PapeSize,
		Skip:  &skip,
	}

	cursor, err := dal.TopCol.Find(ctx, filter, findOpt)
	if err != nil {
		return nil, err
	}

	actions := &[]Action{}
	if err := cursor.All(ctx, actions); err != nil {
		return nil, err
	}

	return actions, nil

}
