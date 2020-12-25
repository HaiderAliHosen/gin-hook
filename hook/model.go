package hook

//Hook struct
type Hook struct {
	ID      string `json:"id" bson:"_id" `
	Name    string `json:"name" bson:"name" binding:"required"`
	Address string `json:"address" bson:"address" binding:"required"`
	Country string `json:"country" bson:"country"`
	Number  int64  `json:"number" bson:"number"`
	Auth    bool   `json:"auth" bson:"auth" binding:"required"`
}
