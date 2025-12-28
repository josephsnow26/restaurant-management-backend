package models

type User struct {
	BaseModel
	FirstName string `bson:"first_name" json:"first_name"`
	LastName  string `bson:"last_name" json:"last_name"`
	Email     string `bson:"email" json:"email"`
	Password  string `bson:"password,omitempty" json:"-"` // hashed password, never return in JSON
}
