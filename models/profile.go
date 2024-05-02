package model

import (
	"context"
	"errors"
	"project/developer-profile-api/db"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	Skills       []string `bson:"skills"`
	Certificates []struct {
		Title string `bson:"title"`
		By    string `bson:"by"`
		URL   string `bson:"url"`
	} `bson:"certificates"`
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

func GetLatestProfile() (*Profile, error) {
	var profile Profile

	// Find the latest document by sorting in descending order of the updatedAt field
	opts := options.Find().SetSort(
		bson.D{
			bson.E{Key: "updatedAt", Value: -1},
		},
	).SetLimit(1)

	cur, err := db.ProfileCollection.Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		return nil, errors.New("error while getting latest profile")
	}
	defer cur.Close(context.TODO())

	// Iterate over the cursor and decode the first (latest) document
	if cur.Next(context.TODO()) {
		err := cur.Decode(&profile)
		if err != nil {
			return nil, errors.New("error while decoding profile")
		}
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return &profile, nil
}

func GetProfileByID(id primitive.ObjectID) (*Profile, error) {
	var profile Profile

	filter := bson.M{"_id": id}

	err := db.ProfileCollection.FindOne(context.TODO(), filter).Decode(&profile)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // Return nil and no error if no document is found
		}
		return nil, err
	}

	return &profile, nil
}

func UpdateByID(filter primitive.D, update primitive.D) (*mongo.UpdateResult, error) {

	result, err := db.ProfileCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return nil, err
	}

	return result, nil
}
