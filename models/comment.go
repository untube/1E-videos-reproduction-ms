package models

import "gopkg.in/mgo.v2/bson"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Comment struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Video_ID string        `bson:"video_id" json:"video_id"`
	User_Id  int           `bson:"user_id" json:"user_id"`
	Message  string        `bson:"message" json:"message"`
}
