package orcid

import (
	"encoding/json"
	"time"
)

// Path is a type alias for ORCID API paths
type Path string

type Record struct {
	OrcidIdentifier   *OrcidIdentifier   `json:"orcid-identifier,omitempty" xml:"orcid-identifier,omitempty"`
	Preferences       *Preferences       `json:"preferences,omitempty" xml:"preferences,omitempty"`
	History           *History           `json:"history,omitempty" xml:"history,omitempty"`
	Person            *Person            `json:"person,omitempty" xml:"person,omitempty"`
	ActivitiesSummary *ActivitiesSummary `json:"activities-summary,omitempty" xml:"activities-summary,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type OrcidIdentifier struct {
	URI  string `json:"uri,omitempty" xml:"uri,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,omitempty"`
	Host string `json:"host,omitempty" xml:"host,omitempty"`
}

type Preferences struct {
	Locale string `json:"locale,omitempty" xml:"locale,omitempty"`
}

type History struct {
	CreationMethod       string  `json:"creation-method,omitempty" xml:"creation-method,omitempty"`
	CompletionDate       *Date   `json:"completion-date,omitempty" xml:"completion-date,omitempty"`
	SubmissionDate       *Date   `json:"submission-date,omitempty" xml:"submission-date,omitempty"`
	LastModifiedDate     *Date   `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Claimed              bool    `json:"claimed,omitempty" xml:"claimed,omitempty"`
	Source               *Source `json:"source,omitempty" xml:"source,omitempty"`
	DeactivationDate     *Date   `json:"deactivation-date,omitempty" xml:"deactivation-date,omitempty"`
	VerifiedEmail        bool    `json:"verified-email,omitempty" xml:"verified-email,omitempty"`
	VerifiedPrimaryEmail bool    `json:"verified-primary-email,omitempty" xml:"verified-primary-email,omitempty"`
}

type Date struct {
	Value time.Time `json:"value,omitempty" xml:"value,omitempty"`
}

// UnmarshalJSON handles both Unix milliseconds and time strings
func (d *Date) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch value := v.(type) {
	case map[string]interface{}:
		if val, ok := value["value"]; ok {
			switch v := val.(type) {
			case float64:
				// Unix milliseconds
				d.Value = time.Unix(0, int64(v)*int64(time.Millisecond))
				return nil
			case string:
				// ISO 8601 string
				t, err := time.Parse(time.RFC3339, v)
				if err != nil {
					return err
				}
				d.Value = t
				return nil
			}
		}
	}
	return nil
}

type Source struct {
	SourceOrcid             *OrcidIdentifier `json:"source-orcid,omitempty" xml:"source-orcid,omitempty"`
	SourceClientID          *SourceClientID  `json:"source-client-id,omitempty" xml:"source-client-id,omitempty"`
	SourceName              *SourceName      `json:"source-name,omitempty" xml:"source-name,omitempty"`
	AssertionOriginOrcid    *OrcidIdentifier `json:"assertion-origin-orcid,omitempty" xml:"assertion-origin-orcid,omitempty"`
	AssertionOriginClientID *SourceClientID  `json:"assertion-origin-client-id,omitempty" xml:"assertion-origin-client-id,omitempty"`
	AssertionOriginName     *SourceName      `json:"assertion-origin-name,omitempty" xml:"assertion-origin-name,omitempty"`
}

type SourceClientID struct {
	URI  string `json:"uri,omitempty" xml:"uri,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,omitempty"`
	Host string `json:"host,omitempty" xml:"host,omitempty"`
}

type SourceName struct {
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

type Person struct {
	Name                *Name                `json:"name,omitempty" xml:"name,omitempty"`
	OtherNames          *OtherNames          `json:"other-names,omitempty" xml:"other-names,omitempty"`
	Biography           *Biography           `json:"biography,omitempty" xml:"biography,omitempty"`
	ResearcherURLs      *ResearcherURLs      `json:"researcher-urls,omitempty" xml:"researcher-urls,omitempty"`
	Emails              *Emails              `json:"emails,omitempty" xml:"emails,omitempty"`
	Addresses           *Addresses           `json:"addresses,omitempty" xml:"addresses,omitempty"`
	Keywords            *Keywords            `json:"keywords,omitempty" xml:"keywords,omitempty"`
	ExternalIdentifiers *ExternalIdentifiers `json:"external-identifiers,omitempty" xml:"external-identifiers,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type Name struct {
	CreatedDate      *Date       `json:"created-date,omitempty" xml:"created-date,omitempty"`
	LastModifiedDate *Date       `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	GivenNames       *GivenNames `json:"given-names,omitempty" xml:"given-names,omitempty"`
	FamilyName       *FamilyName `json:"family-name,omitempty" xml:"family-name,omitempty"`
	CreditName       *CreditName `json:"credit-name,omitempty" xml:"credit-name,omitempty"`
	Source           *Source     `json:"source,omitempty" xml:"source,omitempty"`
	Visibility       string      `json:"visibility,omitempty" xml:"visibility,attr,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type GivenNames struct {
	Value string `json:"value,omitempty" xml:",chardata"`
}

type FamilyName struct {
	Value string `json:"value,omitempty" xml:",chardata"`
}

type CreditName struct {
	Value string `json:"value,omitempty" xml:",chardata"`
}

type OtherNames struct {
	LastModifiedDate *Date        `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	OtherName        []*OtherName `json:"other-name,omitempty" xml:"other-name,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type OtherName struct {
	CreatedDate      *Date   `json:"created-date,omitempty" xml:"created-date,omitempty"`
	LastModifiedDate *Date   `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Source           *Source `json:"source,omitempty" xml:"source,omitempty"`
	Content          string  `json:"content,omitempty" xml:"content,omitempty"`
	Visibility       string  `json:"visibility,omitempty" xml:"visibility,attr,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
	PutCode          int64   `json:"put-code,omitempty" xml:"put-code,attr,omitempty"`
	DisplayIndex     int     `json:"display-index,omitempty" xml:"display-index,attr,omitempty"`
}

type Biography struct {
	CreatedDate      *Date  `json:"created-date,omitempty" xml:"created-date,omitempty"`
	LastModifiedDate *Date  `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Content          string `json:"content,omitempty" xml:"content,omitempty"`
	Visibility       string `json:"visibility,omitempty" xml:"visibility,attr,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type ResearcherURLs struct {
	LastModifiedDate *Date            `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	ResearcherURL    []*ResearcherURL `json:"researcher-url,omitempty" xml:"researcher-url,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type ResearcherURL struct {
	CreatedDate      *Date   `json:"created-date,omitempty" xml:"created-date,omitempty"`
	LastModifiedDate *Date   `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Source           *Source `json:"source,omitempty" xml:"source,omitempty"`
	URLName          string  `json:"url-name,omitempty" xml:"url-name,omitempty"`
	URL              *URL    `json:"url,omitempty" xml:"url,omitempty"`
	Visibility       string  `json:"visibility,omitempty" xml:"visibility,attr,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
	PutCode          int64   `json:"put-code,omitempty" xml:"put-code,attr,omitempty"`
	DisplayIndex     int     `json:"display-index,omitempty" xml:"display-index,attr,omitempty"`
}

type URL struct {
	Value string `json:"value,omitempty" xml:",chardata"`
}

type Emails struct {
	LastModifiedDate *Date    `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Email            []*Email `json:"email,omitempty" xml:"email,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type Email struct {
	CreatedDate      *Date   `json:"created-date,omitempty" xml:"created-date,omitempty"`
	LastModifiedDate *Date   `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Source           *Source `json:"source,omitempty" xml:"source,omitempty"`
	Email            string  `json:"email,omitempty" xml:"email,omitempty"`
	Primary          bool    `json:"primary,omitempty" xml:"primary,omitempty"`
	Verified         bool    `json:"verified,omitempty" xml:"verified,omitempty"`
	Visibility       string  `json:"visibility,omitempty" xml:"visibility,attr,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
	PutCode          int64   `json:"put-code,omitempty" xml:"put-code,attr,omitempty"`
}

type Addresses struct {
	LastModifiedDate *Date      `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Address          []*Address `json:"address,omitempty" xml:"address,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type Address struct {
	CreatedDate      *Date    `json:"created-date,omitempty" xml:"created-date,omitempty"`
	LastModifiedDate *Date    `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Source           *Source  `json:"source,omitempty" xml:"source,omitempty"`
	Country          *Country `json:"country,omitempty" xml:"country,omitempty"`
	Visibility       string   `json:"visibility,omitempty" xml:"visibility,attr,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
	PutCode          int64    `json:"put-code,omitempty" xml:"put-code,attr,omitempty"`
	DisplayIndex     int      `json:"display-index,omitempty" xml:"display-index,attr,omitempty"`
}

type Country struct {
	Value string `json:"value,omitempty" xml:",chardata"`
}

type Keywords struct {
	LastModifiedDate *Date      `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Keyword          []*Keyword `json:"keyword,omitempty" xml:"keyword,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type Keyword struct {
	CreatedDate      *Date   `json:"created-date,omitempty" xml:"created-date,omitempty"`
	LastModifiedDate *Date   `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Source           *Source `json:"source,omitempty" xml:"source,omitempty"`
	Content          string  `json:"content,omitempty" xml:"content,omitempty"`
	Visibility       string  `json:"visibility,omitempty" xml:"visibility,attr,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
	PutCode          int64   `json:"put-code,omitempty" xml:"put-code,attr,omitempty"`
	DisplayIndex     int     `json:"display-index,omitempty" xml:"display-index,attr,omitempty"`
}

type ExternalIdentifiers struct {
	LastModifiedDate   *Date                 `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	ExternalIdentifier []*ExternalIdentifier `json:"external-identifier,omitempty" xml:"external-identifier,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type ExternalIdentifier struct {
	CreatedDate                    *Date   `json:"created-date,omitempty" xml:"created-date,omitempty"`
	LastModifiedDate               *Date   `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Source                         *Source `json:"source,omitempty" xml:"source,omitempty"`
	ExternalIdentifierType         string  `json:"external-id-type,omitempty" xml:"external-id-type,omitempty"`
	ExternalIdentifierValue        string  `json:"external-id-value,omitempty" xml:"external-id-value,omitempty"`
	ExternalIdentifierURL          *URL    `json:"external-id-url,omitempty" xml:"external-id-url,omitempty"`
	ExternalIdentifierRelationship string  `json:"external-id-relationship,omitempty" xml:"external-id-relationship,omitempty"`
	Visibility                     string  `json:"visibility,omitempty" xml:"visibility,attr,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
	PutCode                        int64   `json:"put-code,omitempty" xml:"put-code,attr,omitempty"`
	DisplayIndex                   int     `json:"display-index,omitempty" xml:"display-index,attr,omitempty"`
}

type ActivitiesSummary struct {
	LastModifiedDate  *Date              `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Distinctions      *Distinctions      `json:"distinctions,omitempty" xml:"distinctions,omitempty"`
	Educations        *Educations        `json:"educations,omitempty" xml:"educations,omitempty"`
	Employments       *Employments       `json:"employments,omitempty" xml:"employments,omitempty"`
	Fundings          *Fundings          `json:"fundings,omitempty" xml:"fundings,omitempty"`
	InvitedPositions  *InvitedPositions  `json:"invited-positions,omitempty" xml:"invited-positions,omitempty"`
	Memberships       *Memberships       `json:"memberships,omitempty" xml:"memberships,omitempty"`
	PeerReviews       *PeerReviews       `json:"peer-reviews,omitempty" xml:"peer-reviews,omitempty"`
	Qualifications    *Qualifications    `json:"qualifications,omitempty" xml:"qualifications,omitempty"`
	ResearchResources *ResearchResources `json:"research-resources,omitempty" xml:"research-resources,omitempty"`
	Services          *Services          `json:"services,omitempty" xml:"services,omitempty"`
	Works             *Works             `json:"works,omitempty" xml:"works,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type Works struct {
	LastModifiedDate *Date        `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	WorkGroup        []*WorkGroup `json:"group,omitempty" xml:"group,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type WorkGroup struct {
	LastModifiedDate *Date          `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	ExternalIDs      *ExternalIDs   `json:"external-ids,omitempty" xml:"external-ids,omitempty"`
	WorkSummary      []*WorkSummary `json:"work-summary,omitempty" xml:"work-summary,omitempty"`
}

type ExternalIDs struct {
	ExternalID []*ExternalID `json:"external-id,omitempty" xml:"external-id,omitempty"`
}

type ExternalID struct {
	ExternalIDType         string                `json:"external-id-type,omitempty" xml:"external-id-type,omitempty"`
	ExternalIDValue        string                `json:"external-id-value,omitempty" xml:"external-id-value,omitempty"`
	ExternalIDNormalized   *ExternalIDNormalized `json:"external-id-normalized,omitempty" xml:"external-id-normalized,omitempty"`
	ExternalIDURL          *URL                  `json:"external-id-url,omitempty" xml:"external-id-url,omitempty"`
	ExternalIDRelationship string                `json:"external-id-relationship,omitempty" xml:"external-id-relationship,omitempty"`
}

type ExternalIDNormalized struct {
	Value     string `json:"value,omitempty" xml:",chardata"`
	Transient bool   `json:"transient,omitempty" xml:"transient,attr,omitempty"`
}

type WorkSummary struct {
	PutCode          int64            `json:"put-code,omitempty" xml:"put-code,attr,omitempty"`
	CreatedDate      *Date            `json:"created-date,omitempty" xml:"created-date,omitempty"`
	LastModifiedDate *Date            `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Source           *Source          `json:"source,omitempty" xml:"source,omitempty"`
	Title            *Title           `json:"title,omitempty" xml:"title,omitempty"`
	ExternalIDs      *ExternalIDs     `json:"external-ids,omitempty" xml:"external-ids,omitempty"`
	Type             string           `json:"type,omitempty" xml:"type,omitempty"`
	PublicationDate  *PublicationDate `json:"publication-date,omitempty" xml:"publication-date,omitempty"`
	JournalTitle     JournalTitle     `json:"journal-title,omitempty" xml:"journal-title,omitempty"`
	Visibility       string           `json:"visibility,omitempty" xml:"visibility,attr,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
	DisplayIndex     string           `json:"display-index,omitempty" xml:"display-index,attr,omitempty"`
}

type JournalTitle struct {
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

type Work struct {
	PutCode          int64            `json:"put-code,omitempty" xml:"put-code,attr,omitempty"`
	CreatedDate      *Date            `json:"created-date,omitempty" xml:"created-date,omitempty"`
	LastModifiedDate *Date            `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Source           *Source          `json:"source,omitempty" xml:"source,omitempty"`
	Title            *Title           `json:"title,omitempty" xml:"title,omitempty"`
	Subtitle         *Subtitle        `json:"subtitle,omitempty" xml:"subtitle,omitempty"`
	TranslatedTitle  *TranslatedTitle `json:"translated-title,omitempty" xml:"translated-title,omitempty"`
	JournalTitle     JournalTitle     `json:"journal-title,omitempty" xml:"journal-title,omitempty"`
	ShortDescription string           `json:"short-description,omitempty" xml:"short-description,omitempty"`
	Citation         *Citation        `json:"citation,omitempty" xml:"citation,omitempty"`
	Type             string           `json:"type,omitempty" xml:"type,omitempty"`
	PublicationDate  *PublicationDate `json:"publication-date,omitempty" xml:"publication-date,omitempty"`
	ExternalIDs      *ExternalIDs     `json:"external-ids,omitempty" xml:"external-ids,omitempty"`
	URL              *URL             `json:"url,omitempty" xml:"url,omitempty"`
	Contributors     *Contributors    `json:"contributors,omitempty" xml:"contributors,omitempty"`
	LanguageCode     string           `json:"language-code,omitempty" xml:"language-code,omitempty"`
	Country          *Country         `json:"country,omitempty" xml:"country,omitempty"`
	Visibility       string           `json:"visibility,omitempty" xml:"visibility,attr,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type Title struct {
	Title           *TitleValue      `json:"title,omitempty" xml:"title,omitempty"`
	Subtitle        *Subtitle        `json:"subtitle,omitempty" xml:"subtitle,omitempty"`
	TranslatedTitle *TranslatedTitle `json:"translated-title,omitempty" xml:"translated-title,omitempty"`
}

type TitleValue struct {
	Value string `json:"value,omitempty" xml:",chardata"`
}

type Subtitle struct {
	Value string `json:"value,omitempty" xml:",chardata"`
}

type TranslatedTitle struct {
	Value        string `json:"value,omitempty" xml:",chardata"`
	LanguageCode string `json:"language-code,omitempty" xml:"language-code,attr,omitempty"`
}

type Citation struct {
	CitationType  string `json:"citation-type,omitempty" xml:"citation-type,omitempty"`
	CitationValue string `json:"citation-value,omitempty" xml:"citation-value,omitempty"`
}

type PublicationDate struct {
	Year  *Year  `json:"year,omitempty" xml:"year,omitempty"`
	Month *Month `json:"month,omitempty" xml:"month,omitempty"`
	Day   *Day   `json:"day,omitempty" xml:"day,omitempty"`
}

type Year struct {
	Value string `json:"value,omitempty" xml:",chardata"`
}

type Month struct {
	Value string `json:"value,omitempty" xml:",chardata"`
}

type Day struct {
	Value string `json:"value,omitempty" xml:",chardata"`
}

type Contributors struct {
	Contributor []*Contributor `json:"contributor,omitempty" xml:"contributor,omitempty"`
}

type Contributor struct {
	ContributorOrcid      *ContributorOrcid      `json:"contributor-orcid,omitempty" xml:"contributor-orcid,omitempty"`
	CreditName            *CreditName            `json:"credit-name,omitempty" xml:"credit-name,omitempty"`
	ContributorEmail      string                 `json:"contributor-email,omitempty" xml:"contributor-email,omitempty"`
	ContributorAttributes *ContributorAttributes `json:"contributor-attributes,omitempty" xml:"contributor-attributes,omitempty"`
}

type ContributorOrcid struct {
	URI  string `json:"uri,omitempty" xml:"uri,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,omitempty"`
	Host string `json:"host,omitempty" xml:"host,omitempty"`
}

type ContributorAttributes struct {
	ContributorSequence string `json:"contributor-sequence,omitempty" xml:"contributor-sequence,omitempty"`
	ContributorRole     string `json:"contributor-role,omitempty" xml:"contributor-role,omitempty"`
}

type Educations struct {
	LastModifiedDate *Date               `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	EducationSummary []*EducationSummary `json:"education-summary,omitempty" xml:"education-summary,omitempty"`
	AffiliationGroup []*AffiliationGroup `json:"affiliation-group,omitempty" xml:"affiliation-group,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type EducationSummary struct {
	PutCode          int64         `json:"put-code,omitempty" xml:"put-code,attr,omitempty"`
	CreatedDate      *Date         `json:"created-date,omitempty" xml:"created-date,omitempty"`
	LastModifiedDate *Date         `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Source           *Source       `json:"source,omitempty" xml:"source,omitempty"`
	DepartmentName   string        `json:"department-name,omitempty" xml:"department-name,omitempty"`
	RoleTitle        string        `json:"role-title,omitempty" xml:"role-title,omitempty"`
	StartDate        *FuzzyDate    `json:"start-date,omitempty" xml:"start-date,omitempty"`
	EndDate          *FuzzyDate    `json:"end-date,omitempty" xml:"end-date,omitempty"`
	Organization     *Organization `json:"organization,omitempty" xml:"organization,omitempty"`
	URL              *URL          `json:"url,omitempty" xml:"url,omitempty"`
	ExternalIDs      *ExternalIDs  `json:"external-ids,omitempty" xml:"external-ids,omitempty"`
	DisplayIndex     string        `json:"display-index,omitempty" xml:"display-index,attr,omitempty"`
	Visibility       string        `json:"visibility,omitempty" xml:"visibility,attr,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type AffiliationGroup struct {
	LastModifiedDate *Date         `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	ExternalIDs      *ExternalIDs  `json:"external-ids,omitempty" xml:"external-ids,omitempty"`
	Summaries        []interface{} `json:"summaries,omitempty" xml:"summaries,omitempty"`
}

type FuzzyDate struct {
	Year  *Year  `json:"year,omitempty" xml:"year,omitempty"`
	Month *Month `json:"month,omitempty" xml:"month,omitempty"`
	Day   *Day   `json:"day,omitempty" xml:"day,omitempty"`
}

type Organization struct {
	Name                      string                     `json:"name,omitempty" xml:"name,omitempty"`
	Address                   *OrganizationAddress       `json:"address,omitempty" xml:"address,omitempty"`
	DisambiguatedOrganization *DisambiguatedOrganization `json:"disambiguated-organization,omitempty" xml:"disambiguated-organization,omitempty"`
}

type OrganizationAddress struct {
	City    string `json:"city,omitempty" xml:"city,omitempty"`
	Region  string `json:"region,omitempty" xml:"region,omitempty"`
	Country string `json:"country,omitempty" xml:"country,omitempty"`
}

type DisambiguatedOrganization struct {
	DisambiguatedOrganizationIdentifier string `json:"disambiguated-organization-identifier,omitempty" xml:"disambiguated-organization-identifier,omitempty"`
	DisambiguationSource                string `json:"disambiguation-source,omitempty" xml:"disambiguation-source,omitempty"`
}

type Employments struct {
	LastModifiedDate  *Date                `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	EmploymentSummary []*EmploymentSummary `json:"employment-summary,omitempty" xml:"employment-summary,omitempty"`
	AffiliationGroup  []*AffiliationGroup  `json:"affiliation-group,omitempty" xml:"affiliation-group,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type EmploymentSummary struct {
	PutCode          int64         `json:"put-code,omitempty" xml:"put-code,attr,omitempty"`
	CreatedDate      *Date         `json:"created-date,omitempty" xml:"created-date,omitempty"`
	LastModifiedDate *Date         `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Source           *Source       `json:"source,omitempty" xml:"source,omitempty"`
	DepartmentName   string        `json:"department-name,omitempty" xml:"department-name,omitempty"`
	RoleTitle        string        `json:"role-title,omitempty" xml:"role-title,omitempty"`
	StartDate        *FuzzyDate    `json:"start-date,omitempty" xml:"start-date,omitempty"`
	EndDate          *FuzzyDate    `json:"end-date,omitempty" xml:"end-date,omitempty"`
	Organization     *Organization `json:"organization,omitempty" xml:"organization,omitempty"`
	URL              *URL          `json:"url,omitempty" xml:"url,omitempty"`
	ExternalIDs      *ExternalIDs  `json:"external-ids,omitempty" xml:"external-ids,omitempty"`
	DisplayIndex     string        `json:"display-index,omitempty" xml:"display-index,attr,omitempty"`
	Visibility       string        `json:"visibility,omitempty" xml:"visibility,attr,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type Fundings struct {
	LastModifiedDate *Date           `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	FundingGroup     []*FundingGroup `json:"group,omitempty" xml:"group,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type FundingGroup struct {
	LastModifiedDate *Date             `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	ExternalIDs      *ExternalIDs      `json:"external-ids,omitempty" xml:"external-ids,omitempty"`
	FundingSummary   []*FundingSummary `json:"funding-summary,omitempty" xml:"funding-summary,omitempty"`
}

type FundingSummary struct {
	PutCode          int64         `json:"put-code,omitempty" xml:"put-code,attr,omitempty"`
	CreatedDate      *Date         `json:"created-date,omitempty" xml:"created-date,omitempty"`
	LastModifiedDate *Date         `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Source           *Source       `json:"source,omitempty" xml:"source,omitempty"`
	Title            *Title        `json:"title,omitempty" xml:"title,omitempty"`
	Type             string        `json:"type,omitempty" xml:"type,omitempty"`
	StartDate        *FuzzyDate    `json:"start-date,omitempty" xml:"start-date,omitempty"`
	EndDate          *FuzzyDate    `json:"end-date,omitempty" xml:"end-date,omitempty"`
	Organization     *Organization `json:"organization,omitempty" xml:"organization,omitempty"`
	URL              *URL          `json:"url,omitempty" xml:"url,omitempty"`
	ExternalIDs      *ExternalIDs  `json:"external-ids,omitempty" xml:"external-ids,omitempty"`
	DisplayIndex     string        `json:"display-index,omitempty" xml:"display-index,attr,omitempty"`
	Visibility       string        `json:"visibility,omitempty" xml:"visibility,attr,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type PeerReviews struct {
	LastModifiedDate *Date              `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	PeerReviewGroup  []*PeerReviewGroup `json:"group,omitempty" xml:"group,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type PeerReviewGroup struct {
	LastModifiedDate  *Date                `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	ExternalIDs       *ExternalIDs         `json:"external-ids,omitempty" xml:"external-ids,omitempty"`
	PeerReviewSummary []*PeerReviewSummary `json:"peer-review-summary,omitempty" xml:"peer-review-summary,omitempty"`
}

type PeerReviewSummary struct {
	PutCode              int64         `json:"put-code,omitempty" xml:"put-code,attr,omitempty"`
	CreatedDate          *Date         `json:"created-date,omitempty" xml:"created-date,omitempty"`
	LastModifiedDate     *Date         `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Source               *Source       `json:"source,omitempty" xml:"source,omitempty"`
	ReviewGroupID        string        `json:"review-group-id,omitempty" xml:"review-group-id,omitempty"`
	ReviewType           string        `json:"review-type,omitempty" xml:"review-type,omitempty"`
	ReviewCompletionDate *FuzzyDate    `json:"review-completion-date,omitempty" xml:"review-completion-date,omitempty"`
	ReviewURL            *URL          `json:"review-url,omitempty" xml:"review-url,omitempty"`
	Organization         *Organization `json:"convening-organization,omitempty" xml:"convening-organization,omitempty"`
	ExternalIDs          *ExternalIDs  `json:"external-ids,omitempty" xml:"external-ids,omitempty"`
	DisplayIndex         string        `json:"display-index,omitempty" xml:"display-index,attr,omitempty"`
	Visibility           string        `json:"visibility,omitempty" xml:"visibility,attr,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type Distinctions struct {
	LastModifiedDate   *Date                 `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	DistinctionSummary []*DistinctionSummary `json:"distinction-summary,omitempty" xml:"distinction-summary,omitempty"`
	AffiliationGroup   []*AffiliationGroup   `json:"affiliation-group,omitempty" xml:"affiliation-group,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type DistinctionSummary struct {
	PutCode          int64         `json:"put-code,omitempty" xml:"put-code,attr,omitempty"`
	CreatedDate      *Date         `json:"created-date,omitempty" xml:"created-date,omitempty"`
	LastModifiedDate *Date         `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Source           *Source       `json:"source,omitempty" xml:"source,omitempty"`
	DepartmentName   string        `json:"department-name,omitempty" xml:"department-name,omitempty"`
	RoleTitle        string        `json:"role-title,omitempty" xml:"role-title,omitempty"`
	StartDate        *FuzzyDate    `json:"start-date,omitempty" xml:"start-date,omitempty"`
	EndDate          *FuzzyDate    `json:"end-date,omitempty" xml:"end-date,omitempty"`
	Organization     *Organization `json:"organization,omitempty" xml:"organization,omitempty"`
	URL              *URL          `json:"url,omitempty" xml:"url,omitempty"`
	ExternalIDs      *ExternalIDs  `json:"external-ids,omitempty" xml:"external-ids,omitempty"`
	DisplayIndex     string        `json:"display-index,omitempty" xml:"display-index,attr,omitempty"`
	Visibility       string        `json:"visibility,omitempty" xml:"visibility,attr,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type InvitedPositions struct {
	LastModifiedDate       *Date                     `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	InvitedPositionSummary []*InvitedPositionSummary `json:"invited-position-summary,omitempty" xml:"invited-position-summary,omitempty"`
	AffiliationGroup       []*AffiliationGroup       `json:"affiliation-group,omitempty" xml:"affiliation-group,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type InvitedPositionSummary struct {
	PutCode          int64         `json:"put-code,omitempty" xml:"put-code,attr,omitempty"`
	CreatedDate      *Date         `json:"created-date,omitempty" xml:"created-date,omitempty"`
	LastModifiedDate *Date         `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Source           *Source       `json:"source,omitempty" xml:"source,omitempty"`
	DepartmentName   string        `json:"department-name,omitempty" xml:"department-name,omitempty"`
	RoleTitle        string        `json:"role-title,omitempty" xml:"role-title,omitempty"`
	StartDate        *FuzzyDate    `json:"start-date,omitempty" xml:"start-date,omitempty"`
	EndDate          *FuzzyDate    `json:"end-date,omitempty" xml:"end-date,omitempty"`
	Organization     *Organization `json:"organization,omitempty" xml:"organization,omitempty"`
	URL              *URL          `json:"url,omitempty" xml:"url,omitempty"`
	ExternalIDs      *ExternalIDs  `json:"external-ids,omitempty" xml:"external-ids,omitempty"`
	DisplayIndex     string        `json:"display-index,omitempty" xml:"display-index,attr,omitempty"`
	Visibility       string        `json:"visibility,omitempty" xml:"visibility,attr,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type Memberships struct {
	LastModifiedDate  *Date                `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	MembershipSummary []*MembershipSummary `json:"membership-summary,omitempty" xml:"membership-summary,omitempty"`
	AffiliationGroup  []*AffiliationGroup  `json:"affiliation-group,omitempty" xml:"affiliation-group,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type MembershipSummary struct {
	PutCode          int64         `json:"put-code,omitempty" xml:"put-code,attr,omitempty"`
	CreatedDate      *Date         `json:"created-date,omitempty" xml:"created-date,omitempty"`
	LastModifiedDate *Date         `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Source           *Source       `json:"source,omitempty" xml:"source,omitempty"`
	DepartmentName   string        `json:"department-name,omitempty" xml:"department-name,omitempty"`
	RoleTitle        string        `json:"role-title,omitempty" xml:"role-title,omitempty"`
	StartDate        *FuzzyDate    `json:"start-date,omitempty" xml:"start-date,omitempty"`
	EndDate          *FuzzyDate    `json:"end-date,omitempty" xml:"end-date,omitempty"`
	Organization     *Organization `json:"organization,omitempty" xml:"organization,omitempty"`
	URL              *URL          `json:"url,omitempty" xml:"url,omitempty"`
	ExternalIDs      *ExternalIDs  `json:"external-ids,omitempty" xml:"external-ids,omitempty"`
	DisplayIndex     string        `json:"display-index,omitempty" xml:"display-index,attr,omitempty"`
	Visibility       string        `json:"visibility,omitempty" xml:"visibility,attr,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type Qualifications struct {
	LastModifiedDate     *Date                   `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	QualificationSummary []*QualificationSummary `json:"qualification-summary,omitempty" xml:"qualification-summary,omitempty"`
	AffiliationGroup     []*AffiliationGroup     `json:"affiliation-group,omitempty" xml:"affiliation-group,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type QualificationSummary struct {
	PutCode          int64         `json:"put-code,omitempty" xml:"put-code,attr,omitempty"`
	CreatedDate      *Date         `json:"created-date,omitempty" xml:"created-date,omitempty"`
	LastModifiedDate *Date         `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Source           *Source       `json:"source,omitempty" xml:"source,omitempty"`
	DepartmentName   string        `json:"department-name,omitempty" xml:"department-name,omitempty"`
	RoleTitle        string        `json:"role-title,omitempty" xml:"role-title,omitempty"`
	StartDate        *FuzzyDate    `json:"start-date,omitempty" xml:"start-date,omitempty"`
	EndDate          *FuzzyDate    `json:"end-date,omitempty" xml:"end-date,omitempty"`
	Organization     *Organization `json:"organization,omitempty" xml:"organization,omitempty"`
	URL              *URL          `json:"url,omitempty" xml:"url,omitempty"`
	ExternalIDs      *ExternalIDs  `json:"external-ids,omitempty" xml:"external-ids,omitempty"`
	DisplayIndex     string        `json:"display-index,omitempty" xml:"display-index,attr,omitempty"`
	Visibility       string        `json:"visibility,omitempty" xml:"visibility,attr,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type Services struct {
	LastModifiedDate *Date               `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	ServiceSummary   []*ServiceSummary   `json:"service-summary,omitempty" xml:"service-summary,omitempty"`
	AffiliationGroup []*AffiliationGroup `json:"affiliation-group,omitempty" xml:"affiliation-group,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type ServiceSummary struct {
	PutCode          int64         `json:"put-code,omitempty" xml:"put-code,attr,omitempty"`
	CreatedDate      *Date         `json:"created-date,omitempty" xml:"created-date,omitempty"`
	LastModifiedDate *Date         `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Source           *Source       `json:"source,omitempty" xml:"source,omitempty"`
	DepartmentName   string        `json:"department-name,omitempty" xml:"department-name,omitempty"`
	RoleTitle        string        `json:"role-title,omitempty" xml:"role-title,omitempty"`
	StartDate        *FuzzyDate    `json:"start-date,omitempty" xml:"start-date,omitempty"`
	EndDate          *FuzzyDate    `json:"end-date,omitempty" xml:"end-date,omitempty"`
	Organization     *Organization `json:"organization,omitempty" xml:"organization,omitempty"`
	URL              *URL          `json:"url,omitempty" xml:"url,omitempty"`
	ExternalIDs      *ExternalIDs  `json:"external-ids,omitempty" xml:"external-ids,omitempty"`
	DisplayIndex     string        `json:"display-index,omitempty" xml:"display-index,attr,omitempty"`
	Visibility       string        `json:"visibility,omitempty" xml:"visibility,attr,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type ResearchResources struct {
	LastModifiedDate      *Date                    `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	ResearchResourceGroup []*ResearchResourceGroup `json:"group,omitempty" xml:"group,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}

type ResearchResourceGroup struct {
	LastModifiedDate        *Date                      `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	ExternalIDs             *ExternalIDs               `json:"external-ids,omitempty" xml:"external-ids,omitempty"`
	ResearchResourceSummary []*ResearchResourceSummary `json:"research-resource-summary,omitempty" xml:"research-resource-summary,omitempty"`
}

type ResearchResourceSummary struct {
	PutCode          int64        `json:"put-code,omitempty" xml:"put-code,attr,omitempty"`
	CreatedDate      *Date        `json:"created-date,omitempty" xml:"created-date,omitempty"`
	LastModifiedDate *Date        `json:"last-modified-date,omitempty" xml:"last-modified-date,omitempty"`
	Source           *Source      `json:"source,omitempty" xml:"source,omitempty"`
	Title            string       `json:"title,omitempty" xml:"title,omitempty"`
	ExternalIDs      *ExternalIDs `json:"external-ids,omitempty" xml:"external-ids,omitempty"`
	DisplayIndex     string       `json:"display-index,omitempty" xml:"display-index,attr,omitempty"`
	Visibility       string       `json:"visibility,omitempty" xml:"visibility,attr,omitempty"`
	Path Path `json:"path,omitempty" xml:"path,attr,omitempty"`
}
