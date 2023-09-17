package helper

import (
	"context"

	"github.com/abdulmanafc2001/mongodb/database"
	"github.com/abdulmanafc2001/mongodb/models"
	"gopkg.in/mgo.v2/bson"
)

func AddOneMovie(movie models.Movie) (interface{}, error) {
	res, err := database.Collections.InsertOne(context.Background(), movie)
	if err != nil {
		return nil, err
	}
	return res, nil

}

func ListAllMovies() ([]bson.M,error) {
	cur , err := database.Collections.Find(context.Background(), bson.M{})
	if err != nil {
		return nil,err
	}
	var movies []bson.M

	for cur.Next(context.Background()) {
		var movie bson.M
		if err = cur.Decode(&movie); err != nil {
			return nil,err
		}
		movies = append(movies, movie)
	}
	return movies,nil
}
