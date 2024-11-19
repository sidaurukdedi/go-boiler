package user

import (
	"context"

	"github.com/sidaurukdedi/go-boiler/internal/entity"
	"github.com/sidaurukdedi/go-boiler/pkg/exception"
	"github.com/sidaurukdedi/go-boiler/pkg/mongodb"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	Save(ctx context.Context, user entity.User) (err error)
	Update(ctx context.Context, userId string, user entity.User) (err error)
	FindByID(ctx context.Context, userId string) (user entity.User, err error)
}

type userRepository struct {
	logger         *logrus.Logger
	db             mongodb.Database
	collectionName string
}

// NewUserRepository is a constructor.
func NewUserRepository(logger *logrus.Logger, db mongodb.Database, collectionName string) UserRepository {
	return &userRepository{logger: logger, db: db, collectionName: collectionName}
}

func (r *userRepository) Save(ctx context.Context, user entity.User) (err error) {
	_, err = r.db.Collection(r.collectionName).InsertOne(ctx, user)
	if err != nil {
		r.logger.Error(err)
		return
	}
	return
}

func (r *userRepository) Update(ctx context.Context, userId string, user entity.User) (err error) {
	filter := bson.M{
		"id": userId,
	}

	data := map[string]interface{}{
		"$set": user,
	}

	result, err := r.db.Collection(r.collectionName).UpdateOne(ctx, filter, data, options.Update().SetUpsert(true))
	if err != nil {
		r.logger.Error(err)
		err = exception.ErrInternalServer
		return
	}

	if result.MatchedCount < 1 {
		err = exception.ErrNotFound
		return
	}
	return
}

func (r *userRepository) FindByID(ctx context.Context, userId string) (user entity.User, err error) {
	filter := bson.M{
		"id": userId,
	}

	if err = r.db.Collection(r.collectionName).FindOne(ctx, filter).Decode(&user); err != nil {
		if err != mongo.ErrNoDocuments {
			r.logger.Error(err)
			err = exception.ErrInternalServer
			return
		}
		err = exception.ErrNotFound
		return
	}

	return
}
