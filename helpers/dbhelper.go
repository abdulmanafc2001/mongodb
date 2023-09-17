package helper

import (
	"context"

	"github.com/abdulmanafc2001/mongodb/database"
	"github.com/abdulmanafc2001/mongodb/models"
)

func AddOneMovie(movie models.Movie) (interface{}, error) {
	res, err := database.Collections.InsertOne(context.Background(), movie)
	if err != nil {
		return nil, err
	}
	return res, nil

}
