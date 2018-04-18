package models

import "gopkg.in/mgo.v2/bson"

type Character struct {
	ID     bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
	Name   string        `bson:"name" json:"name"`
	Life   int16         `bson:"life" json:"life"`
	Strong int16         `bson:"strong"`
}
