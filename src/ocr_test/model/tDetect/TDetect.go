package tDetect

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"ocr_test/manager/mongoManager"
)

type TDetect struct {
	Id         bson.ObjectId `bson:"_id"`
	Content    []string      `bson:"content"`
	ImageFile  string        `bson:"image_file"`
	DetectTime string         `bson:"detect_time"`
}

var collection = "t_detect"

/* Get all the api */
func FindAll() ([]TDetect, error) {
	var result []TDetect
	query := bson.M{}
	exop := func(c *mgo.Collection) error {
		return c.Find(query).All(&result)
	}
	err := mongoManager.WitchCollection(collection, exop)
	if err != nil {
		return nil, err
	}
	return result, nil
}

/*
	insert the object
*/
func (object TDetect) Insert() (string, error) {

	object.Id = bson.NewObjectId()
	exop := func(c *mgo.Collection) error {
		return c.Insert(object)
	}

	err := mongoManager.WitchCollection(collection, exop)
	if err != nil {
		return "", err
	}
	return object.Id.Hex(), nil
}
