package model

import (
	"time"
)

type PidModel struct {
	ID                int
	StartTime         time.Time
	LastUpdate        time.Time
	CurrentIterations int
	CurrentSpeed 	  string
	TotalIterations   int
	Percentage        float64
	LapsedTime		  int
	ExpectedTime	  time.Time
	Logs 			  []string
}