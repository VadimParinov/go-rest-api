package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"rest-api/cmd/internal/user"
)

type db struct {
	collection *mongo.Collection
}

func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	return "", nil
}

func (d *db) FindOne(ctx context.Context, id string) (u user.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, err
	}
	filter := bson.M{"_id": oid}

	result := d.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		return u, err
	}
	if err = result.Decode(&u); err != nil {
		return u, err
	}
	return u, nil
}

func (d *db) FindAll(ctx context.Context) (u []user.User, err error) {

	cursor, err := d.collection.Find(ctx, bson.M{})
	if cursor.Err() != nil {
		return u, err
	}
	if err = cursor.All(ctx, &u); err != nil {
		return u, err
	}
	return u, nil
}

func (d *db) Update(ctx context.Context, user user.User) error {
	objectId, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectId}
	update, err := bson.Marshal(user)
	if err != nil {
		return err
	}
	var updateUserObj bson.M
	err = bson.Unmarshal(update, &updateUserObj)
	if err != nil {
		return err
	}
	delete(updateUserObj, "_id")
	result, err := d.collection.UpdateOne(ctx, filter, bson.M{"$set": updateUserObj})
	if result.MatchedCount == 0 {
		return fmt.Errorf("user with id %s not found", user.ID)
	}
	return nil
}

func (d *db) Delete(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objectId}
	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("user with id %s not found", id)
	}
	return nil
}

func NewStorage(database *mongo.Database, collection string) user.Storage {
	return &db{
		collection: database.Collection(collection),
	}
}
