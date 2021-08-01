package model

type User struct {
	BaseModel
	Name string `json:"name" bson:"name"`
}
