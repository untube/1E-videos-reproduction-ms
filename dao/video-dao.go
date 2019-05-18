package dao

import (
	"log"

	. "video-reproduction-ms/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type VideosDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

// Establish a connection to database
func (m *VideosDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of videos
func (m *VideosDAO) FindAllVideos() ([]Video, error) {
	var videos []Video
	err := db.C("videos").Find(bson.M{}).All(&videos)
	return videos, err
}

// Find a video by its id
func (m *VideosDAO) FindVideoById(id string) (Video, error) {
	var video Video
	err := db.C("videos").FindId(bson.ObjectIdHex(id)).One(&video)
	return video, err
}

// Insert a video into database
func (m *VideosDAO) InsertVideo(video Video) error {
	err := db.C("videos").Insert(&video)
	return err
}

// Delete an existing video
func (m *VideosDAO) DeleteVideo(video Video) error {
	err := db.C("videos").Remove(&video)
	return err
}

// Update an existing video
func (m *VideosDAO) UpdateVideo(video Video) error {
	err := db.C("videos").UpdateId(video.ID, &video)
	return err
}

// Find list of comments
func (m *VideosDAO) FindCommentsByVideoId(id string) ([]Comment, error) {

	var comments []Comment
	err := db.C("comments").Find(bson.M{"video_id": id}).All(&comments)
	return comments, err
}

// Find list of comments
func (m *VideosDAO) FindAllComments() ([]Comment, error) {
	var comments []Comment
	err := db.C("comments").Find(bson.M{}).
		All(&comments)
	return comments, err
}

// Find a comment by its id
func (m *VideosDAO) FindCommentById(id string) (Comment, error) {
	var comment Comment
	err := db.C("comments").FindId(bson.ObjectIdHex(id)).One(&comment)
	return comment, err
}

// Insert a comment into database
func (m *VideosDAO) InsertComment(comment Comment) error {
	err := db.C("comments").Insert(&comment)
	return err
}

// Delete an existing comment
func (m *VideosDAO) DeleteComment(comment Comment) error {
	err := db.C("comments").Remove(&comment)
	return err
}

// Update an existing comment
func (m *VideosDAO) UpdateComment(comment Comment) error {
	err := db.C("comments").UpdateId(comment.ID, &comment)
	return err
}

// Find list of categories
func (m *VideosDAO) FindAllCategories() ([]Category, error) {
	var categories []Category
	err := db.C("categories").Find(bson.M{}).All(&categories)
	return categories, err
}

// Find a category by its id
func (m *VideosDAO) FindCategoryById(id string) (Category, error) {
	var category Category
	err := db.C("categories").FindId(bson.ObjectIdHex(id)).One(&category)
	return category, err
}

// Insert a category into database
func (m *VideosDAO) InsertCategory(category Category) error {
	err := db.C("categories").Insert(&category)
	return err
}

// Delete an existing category
func (m *VideosDAO) DeleteCategory(category Category) error {
	err := db.C("categories").Remove(&category)
	return err
}

// Update an existing video
func (m *VideosDAO) UpdateCategory(category Category) error {
	err := db.C("categories").UpdateId(category.ID, &category)
	return err
}
