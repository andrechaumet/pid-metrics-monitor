package model

import (
	"time"
)

/*type MetricModel struct {
	ID                int
	StartTime         time.Time
	LastUpdate        time.Time
	CurrentIterations int64
	CurrentSpeed 	  string
	TotalIterations   int64
	Percentage        float64
	LapsedTime		  int64
	ExpectedTime	  time.Time	
}

type PidModel struct {
	ID                int
	StartTime         time.Time
	LastUpdate        time.Time
	CurrentIterations int64
	TotalIterations   int64
	Logs 			  []string
}*/

type PidModel struct {
	ID                int
	StartTime         time.Time
	LastUpdate        time.Time
	CurrentIterations int64
	CurrentSpeed 	  string
	TotalIterations   int64
	Percentage        float64
	LapsedTime		  int64
	ExpectedTime	  time.Time
	Logs 			  []string
}
