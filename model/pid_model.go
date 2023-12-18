package model

import (
	"time"
)

type PidModel struct {
	ID                int
	StartTime         time.Time
	LastUpdate        time.Time
	CurrentIterations int64
	TotalIterations   int64
	Logs 			  []string
}