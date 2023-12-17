package model

import (
	"time"
)

type MetricModel struct {
	ID                int64
	StartTime         time.Time
	LastUpdate        time.Time
	CurrentIterations int64
	CurrentSpeed 	  string
	TotalIterations   int64
	Percentage        float64
	LapsedTime		  int64
	ExpectedTime	  time.Time	
}
