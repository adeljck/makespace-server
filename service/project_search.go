package service

import "time"

type ProjectSearch struct {
	StartTime time.Time `json:"start_time"`
	StopTime  time.Time `json:"stop_time"`

}
