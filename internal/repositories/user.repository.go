package repositories

import (
	"context"

	"github.com/khoand3012/go-ieltsgrader/db"
	"github.com/khoand3012/go-ieltsgrader/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	database   db.Database
	collection string
}

func NewUserRepository(db db.Database, collection string) domain.UserRepository {
	return &UserRepository{
		database:   db,
		collection: collection,
	}
}

func (u *UserRepository) Create(c context.Context, user *domain.User) error {
	collection := u.database.Collection(u.collection)
	_, err := collection.InsertOne(c, user)

	return err
}

func (u *UserRepository) Fetch(c context.Context) ([]domain.User, error) {
	collection := u.database.Collection(u.collection)

	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)
	if err != nil {
		return nil, err
	}

	var users []domain.User
	err = cursor.All(c, &users)
	if users == nil {
		return []domain.User{}, err
	}

	return users, nil
}

func (u *UserRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	collection := u.database.Collection(u.collection)
	var user domain.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)

	return user, err
}

func (u *UserRepository) GetByID(c context.Context, id string) (domain.User, error) {
	collection := u.database.Collection(u.collection)
	var user domain.User

	idHex, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return user, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&user)

	return user, err
}
