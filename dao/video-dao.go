package dao

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"strconv"

	. "video-reproduction-ms/models"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type VideosDAO struct {
	Server   string
	Database string
}

var DB *mgo.Database

// Establish a connection to database
func (m *VideosDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	DB = session.DB(m.Database)
}

// Find list of videos
func (m *VideosDAO) FindAllVideos() ([]Video, error) {
	var videos []Video
	err := DB.C("videos").Find(bson.M{}).All(&videos)
	return videos, err
}

// Find a video by its id
func (m *VideosDAO) FindVideoById(id string) (Video, error) {
	var video Video
	err := DB.C("videos").FindId(bson.ObjectIdHex(id)).One(&video)
	return video, err
}

//Find a videos by Categories
func (m *VideosDAO) FindVideoByCategory(id string) ([]Video, error) {
	var videos []Video
	err := DB.C("videos").Find(bson.M{"category_id": id}).All(&videos)
	return videos, err
}

// Insert a video into database
func (m *VideosDAO) InsertVideo(video Video) error {
	err := DB.C("videos").Insert(&video)
	return err
}

// Delete an existing video
func (m *VideosDAO) DeleteVideo(id string) error {
	var video Video
	err := DB.C("videos").FindId(bson.ObjectIdHex(id)).One(&video)
	err = DB.C("videos").Remove(&video)
	return err
}

// Update an existing video
func (m *VideosDAO) UpdateVideo(video Video) error {
	err := DB.C("videos").UpdateId(video.ID, &video)
	return err
}

// Find list of comments
func (m *VideosDAO) FindCommentsByVideoId(id string) ([]Comment, error) {

	var comments []Comment
	err := DB.C("comments").Find(bson.M{"video_id": id}).All(&comments)
	return comments, err
}

//Find Videos By UserId
func (m *VideosDAO) FindVideosByUserId(id string) ([]Video, error) {

	user_id, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(user_id)
	}
	var videos []Video
	err = DB.C("videos").Find(bson.M{"user_id": user_id}).All(&videos)
	return videos, err
}

// Find list of comments
func (m *VideosDAO) FindAllComments() ([]Comment, error) {
	var comments []Comment
	err := DB.C("comments").Find(bson.M{}).
		All(&comments)
	return comments, err
}

// Find a comment by its id
func (m *VideosDAO) FindCommentById(id string) (Comment, error) {
	var comment Comment
	err := DB.C("comments").FindId(bson.ObjectIdHex(id)).One(&comment)
	return comment, err
}

// Insert a comment into database
func (m *VideosDAO) InsertComment(comment Comment) error {
	err := DB.C("comments").Insert(&comment)
	return err
}

// Delete an existing comment
func (m *VideosDAO) DeleteComment(comment Comment) error {
	err := DB.C("comments").Remove(&comment)
	return err
}

// Update an existing comment
func (m *VideosDAO) UpdateComment(comment Comment) error {
	err := DB.C("comments").UpdateId(comment.ID, &comment)
	return err
}

// Find list of categories
func (m *VideosDAO) FindAllCategories() ([]Category, error) {
	var categories []Category
	err := DB.C("categories").Find(bson.M{}).All(&categories)
	return categories, err
}

// Find a category by its id
func (m *VideosDAO) FindCategoryById(id string) (Category, error) {
	var category Category
	err := DB.C("categories").FindId(bson.ObjectIdHex(id)).One(&category)
	return category, err
}

func (m *VideosDAO) FindVideosByName(title string) ([]Video, error) {

	var videos []Video
	err := DB.C("videos").Find(bson.M{"title": bson.M{"$regex": bson.RegEx{Pattern: title, Options: "i"}}}).All(&videos)
	return videos, err

}

// Insert a category into database
func (m *VideosDAO) InsertCategory(category Category) error {
	err := DB.C("categories").Insert(&category)
	return err
}

// Delete an existing category
func (m *VideosDAO) DeleteCategory(category Category) error {
	err := DB.C("categories").Remove(&category)
	return err
}

// Update an existing video
func (m *VideosDAO) UpdateCategory(category Category) error {
	err := DB.C("categories").UpdateId(category.ID, &category)
	return err
}

// Upload video to gridFS
func (m *VideosDAO) UploadVideo(file multipart.File) error {

	data, err := ioutil.ReadAll(file)
	// ... check err value for nil

	// Specify the Mongodb database
	my_db := DB

	// Create the file in the Mongodb Gridfs instance
	my_file, err := my_db.GridFS("fs").Create("new2.mp4")
	// ... check err value for nil

	// Write the file to the Mongodb Gridfs instance
	n, err := my_file.Write(data)
	// ... check err value for nil

	// Close the file
	err = my_file.Close()
	// ... check err value for nil

	// Write a log type message
	fmt.Println(n)

	fmt.Printf("%d bytes written to the Mongodb instance\n", my_file.Name)

	return err
}
