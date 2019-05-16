package dao

import (
	"log"

	. "VideoPlayer-ms/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type VideosDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "videos"
)

// Establish a connection to database
func (m *VideosDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of videos
func (m *VideosDAO) FindAll() ([]Video, error) {
	var videos []Video
	err := db.C(COLLECTION).Find(bson.M{}).All(&videos)
	return videos, err
}

// Find a video by its id
func (m *VideosDAO) FindById(id string) (Video, error) {
	var video Video
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&video)
	return video, err
}

// Insert a video into database
func (m *VideosDAO) Insert(video Video) error {
	err := db.C(COLLECTION).Insert(&video)
	return err
}

// Delete an existing video
func (m *VideosDAO) Delete(video Video) error {
	err := db.C(COLLECTION).Remove(&video)
	return err
}

// Update an existing video
func (m *VideosDAO) Update(video Video) error {
	err := db.C(COLLECTION).UpdateId(video.ID, &video)
	return err
}
