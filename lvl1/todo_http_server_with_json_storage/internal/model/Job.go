package model

import "time"

type Job struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	FinishAt    time.Time `json:"finishAt"`
	Status      string    `json:"status"`
}
