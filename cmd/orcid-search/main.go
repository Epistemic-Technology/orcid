package main

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"os"

	"github.com/Epistemic-Technology/orcid/orcid"
)

func main() {
	var (
		bearerToken string
		searchQuery string
		orcidID     string
		sandbox     bool
		useXML      bool
		raw         bool
		rows        int
		start       int
	)

	flag.StringVar(&bearerToken, "token", "", "Bearer token for ORCID API authentication (required)")
	flag.StringVar(&bearerToken, "t", "", "Bearer token for ORCID API authentication (shorthand)")
	flag.StringVar(&searchQuery, "query", "", "Search query string")
	flag.StringVar(&searchQuery, "q", "", "Search query string (shorthand)")
	flag.StringVar(&orcidID, "orcid", "", "ORCID ID to retrieve")
	flag.StringVar(&orcidID, "o", "", "ORCID ID to retrieve (shorthand)")
	flag.BoolVar(&sandbox, "sandbox", false, "Use ORCID sandbox instead of production")
	flag.BoolVar(&useXML, "xml", false, "Output XML instead of JSON")
	flag.BoolVar(&raw, "raw", false, "Output raw response (only works with -o flag)")
	flag.IntVar(&rows, "rows", 10, "Number of results to return (for search)")
	flag.IntVar(&start, "start", 0, "Starting position for pagination (for search)")
	flag.Parse()

	// Check for required parameters
	if bearerToken == "" {
		fmt.Fprintf(os.Stderr, "Error: Bearer token is required. Use -token or -t flag\n")
		flag.Usage()
		os.Exit(1)
	}

	// Check that either search query or orcid ID is provided, but not both
	if searchQuery == "" && orcidID == "" {
		fmt.Fprintf(os.Stderr, "Error: Either search query (-q) or ORCID ID (-o) is required\n")
		flag.Usage()
		os.Exit(1)
	}

	if searchQuery != "" && orcidID != "" {
		fmt.Fprintf(os.Stderr, "Error: Cannot use both search query (-q) and ORCID ID (-o) at the same time\n")
		flag.Usage()
		os.Exit(1)
	}

	// Configure client options
	var clientOpts []orcid.ClientOption

	// Set API URL based on sandbox flag
	if sandbox {
		clientOpts = append(clientOpts, orcid.WithAPIURL(orcid.PublicSandboxHost))
	} else {
		clientOpts = append(clientOpts, orcid.WithAPIURL(orcid.PublicHost))
	}

	// Set content type based on xml flag
	if useXML {
		clientOpts = append(clientOpts, orcid.WithContentType(orcid.ContentTypeXML))
	} else {
		clientOpts = append(clientOpts, orcid.WithContentType(orcid.ContentTypeJSON))
	}

	// Add bearer token
	clientOpts = append(clientOpts, orcid.WithBearerToken(bearerToken))

	// Create client
	client := orcid.NewClient(clientOpts...)
	ctx := context.Background()

	var output []byte

	if searchQuery != "" {
		// Perform search
		params := orcid.SearchParams{
			Query: searchQuery,
			Start: start,
			Rows:  rows,
		}

		results, err := client.Search(ctx, params)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error performing search: %v\n", err)
			os.Exit(1)
		}

		// Marshal search results
		if useXML {
			output, err = xml.MarshalIndent(results, "", "  ")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error marshaling XML: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(xml.Header + string(output))
		} else {
			output, err = json.MarshalIndent(results, "", "  ")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(string(output))
		}
	} else {
		// Get ORCID record
		if raw {
			// Get raw response
			rawData, err := client.GetRecordRaw(ctx, orcidID)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error retrieving ORCID record: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(string(rawData))
		} else {
			// Get parsed record
			record, err := client.GetRecord(ctx, orcidID)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error retrieving ORCID record: %v\n", err)
				os.Exit(1)
			}

			// Marshal record
			if useXML {
				output, err = xml.MarshalIndent(record, "", "  ")
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error marshaling XML: %v\n", err)
					os.Exit(1)
				}
				fmt.Println(xml.Header + string(output))
			} else {
				output, err = json.MarshalIndent(record, "", "  ")
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
					os.Exit(1)
				}
				fmt.Println(string(output))
			}
		}
	}
}
