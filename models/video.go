package models

import "gopkg.in/mgo.v2/bson"

// Represents a video, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Video struct {
	ID           bson.ObjectId `bson:"_id" json:"id"`
	User_ID      int           `bson:"user_id" json:"user_id"`
	Category_ID  string        `bson:"category_id" json:"category_id"`
	Title        string        `bson:"title" json:"title"`
	Description  string        `bson:"description" json:"description"`
	Destination  string        `bson:"destination" json:"destination"`
	Views        int           `bson:"views" json:"views"`
	Size         float32       `bson:"size" json:"size"`
	FieldName    string        `bson:"fieldname" json:"fieldname"`
	OriginalName string        `bson:"originalname" json:"originalname"`
	Encoding     string        `bson:"encoding" json:"encoding"`
	Mimetype     string        `bson:"mimetype" json:"mimetype"`
	Filename     string        `bson:"filename" json:"filename"`
}
