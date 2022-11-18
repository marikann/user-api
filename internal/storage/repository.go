package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
	"user-api/internal/errors"
	mongoTypes "user-api/internal/storage/types"
)

type Repository struct {
	mc *mongo.Collection
}

func NewRepository(mc *mongo.Collection) *Repository {
	return &Repository{mc: mc}
}

func (r Repository) Create(document interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.mc.InsertOne(ctx, document)

	if err != nil {
		return errors.InsertUser.WrapDesc(err.Error())
	}
	return nil
}

func (r Repository) FindOne(filter bson.M) (user *mongoTypes.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if err := r.mc.FindOne(ctx, filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.FindOne.WrapDesc(err.Error())
	}

	return user, nil
}

func (r Repository) FindOneAndDelete(filter bson.M) (user *mongoTypes.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if err := r.mc.FindOneAndDelete(ctx, filter).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.FindOneAndDelete.WrapDesc(err.Error())
	}

	return user, nil
}

func (r Repository) FindOneAndUpdate(filter bson.M, update bson.D) (user *mongoTypes.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	if err := r.mc.FindOneAndUpdate(ctx, filter, update).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.FindOneAndUpdate.WrapDesc(err.Error())
	}

	return user, nil
}

func (r Repository) Find(limit, offset int64, filter interface{}) (users []mongoTypes.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	p := options.Find().SetLimit(limit).SetSkip(offset)

	cursor, err := r.mc.Find(ctx, filter, p)

	if err != nil {
		return nil, errors.Cursor.WrapDesc(err.Error())
	}

	if err := cursor.All(context.Background(), &users); err != nil {
		return nil, errors.CursorAll.WrapDesc(err.Error())
	}

	return users, nil

}

func (r Repository) CountDocuments(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	totalCounts, err := r.mc.CountDocuments(ctx, filter)
	if err != nil {
		return 0, errors.CountDocuments.WrapDesc(err.Error())
	}
	return totalCounts, nil
}
