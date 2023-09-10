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
	RateLimit   float64            `json:"rate_limit" bson:"rate_limit"`
	IsAuth      bool               `json:"is_auth" bson:"is_auth"`
	Path        string             `json:"path" bson:"path"`
	Proxy       string             `json:"proxy" bson:"proxy"`
	Timeout     int                `json:"timeout" bson:"timeout"`
	Version     string             `json:"version" bson:"version"`
}

func (a *Action) Validate() error {
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

	if a.RateLimit < 1 {
		a.RateLimit = 1
	}

	if a.Timeout < 1 {
		a.Timeout = 30000

	}
	return nil
}

// InsertOne .
func (a *Action) InsertOne(ctx context.Context) error {
	a.ID = primitive.NewObjectID()
	_, err := dal.TopCol.InsertOne(ctx, a)
	return err
}

// GetAction .
func (a *Action) GetAction(ctx context.Context) error {
	fliter := bson.M{
		"_id": a.ID,
	}

	return dal.TopCol.FindOne(ctx, fliter).Decode(a)
}

// UpdateAction .
func (a *Action) UpdateAction(ctx context.Context) error {
	fliter := bson.M{
		"_id": a.ID,
	}

	_, err := dal.TopCol.ReplaceOne(ctx, fliter, a)
	return err
}

// UpdateAction .
func (a *Action) FindOne(ctx context.Context) error {
	fliter := bson.M{
		"name":         a.Name,
		"service_name": a.ServiceName,
		"version":      a.Version,
	}

	return dal.TopCol.FindOne(ctx, fliter).Decode(a)
}

// ListActions
type ListOption struct {
	PapeSize int64
	PageNum  int64
}

func ListActions(ctx context.Context, opt ListOption) (*[]Action, int, error) {
	if opt.PapeSize > 0 {
		return listActionLimit(ctx, opt)
	}

	return listAction(ctx)
}

func listAction(ctx context.Context) (*[]Action, int, error) {
	actions := &[]Action{}
	filter := bson.M{}
	cursor, err := dal.TopCol.Find(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	if err := cursor.All(ctx, actions); err != nil {
		return nil, 0, err
	}

	return actions, len(*actions), nil
}

func listActionLimit(ctx context.Context, opt ListOption) (*[]Action, int, error) {
	filter := bson.M{}
	count, err := dal.TopCol.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	skip := (opt.PageNum - 1) * opt.PapeSize
	findOpt := &options.FindOptions{
		Limit: &opt.PapeSize,
		Skip:  &skip,
	}

	cursor, err := dal.TopCol.Find(ctx, filter, findOpt)
	if err != nil {
		return nil, 0, err
	}

	actions := &[]Action{}
	if err := cursor.All(ctx, actions); err != nil {
		return nil, 0, err
	}

	return actions, int(count), nil
}

// DeleteByIDs
func DeleteByIDs(ctx context.Context, ids []primitive.ObjectID) error {
	filter := bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	}
	_, err := dal.TopCol.DeleteMany(ctx, filter)
	return err
}

// DeleteByID
func DeleteByID(ctx context.Context, id primitive.ObjectID) error {
	filter := bson.M{
		"_id": id,
	}

	_, err := dal.TopCol.DeleteOne(ctx, filter)
	return err
}
