package model

import (
	"time"
)

type Status int

const (
	Running Status = iota
	Completed
	Failed
)

type PidModel struct {
	ID                int
	Name              string
	StartTime         *time.Time
	LastUpdate        *time.Time
	CurrentIterations int
	TotalIterations   int
	Status            Status
	Logs              []string
}

func (p *PidModel) CurrentSpeed() float64 {
	if p.StartTime == nil || p.LastUpdate == nil {
		return 0
	}
	timeDifference := p.LastUpdate.Sub(*p.StartTime).Seconds()
	if timeDifference == 0 {
		return 0
	}
	return float64(p.CurrentIterations) / timeDifference
}

func (p *PidModel) ExpectedTime() float64 {
	currentSpeed := p.CurrentSpeed()
	if currentSpeed == 0 {
		return 0
	}
	return float64(p.TotalIterations-p.CurrentIterations) / currentSpeed
}

func (p *PidModel) Percentage() float64 {
	if p.TotalIterations == 0 {
		return 0.0
	}
	return float64(p.CurrentIterations) / float64(p.TotalIterations) * 100.0
}

func (p *PidModel) LapsedTime() int {
	if p.StartTime == nil {
		return 0
	}
	return int(time.Since(*p.StartTime).Seconds())
}
