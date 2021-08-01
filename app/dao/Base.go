package dao

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type BaseDao struct {
	Id        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
	CreatedBy string        `json:"created_by" bson:"created_by"`
	UpdatedBy string        `json:"updated_by" bson:"updated_by"`
}
