package dao

import (
	"log"

	"github.com/xmin-github/go-restapi-example/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MoviesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "movies"
)

// Establish a connection to database
func (m *MoviesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of movies
func (m *MoviesDAO) FindAll() ([]models.Movie, error) {
	var movies []models.Movie
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

// Find a movie by its id
func (m *MoviesDAO) FindById(id string) (models.Movie, error) {
	var movie models.Movie
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

// Insert a movie into database
func (m *MoviesDAO) Insert(movie models.Movie) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

// Delete an existing movie
func (m *MoviesDAO) Delete(movie models.Movie) error {
	err := db.C(COLLECTION).Remove(&movie)
	return err
}

// Update an existing movie
func (m *MoviesDAO) Update(movie models.Movie) error {
	err := db.C(COLLECTION).UpdateId(movie.ID, &movie)
	return err
}
