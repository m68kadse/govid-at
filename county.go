package govidat

import (
	"fmt"
	"strconv"
)

type County struct {
	Id             string
	Name           string
	Population     int64
	Cases          int64
	Deaths         int64
	Cases7d        int64
	Municipalities []Municipality
}

func countyFromRecord(rec []string) (*County, error) {
	if len(rec) != 6 {
		return nil, fmt.Errorf("record has incorrect number of fields: %v", rec)
	}
	county := new(County)

	//record format:
	//Name;id;Residents;Cases;Deaths;Cases7d
	county.Name = rec[0]
	county.Id = rec[1]
	if population, err := strconv.ParseInt(rec[2], 10, 64); err != nil {
		return nil, fmt.Errorf("cannot parse record, population not a valid integer: %v", rec)
	} else {
		county.Population = population
	}
	if cases, err := strconv.ParseInt(rec[3], 10, 64); err != nil {
		return nil, fmt.Errorf("cannot parse record, cases not a valid integer: %v", rec)
	} else {
		county.Cases = cases
	}
	if deaths, err := strconv.ParseInt(rec[4], 10, 64); err != nil {
		return nil, fmt.Errorf("cannot parse record, deaths not a valid integer: %v", rec)
	} else {
		county.Deaths = deaths
	}
	if cases7d, err := strconv.ParseInt(rec[5], 10, 64); err != nil {
		return nil, fmt.Errorf("cannot parse record, 7 day cases not a valid integer: %v", rec)
	} else {
		county.Cases7d = cases7d
	}

	return county, nil
}

func (c County) Dose1() int64 {
	var total int64
	for _, m := range c.Municipalities {
		total += m.Dose1
	}
	return total
}

func (c County) Dose2() int64 {
	var total int64
	for _, m := range c.Municipalities {
		total += m.Dose2
	}
	return total
}

func (c County) Dose3() int64 {
	var total int64
	for _, m := range c.Municipalities {
		total += m.Dose3
	}
	return total
}

func (c County) Certs() int64 {
	var total int64
	for _, m := range c.Municipalities {
		total += m.Certs
	}
	return total
}

func (c County) PercentVaccinated() float64 {
	return float64(c.Certs()*100) / float64(c.Population)
}

func (c County) Cases7dPer100k() int64 {
	return int64(float64(c.Cases7d) / (float64(c.Population) / 100000))
}
