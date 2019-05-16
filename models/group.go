package models

import "gopkg.in/mgo.v2/bson"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Group struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Category    string        `bson:"category" json:"category"`
	Thumbnail   string        `bson:"thumbnail" json:"thumbnail"`
	Description string        `bson:"description" json:"description"`
}
