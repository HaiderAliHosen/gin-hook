package hook

import "time"

//Hook struct
type Hook struct {
	ID      string    `json:"id" bson:"_id" `
	Title   string    `json:"title" bson:"title" binding:"required"`
	Date    time.Time `json:"date" bson:"date" binding:"required"`
	Content string    `json:"content" bson:"content" binding:"required"`
	// Address string    `json:"address" bson:"address" binding:"required"`
	// Country string    `json:"country" bson:"country"`
	// Number  int64     `json:"number" bson:"number"`
	Auth bool `json:"auth" bson:"auth" `
}
