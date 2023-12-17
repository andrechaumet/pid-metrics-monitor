package model

import (
	"time"
)

type TaskModel struct {
	ID                int64
	StartTime         time.Time
	LastUpdate        time.Time
	CurrentIterations int64
	TotalIterations   int64
	Logs 			  []string
}