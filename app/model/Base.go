package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type BaseModel struct {
	Id        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at"`
	CreatedBy string        `json:"created_by" bson:"created_by"`
	UpdatedBy string        `json:"updated_by" bson:"updated_by"`
}
