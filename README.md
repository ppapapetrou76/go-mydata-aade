[![Codacy Badge](https://app.codacy.com/project/badge/Grade/835877cc4dbf479b92178bbe1e3c0fdc)](https://www.codacy.com/gh/ppapapetrou76/go-mydata-aade/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=ppapapetrou76/go-mydata-aade&amp;utm_campaign=Badge_Grade)
[![codebeat badge](https://codebeat.co/badges/91b671db-b1f8-49cc-b299-fc772f45ff52)](https://codebeat.co/projects/github-com-ppapapetrou76-go-mydata-aade-main)
[![CircleCI](https://circleci.com/gh/ppapapetrou76/go-mydata-aade/tree/main.svg?style=svg)](https://circleci.com/gh/ppapapetrou76/go-mydata-aade/tree/main)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ppapapetrou76_go-mydata-aade&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=ppapapetrou76_go-mydata-aade)
[![codecov](https://codecov.io/gh/ppapapetrou76/go-mydata-aade/branch/main/graph/badge.svg?token=CX3I6LDF3J)](https://codecov.io/gh/ppapapetrou76/go-mydata-aade)
[![Go Report Card](https://goreportcard.com/badge/github.com/ppapapetrou76/go-mydata-aade)](https://goreportcard.com/report/github.com/ppapapetrou76/go-mydata-aade)
[![GoDoc](https://godoc.org/github.com/ppapapetrou76/go-mydata-aade?status.svg)](https://pkg.go.dev/github.com/ppapapetrou76/go-mydata-aade)

# go-myDATA-aade
A Go based SDK to communicate with the [myDATA Rest API](https://mydata-prod-apim.portal.azure-api.net/docs/services/mydata-prod-api-func/operations/post-cancelinvoice) provided by the Greek Government

## Supported methods
### ERP Methods
- [x] RequestDocs (Returns all invoices submitted by other parties and the receiver is the authorized entity)
- [x] RequestTransmittedDocs (Returns all invoices submitted by the authorized entity to third-parties)
- [x] CancelInvoice (Cancels an already issued invoice without issuing a new one)
- [x] SendInvoices (Submits one or more invoices for a given entity)

### Provider Methods

## To be implemented soon
### ERP Methods
- [ ] SendExpensesClassification
- [ ] SendIncomeClassification

### Provider Methods
- [ ] SendInvoices
- [ ] RequestTransmittedDocs
- [ ] RequestReceiverInfo
