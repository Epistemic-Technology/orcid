package orcid_test

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Epistemic-Technology/orcid/orcid"
)

func ExampleClient_GetRecord() {
	// For public API access, you need a bearer token
	// Get your token from https://orcid.org/developer-tools
	client := orcid.NewClient(
		orcid.WithBearerToken("your-bearer-token-here"),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	record, err := client.GetRecord(ctx, "0000-0002-1825-0097")
	if err != nil {
		log.Fatal(err)
	}

	if record.Person != nil && record.Person.Name != nil {
		if record.Person.Name.GivenNames != nil {
			fmt.Printf("Given Names: %s\n", record.Person.Name.GivenNames.Value)
		}
		if record.Person.Name.FamilyName != nil {
			fmt.Printf("Family Name: %s\n", record.Person.Name.FamilyName.Value)
		}
	}
}

func ExampleClient_Search() {
	client := orcid.NewClient()

	ctx := context.Background()

	params := orcid.SearchParams{
		Query: "family-name:Smith AND given-names:John",
		Start: 0,
		Rows:  10,
	}

	results, err := client.Search(ctx, params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d results\n", results.NumFound)
	for _, result := range results.Results {
		if result.OrcidIdentifier != nil {
			fmt.Printf("ORCID: %s\n", result.OrcidIdentifier.Path)
		}
	}
}

func ExampleSearchQuery() {
	client := orcid.NewClient()
	ctx := context.Background()

	query := orcid.NewSearchQuery().
		FamilyName("Einstein").
		And().
		GivenNames("Albert").
		WithRows(5)

	results, err := client.SearchWithQuery(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d results for Albert Einstein\n", results.NumFound)
}

func ExampleClient_SearchIter() {
	client := orcid.NewClient()
	ctx := context.Background()

	query := orcid.NewSearchQuery().
		Keyword("machine learning").
		WithRows(100)

	iter := client.SearchIterWithQuery(ctx, query)

	count := 0
	for iter.Next() {
		record := iter.Value()
		if record != nil && record.OrcidIdentifier != nil {
			count++
			if count > 5 {
				break
			}
			fmt.Printf("ORCID: %s\n", record.OrcidIdentifier.Path)
		}
	}

	if err := iter.Error(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total results available: %d\n", iter.TotalResults())
}

func ExampleClient_GetWorks() {
	client := orcid.NewClient()
	ctx := context.Background()

	works, err := client.GetWorks(ctx, "0000-0002-1825-0097")
	if err != nil {
		log.Fatal(err)
	}

	for _, group := range works.WorkGroup {
		for _, summary := range group.WorkSummary {
			if summary.Title != nil && summary.Title.Title != nil {
				fmt.Printf("Work: %s (Type: %s)\n",
					summary.Title.Title.Value,
					summary.Type)
			}
		}
	}
}

func ExampleValidateOrcidID() {
	orcidIDs := []string{
		"0000-0002-1825-0097",
		"https://orcid.org/0000-0002-1825-0097",
		"0000-0000-0000-0000",
		"invalid-orcid",
	}

	for _, id := range orcidIDs {
		err := orcid.ValidateOrcidID(id)
		if err != nil {
			fmt.Printf("%s: Invalid - %v\n", id, err)
		} else {
			fmt.Printf("%s: Valid\n", id)
		}
	}
}

func ExampleClient_withOptions() {
	client := orcid.NewClient(
		orcid.WithBearerToken("your-bearer-token-here"),
		orcid.WithTimeout(60*time.Second),
		orcid.WithRateLimit(5),
		orcid.WithMaxRetries(5),
		orcid.WithContentType(orcid.ContentTypeJSON),
		orcid.WithUserAgent("MyApp/1.0"),
	)

	ctx := context.Background()
	person, err := client.GetPerson(ctx, "0000-0002-1825-0097")
	if err != nil {
		log.Fatal(err)
	}

	if person.Biography != nil {
		fmt.Printf("Biography: %s\n", person.Biography.Content)
	}
}

func ExampleSearchQuery_complex() {
	client := orcid.NewClient()
	ctx := context.Background()

	query := orcid.NewSearchQuery().
		AffiliationOrganization("MIT").
		And().
		RawQuery("(keyword:physics OR keyword:\"quantum computing\")").
		And().
		Not().
		FamilyName("Test").
		WithRows(20)

	results, err := client.SearchWithQuery(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Found %d MIT researchers in physics/quantum computing\n", results.NumFound)
}

func ExampleClient_GetEducations() {
	client := orcid.NewClient()
	ctx := context.Background()

	educations, err := client.GetEducations(ctx, "0000-0002-1825-0097")
	if err != nil {
		log.Fatal(err)
	}

	for _, edu := range educations.EducationSummary {
		if edu.Organization != nil {
			fmt.Printf("Education: %s", edu.Organization.Name)
			if edu.RoleTitle != "" {
				fmt.Printf(" - %s", edu.RoleTitle)
			}
			if edu.StartDate != nil && edu.StartDate.Year != nil {
				fmt.Printf(" (Started: %s)", edu.StartDate.Year.Value)
			}
			fmt.Println()
		}
	}
}

func ExampleClient_GetEmployments() {
	client := orcid.NewClient()
	ctx := context.Background()

	employments, err := client.GetEmployments(ctx, "0000-0002-1825-0097")
	if err != nil {
		log.Fatal(err)
	}

	for _, emp := range employments.EmploymentSummary {
		if emp.Organization != nil {
			fmt.Printf("Employment: %s", emp.Organization.Name)
			if emp.RoleTitle != "" {
				fmt.Printf(" - %s", emp.RoleTitle)
			}
			if emp.StartDate != nil && emp.StartDate.Year != nil {
				fmt.Printf(" (Since: %s)", emp.StartDate.Year.Value)
			}
			fmt.Println()
		}
	}
}
