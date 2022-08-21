package bunny

import "context"

// DNSZoneAddOptions are the request parameters for the Get DNS Zone API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/dnszonepublic_add
type DNSZoneAddOptions struct {
	ID int64 `json:"Id,omitempty"`

	Domain                        string      `json:"Domain,omitempty"`
	Records                       []DNSRecord `json:"Records,omitempty"`
	DateModified                  string      `json:"DateModified,omitempty"` // should be time
	DateCreated                   string      `json:"DateCreated,omitempty"`  // should be time
	NameserversDetected           bool        `json:"NameserversDetected,omitempty"`
	CustomNameserversEnabled      bool        `json:"CustomNameserversEnabled,omitempty"`
	Nameserver1                   string      `json:"Nameserver1,omitempty"`
	Nameserver2                   string      `json:"Nameserver2,omitempty"`
	SoaEmail                      string      `json:"SoaEmail,omitempty"`
	NameserversNextCheck          string      `json:"NameserversNextCheck,omitempty"` // should be time
	LoggingEnabled                bool        `json:"LoggingEnabled,omitempty"`
	LoggingIPAnonymizationEnabled bool        `json:"LoggingIPAnonymizationEnabled,omitempty"`
	LogAnonymizationType          int         `json:"LogAnonymizationType,omitempty"`
}

// Add creates a new DNS Zone.
// opts and the non-optional parameters in the struct must be specified for a successful request.
// On success the created DNSZone is returned.
//
// Bunny.net API docs: https://docs.bunny.net/reference/dnszonepublic_add
func (s *DNSZoneService) Add(ctx context.Context, opts *DNSZoneAddOptions) (*DNSZone, error) {
	return resourcePostWithResponse[DNSZone](
		ctx,
		s.client,
		"/dnszone",
		opts,
	)
}
