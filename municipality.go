package govidat

import (
	"time"
)

type Municipality struct {
	id                 string
	Updated            time.Time
	Name               string
	Population         int64
	Dose1              int64
	Dose2              int64
	Dose3              int64
	Certs              int64
	Percent_vaccinated float64
}
