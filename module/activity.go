package module

import "time"

type Activity struct {
	StartTime        time.Time `bson:"start_time" json:"start_time"`
	StopTime         time.Time `bson:"stop_time" json:"stop_time"`
	CreateTime       time.Time `bson:"create_time" json:"create_time"`
	ActivityName     string    `bson:"activity_name" json:"acvitity_name"`
	Creator          string    `bson:"creator" json:"creator"`
	Status           int       `bson:"status" json:"status"`
	Info             string    `bson:"info" json:"info"`
	ShortInfo        string    `bson:"short_info" json:"short_info"`
	Attach           []Attach  `bson:"attach" json:"attach"`
	Avatar           string    `bson:"avatar" json:"avatar"`
	Contact          Contact   `bson:"contact" json:"contact"`
	CommentCount     int       `bson:"comment_count" json:"comment_count"`
	ActivityType     string    `bson:"activity_type" json:"activity_type"`
	ActivityLocation string    `bson:"activity_location" json:"activity_location"`
	Joiner           []string  `bson:"joiner" json:"joiner"`
}
