package orcid

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	client := NewClient()
	if client == nil {
		t.Fatal("Expected non-nil client")
	}
	if client.apiURL != DefaultAPIURL {
		t.Errorf("Expected apiURL %s, got %s", DefaultAPIURL, client.apiURL)
	}
	if client.timeout != DefaultTimeout {
		t.Errorf("Expected timeout %v, got %v", DefaultTimeout, client.timeout)
	}
	if client.maxRetries != DefaultMaxRetries {
		t.Errorf("Expected maxRetries %d, got %d", DefaultMaxRetries, client.maxRetries)
	}
	if client.contentType != ContentTypeJSON {
		t.Errorf("Expected contentType %s, got %s", ContentTypeJSON, client.contentType)
	}
}

func TestClientOptions(t *testing.T) {
	customHTTPClient := &http.Client{Timeout: 5 * time.Second}
	customTimeout := 10 * time.Second

	client := NewClient(
		WithHTTPClient(customHTTPClient),
		WithAPIURL("https://custom.orcid.org/v3.0"),
		WithTimeout(customTimeout),
		WithMaxRetries(5),
		WithRateLimit(20),
		WithUserAgent("TestAgent/1.0"),
		WithContentType(ContentTypeXML),
		WithBearerToken("test-token-123"),
	)

	if client.apiURL != "https://custom.orcid.org/v3.0" {
		t.Errorf("Expected apiURL %s, got %s", "https://custom.orcid.org/v3.0", client.apiURL)
	}
	if client.timeout != customTimeout {
		t.Errorf("Expected timeout %v, got %v", customTimeout, client.timeout)
	}
	if client.maxRetries != 5 {
		t.Errorf("Expected maxRetries %d, got %d", 5, client.maxRetries)
	}
	if client.rateLimit != 20 {
		t.Errorf("Expected rateLimit %d, got %d", 20, client.rateLimit)
	}
	if client.userAgent != "TestAgent/1.0" {
		t.Errorf("Expected userAgent %s, got %s", "TestAgent/1.0", client.userAgent)
	}
	if client.contentType != ContentTypeXML {
		t.Errorf("Expected contentType %s, got %s", ContentTypeXML, client.contentType)
	}
	if client.rateLimiter == nil {
		t.Fatal("Expected non-nil rateLimiter")
	}
	if client.bearerToken != "test-token-123" {
		t.Errorf("Expected bearerToken %s, got %s", "test-token-123", client.bearerToken)
	}
}

func TestMissingBearerToken(t *testing.T) {
	client := NewClient()
	ctx := context.Background()

	_, err := client.GetRecord(ctx, "0000-0002-1825-0097")
	if err == nil {
		t.Fatal("Expected error for missing bearer token")
	}
	if !strings.Contains(err.Error(), "bearer token is required") {
		t.Errorf("Expected error about missing bearer token, got: %v", err)
	}
}

func TestBearerTokenAuthentication(t *testing.T) {
	expectedToken := "test-bearer-token-abc123"
	authHeaderReceived := ""

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeaderReceived = r.Header.Get("Authorization")
		if authHeaderReceived != "Bearer "+expectedToken {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"orcid-identifier": {"path": "0000-0002-1825-0097"}}`))
	}))
	defer server.Close()

	// Test with bearer token
	client := NewClient(
		WithAPIURL(server.URL+"/v3.0"),
		WithBearerToken(expectedToken),
	)
	ctx := context.Background()

	record, err := client.GetRecord(ctx, "0000-0002-1825-0097")
	if err != nil {
		t.Fatalf("Unexpected error with bearer token: %v", err)
	}
	if record == nil {
		t.Fatal("Expected non-nil record")
	}
	if authHeaderReceived != "Bearer "+expectedToken {
		t.Errorf("Expected Authorization header 'Bearer %s', got '%s'", expectedToken, authHeaderReceived)
	}

	// Test without bearer token (should fail with our validation)
	clientNoToken := NewClient(WithAPIURL(server.URL + "/v3.0"))
	_, err = clientNoToken.GetRecord(ctx, "0000-0002-1825-0097")
	if err == nil {
		t.Fatal("Expected error without bearer token")
	}
	if !strings.Contains(err.Error(), "bearer token is required") {
		t.Errorf("Expected bearer token required error, got: %v", err)
	}
}

func TestGetRecord(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3.0/0000-0002-1825-0097/record" {
			t.Errorf("Expected path %s, got %s", "/v3.0/0000-0002-1825-0097/record", r.URL.Path)
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept header %s, got %s", "application/json", r.Header.Get("Accept"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"orcid-identifier": {
				"uri": "https://orcid.org/0000-0002-1825-0097",
				"path": "0000-0002-1825-0097",
				"host": "orcid.org"
			},
			"person": {
				"name": {
					"given-names": {"value": "John"},
					"family-name": {"value": "Doe"}
				}
			}
		}`))
	}))
	defer server.Close()

	client := NewClient(
		WithAPIURL(server.URL+"/v3.0"),
		WithBearerToken("test-token"),
	)
	ctx := context.Background()

	record, err := client.GetRecord(ctx, "0000-0002-1825-0097")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if record == nil {
		t.Fatal("Expected non-nil record")
	}

	if record.OrcidIdentifier.Path != "0000-0002-1825-0097" {
		t.Errorf("Expected path %s, got %s", "0000-0002-1825-0097", record.OrcidIdentifier.Path)
	}
	if record.Person.Name.GivenNames.Value != "John" {
		t.Errorf("Expected given name %s, got %s", "John", record.Person.Name.GivenNames.Value)
	}
	if record.Person.Name.FamilyName.Value != "Doe" {
		t.Errorf("Expected family name %s, got %s", "Doe", record.Person.Name.FamilyName.Value)
	}
}

func TestGetWorks(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3.0/0000-0002-1825-0097/works" {
			t.Errorf("Expected path %s, got %s", "/v3.0/0000-0002-1825-0097/works", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"last-modified-date": {"value": 1609459200000},
			"group": [{
				"work-summary": [{
					"put-code": 12345,
					"title": {
						"title": {"value": "Test Publication"}
					},
					"type": "journal-article",
					"visibility": "public"
				}]
			}],
			"path": "/0000-0002-1825-0097/works"
		}`))
	}))
	defer server.Close()

	client := NewClient(
		WithAPIURL(server.URL+"/v3.0"),
		WithBearerToken("test-token"),
	)
	ctx := context.Background()

	works, err := client.GetWorks(ctx, "0000-0002-1825-0097")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if works == nil {
		t.Fatal("Expected non-nil works")
	}

	if len(works.WorkGroup) != 1 {
		t.Errorf("Expected 1 work group, got %d", len(works.WorkGroup))
	}
	if len(works.WorkGroup[0].WorkSummary) != 1 {
		t.Errorf("Expected 1 work summary, got %d", len(works.WorkGroup[0].WorkSummary))
	}
	if works.WorkGroup[0].WorkSummary[0].Title.Title.Value != "Test Publication" {
		t.Errorf("Expected title %s, got %s", "Test Publication", works.WorkGroup[0].WorkSummary[0].Title.Title.Value)
	}
	if works.WorkGroup[0].WorkSummary[0].Type != "journal-article" {
		t.Errorf("Expected type %s, got %s", "journal-article", works.WorkGroup[0].WorkSummary[0].Type)
	}
}

func TestSearch(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v3.0/search" {
			t.Errorf("Expected path %s, got %s", "/v3.0/search", r.URL.Path)
		}
		if r.URL.Query().Get("q") != "family-name:Smith" {
			t.Errorf("Expected query %s, got %s", "family-name:Smith", r.URL.Query().Get("q"))
		}
		if r.URL.Query().Get("rows") != "10" {
			t.Errorf("Expected rows %s, got %s", "10", r.URL.Query().Get("rows"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"num-found": 100,
			"start": 0,
			"num-rows": 10,
			"result": [
				{
					"orcid-identifier": {
						"uri": "https://orcid.org/0000-0001-0000-0000",
						"path": "0000-0001-0000-0000",
						"host": "orcid.org"
					}
				}
			]
		}`))
	}))
	defer server.Close()

	client := NewClient(
		WithAPIURL(server.URL+"/v3.0"),
		WithBearerToken("test-token"),
	)
	ctx := context.Background()

	params := SearchParams{
		Query: "family-name:Smith",
		Start: 0,
		Rows:  10,
	}

	result, err := client.Search(ctx, params)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("Expected non-nil result")
	}

	if result.NumFound != 100 {
		t.Errorf("Expected NumFound %d, got %d", 100, result.NumFound)
	}
	if result.Start != 0 {
		t.Errorf("Expected Start %d, got %d", 0, result.Start)
	}
	if result.NumRows != 10 {
		t.Errorf("Expected NumRows %d, got %d", 10, result.NumRows)
	}
	if len(result.Results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(result.Results))
	}
	if result.Results[0].OrcidIdentifier.Path != "0000-0001-0000-0000" {
		t.Errorf("Expected path %s, got %s", "0000-0001-0000-0000", result.Results[0].OrcidIdentifier.Path)
	}
}

func TestSearchWithQuery(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedQuery := "family-name:Einstein AND given-names:Albert"
		if r.URL.Query().Get("q") != expectedQuery {
			t.Errorf("Expected query %s, got %s", expectedQuery, r.URL.Query().Get("q"))
		}
		if r.URL.Query().Get("rows") != "5" {
			t.Errorf("Expected rows %s, got %s", "5", r.URL.Query().Get("rows"))
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"num-found": 1, "start": 0, "num-rows": 5, "result": []}`))
	}))
	defer server.Close()

	client := NewClient(
		WithAPIURL(server.URL+"/v3.0"),
		WithBearerToken("test-token"),
	)
	ctx := context.Background()

	query := NewSearchQuery().
		FamilyName("Einstein").
		And().
		GivenNames("Albert").
		WithRows(5)

	result, err := client.SearchWithQuery(ctx, query)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result == nil {
		t.Fatal("Expected non-nil result")
	}

	if result.NumFound != 1 {
		t.Errorf("Expected NumFound %d, got %d", 1, result.NumFound)
	}
}

func TestRetryLogic(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts < 3 {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"orcid-identifier": {"path": "0000-0002-1825-0097"}}`))
	}))
	defer server.Close()

	client := NewClient(
		WithAPIURL(server.URL+"/v3.0"),
		WithBearerToken("test-token"),
		WithMaxRetries(3),
	)
	ctx := context.Background()

	record, err := client.GetRecord(ctx, "0000-0002-1825-0097")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if record == nil {
		t.Fatal("Expected non-nil record")
	}

	if attempts != 3 {
		t.Errorf("Expected 3 attempts, got %d", attempts)
	}
	if record.OrcidIdentifier.Path != "0000-0002-1825-0097" {
		t.Errorf("Expected path %s, got %s", "0000-0002-1825-0097", record.OrcidIdentifier.Path)
	}
}

func TestContextCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(
		WithAPIURL(server.URL+"/v3.0"),
		WithBearerToken("test-token"),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	_, err := client.GetRecord(ctx, "0000-0002-1825-0097")
	if err == nil {
		t.Error("Expected error for cancelled context")
	}
}

func TestValidateOrcidID(t *testing.T) {
	tests := []struct {
		name    string
		orcidID string
		wantErr bool
	}{
		{
			name:    "Valid ORCID",
			orcidID: "0000-0002-1825-0097",
			wantErr: false,
		},
		{
			name:    "Valid ORCID with URL",
			orcidID: "https://orcid.org/0000-0002-1825-0097",
			wantErr: false,
		},
		{
			name:    "Invalid length",
			orcidID: "0000-0002-1825",
			wantErr: true,
		},
		{
			name:    "Invalid characters",
			orcidID: "XXXX-0002-1825-0097",
			wantErr: true,
		},
		{
			name:    "Invalid checksum",
			orcidID: "0000-0002-1825-0099",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateOrcidID(tt.orcidID)
			if tt.wantErr {
				if err == nil {
					t.Error("Expected error")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestFormatOrcidID(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Already formatted",
			input:    "0000-0002-1825-0097",
			expected: "0000-0002-1825-0097",
		},
		{
			name:     "Without hyphens",
			input:    "0000000218250097",
			expected: "0000-0002-1825-0097",
		},
		{
			name:     "With URL",
			input:    "https://orcid.org/0000-0002-1825-0097",
			expected: "0000-0002-1825-0097",
		},
		{
			name:     "Lowercase",
			input:    "0000-0002-1825-009x",
			expected: "0000-0002-1825-009X",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatOrcidID(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestSearchIterator(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++

		start := r.URL.Query().Get("start")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if start == "" || start == "0" {
			w.Write([]byte(`{
				"num-found": 25,
				"start": 0,
				"num-rows": 10,
				"result": [
					{"orcid-identifier": {"path": "0000-0000-0000-0001"}},
					{"orcid-identifier": {"path": "0000-0000-0000-0002"}}
				]
			}`))
		} else if start == "10" {
			w.Write([]byte(`{
				"num-found": 25,
				"start": 10,
				"num-rows": 10,
				"result": [
					{"orcid-identifier": {"path": "0000-0000-0000-0003"}},
					{"orcid-identifier": {"path": "0000-0000-0000-0004"}}
				]
			}`))
		} else {
			w.Write([]byte(`{
				"num-found": 25,
				"start": 20,
				"num-rows": 10,
				"result": [
					{"orcid-identifier": {"path": "0000-0000-0000-0005"}}
				]
			}`))
		}
	}))
	defer server.Close()

	client := NewClient(
		WithAPIURL(server.URL+"/v3.0"),
		WithBearerToken("test-token"),
	)
	ctx := context.Background()

	params := SearchParams{
		Query: "test",
		Rows:  10,
	}

	iter := client.SearchIter(ctx, params)

	var results []string
	for iter.Next() {
		record := iter.Value()
		if record != nil && record.OrcidIdentifier != nil {
			results = append(results, string(record.OrcidIdentifier.Path))
		}
	}

	if iter.Error() != nil {
		t.Fatalf("Unexpected error: %v", iter.Error())
	}
	if iter.TotalResults() != 25 {
		t.Errorf("Expected 25 total results, got %d", iter.TotalResults())
	}
	if len(results) != 5 {
		t.Errorf("Expected 5 results, got %d", len(results))
	}
	if callCount != 3 {
		t.Errorf("Expected 3 API calls, got %d", callCount)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		responseBody   string
		expectedErrMsg string
	}{
		{
			name:           "404 Not Found",
			statusCode:     http.StatusNotFound,
			responseBody:   `{"error":"Not found"}`,
			expectedErrMsg: "404",
		},
		{
			name:           "500 Internal Server Error",
			statusCode:     http.StatusInternalServerError,
			responseBody:   `{"error":"Internal server error"}`,
			expectedErrMsg: "500",
		},
		{
			name:           "401 Unauthorized",
			statusCode:     http.StatusUnauthorized,
			responseBody:   `{"error":"Unauthorized"}`,
			expectedErrMsg: "401",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				w.Write([]byte(tt.responseBody))
			}))
			defer server.Close()

			client := NewClient(
				WithAPIURL(server.URL+"/v3.0"),
				WithBearerToken("test-token"),
			)
			ctx := context.Background()

			_, err := client.GetRecord(ctx, "0000-0002-1825-0097")
			if err == nil {
				t.Fatal("Expected error")
			}
			if !strings.Contains(err.Error(), tt.expectedErrMsg) {
				t.Errorf("Expected error message to contain %s, got %v", tt.expectedErrMsg, err)
			}
		})
	}
}

func TestGetRecordRaw(t *testing.T) {
	expectedJSON := `{
		"orcid-identifier": {
			"uri": "https://orcid.org/0000-0002-1825-0097",
			"path": "0000-0002-1825-0097",
			"host": "orcid.org"
		},
		"person": {
			"name": {
				"given-names": {"value": "John"},
				"family-name": {"value": "Doe"}
			}
		}
	}`

	expectedXML := `<?xml version="1.0" encoding="UTF-8"?>
<record:record xmlns:record="http://www.orcid.org/ns/record">
	<record:orcid-identifier>
		<common:uri>https://orcid.org/0000-0002-1825-0097</common:uri>
		<common:path>0000-0002-1825-0097</common:path>
		<common:host>orcid.org</common:host>
	</record:orcid-identifier>
</record:record>`

	t.Run("JSON response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Accept") != "application/json" {
				t.Errorf("Expected Accept header %s, got %s", "application/json", r.Header.Get("Accept"))
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(expectedJSON))
		}))
		defer server.Close()

		client := NewClient(
			WithAPIURL(server.URL+"/v3.0"),
			WithBearerToken("test-token"),
			WithContentType(ContentTypeJSON),
		)
		ctx := context.Background()

		rawData, err := client.GetRecordRaw(ctx, "0000-0002-1825-0097")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if string(rawData) != expectedJSON {
			t.Errorf("Expected raw JSON:\n%s\nGot:\n%s", expectedJSON, string(rawData))
		}
	})

	t.Run("XML response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Accept") != "application/vnd.orcid+xml" {
				t.Errorf("Expected Accept header %s, got %s", "application/vnd.orcid+xml", r.Header.Get("Accept"))
			}
			w.Header().Set("Content-Type", "application/vnd.orcid+xml")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(expectedXML))
		}))
		defer server.Close()

		client := NewClient(
			WithAPIURL(server.URL+"/v3.0"),
			WithBearerToken("test-token"),
			WithContentType(ContentTypeXML),
		)
		ctx := context.Background()

		rawData, err := client.GetRecordRaw(ctx, "0000-0002-1825-0097")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if string(rawData) != expectedXML {
			t.Errorf("Expected raw XML:\n%s\nGot:\n%s", expectedXML, string(rawData))
		}
	})

	t.Run("Error handling", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error":"Record not found"}`))
		}))
		defer server.Close()

		client := NewClient(
			WithAPIURL(server.URL+"/v3.0"),
			WithBearerToken("test-token"),
		)
		ctx := context.Background()

		_, err := client.GetRecordRaw(ctx, "invalid-orcid")
		if err == nil {
			t.Fatal("Expected error for 404 response")
		}
		if !strings.Contains(err.Error(), "404") {
			t.Errorf("Expected error message to contain 404, got: %v", err)
		}
	})
}

func TestGetByPath(t *testing.T) {
	tests := []struct {
		name          string
		path          Path
		expectedPath  string
		responseBody  string
		expectError   bool
		errorContains string
	}{
		{
			name:         "Get Record by ORCID only",
			path:         Path("/0000-0002-1825-0097"),
			expectedPath: "/v3.0/0000-0002-1825-0097/record",
			responseBody: `{"orcid-identifier": {"path": "0000-0002-1825-0097"}}`,
			expectError:  false,
		},
		{
			name:         "Get Record with explicit record path",
			path:         Path("/0000-0002-1825-0097/record"),
			expectedPath: "/v3.0/0000-0002-1825-0097/record",
			responseBody: `{"orcid-identifier": {"path": "0000-0002-1825-0097"}}`,
			expectError:  false,
		},
		{
			name:         "Get Person",
			path:         Path("/0000-0002-1825-0097/person"),
			expectedPath: "/v3.0/0000-0002-1825-0097/person",
			responseBody: `{"name": {"given-names": {"value": "John"}}}`,
			expectError:  false,
		},
		{
			name:         "Get Works",
			path:         Path("/0000-0002-1825-0097/works"),
			expectedPath: "/v3.0/0000-0002-1825-0097/works",
			responseBody: `{"path": "/0000-0002-1825-0097/works"}`,
			expectError:  false,
		},
		{
			name:         "Get Qualifications",
			path:         Path("/0000-0003-1401-2056/qualifications"),
			expectedPath: "/v3.0/0000-0003-1401-2056/qualifications",
			responseBody: `{"path": "/0000-0003-1401-2056/qualifications"}`,
			expectError:  false,
		},
		{
			name:         "Get specific work with put-code",
			path:         Path("/0000-0002-1825-0097/work/123456"),
			expectedPath: "/v3.0/0000-0002-1825-0097/work/123456",
			responseBody: `{"put-code": 123456, "title": {"title": {"value": "Test Work"}}}`,
			expectError:  false,
		},
		{
			name:         "Get Educations",
			path:         Path("/0000-0002-1825-0097/educations"),
			expectedPath: "/v3.0/0000-0002-1825-0097/educations",
			responseBody: `{"path": "/0000-0002-1825-0097/educations"}`,
			expectError:  false,
		},
		{
			name:         "Get Employments",
			path:         Path("/0000-0002-1825-0097/employments"),
			expectedPath: "/v3.0/0000-0002-1825-0097/employments",
			responseBody: `{"path": "/0000-0002-1825-0097/employments"}`,
			expectError:  false,
		},
		{
			name:         "Get Fundings",
			path:         Path("/0000-0002-1825-0097/fundings"),
			expectedPath: "/v3.0/0000-0002-1825-0097/fundings",
			responseBody: `{"path": "/0000-0002-1825-0097/fundings"}`,
			expectError:  false,
		},
		{
			name:         "Get Peer Reviews",
			path:         Path("/0000-0002-1825-0097/peer-reviews"),
			expectedPath: "/v3.0/0000-0002-1825-0097/peer-reviews",
			responseBody: `{"path": "/0000-0002-1825-0097/peer-reviews"}`,
			expectError:  false,
		},
		{
			name:         "Get Distinctions",
			path:         Path("/0000-0002-1825-0097/distinctions"),
			expectedPath: "/v3.0/0000-0002-1825-0097/distinctions",
			responseBody: `{"path": "/0000-0002-1825-0097/distinctions"}`,
			expectError:  false,
		},
		{
			name:         "Get Invited Positions",
			path:         Path("/0000-0002-1825-0097/invited-positions"),
			expectedPath: "/v3.0/0000-0002-1825-0097/invited-positions",
			responseBody: `{"path": "/0000-0002-1825-0097/invited-positions"}`,
			expectError:  false,
		},
		{
			name:         "Get Memberships",
			path:         Path("/0000-0002-1825-0097/memberships"),
			expectedPath: "/v3.0/0000-0002-1825-0097/memberships",
			responseBody: `{"path": "/0000-0002-1825-0097/memberships"}`,
			expectError:  false,
		},
		{
			name:         "Get Services",
			path:         Path("/0000-0002-1825-0097/services"),
			expectedPath: "/v3.0/0000-0002-1825-0097/services",
			responseBody: `{"path": "/0000-0002-1825-0097/services"}`,
			expectError:  false,
		},
		{
			name:         "Get Research Resources",
			path:         Path("/0000-0002-1825-0097/research-resources"),
			expectedPath: "/v3.0/0000-0002-1825-0097/research-resources",
			responseBody: `{"path": "/0000-0002-1825-0097/research-resources"}`,
			expectError:  false,
		},
		{
			name:          "Invalid empty path",
			path:          Path(""),
			expectError:   true,
			errorContains: "invalid path",
		},
		{
			name:          "Invalid path with just slash",
			path:          Path("/"),
			expectError:   true,
			errorContains: "invalid path",
		},
		{
			name:          "Work path without put-code",
			path:          Path("/0000-0002-1825-0097/work"),
			expectError:   true,
			errorContains: "work path requires put-code",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actualPath string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				actualPath = r.URL.Path

				// Check authorization header
				if r.Header.Get("Authorization") != "Bearer test-token" {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(tt.responseBody))
			}))
			defer server.Close()

			client := NewClient(
				WithAPIURL(server.URL+"/v3.0"),
				WithBearerToken("test-token"),
			)
			ctx := context.Background()

			result, err := client.GetByPath(ctx, tt.path)

			if tt.expectError {
				if err == nil {
					t.Fatal("Expected error but got none")
				}
				if tt.errorContains != "" && !strings.Contains(err.Error(), tt.errorContains) {
					t.Errorf("Expected error containing '%s', got: %v", tt.errorContains, err)
				}
			} else {
				if err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
				if result == nil {
					t.Fatal("Expected non-nil result")
				}
				if actualPath != tt.expectedPath {
					t.Errorf("Expected path %s, got %s", tt.expectedPath, actualPath)
				}
			}
		})
	}
}

func TestGetByPathBiographyFromPerson(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// GetByPath for biography should fetch the person record
		if r.URL.Path != "/v3.0/0000-0002-1825-0097/person" {
			t.Errorf("Expected path %s, got %s", "/v3.0/0000-0002-1825-0097/person", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"biography": {
				"content": "Test biography",
				"visibility": "public",
				"path": "/0000-0002-1825-0097/biography"
			}
		}`))
	}))
	defer server.Close()

	client := NewClient(
		WithAPIURL(server.URL+"/v3.0"),
		WithBearerToken("test-token"),
	)
	ctx := context.Background()

	result, err := client.GetByPath(ctx, Path("/0000-0002-1825-0097/biography"))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	biography, ok := result.(*Biography)
	if !ok {
		t.Fatalf("Expected *Biography, got %T", result)
	}
	if biography.Content != "Test biography" {
		t.Errorf("Expected biography content 'Test biography', got '%s'", biography.Content)
	}
}

func TestGetByPathActivitiesFromRecord(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// GetByPath for activities should fetch the full record
		if r.URL.Path != "/v3.0/0000-0002-1825-0097/record" {
			t.Errorf("Expected path %s, got %s", "/v3.0/0000-0002-1825-0097/record", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"orcid-identifier": {"path": "0000-0002-1825-0097"},
			"activities-summary": {
				"works": {
					"path": "/0000-0002-1825-0097/works"
				}
			}
		}`))
	}))
	defer server.Close()

	client := NewClient(
		WithAPIURL(server.URL+"/v3.0"),
		WithBearerToken("test-token"),
	)
	ctx := context.Background()

	result, err := client.GetByPath(ctx, Path("/0000-0002-1825-0097/activities"))
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	activities, ok := result.(*ActivitiesSummary)
	if !ok {
		t.Fatalf("Expected *ActivitiesSummary, got %T", result)
	}
	if activities.Works == nil {
		t.Error("Expected non-nil Works in activities summary")
	}
}
