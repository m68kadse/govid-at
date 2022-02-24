package govidat

import (
	"encoding/csv"
	"net/http"
)

const (
	COUNTY_CASE_URL               = "https://covid19-dashboard.ages.at/data/CovidFaelle_GKZ.csv"
	MUNICIPALITY_VACCINATIONS_URL = "https://info.gesundheitsministerium.gv.at/data/COVID19_vaccination_municipalities.csv"
)

func LoadData() ([]County, error) {
	ccResponse, ccResponseErr := http.Get(COUNTY_CASE_URL)
	if ccResponseErr != nil {
		return nil, ccResponseErr
	}
	ccReader := csv.NewReader(ccResponse.Body)
	ccReader.Comma = ';'
	ccRecords, ccRecErr := ccReader.ReadAll()
	if ccRecErr != nil {
		return nil, ccRecErr
	}

	mvResponse, mvErr := http.Get(MUNICIPALITY_VACCINATIONS_URL)
	if mvErr != nil {
		return nil, mvErr
	}
	mvReader := csv.NewReader(mvResponse.Body)
	mvReader.Comma = ';'
	mvRecords, mvRecErr := mvReader.ReadAll()
	if mvRecErr != nil {
		return nil, mvRecErr
	}

	//remove field name records
	ccRecords = ccRecords[1:]
	mvRecords = mvRecords[1:]

	counties := make(map[string]County)
	for _, rec := range ccRecords {
		if county, err := countyFromRecord(rec); err != nil {
			return nil, err
		} else {
			counties[county.Id] = *county
		}
	}

	for _, rec := range mvRecords {
		if muni, err := municipalityFromRecord(rec); err != nil {
			return nil, err
		} else {
			if muni.Id[0] == '9' { //special case for Vienna
				cty := counties["900"]
				cty.Municipalities = append(cty.Municipalities, *muni)
				counties["900"] = cty
			} else {
				cty := counties[muni.Id[0:3]]
				cty.Municipalities = append(cty.Municipalities, *muni)
				counties[muni.Id[0:3]] = cty
			}
		}
	}

	countyList := make([]County, 0, 94)
	for _, c := range counties {
		countyList = append(countyList, c)
	}

	return countyList, nil

}
