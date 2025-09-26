package orcid

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *Client) GetRecord(ctx context.Context, orcidID string) (*Record, error) {
	url := fmt.Sprintf("%s/%s/record", c.apiURL, orcidID)

	resp, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var record Record
	if err := c.unmarshalResponse(data, &record); err != nil {
		return nil, err
	}

	return &record, nil
}

func (c *Client) GetRecordRaw(ctx context.Context, orcidID string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s/record", c.apiURL, orcidID)

	resp, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (c *Client) GetPerson(ctx context.Context, orcidID string) (*Person, error) {
	url := fmt.Sprintf("%s/%s/person", c.apiURL, orcidID)

	resp, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var person Person
	if err := c.unmarshalResponse(data, &person); err != nil {
		return nil, err
	}

	return &person, nil
}

func (c *Client) GetWorks(ctx context.Context, orcidID string) (*Works, error) {
	url := fmt.Sprintf("%s/%s/works", c.apiURL, orcidID)

	resp, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var works Works
	if err := c.unmarshalResponse(data, &works); err != nil {
		return nil, err
	}

	return &works, nil
}

func (c *Client) GetWork(ctx context.Context, orcidID string, putCode string) (*Work, error) {
	url := fmt.Sprintf("%s/%s/work/%s", c.apiURL, orcidID, putCode)

	resp, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var work Work
	if err := c.unmarshalResponse(data, &work); err != nil {
		return nil, err
	}

	return &work, nil
}

func (c *Client) GetEducations(ctx context.Context, orcidID string) (*Educations, error) {
	url := fmt.Sprintf("%s/%s/educations", c.apiURL, orcidID)

	resp, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var educations Educations
	if err := c.unmarshalResponse(data, &educations); err != nil {
		return nil, err
	}

	return &educations, nil
}

func (c *Client) GetEmployments(ctx context.Context, orcidID string) (*Employments, error) {
	url := fmt.Sprintf("%s/%s/employments", c.apiURL, orcidID)

	resp, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var employments Employments
	if err := c.unmarshalResponse(data, &employments); err != nil {
		return nil, err
	}

	return &employments, nil
}

func (c *Client) GetFundings(ctx context.Context, orcidID string) (*Fundings, error) {
	url := fmt.Sprintf("%s/%s/fundings", c.apiURL, orcidID)

	resp, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var fundings Fundings
	if err := c.unmarshalResponse(data, &fundings); err != nil {
		return nil, err
	}

	return &fundings, nil
}

func (c *Client) GetPeerReviews(ctx context.Context, orcidID string) (*PeerReviews, error) {
	url := fmt.Sprintf("%s/%s/peer-reviews", c.apiURL, orcidID)

	resp, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var peerReviews PeerReviews
	if err := c.unmarshalResponse(data, &peerReviews); err != nil {
		return nil, err
	}

	return &peerReviews, nil
}

func (c *Client) GetDistinctions(ctx context.Context, orcidID string) (*Distinctions, error) {
	url := fmt.Sprintf("%s/%s/distinctions", c.apiURL, orcidID)

	resp, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var distinctions Distinctions
	if err := c.unmarshalResponse(data, &distinctions); err != nil {
		return nil, err
	}

	return &distinctions, nil
}

func (c *Client) GetInvitedPositions(ctx context.Context, orcidID string) (*InvitedPositions, error) {
	url := fmt.Sprintf("%s/%s/invited-positions", c.apiURL, orcidID)

	resp, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var invitedPositions InvitedPositions
	if err := c.unmarshalResponse(data, &invitedPositions); err != nil {
		return nil, err
	}

	return &invitedPositions, nil
}

func (c *Client) GetMemberships(ctx context.Context, orcidID string) (*Memberships, error) {
	url := fmt.Sprintf("%s/%s/memberships", c.apiURL, orcidID)

	resp, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var memberships Memberships
	if err := c.unmarshalResponse(data, &memberships); err != nil {
		return nil, err
	}

	return &memberships, nil
}

func (c *Client) GetQualifications(ctx context.Context, orcidID string) (*Qualifications, error) {
	url := fmt.Sprintf("%s/%s/qualifications", c.apiURL, orcidID)

	resp, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var qualifications Qualifications
	if err := c.unmarshalResponse(data, &qualifications); err != nil {
		return nil, err
	}

	return &qualifications, nil
}

func (c *Client) GetServices(ctx context.Context, orcidID string) (*Services, error) {
	url := fmt.Sprintf("%s/%s/services", c.apiURL, orcidID)

	resp, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var services Services
	if err := c.unmarshalResponse(data, &services); err != nil {
		return nil, err
	}

	return &services, nil
}

func (c *Client) GetResearchResources(ctx context.Context, orcidID string) (*ResearchResources, error) {
	url := fmt.Sprintf("%s/%s/research-resources", c.apiURL, orcidID)

	resp, err := c.doRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var researchResources ResearchResources
	if err := c.unmarshalResponse(data, &researchResources); err != nil {
		return nil, err
	}

	return &researchResources, nil
}

// GetByPath fetches a resource by its path.
// Path values are returned by various ORCID API endpoints and can be used
// to directly fetch specific resources.
//
// Examples:
//   - "/0000-0003-1401-2056/qualifications" -> calls GetQualifications
//   - "/0000-0003-1401-2056/works" -> calls GetWorks
//   - "/0000-0003-1401-2056/person" -> calls GetPerson
//   - "/0000-0003-1401-2056" or "/0000-0003-1401-2056/record" -> calls GetRecord
func (c *Client) GetByPath(ctx context.Context, path Path) (interface{}, error) {
	pathStr := string(path)

	// Extract ORCID ID and resource type from path
	// Path format: /{orcid-id}/{resource-type} or /{orcid-id}
	pathStr = strings.TrimPrefix(pathStr, "/")
	parts := strings.SplitN(pathStr, "/", 2)

	if len(parts) == 0 || parts[0] == "" {
		return nil, fmt.Errorf("invalid path: %s", path)
	}

	orcidID := parts[0]

	// If no resource type specified, or if it's "record", get the full record
	if len(parts) == 1 || (len(parts) == 2 && parts[1] == "record") {
		return c.GetRecord(ctx, orcidID)
	}

	// Route to appropriate method based on resource type
	resourceType := parts[1]

	// Handle paths with put-codes (e.g., "/0000-0003-1401-2056/work/92636200")
	resourceParts := strings.SplitN(resourceType, "/", 2)
	baseResource := resourceParts[0]

	switch baseResource {
	case "person":
		return c.GetPerson(ctx, orcidID)
	case "works":
		return c.GetWorks(ctx, orcidID)
	case "work":
		if len(resourceParts) == 2 {
			return c.GetWork(ctx, orcidID, resourceParts[1])
		}
		return nil, fmt.Errorf("work path requires put-code: %s", path)
	case "educations":
		return c.GetEducations(ctx, orcidID)
	case "employments":
		return c.GetEmployments(ctx, orcidID)
	case "fundings":
		return c.GetFundings(ctx, orcidID)
	case "peer-reviews":
		return c.GetPeerReviews(ctx, orcidID)
	case "distinctions":
		return c.GetDistinctions(ctx, orcidID)
	case "invited-positions":
		return c.GetInvitedPositions(ctx, orcidID)
	case "memberships":
		return c.GetMemberships(ctx, orcidID)
	case "qualifications":
		return c.GetQualifications(ctx, orcidID)
	case "services":
		return c.GetServices(ctx, orcidID)
	case "research-resources":
		return c.GetResearchResources(ctx, orcidID)
	case "activities":
		// Activities summary is part of the record
		record, err := c.GetRecord(ctx, orcidID)
		if err != nil {
			return nil, err
		}
		return record.ActivitiesSummary, nil
	case "biography", "other-names", "researcher-urls", "email", "address", "keywords", "external-identifiers":
		// These are part of the person record
		person, err := c.GetPerson(ctx, orcidID)
		if err != nil {
			return nil, err
		}
		switch baseResource {
		case "biography":
			return person.Biography, nil
		case "other-names":
			return person.OtherNames, nil
		case "researcher-urls":
			return person.ResearcherURLs, nil
		case "email":
			return person.Emails, nil
		case "address":
			return person.Addresses, nil
		case "keywords":
			return person.Keywords, nil
		case "external-identifiers":
			return person.ExternalIdentifiers, nil
		}
	default:
		// Try to handle specific put-code paths for other resources
		// e.g., "/0000-0003-1401-2056/keywords/1925453"
		person, err := c.GetPerson(ctx, orcidID)
		if err != nil {
			return nil, fmt.Errorf("unsupported resource type in path: %s", path)
		}

		// Check if this is a specific item within person data
		switch baseResource {
		case "keywords":
			return person.Keywords, nil
		case "address":
			return person.Addresses, nil
		case "researcher-urls":
			return person.ResearcherURLs, nil
		default:
			return nil, fmt.Errorf("unsupported resource type in path: %s", path)
		}
	}

	return nil, fmt.Errorf("unsupported path: %s", path)
}
