package model

import (
	"context"
	"errors"
	"project/developer-profile-api/db"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Profile struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name struct {
		FirstName string `bson:"firstName"`
		LastName  string `bson:"lastName"`
	} `bson:"name"`
	Bio      string `bson:"bio"`
	Location struct {
		State   string `bson:"state"`
		Country string `bson:"country"`
	} `bson:"location"`
	SocialLinks []struct {
		Platform string `bson:"platform"`
		URL      string `bson:"url"`
	} `bson:"socialLinks"`
	Skills    []string `bson:"skills"`
	Education []struct {
		Institution  string `bson:"institution"`
		Degree       string `bson:"degree"`
		FieldOfStudy string `bson:"fieldOfStudy"`
		StartYear    int    `bson:"startYear"`
		EndYear      int    `bson:"endYear"`
	} `bson:"education"`
	Work []struct {
		Company     string    `bson:"company"`
		Title       string    `bson:"title"`
		StartDate   time.Time `bson:"startDate"`
		EndDate     time.Time `bson:"endDate"`
		Description string    `bson:"description"`
	} `bson:"work"`
	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

func (p *Profile) Save() (*mongo.InsertOneResult, error) {

	p.CreatedAt = time.Now()

	result, err := db.ProfileCollection.InsertOne(context.TODO(), p)

	if err != nil {
		return nil, errors.New("error while inserting profile data")
	}

	return result, nil
}
