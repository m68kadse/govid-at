package govidat

import (
	"fmt"
	"strconv"
)

type County struct {
	Id             string
	Name           string
	Residents      int64
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
	//Name;id;Residents;Cases;Deaths;Incidence7d
	county.Name = rec[0]
	county.Id = rec[1]
	if residents, err := strconv.ParseInt(rec[2], 10, 64); err != nil {
		return nil, fmt.Errorf("cannot parse record, residents not a valid integer: %v", rec)
	} else {
		county.Residents = residents
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
