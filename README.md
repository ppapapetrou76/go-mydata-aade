[![Codacy Badge](https://api.codacy.com/project/badge/Grade/70aaf3cfcd9d46f08ba1de5eb4156577)](https://app.codacy.com/manual/ppapapetrou76/go-data-gov-gr-sdk?utm_source=github.com&utm_medium=referral&utm_content=ppapapetrou76/go-testing&utm_campaign=Badge_Grade_Dashboard)
[![codebeat badge](https://codebeat.co/badges/def76d1b-0889-4ff4-908c-b333349bd136)](https://codebeat.co/projects/github-com-ppapapetrou76-go-data-gov-gr-sdk-main)
[![Fluent Go Testing](https://circleci.com/gh/circleci/circleci-docs.svg?style=shield)](https://app.circleci.com/pipelines/github/ppapapetrou76/go-data-gov-gr-sdk?branch=master)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ppapapetrou76_go-data-gov-gr-sdk&metric=alert_status)](https://sonarcloud.io/dashboard?id=ppapapetrou76_go-data-gov-gr-sdk)
[![codecov](https://codecov.io/gh/ppapapetrou76/go-data-gov-gr-sdk/branch/main/graph/badge.svg?token=CX3I6LDF3J)](https://codecov.io/gh/ppapapetrou76/go-data-gov-gr-sdk)
[![Go Report Card](https://goreportcard.com/badge/github.com/ppapapetrou76/go-data-gov-gr-sdk)](https://goreportcard.com/report/github.com/ppapapetrou76/go-data-gov-gr-sdk)
[![GoDoc](https://godoc.org/github.com/ppapapetrou76/go-data-gov-gr-sdk?status.svg)](https://pkg.go.dev/github.com/ppapapetrou76/go-data-gov-gr-sdk)

# go-data-gov-gr-sdk
A Go based SDK to access the public data provided by the Greek Government and are available at https://www.data.gov.gr/

## Quick Start

### Get your API Token
Submit the form found [here](https://www.data.gov.gr/token/)
You will receive by email an API token. You will use it to access the data in the example below

### Add the SDK as a dependency to your project 

`go get github.com/ppapapetrou76/go-data-gov-gr-sdk`

or if you are using go modules ( recommended ) 

`go mod download github.com/ppapapetrou76/go-data-gov-gr-sdk` 

### Implement a client to read some data
```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/ppapapetrou76/go-data-gov-gr-sdk/api"
	"github.com/ppapapetrou76/go-data-gov-gr-sdk/pkg/vaccination"
)

func main() {
	// Fetches the vaccination data for the last 6 days for all areas
	client := api.NewClient("<YOUR_API_TOKEN_HERE>")
	vaccinationData, err := vaccination.Get(client,
		api.NewDefaultGetParams(api.SetDateFrom(time.Now().Add(-fiveDays))),
	)
	if err != nil {
		panic(err)
	}
	// Filter by a specific region
	for _, d := range vaccinationData.FilterByArea("ΘΕΣΣΑΛΟΝΙΚΗΣ") {
		fmt.Fprintf(os.Stdout, "Area:%s, Vaccinations on %v:%d\n", d.Area, d.ReferenceDate, d.DayTotal)
	}
}
```

## Implemented endpoints
  * COVID-19 vaccination statistics ( https://www.data.gov.gr/datasets/mdg_emvolio ) 

## To be implemented soon
- [ ] Business and Economy (Small businesses, industry, imports, exports and trade)
- [ ] Crime and Justice (Courts, police, prison, offenders, borders and immigration)
- [ ] Education (Students, training, universities, quaifications)
- [ ] Environment (Weather, flooding, rivers, air quality, geology and agriculture) 
- [ ] Health (Includes smoking, drugs, alcohol, medicine performance and hospitals)
- [ ] Society (Employment, benefits, household finances, poverty and population)
- [ ] Technology (Internet, technology and digital adoption)
- [ ] Telecommunication (Telecommunications data, television and radio) 
- [ ] Transport (Airports, roads, freight, electric vehicles, parking, buses and footpaths)


## data.gov.gr clients / SDK in other languages

- [x] Python [https://github.com/ilias-ant/pydatagovgr](https://github.com/ilias-ant/pydatagovgr)
- [ ] Java
- [ ] Javascript
- [ ] C++
- [ ] C#
