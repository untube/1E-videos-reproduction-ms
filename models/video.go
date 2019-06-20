package models

import "gopkg.in/mgo.v2/bson"

// Represents a video, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Video struct {
	ID           bson.ObjectId `bson:"_id" json:"id"`
	User_ID      int           `bson:"user_id" json:"user_id"`
	Category_ID  string        `bson:"category_id" json:"category_id"`
	Video_ID     string        `bson:"video_id" json:"video_id"`
	Views        int           `bson:"views" json:"views"`
	Title        string        `bson:"title" json:"title"`
	Description  string        `bson:"description" json:"description"`
	OriginalName string        `bson:"originalname" json:"originalname"`
	Filename     string        `bson:"filename" json:"filename"`
	Image        string        `bson:"image" json:"image"`
}
