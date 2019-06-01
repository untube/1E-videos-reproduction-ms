package models

import "gopkg.in/mgo.v2/bson"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Category struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Category    string        `bson:"" json:"category"`
	Description string        `bson:"description" json:"description"`
	Category_ID string        `bson:"category_id" json:"category_id"`
}
