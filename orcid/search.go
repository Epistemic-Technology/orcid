package orcid

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type SearchParams struct {
	Query string
	Start int
	Rows  int
}

type SearchQuery struct {
	params     SearchParams
	queryParts []string
}

func NewSearchQuery() *SearchQuery {
	return &SearchQuery{
		queryParts: []string{},
		params: SearchParams{
			Rows: 10,
		},
	}
}

func (sq *SearchQuery) ORCID(orcid string) *SearchQuery {
	sq.queryParts = append(sq.queryParts, fmt.Sprintf("orcid:%s", orcid))
	return sq
}

func (sq *SearchQuery) Email(email string) *SearchQuery {
	sq.queryParts = append(sq.queryParts, fmt.Sprintf("email:%s", email))
	return sq
}

func (sq *SearchQuery) FamilyName(name string) *SearchQuery {
	if strings.Contains(name, " ") {
		name = fmt.Sprintf("\"%s\"", name)
	}
	sq.queryParts = append(sq.queryParts, fmt.Sprintf("family-name:%s", name))
	return sq
}

func (sq *SearchQuery) GivenNames(names string) *SearchQuery {
	if strings.Contains(names, " ") {
		names = fmt.Sprintf("\"%s\"", names)
	}
	sq.queryParts = append(sq.queryParts, fmt.Sprintf("given-names:%s", names))
	return sq
}

func (sq *SearchQuery) CreditName(name string) *SearchQuery {
	if strings.Contains(name, " ") {
		name = fmt.Sprintf("\"%s\"", name)
	}
	sq.queryParts = append(sq.queryParts, fmt.Sprintf("credit-name:%s", name))
	return sq
}

func (sq *SearchQuery) OtherNames(names string) *SearchQuery {
	if strings.Contains(names, " ") {
		names = fmt.Sprintf("\"%s\"", names)
	}
	sq.queryParts = append(sq.queryParts, fmt.Sprintf("other-names:%s", names))
	return sq
}

func (sq *SearchQuery) Keyword(keyword string) *SearchQuery {
	if strings.Contains(keyword, " ") {
		keyword = fmt.Sprintf("\"%s\"", keyword)
	}
	sq.queryParts = append(sq.queryParts, fmt.Sprintf("keyword:%s", keyword))
	return sq
}

func (sq *SearchQuery) ExternalIdentifier(identifier string) *SearchQuery {
	sq.queryParts = append(sq.queryParts, fmt.Sprintf("external-identifier-type-and-value:%s", identifier))
	return sq
}

func (sq *SearchQuery) DOI(doi string) *SearchQuery {
	sq.queryParts = append(sq.queryParts, fmt.Sprintf("doi-self:%s", doi))
	return sq
}

func (sq *SearchQuery) PersonalDetails(details string) *SearchQuery {
	if strings.Contains(details, " ") {
		details = fmt.Sprintf("\"%s\"", details)
	}
	sq.queryParts = append(sq.queryParts, fmt.Sprintf("personal-details:%s", details))
	return sq
}

func (sq *SearchQuery) Biography(bio string) *SearchQuery {
	if strings.Contains(bio, " ") {
		bio = fmt.Sprintf("\"%s\"", bio)
	}
	sq.queryParts = append(sq.queryParts, fmt.Sprintf("biography:%s", bio))
	return sq
}

func (sq *SearchQuery) WorkTitle(title string) *SearchQuery {
	if strings.Contains(title, " ") {
		title = fmt.Sprintf("\"%s\"", title)
	}
	sq.queryParts = append(sq.queryParts, fmt.Sprintf("work-titles:%s", title))
	return sq
}

func (sq *SearchQuery) FundingTitle(title string) *SearchQuery {
	if strings.Contains(title, " ") {
		title = fmt.Sprintf("\"%s\"", title)
	}
	sq.queryParts = append(sq.queryParts, fmt.Sprintf("funding-titles:%s", title))
	return sq
}

func (sq *SearchQuery) AffiliationOrganization(org string) *SearchQuery {
	if strings.Contains(org, " ") {
		org = fmt.Sprintf("\"%s\"", org)
	}
	sq.queryParts = append(sq.queryParts, fmt.Sprintf("affiliation-org-name:%s", org))
	return sq
}

func (sq *SearchQuery) RINGGOLD(id string) *SearchQuery {
	sq.queryParts = append(sq.queryParts, fmt.Sprintf("ringgold-org-id:%s", id))
	return sq
}

func (sq *SearchQuery) GRID(id string) *SearchQuery {
	sq.queryParts = append(sq.queryParts, fmt.Sprintf("grid-org-id:%s", id))
	return sq
}

func (sq *SearchQuery) ROR(id string) *SearchQuery {
	sq.queryParts = append(sq.queryParts, fmt.Sprintf("ror-org-id:%s", id))
	return sq
}

func (sq *SearchQuery) FundRef(id string) *SearchQuery {
	sq.queryParts = append(sq.queryParts, fmt.Sprintf("fundref-org-id:%s", id))
	return sq
}

func (sq *SearchQuery) RawQuery(query string) *SearchQuery {
	sq.queryParts = append(sq.queryParts, query)
	return sq
}

func (sq *SearchQuery) WithStart(start int) *SearchQuery {
	sq.params.Start = start
	return sq
}

func (sq *SearchQuery) WithRows(rows int) *SearchQuery {
	sq.params.Rows = rows
	return sq
}

func (sq *SearchQuery) And() *SearchQuery {
	sq.queryParts = append(sq.queryParts, "AND")
	return sq
}

func (sq *SearchQuery) Or() *SearchQuery {
	sq.queryParts = append(sq.queryParts, "OR")
	return sq
}

func (sq *SearchQuery) Not() *SearchQuery {
	sq.queryParts = append(sq.queryParts, "NOT")
	return sq
}

func (sq *SearchQuery) Build() SearchParams {
	sq.params.Query = strings.Join(sq.queryParts, " ")
	return sq.params
}

type SearchResult struct {
	NumFound int             `json:"num-found" xml:"num-found,attr"`
	Start    int             `json:"start" xml:"start,attr"`
	NumRows  int             `json:"num-rows" xml:"num-rows,attr"`
	Results  []*SearchRecord `json:"result,omitempty" xml:"result,omitempty"`
}

type SearchRecord struct {
	OrcidIdentifier *OrcidIdentifier `json:"orcid-identifier,omitempty" xml:"orcid-identifier,omitempty"`
}

func (c *Client) Search(ctx context.Context, params SearchParams) (*SearchResult, error) {
	searchURL := c.buildSearchURL(params)

	resp, err := c.doRequest(ctx, http.MethodGet, searchURL, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result SearchResult
	if err := c.unmarshalResponse(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) SearchWithQuery(ctx context.Context, query *SearchQuery) (*SearchResult, error) {
	params := query.Build()
	return c.Search(ctx, params)
}

type SearchIterator struct {
	client       *Client
	params       SearchParams
	currentBatch *SearchResult
	currentIndex int
	totalResults int
	ctx          context.Context
	err          error
}

func (c *Client) SearchIter(ctx context.Context, params SearchParams) *SearchIterator {
	return &SearchIterator{
		client:       c,
		params:       params,
		ctx:          ctx,
		currentIndex: -1,
	}
}

func (c *Client) SearchIterWithQuery(ctx context.Context, query *SearchQuery) *SearchIterator {
	params := query.Build()
	return c.SearchIter(ctx, params)
}

func (si *SearchIterator) Next() bool {
	if si.err != nil {
		return false
	}

	select {
	case <-si.ctx.Done():
		si.err = si.ctx.Err()
		return false
	default:
	}

	// Check if we need to fetch the next batch
	if si.currentBatch == nil || si.currentIndex >= len(si.currentBatch.Results)-1 {
		// Check if we've fetched all available results
		if si.currentBatch != nil {
			// Check if we've reached the total results or no more results in current batch
			if si.params.Start+si.params.Rows >= si.totalResults {
				return false
			}
			// Move start position by the requested rows (not actual fetched)
			si.params.Start += si.params.Rows
		}

		result, err := si.client.Search(si.ctx, si.params)
		if err != nil {
			si.err = err
			return false
		}

		si.currentBatch = result
		si.totalResults = result.NumFound
		si.currentIndex = -1

		if len(result.Results) == 0 {
			return false
		}
	}

	si.currentIndex++
	return si.currentIndex < len(si.currentBatch.Results)
}

func (si *SearchIterator) Value() *SearchRecord {
	if si.currentBatch == nil || si.currentIndex < 0 || si.currentIndex >= len(si.currentBatch.Results) {
		return nil
	}
	return si.currentBatch.Results[si.currentIndex]
}

func (si *SearchIterator) Error() error {
	return si.err
}

func (si *SearchIterator) TotalResults() int {
	return si.totalResults
}

func (c *Client) ExpandedSearch(ctx context.Context, query string) (*ExpandedSearchResult, error) {
	searchURL := fmt.Sprintf("%s/expanded-search/?q=%s", c.apiURL, url.QueryEscape(query))

	resp, err := c.doRequest(ctx, http.MethodGet, searchURL, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result ExpandedSearchResult
	if err := c.unmarshalResponse(data, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

type ExpandedSearchResult struct {
	NumFound        int                     `json:"num-found" xml:"num-found,attr"`
	ExpandedResults []*ExpandedSearchRecord `json:"expanded-result,omitempty" xml:"expanded-result,omitempty"`
}

type ExpandedSearchRecord struct {
	OrcidID         string   `json:"orcid-id,omitempty" xml:"orcid-id,omitempty"`
	GivenNames      string   `json:"given-names,omitempty" xml:"given-names,omitempty"`
	FamilyNames     string   `json:"family-names,omitempty" xml:"family-names,omitempty"`
	CreditName      string   `json:"credit-name,omitempty" xml:"credit-name,omitempty"`
	Email           []string `json:"email,omitempty" xml:"email,omitempty"`
	InstitutionName []string `json:"institution-name,omitempty" xml:"institution-name,omitempty"`
}

func ParseOrcidID(input string) string {
	input = strings.TrimSpace(input)

	if strings.HasPrefix(input, "http://") || strings.HasPrefix(input, "https://") {
		parts := strings.Split(input, "/")
		if len(parts) > 0 {
			return parts[len(parts)-1]
		}
	}

	parts := strings.Split(input, "/")
	return parts[len(parts)-1]
}

func FormatOrcidID(orcid string) string {
	orcid = ParseOrcidID(orcid)
	orcid = strings.ToUpper(strings.ReplaceAll(orcid, "-", ""))

	if len(orcid) == 16 {
		return fmt.Sprintf("%s-%s-%s-%s",
			orcid[0:4], orcid[4:8], orcid[8:12], orcid[12:16])
	}

	return orcid
}

func ValidateOrcidID(orcid string) error {
	orcid = ParseOrcidID(orcid)
	orcid = strings.ReplaceAll(orcid, "-", "")

	if len(orcid) != 16 {
		return fmt.Errorf("invalid ORCID iD length: %d", len(orcid))
	}

	for i := 0; i < 15; i++ {
		if orcid[i] < '0' || orcid[i] > '9' {
			return fmt.Errorf("invalid character in ORCID iD at position %d", i)
		}
	}

	lastChar := orcid[15]
	if (lastChar < '0' || lastChar > '9') && lastChar != 'X' {
		return fmt.Errorf("invalid check digit in ORCID iD")
	}

	if !isValidChecksum(orcid) {
		return fmt.Errorf("invalid ORCID iD checksum")
	}

	return nil
}

func isValidChecksum(orcid string) bool {
	total := 0
	for i := 0; i < 15; i++ {
		digit, _ := strconv.Atoi(string(orcid[i]))
		total = (total + digit) * 2
	}

	remainder := total % 11
	result := (12 - remainder) % 11

	checkDigit := string(orcid[15])
	if result == 10 {
		return checkDigit == "X"
	}

	return checkDigit == strconv.Itoa(result)
}
