package model

import (
	"context"
	"errors"
	"project/developer-profile-api/db"
	"project/developer-profile-api/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	Role      string             `bson:"role"`
	CreatedAt time.Time          `bson:"createdAt"`
}

func (u *User) Save() (*mongo.InsertOneResult, error) {

	u.CreatedAt = time.Now()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return nil, errors.New("error while hashing password")
	}

	u.Password = hashedPassword

	result, err := db.UserCollection.InsertOne(context.TODO(), u)

	if err != nil {
		return nil, errors.New("error while inserting user data")
	}

	return result, nil
}

func (u *User) ValidateCredentials() error {

	found, err := GetUserByEmail(u.Email)

	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, found.Password)

	if !passwordIsValid {
		return errors.New("invalid password")
	}

	return nil
}

func GetUserByEmail(email string) (*User, error) {

	var user User

	filter := bson.M{"email": email}

	err := db.UserCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
