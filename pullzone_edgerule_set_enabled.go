package bunny

import (
	"context"
	"fmt"
)

// edgeRuleEnableOptions represents the message that is sent to Add/Update Edge Rule endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_addedgerule
type edgeRuleEnableOptions struct {
	// ID is the pull zone id
	ID    *int64 `json:"Id,omitempty"`
	Value *bool  `json:"Value,omitempty"`
}

// SetEdgeRuleEnabled enables or disables an Edge Rule of a Pull Zone.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_addedgerule
func (s *PullZoneService) SetEdgeRuleEnabled(ctx context.Context, pullZoneID int64, edgeRuleGUID string, enabled bool) error {
	msg := edgeRuleEnableOptions{
		ID:    &pullZoneID,
		Value: &enabled,
	}

	req, err := s.client.newPostRequest(fmt.Sprintf("pullzone/%d/edgerules/%s/setEdgeRuleEnabled", pullZoneID, edgeRuleGUID), &msg)
	if err != nil {
		return err
	}

	return s.client.sendRequest(ctx, req, nil)
}
