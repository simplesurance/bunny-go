package bunny

import (
	"context"
	"fmt"
)

// Constants for the Type field of a DNS Record
const (
	DNSRecordTypeA     int = 0
	DNSRecordTypeAAAA  int = 1
	DNSRecordTypeCNAME int = 2
	DNSRecordTypeTXT   int = 3
	DNSRecordTypeMX    int = 4
	DNSRecordTypeRDR   int = 5 // Bunny.NET Redirect custom record
	DNSRecordTypePZ    int = 7 // Bunny.NET Pull Zone custom record
	DNSRecordTypeSRV   int = 8
	DNSRecordTypeCAA   int = 9
	DNSRecordTypePTR   int = 10
	DNSRecordTypeSCR   int = 11 // Bunny.NET Script custom record
	DNSRecordTypeNS    int = 12
)

// DNSZone represents the response of the the List and Get DNS Zone API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/dnszonepublic_index2 https://docs.bunny.net/reference/dnszonepublic_index
type DNSZone struct {
	ID *int64 `json:"Id,omitempty"`

	Domain                        *string     `json:"Domain,omitempty"`
	Records                       []DNSRecord `json:"Records,omitempty"`
	DateModified                  *string     `json:"DateModified,omitempty"` // should be time
	DateCreated                   *string     `json:"DateCreated,omitempty"`  // should be time
	NameserversDetected           *bool       `json:"NameserversDetected,omitempty"`
	CustomNameserversEnabled      *bool       `json:"CustomNameserversEnabled,omitempty"`
	Nameserver1                   *string     `json:"Nameserver1,omitempty"`
	Nameserver2                   *string     `json:"Nameserver2,omitempty"`
	SoaEmail                      *string     `json:"SoaEmail,omitempty"`
	NameserversNextCheck          *string     `json:"NameserversNextCheck,omitempty"` // should be time
	LoggingEnabled                *bool       `json:"LoggingEnabled,omitempty"`
	LoggingIPAnonymizationEnabled *bool       `json:"LoggingIPAnonymizationEnabled,omitempty"`
	LogAnonymizationType          *int        `json:"LogAnonymizationType,omitempty"`
}

type DNSRecord struct {
	ID                     *int64                  `json:"Id,omitempty"`
	Type                   *int                    `json:"Type,omitempty"`
	Ttl                    *int32                  `json:"Ttl,omitempty"`
	Value                  *string                 `json:"Value,omitempty"`
	Name                   *string                 `json:"Name,omitempty"`
	Weight                 *int32                  `json:"Weight,omitempty"`
	Priority               *int32                  `json:"Priority,omitempty"`
	Port                   *int32                  `json:"Port,omitempty"`
	Flags                  *int                    `json:"Flags,omitempty"`
	Tag                    *string                 `json:"Tag,omitempty"`
	Accelerated            *bool                   `json:"Accelerated,omitempty"`
	AcceleratedPullZoneId  *int64                  `json:"AcceleratedPullZoneId,omitempty"`
	LinkName               *string                 `json:"LinkName,omitempty"`
	IPGeoLocationInfo      *IPGeoLocationInfo      `json:"IPGeoLocationInfo,omitempty"`
	MonitorStatus          *int                    `json:"MonitorStatus,omitempty"`
	MonitorType            *int                    `json:"MonitorType,omitempty"`
	GeolocationLatitude    *float64                `json:"GeolocationLatitude,omitempty"`
	GeolocationLongitude   *float64                `json:"GeolocationLongitude,omitempty"`
	EnvironmentalVariables []EnvironmentalVariable `json:"EnvironmentalVariables,omitempty"`
	LatencyZone            *string                 `json:"LatencyZone,omitempty"`
	SmartRoutingType       *int                    `json:"SmartRoutingType,omitempty"`
	Disabled               *bool                   `json:"Disabled,omitempty"`
}

type IPGeoLocationInfo struct {
	CountryCode      *string `json:"CountryCode,omitempty"`
	Country          *string `json:"Country,omitempty"`
	ASN              *int64  `json:"ASN,omitempty"`
	OrganizationName *string `json:"OrganizationName,omitempty"`
	City             *string `json:"City,omitempty"`
}

type EnvironmentalVariable struct {
	Name  *string `json:"Name,omitempty"`
	Value *string `json:"Value,omitempty"`
}

// Get retrieves the DNS Zone with the given id.
//
// Bunny.net API docs: https://docs.bunny.net/reference/dnszonepublic_index2
func (s *DNSZoneService) Get(ctx context.Context, id int64) (*DNSZone, error) {
	path := fmt.Sprintf("dnszone/%d", id)
	return resourceGet[DNSZone](ctx, s.client, path)
}
