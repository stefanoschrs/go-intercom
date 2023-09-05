package intercom

import "fmt"

// ContactService handles interactions with the API through a ContactRepository.
type ContactService struct {
	Repository ContactRepository
}

// ContactList holds a list of Contacts and paging information
type ContactList struct {
	Pages       PageParams
	Contacts    []Contact
	ScrollParam string `json:"scroll_param,omitempty"`
}

// Contact represents a Contact within Intercom.
// Not all of the fields are writeable to the API, non-writeable fields are
// stripped out from the request. Please see the API documentation for details.
type Contact struct {
	ID                     string                 `json:"id,omitempty"`
	Email                  string                 `json:"email,omitempty"`
	Phone                  string                 `json:"phone,omitempty"`
	UserID                 string                 `json:"user_id,omitempty"`
	Name                   string                 `json:"name,omitempty"`
	Avatar                 *UserAvatar            `json:"avatar,omitempty"`
	LocationData           *LocationData          `json:"location_data,omitempty"`
	LastRequestAt          int64                  `json:"last_request_at,omitempty"`
	CreatedAt              int64                  `json:"created_at,omitempty"`
	UpdatedAt              int64                  `json:"updated_at,omitempty"`
	SessionCount           int64                  `json:"session_count,omitempty"`
	LastSeenIP             string                 `json:"last_seen_ip,omitempty"`
	SocialProfiles         *SocialProfileList     `json:"social_profiles,omitempty"`
	UnsubscribedFromEmails *bool                  `json:"unsubscribed_from_emails,omitempty"`
	UserAgentData          string                 `json:"user_agent_data,omitempty"`
	Tags                   *TagList               `json:"tags,omitempty"`
	Segments               *SegmentList           `json:"segments,omitempty"`
	Companies              *CompanyList           `json:"companies,omitempty"`
	CustomAttributes       map[string]interface{} `json:"custom_attributes,omitempty"`
	UpdateLastRequestAt    *bool                  `json:"update_last_request_at,omitempty"`
	NewSession             *bool                  `json:"new_session,omitempty"`
}

type contactListParams struct {
	PageParams
	SegmentID string `url:"segment_id,omitempty"`
	TagID     string `url:"tag_id,omitempty"`
	Email     string `url:"email,omitempty"`
}

type SearchContact struct {
	Type           string      `json:"type"`
	Id             string      `json:"id"`
	WorkspaceId    string      `json:"workspace_id"`
	ExternalId     string      `json:"external_id"`
	Role           string      `json:"role"`
	Email          string      `json:"email"`
	Phone          string      `json:"phone"`
	Name           string      `json:"name"`
	Avatar         *UserAvatar `json:"avatar"`
	OwnerId        any         `json:"owner_id"`
	SocialProfiles struct {
		Type string          `json:"type"`
		Data []SocialProfile `json:"data"`
	} `json:"social_profiles"`
	HasHardBounced         bool   `json:"has_hard_bounced"`
	MarkedEmailAsSpam      bool   `json:"marked_email_as_spam"`
	UnsubscribedFromEmails bool   `json:"unsubscribed_from_emails"`
	CreatedAt              int    `json:"created_at"`
	UpdatedAt              int    `json:"updated_at"`
	SignedUpAt             int    `json:"signed_up_at"`
	LastSeenAt             int    `json:"last_seen_at"`
	LastRepliedAt          any    `json:"last_replied_at"`
	LastContactedAt        int    `json:"last_contacted_at"`
	LastEmailOpenedAt      any    `json:"last_email_opened_at"`
	LastEmailClickedAt     any    `json:"last_email_clicked_at"`
	LanguageOverride       any    `json:"language_override"`
	Browser                string `json:"browser"`
	BrowserVersion         string `json:"browser_version"`
	BrowserLanguage        string `json:"browser_language"`
	Os                     string `json:"os"`
	Location               struct {
		Type          string `json:"type"`
		Country       string `json:"country"`
		Region        string `json:"region"`
		City          string `json:"city"`
		CountryCode   string `json:"country_code"`
		ContinentCode string `json:"continent_code"`
	} `json:"location"`
	AndroidAppName    any                    `json:"android_app_name"`
	AndroidAppVersion any                    `json:"android_app_version"`
	AndroidDevice     any                    `json:"android_device"`
	AndroidOsVersion  any                    `json:"android_os_version"`
	AndroidSdkVersion any                    `json:"android_sdk_version"`
	AndroidLastSeenAt any                    `json:"android_last_seen_at"`
	IosAppName        any                    `json:"ios_app_name"`
	IosAppVersion     any                    `json:"ios_app_version"`
	IosDevice         any                    `json:"ios_device"`
	IosOsVersion      any                    `json:"ios_os_version"`
	IosSdkVersion     any                    `json:"ios_sdk_version"`
	IosLastSeenAt     any                    `json:"ios_last_seen_at"`
	CustomAttributes  map[string]interface{} `json:"custom_attributes"`
	Tags              struct {
		Type       string `json:"type"`
		Data       []Tag  `json:"data"`
		Url        string `json:"url"`
		TotalCount int    `json:"total_count"`
		HasMore    bool   `json:"has_more"`
	} `json:"tags"`
	Notes struct {
		Type       string `json:"type"`
		Data       []any  `json:"data"`
		Url        string `json:"url"`
		TotalCount int    `json:"total_count"`
		HasMore    bool   `json:"has_more"`
	} `json:"notes"`
	Companies struct {
		Type       string    `json:"type"`
		Data       []Company `json:"data"`
		Url        string    `json:"url"`
		TotalCount int       `json:"total_count"`
		HasMore    bool      `json:"has_more"`
	} `json:"companies"`
	OptedOutSubscriptionTypes struct {
		Type       string `json:"type"`
		Data       []any  `json:"data"`
		Url        string `json:"url"`
		TotalCount int    `json:"total_count"`
		HasMore    bool   `json:"has_more"`
	} `json:"opted_out_subscription_types"`
	OptedInSubscriptionTypes struct {
		Type       string `json:"type"`
		Data       []any  `json:"data"`
		Url        string `json:"url"`
		TotalCount int    `json:"total_count"`
		HasMore    bool   `json:"has_more"`
	} `json:"opted_in_subscription_types"`
	UtmCampaign         any    `json:"utm_campaign"`
	UtmContent          any    `json:"utm_content"`
	UtmMedium           any    `json:"utm_medium"`
	UtmSource           any    `json:"utm_source"`
	UtmTerm             any    `json:"utm_term"`
	Referrer            string `json:"referrer"`
	SmsConsent          bool   `json:"sms_consent"`
	UnsubscribedFromSms bool   `json:"unsubscribed_from_sms"`
}

type ContactSearchParams struct {
	Query map[string]string `json:"query,omitempty"`
}

type contactSearchResult struct {
	Type       string `json:"type"`
	TotalCount int    `json:"total_count"`
	Pages      struct {
		Type       string `json:"type"`
		Page       int    `json:"page"`
		PerPage    int    `json:"per_page"`
		TotalPages int    `json:"total_pages"`
	} `json:"pages"`
	Data []SearchContact `json:"data"`
}

// Search looks up a Contact by their Intercom ID.
func (c *ContactService) Search(params ContactSearchParams) (contactSearchResult, error) {
	return c.Repository.search(params)
}

// FindByID looks up a Contact by their Intercom ID.
func (c *ContactService) FindByID(id string) (Contact, error) {
	return c.findWithIdentifiers(UserIdentifiers{ID: id})
}

// FindByUserID looks up a Contact by their UserID (automatically generated server side).
func (c *ContactService) FindByUserID(userID string) (Contact, error) {
	return c.findWithIdentifiers(UserIdentifiers{UserID: userID})
}

func (c *ContactService) findWithIdentifiers(identifiers UserIdentifiers) (Contact, error) {
	return c.Repository.find(identifiers)
}

// List all Contacts for App.
func (c *ContactService) List(params PageParams) (ContactList, error) {
	return c.Repository.list(contactListParams{PageParams: params})
}

// List all Contacts for App via Scroll API
func (c *ContactService) Scroll(scrollParam string) (ContactList, error) {
	return c.Repository.scroll(scrollParam)
}

// ListByEmail looks up a list of Contacts by their Email.
func (c *ContactService) ListByEmail(email string, params PageParams) (ContactList, error) {
	return c.Repository.list(contactListParams{PageParams: params, Email: email})
}

// List Contacts by Segment.
func (c *ContactService) ListBySegment(segmentID string, params PageParams) (ContactList, error) {
	return c.Repository.list(contactListParams{PageParams: params, SegmentID: segmentID})
}

// List Contacts By Tag.
func (c *ContactService) ListByTag(tagID string, params PageParams) (ContactList, error) {
	return c.Repository.list(contactListParams{PageParams: params, TagID: tagID})
}

// Create Contact
func (c *ContactService) Create(contact *Contact) (Contact, error) {
	return c.Repository.create(contact)
}

// Update Contact
func (c *ContactService) Update(contact *Contact) (Contact, error) {
	return c.Repository.update(contact)
}

// Convert Contact to User
func (c *ContactService) Convert(contact *Contact, user *User) (User, error) {
	return c.Repository.convert(contact, user)
}

// Delete Contact
func (c *ContactService) Delete(contact *Contact) (Contact, error) {
	return c.Repository.delete(contact.ID)
}

// MessageAddress gets the address for a Contact in order to message them
func (c Contact) MessageAddress() MessageAddress {
	return MessageAddress{
		Type:   "contact",
		ID:     c.ID,
		Email:  c.Email,
		UserID: c.UserID,
	}
}

func (c Contact) String() string {
	return fmt.Sprintf("[intercom] contact { id: %s name: %s, user_id: %s, email: %s }", c.ID, c.Name, c.UserID, c.Email)
}
