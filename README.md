# ORCID Go Client

A Go client library for interacting with the ORCID Public API v3.0.

## Installation

```bash
go get github.com/Epistemic-Technology/orcid
```

## Features

- Full support for ORCID Public API v3.0 endpoints
- Automatic retry logic with exponential backoff
- Built-in rate limiting
- Search with fluent query builder and pagination
- Support for both JSON and XML formats
- Context-aware for cancellation and timeouts

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/Epistemic-Technology/orcid/orcid"
)

func main() {
    client := orcid.NewClient()

    ctx := context.Background()
    record, err := client.GetRecord(ctx, "0000-0002-1825-0097")
    if err != nil {
        log.Fatal(err)
    }

    if record.Person != nil && record.Person.Name != nil {
        fmt.Printf("Name: %s %s\n",
            record.Person.Name.GivenNames.Value,
            record.Person.Name.FamilyName.Value)
    }
}
```

## Client Configuration

```go
client := orcid.NewClient(
    orcid.WithTimeout(60*time.Second),
    orcid.WithRateLimit(10), // 10 requests per second
    orcid.WithMaxRetries(5),
    orcid.WithContentType(orcid.ContentTypeJSON),
    orcid.WithUserAgent("MyApp/1.0"),
)
```

## Search

```go
// Simple search
params := orcid.SearchParams{
    Query: "family-name:Smith AND given-names:John",
    Start: 0,
    Rows:  10,
}
results, err := client.Search(ctx, params)

// Query builder
query := orcid.NewSearchQuery().
    FamilyName("Einstein").
    And().
    GivenNames("Albert").
    WithRows(5)
results, err := client.SearchWithQuery(ctx, query)

// Iterator for large result sets
iter := client.SearchIterWithQuery(ctx, orcid.NewSearchQuery().
    Keyword("machine learning").
    WithRows(100))

for iter.Next() {
    record := iter.Value()
    fmt.Printf("ORCID: %s\n", record.OrcidIdentifier.Path)
}
if err := iter.Error(); err != nil {
    log.Fatal(err)
}
```

## API Methods

### Core Methods
- `GetRecord(ctx, orcidID)` - Complete record
- `GetPerson(ctx, orcidID)` - Person details
- `GetWorks(ctx, orcidID)` - Works/publications
- `GetWork(ctx, orcidID, putCode)` - Specific work

### Affiliations
- `GetEducations(ctx, orcidID)`
- `GetEmployments(ctx, orcidID)`
- `GetDistinctions(ctx, orcidID)`
- `GetInvitedPositions(ctx, orcidID)`
- `GetMemberships(ctx, orcidID)`
- `GetQualifications(ctx, orcidID)`
- `GetServices(ctx, orcidID)`

### Activities
- `GetFundings(ctx, orcidID)`
- `GetPeerReviews(ctx, orcidID)`
- `GetResearchResources(ctx, orcidID)`

## Search Query Builder

Supported fields: `ORCID()`, `Email()`, `FamilyName()`, `GivenNames()`, `CreditName()`, `OtherNames()`, `Keyword()`, `ExternalIdentifier()`, `DOI()`, `WorkTitle()`, `FundingTitle()`, `AffiliationOrganization()`, `RINGGOLD()`, `GRID()`, `ROR()`, `FundRef()`

Combine with `.And()`, `.Or()`, `.Not()`

## License

MIT
