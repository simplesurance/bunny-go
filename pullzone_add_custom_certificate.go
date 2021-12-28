package bunny

import (
	"context"
	"fmt"
)

type pullZoneAddCustomCertificateOptions struct {
	Hostname       string `json:"Hostname"`
	Certificate    []byte `json:"Certificate"`
	CertificateKey []byte `json:"CertificateKey"`
}

// AddCustomCertificate represents the Add Custom Certificate API Endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_addcertificate
func (s *PullZoneService) AddCustomCertificate(ctx context.Context, pullZoneID int64, hostname string, certificates []byte, key []byte) error {
	params := pullZoneAddCustomCertificateOptions{
		Hostname:       hostname,
		Certificate:    certificates,
		CertificateKey: key,
	}

	req, err := s.client.newPostRequest(fmt.Sprintf("/pullzone/%d/addCertificate", pullZoneID), &params)
	if err != nil {
		return err
	}

	return s.client.sendRequest(ctx, req, nil)
}
