package govidat

import (
	"fmt"
	"strconv"
	"time"
)

type Municipality struct {
	Id                string
	Updated           time.Time
	Name              string
	Population        int64
	Dose1             int64
	Dose2             int64
	Dose3             int64
	Certs             int64
	PercentVaccinated float64
}

const (
	timestampFormat = "2006-01-02T15:04:05-07:00"
)

func municipalityFromRecord(rec []string) (*Municipality, error) {
	if len(rec) != 9 {
		return nil, fmt.Errorf("record has incorrect number of records: %v", rec)
	}
	muni := new(Municipality)
	//record format:
	//Updated, Id, Name, Population, Dose1, Dose2, Dose3, Certs, PercentVaccinated
	if updated, err := time.Parse(timestampFormat, rec[0]); err != nil {
		return nil, fmt.Errorf("record has invalid timestamp: %v", rec)
	} else {
		muni.Updated = updated
	}
	muni.Id = rec[1]
	muni.Name = rec[2]
	if population, err := strconv.ParseInt(rec[3], 10, 64); err != nil {
		return nil, fmt.Errorf("population is not a valid integer: %v", rec)
	} else {
		muni.Population = population
	}
	if d1, err := strconv.ParseInt(rec[4], 10, 64); err != nil {
		return nil, fmt.Errorf("dose1 is not a valid integer: %v", rec)
	} else {
		muni.Dose1 = d1
	}
	if d2, err := strconv.ParseInt(rec[5], 10, 64); err != nil {
		return nil, fmt.Errorf("dose2 is not a valid integer: %v", rec)
	} else {
		muni.Dose2 = d2
	}
	if d3, err := strconv.ParseInt(rec[6], 10, 64); err != nil {
		return nil, fmt.Errorf("dose3 is not a valid integer: %v", rec)
	} else {
		muni.Dose3 = d3
	}
	if certs, err := strconv.ParseInt(rec[7], 10, 64); err != nil {
		return nil, fmt.Errorf("certs is not a valid integer: %v", rec)
	} else {
		muni.Certs = certs
	}
	if percent_vax, err := strconv.ParseFloat(rec[8], 64); err != nil {
		return nil, fmt.Errorf("percentVaccinated is not a valid float: %v", rec)
	} else {
		muni.PercentVaccinated = percent_vax
	}
	return muni, nil
}
