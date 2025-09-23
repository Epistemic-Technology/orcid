package orcid

import (
	"time"
)

// Record represents the complete ORCID record
type Record struct {
	OrcidIdentifier   *OrcidIdentifier   `json:"orcid-identifier,omitempty"`
	Preferences       *Preferences       `json:"preferences,omitempty"`
	History           *History           `json:"history,omitempty"`
	Person            *Person            `json:"person,omitempty"`
	ActivitiesSummary *ActivitiesSummary `json:"activities-summary,omitempty"`
	Path              string             `json:"path,omitempty"`
}

// OrcidIdentifier represents the ORCID identifier information
type OrcidIdentifier struct {
	URI  string `json:"uri"`
	Path string `json:"path"`
	Host string `json:"host"`
}

// Preferences represents user preferences
type Preferences struct {
	Locale string `json:"locale"`
}

// History represents the history of the ORCID record
type History struct {
	CreationMethod       string     `json:"creation-method,omitempty"`
	CompletionDate       *DateValue `json:"completion-date,omitempty"`
	SubmissionDate       *DateValue `json:"submission-date,omitempty"`
	LastModifiedDate     *DateValue `json:"last-modified-date,omitempty"`
	Claimed              bool       `json:"claimed"`
	Source               *Source    `json:"source,omitempty"`
	DeactivationDate     *DateValue `json:"deactivation-date,omitempty"`
	VerifiedEmail        bool       `json:"verified-email"`
	VerifiedPrimaryEmail bool       `json:"verified-primary-email"`
}

// DateValue represents a date/timestamp value
type DateValue struct {
	Value int64 `json:"value"`
}

// Person represents personal information
type Person struct {
	LastModifiedDate    *DateValue           `json:"last-modified-date,omitempty"`
	Name                *Name                `json:"name,omitempty"`
	OtherNames          *OtherNames          `json:"other-names,omitempty"`
	Biography           *Biography           `json:"biography,omitempty"`
	ResearcherUrls      *ResearcherUrls      `json:"researcher-urls,omitempty"`
	Emails              *Emails              `json:"emails,omitempty"`
	Addresses           *Addresses           `json:"addresses,omitempty"`
	Keywords            *Keywords            `json:"keywords,omitempty"`
	ExternalIdentifiers *ExternalIdentifiers `json:"external-identifiers,omitempty"`
	Path                string               `json:"path,omitempty"`
}

// Name represents a person's name information
type Name struct {
	CreatedDate      *DateValue   `json:"created-date,omitempty"`
	LastModifiedDate *DateValue   `json:"last-modified-date,omitempty"`
	GivenNames       *StringValue `json:"given-names,omitempty"`
	FamilyName       *StringValue `json:"family-name,omitempty"`
	CreditName       *StringValue `json:"credit-name,omitempty"`
	Source           *Source      `json:"source,omitempty"`
	Visibility       string       `json:"visibility,omitempty"`
	Path             string       `json:"path,omitempty"`
}

// StringValue represents a string value wrapper
type StringValue struct {
	Value string `json:"value"`
}

// OtherNames represents alternative names
type OtherNames struct {
	LastModifiedDate *DateValue   `json:"last-modified-date,omitempty"`
	OtherName        []*OtherName `json:"other-name,omitempty"`
	Path             string       `json:"path,omitempty"`
}

// OtherName represents a single other name
type OtherName struct {
	CreatedDate      *DateValue `json:"created-date,omitempty"`
	LastModifiedDate *DateValue `json:"last-modified-date,omitempty"`
	Source           *Source    `json:"source,omitempty"`
	Content          string     `json:"content"`
	Visibility       string     `json:"visibility,omitempty"`
	Path             string     `json:"path,omitempty"`
	PutCode          int64      `json:"put-code,omitempty"`
	DisplayIndex     any        `json:"display-index,omitempty"`
}

// Biography represents biographical information
type Biography struct {
	CreatedDate      *DateValue `json:"created-date,omitempty"`
	LastModifiedDate *DateValue `json:"last-modified-date,omitempty"`
	Content          string     `json:"content"`
	Visibility       string     `json:"visibility,omitempty"`
	Path             string     `json:"path,omitempty"`
}

// ResearcherUrls represents researcher URLs
type ResearcherUrls struct {
	LastModifiedDate *DateValue       `json:"last-modified-date,omitempty"`
	ResearcherUrl    []*ResearcherUrl `json:"researcher-url,omitempty"`
	Path             string           `json:"path,omitempty"`
}

// ResearcherUrl represents a single researcher URL
type ResearcherUrl struct {
	CreatedDate      *DateValue   `json:"created-date,omitempty"`
	LastModifiedDate *DateValue   `json:"last-modified-date,omitempty"`
	Source           *Source      `json:"source,omitempty"`
	UrlName          string       `json:"url-name"`
	Url              *StringValue `json:"url"`
	Visibility       string       `json:"visibility,omitempty"`
	Path             string       `json:"path,omitempty"`
	PutCode          int64        `json:"put-code,omitempty"`
	DisplayIndex     any          `json:"display-index,omitempty"`
}

// Emails represents email addresses
type Emails struct {
	LastModifiedDate *DateValue `json:"last-modified-date,omitempty"`
	Email            []*Email   `json:"email,omitempty"`
	Path             string     `json:"path,omitempty"`
}

// Email represents a single email address
type Email struct {
	CreatedDate      *DateValue `json:"created-date,omitempty"`
	LastModifiedDate *DateValue `json:"last-modified-date,omitempty"`
	Source           *Source    `json:"source,omitempty"`
	Email            string     `json:"email"`
	Primary          bool       `json:"primary,omitempty"`
	Verified         bool       `json:"verified,omitempty"`
	Visibility       string     `json:"visibility,omitempty"`
	Path             string     `json:"path,omitempty"`
	PutCode          int64      `json:"put-code,omitempty"`
}

// Addresses represents addresses
type Addresses struct {
	LastModifiedDate *DateValue `json:"last-modified-date,omitempty"`
	Address          []*Address `json:"address,omitempty"`
	Path             string     `json:"path,omitempty"`
}

// Address represents a single address
type Address struct {
	CreatedDate      *DateValue   `json:"created-date,omitempty"`
	LastModifiedDate *DateValue   `json:"last-modified-date,omitempty"`
	Source           *Source      `json:"source,omitempty"`
	Country          *StringValue `json:"country,omitempty"`
	Visibility       string       `json:"visibility,omitempty"`
	Path             string       `json:"path,omitempty"`
	PutCode          int64        `json:"put-code,omitempty"`
	DisplayIndex     any          `json:"display-index,omitempty"`
}

// Keywords represents keywords
type Keywords struct {
	LastModifiedDate *DateValue `json:"last-modified-date,omitempty"`
	Keyword          []*Keyword `json:"keyword,omitempty"`
	Path             string     `json:"path,omitempty"`
}

// Keyword represents a single keyword
type Keyword struct {
	CreatedDate      *DateValue `json:"created-date,omitempty"`
	LastModifiedDate *DateValue `json:"last-modified-date,omitempty"`
	Source           *Source    `json:"source,omitempty"`
	Content          string     `json:"content"`
	Visibility       string     `json:"visibility,omitempty"`
	Path             string     `json:"path,omitempty"`
	PutCode          int64      `json:"put-code,omitempty"`
	DisplayIndex     any        `json:"display-index,omitempty"`
}

// ExternalIdentifiers represents external identifiers
type ExternalIdentifiers struct {
	LastModifiedDate   *DateValue            `json:"last-modified-date,omitempty"`
	ExternalIdentifier []*ExternalIdentifier `json:"external-identifier,omitempty"`
	Path               string                `json:"path,omitempty"`
}

// ExternalIdentifier represents a single external identifier
type ExternalIdentifier struct {
	CreatedDate            *DateValue   `json:"created-date,omitempty"`
	LastModifiedDate       *DateValue   `json:"last-modified-date,omitempty"`
	Source                 *Source      `json:"source,omitempty"`
	ExternalIdType         string       `json:"external-id-type"`
	ExternalIdValue        string       `json:"external-id-value"`
	ExternalIdUrl          *StringValue `json:"external-id-url,omitempty"`
	ExternalIdRelationship string       `json:"external-id-relationship,omitempty"`
	Visibility             string       `json:"visibility,omitempty"`
	Path                   string       `json:"path,omitempty"`
	PutCode                int64        `json:"put-code,omitempty"`
	DisplayIndex           any          `json:"display-index,omitempty"`
}

// Source represents the source of information
type Source struct {
	SourceOrcid             *OrcidIdentifier `json:"source-orcid,omitempty"`
	SourceClientId          *ClientId        `json:"source-client-id,omitempty"`
	SourceName              *StringValue     `json:"source-name,omitempty"`
	AssertionOriginOrcid    *OrcidIdentifier `json:"assertion-origin-orcid,omitempty"`
	AssertionOriginClientId *ClientId        `json:"assertion-origin-client-id,omitempty"`
	AssertionOriginName     *StringValue     `json:"assertion-origin-name,omitempty"`
}

// ClientId represents a client identifier
type ClientId struct {
	URI  string `json:"uri"`
	Path string `json:"path"`
	Host string `json:"host"`
}

// ActivitiesSummary represents a summary of activities
type ActivitiesSummary struct {
	LastModifiedDate  *DateValue         `json:"last-modified-date,omitempty"`
	Distinctions      *Distinctions      `json:"distinctions,omitempty"`
	Educations        *Educations        `json:"educations,omitempty"`
	Employments       *Employments       `json:"employments,omitempty"`
	Fundings          *Fundings          `json:"fundings,omitempty"`
	InvitedPositions  *InvitedPositions  `json:"invited-positions,omitempty"`
	Memberships       *Memberships       `json:"memberships,omitempty"`
	PeerReviews       *PeerReviews       `json:"peer-reviews,omitempty"`
	Qualifications    *Qualifications    `json:"qualifications,omitempty"`
	ResearchResources *ResearchResources `json:"research-resources,omitempty"`
	Services          *Services          `json:"services,omitempty"`
	Works             *Works             `json:"works,omitempty"`
	Path              string             `json:"path,omitempty"`
}

// Distinctions represents distinctions
type Distinctions struct {
	LastModifiedDate *DateValue          `json:"last-modified-date,omitempty"`
	AffiliationGroup []*AffiliationGroup `json:"affiliation-group,omitempty"`
	Path             string              `json:"path,omitempty"`
}

// Educations represents education affiliations
type Educations struct {
	LastModifiedDate *DateValue          `json:"last-modified-date,omitempty"`
	AffiliationGroup []*AffiliationGroup `json:"affiliation-group,omitempty"`
	Path             string              `json:"path,omitempty"`
}

// Employments represents employment affiliations
type Employments struct {
	LastModifiedDate *DateValue          `json:"last-modified-date,omitempty"`
	AffiliationGroup []*AffiliationGroup `json:"affiliation-group,omitempty"`
	Path             string              `json:"path,omitempty"`
}

// InvitedPositions represents invited positions
type InvitedPositions struct {
	LastModifiedDate *DateValue          `json:"last-modified-date,omitempty"`
	AffiliationGroup []*AffiliationGroup `json:"affiliation-group,omitempty"`
	Path             string              `json:"path,omitempty"`
}

// Memberships represents memberships
type Memberships struct {
	LastModifiedDate *DateValue          `json:"last-modified-date,omitempty"`
	AffiliationGroup []*AffiliationGroup `json:"affiliation-group,omitempty"`
	Path             string              `json:"path,omitempty"`
}

// Qualifications represents qualifications
type Qualifications struct {
	LastModifiedDate *DateValue          `json:"last-modified-date,omitempty"`
	AffiliationGroup []*AffiliationGroup `json:"affiliation-group,omitempty"`
	Path             string              `json:"path,omitempty"`
}

// Services represents services
type Services struct {
	LastModifiedDate *DateValue          `json:"last-modified-date,omitempty"`
	AffiliationGroup []*AffiliationGroup `json:"affiliation-group,omitempty"`
	Path             string              `json:"path,omitempty"`
}

// AffiliationGroup represents a group of affiliations
type AffiliationGroup struct {
	LastModifiedDate *DateValue                `json:"last-modified-date,omitempty"`
	ExternalIds      *ExternalIds              `json:"external-ids,omitempty"`
	Summaries        []*AffiliationSummaryWrap `json:"summaries,omitempty"`
}

// AffiliationSummaryWrap wraps the different types of affiliation summaries
type AffiliationSummaryWrap struct {
	EducationSummary       *AffiliationSummary `json:"education-summary,omitempty"`
	EmploymentSummary      *AffiliationSummary `json:"employment-summary,omitempty"`
	DistinctionSummary     *AffiliationSummary `json:"distinction-summary,omitempty"`
	InvitedPositionSummary *AffiliationSummary `json:"invited-position-summary,omitempty"`
	MembershipSummary      *AffiliationSummary `json:"membership-summary,omitempty"`
	QualificationSummary   *AffiliationSummary `json:"qualification-summary,omitempty"`
	ServiceSummary         *AffiliationSummary `json:"service-summary,omitempty"`
}

// AffiliationSummary represents a summary of an affiliation
type AffiliationSummary struct {
	CreatedDate      *DateValue    `json:"created-date,omitempty"`
	LastModifiedDate *DateValue    `json:"last-modified-date,omitempty"`
	Source           *Source       `json:"source,omitempty"`
	PutCode          int64         `json:"put-code,omitempty"`
	DepartmentName   string        `json:"department-name,omitempty"`
	RoleTitle        string        `json:"role-title,omitempty"`
	StartDate        *FuzzyDate    `json:"start-date,omitempty"`
	EndDate          *FuzzyDate    `json:"end-date,omitempty"`
	Organization     *Organization `json:"organization,omitempty"`
	Url              *StringValue  `json:"url,omitempty"`
	ExternalIds      *ExternalIds  `json:"external-ids,omitempty"`
	DisplayIndex     any           `json:"display-index,omitempty"`
	Visibility       string        `json:"visibility,omitempty"`
	Path             string        `json:"path,omitempty"`
}

// FuzzyDate represents a date that may be incomplete
type FuzzyDate struct {
	Year  *YearValue  `json:"year,omitempty"`
	Month *MonthValue `json:"month,omitempty"`
	Day   *DayValue   `json:"day,omitempty"`
}

// YearValue represents a year value (can be string or int in JSON)
type YearValue struct {
	Value string `json:"value"`
}

// MonthValue represents a month value (can be string or int in JSON)
type MonthValue struct {
	Value string `json:"value"`
}

// DayValue represents a day value (can be string or int in JSON)
type DayValue struct {
	Value string `json:"value"`
}

// Organization represents an organization
type Organization struct {
	Name                      string                     `json:"name"`
	Address                   *OrganizationAddress       `json:"address,omitempty"`
	DisambiguatedOrganization *DisambiguatedOrganization `json:"disambiguated-organization,omitempty"`
}

// OrganizationAddress represents an organization's address
type OrganizationAddress struct {
	City    string `json:"city,omitempty"`
	Region  string `json:"region,omitempty"`
	Country string `json:"country,omitempty"`
}

// DisambiguatedOrganization represents a disambiguated organization
type DisambiguatedOrganization struct {
	DisambiguatedOrganizationIdentifier string `json:"disambiguated-organization-identifier"`
	DisambiguationSource                string `json:"disambiguation-source"`
}

// ExternalIds represents a collection of external IDs
type ExternalIds struct {
	ExternalId []*ExternalId `json:"external-id,omitempty"`
}

// ExternalId represents a single external ID
type ExternalId struct {
	ExternalIdType            string                `json:"external-id-type"`
	ExternalIdValue           string                `json:"external-id-value"`
	ExternalIdNormalized      *ExternalIdNormalized `json:"external-id-normalized,omitempty"`
	ExternalIdNormalizedError string                `json:"external-id-normalized-error,omitempty"`
	ExternalIdUrl             *StringValue          `json:"external-id-url,omitempty"`
	ExternalIdRelationship    string                `json:"external-id-relationship,omitempty"`
}

// ExternalIdNormalized represents a normalized external ID
type ExternalIdNormalized struct {
	Value     string `json:"value"`
	Transient bool   `json:"transient,omitempty"`
}

// Fundings represents funding information
type Fundings struct {
	LastModifiedDate *DateValue      `json:"last-modified-date,omitempty"`
	Group            []*FundingGroup `json:"group,omitempty"`
	Path             string          `json:"path,omitempty"`
}

// FundingGroup represents a group of funding items
type FundingGroup struct {
	LastModifiedDate *DateValue        `json:"last-modified-date,omitempty"`
	ExternalIds      *ExternalIds      `json:"external-ids,omitempty"`
	FundingSummary   []*FundingSummary `json:"funding-summary,omitempty"`
}

// FundingSummary represents a summary of a funding item
type FundingSummary struct {
	CreatedDate      *DateValue    `json:"created-date,omitempty"`
	LastModifiedDate *DateValue    `json:"last-modified-date,omitempty"`
	Source           *Source       `json:"source,omitempty"`
	PutCode          int64         `json:"put-code,omitempty"`
	Title            *FundingTitle `json:"title,omitempty"`
	Type             string        `json:"type,omitempty"`
	StartDate        *FuzzyDate    `json:"start-date,omitempty"`
	EndDate          *FuzzyDate    `json:"end-date,omitempty"`
	Organization     *Organization `json:"organization,omitempty"`
	Url              *StringValue  `json:"url,omitempty"`
	ExternalIds      *ExternalIds  `json:"external-ids,omitempty"`
	DisplayIndex     any           `json:"display-index,omitempty"`
	Visibility       string        `json:"visibility,omitempty"`
	Path             string        `json:"path,omitempty"`
}

// FundingTitle represents a funding title
type FundingTitle struct {
	Title           *StringValue     `json:"title,omitempty"`
	TranslatedTitle *TranslatedTitle `json:"translated-title,omitempty"`
}

// TranslatedTitle represents a translated title
type TranslatedTitle struct {
	Value        string `json:"value"`
	LanguageCode string `json:"language-code,omitempty"`
}

// PeerReviews represents peer review information
type PeerReviews struct {
	LastModifiedDate *DateValue         `json:"last-modified-date,omitempty"`
	Group            []*PeerReviewGroup `json:"group,omitempty"`
	Path             string             `json:"path,omitempty"`
}

// PeerReviewGroup represents a group of peer reviews
type PeerReviewGroup struct {
	LastModifiedDate *DateValue              `json:"last-modified-date,omitempty"`
	ExternalIds      *ExternalIds            `json:"external-ids,omitempty"`
	PeerReviewGroup  []*PeerReviewGroupInner `json:"peer-review-group,omitempty"`
}

// PeerReviewGroupInner represents an inner peer review group
type PeerReviewGroupInner struct {
	LastModifiedDate  *DateValue           `json:"last-modified-date,omitempty"`
	ExternalIds       *ExternalIds         `json:"external-ids,omitempty"`
	PeerReviewSummary []*PeerReviewSummary `json:"peer-review-summary,omitempty"`
}

// PeerReviewSummary represents a summary of a peer review
type PeerReviewSummary struct {
	CreatedDate           *DateValue    `json:"created-date,omitempty"`
	LastModifiedDate      *DateValue    `json:"last-modified-date,omitempty"`
	Source                *Source       `json:"source,omitempty"`
	ReviewerRole          string        `json:"reviewer-role,omitempty"`
	ExternalIds           *ExternalIds  `json:"external-ids,omitempty"`
	ReviewUrl             *StringValue  `json:"review-url,omitempty"`
	ReviewType            string        `json:"review-type,omitempty"`
	CompletionDate        *FuzzyDate    `json:"completion-date,omitempty"`
	ReviewGroupId         string        `json:"review-group-id,omitempty"`
	ConveningOrganization *Organization `json:"convening-organization,omitempty"`
	Visibility            string        `json:"visibility,omitempty"`
	PutCode               int64         `json:"put-code,omitempty"`
	Path                  string        `json:"path,omitempty"`
	DisplayIndex          any           `json:"display-index,omitempty"`
}

// ResearchResources represents research resources
type ResearchResources struct {
	LastModifiedDate *DateValue               `json:"last-modified-date,omitempty"`
	Group            []*ResearchResourceGroup `json:"group,omitempty"`
	Path             string                   `json:"path,omitempty"`
}

// ResearchResourceGroup represents a group of research resources
type ResearchResourceGroup struct {
	LastModifiedDate        *DateValue                 `json:"last-modified-date,omitempty"`
	ExternalIds             *ExternalIds               `json:"external-ids,omitempty"`
	ResearchResourceSummary []*ResearchResourceSummary `json:"research-resource-summary,omitempty"`
}

// ResearchResourceSummary represents a summary of a research resource
type ResearchResourceSummary struct {
	CreatedDate      *DateValue `json:"created-date,omitempty"`
	LastModifiedDate *DateValue `json:"last-modified-date,omitempty"`
	Source           *Source    `json:"source,omitempty"`
	PutCode          int64      `json:"put-code,omitempty"`
	DisplayIndex     any        `json:"display-index,omitempty"`
	Visibility       string     `json:"visibility,omitempty"`
	Path             string     `json:"path,omitempty"`
}

// Works represents works (publications)
type Works struct {
	LastModifiedDate *DateValue   `json:"last-modified-date,omitempty"`
	Group            []*WorkGroup `json:"group,omitempty"`
	Path             string       `json:"path,omitempty"`
}

// WorkGroup represents a group of works
type WorkGroup struct {
	LastModifiedDate *DateValue     `json:"last-modified-date,omitempty"`
	ExternalIds      *ExternalIds   `json:"external-ids,omitempty"`
	WorkSummary      []*WorkSummary `json:"work-summary,omitempty"`
}

// WorkSummary represents a summary of a work
type WorkSummary struct {
	PutCode          int64        `json:"put-code,omitempty"`
	CreatedDate      *DateValue   `json:"created-date,omitempty"`
	LastModifiedDate *DateValue   `json:"last-modified-date,omitempty"`
	Source           *Source      `json:"source,omitempty"`
	Title            *WorkTitle   `json:"title,omitempty"`
	ExternalIds      *ExternalIds `json:"external-ids,omitempty"`
	Url              *StringValue `json:"url,omitempty"`
	Type             string       `json:"type,omitempty"`
	PublicationDate  *FuzzyDate   `json:"publication-date,omitempty"`
	JournalTitle     *StringValue `json:"journal-title,omitempty"`
	Visibility       string       `json:"visibility,omitempty"`
	Path             string       `json:"path,omitempty"`
	DisplayIndex     any          `json:"display-index,omitempty"`
}

// WorkTitle represents a work title
type WorkTitle struct {
	Title           *StringValue     `json:"title,omitempty"`
	Subtitle        *StringValue     `json:"subtitle,omitempty"`
	TranslatedTitle *TranslatedTitle `json:"translated-title,omitempty"`
}

// Time returns the time.Time representation of a DateValue
func (d *DateValue) Time() time.Time {
	if d == nil {
		return time.Time{}
	}
	return time.Unix(0, d.Value*int64(time.Millisecond))
}

// String returns the string representation of a StringValue
func (s *StringValue) String() string {
	if s == nil {
		return ""
	}
	return s.Value
}
