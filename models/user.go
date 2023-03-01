package models

// - get stats: GET /users/<userID>/stats
type User struct {
	Name string // json:"name" bson:"user_name"
 	userID string
}