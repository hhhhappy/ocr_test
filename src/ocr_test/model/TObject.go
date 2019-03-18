package model

import "gopkg.in/mgo.v2/bson"

type TObject struct {
	Id bson.ObjectId `bson:"_id"`
}