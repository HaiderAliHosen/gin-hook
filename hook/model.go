package hook

//Hook struct
type Hook struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	Auth bool   `json:"auth" bson:"auth"`
}
