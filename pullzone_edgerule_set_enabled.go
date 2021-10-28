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

// Enable enables or disables an Edge Rule of a Pull Zone.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_addedgerule
func (s *EdgeRuleService) Enable(ctx context.Context, edgeRuleGUID string, enabled bool) error {
	msg := edgeRuleEnableOptions{
		ID:    &s.pullZoneID,
		Value: &enabled,
	}

	req, err := s.client.newPostRequest(fmt.Sprintf("pullzone/%d/edgerules/%s/setEdgeRuleEnabled", s.pullZoneID, edgeRuleGUID), &msg)
	if err != nil {
		return err
	}

	return s.client.sendRequest(ctx, req, nil)
}
