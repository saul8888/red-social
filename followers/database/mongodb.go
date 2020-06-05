package database

import (
	"context"
	"time"

	"github.com/orderforme/user/config"
	"github.com/orderforme/user/followers/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const (
	mongoDefaultTimeOut   = 5 * time.Second
	defaultCollectionName = "followers"
)

type MongoDB interface {
	ConnectDB() error
	DisconnectDB() error
	CreateNew(params interface{}) error
	Delete(userID primitive.ObjectID, followID string) error
	GetFollow(userID primitive.ObjectID, followID string) (model.Follow, error)
	//GetAll(params *model.GetLimit) ([]*model.User, error)
	GetCantTotal() (int, error)
}

type Mongodb struct {
	client  *mongo.Client
	context context.Context
	db      *mongo.Database
}

// Connect method
func (repo *Mongodb) ConnectDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.AppConfig.DatabaseURI))
	defer cancel()
	if err != nil {
		return err
	}

	database := client.Database(config.AppConfig.DatabaseName)

	repo.client = client
	repo.context = ctx
	repo.db = database

	return nil
}

// Disconnect method
func (repo *Mongodb) DisconnectDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()
	err := repo.client.Disconnect(ctx)
	return err
}

// GetFollow method
func (repo *Mongodb) GetFollow(userID primitive.ObjectID, followID string) (model.Follow, error) {
	collection := repo.db.Collection(defaultCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()

	follow := model.Follow{}

	objectID, err := primitive.ObjectIDFromHex(followID)
	if err != nil {
		return follow, err
	}

	err = collection.FindOne(ctx, bson.M{"userId": userID, "followingId": objectID}).Decode(&follow)
	if err != nil {
		return follow, err
	}

	return follow, nil
}

// CreateNew method
func (repo *Mongodb) CreateNew(params interface{}) error {
	//connect collection
	collection := repo.db.Collection(defaultCollectionName)

	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()

	_, err := collection.InsertOne(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

// Delete method
func (repo *Mongodb) Delete(userID primitive.ObjectID, followID string) error {
	collection := repo.db.Collection(defaultCollectionName)

	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(followID)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"userId": userID, "followingId": objectID})

	return err
}

/*
func (repo *Mongodb) GetAll(ID string, params *model.GetLimit, search string) ([]*model.Public, error) {
	collection := repo.db.Collection(defaultCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()

	publics := []*model.Public{}
	options := options.Find()
	options.SetSkip(int64(params.Offset))
	options.SetLimit(int64(params.Limit))

	query := bson.M{
		""
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := collection.Find(ctx, query, options)
	if err != nil {
		return publics, err
	}

	if err = cursor.All(ctx, &publics); err != nil {
		return publics, err
	}

}
*/

//GetCantTotal method
func (repo *Mongodb) GetCantTotal() (int, error) {
	//connect collection
	collection := repo.db.Collection(defaultCollectionName)

	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()

	total, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return int(total), nil
}
