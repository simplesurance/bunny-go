package bunny

import (
	"context"
	"fmt"
)

// AddOrUpdateEdgeRule adds or updates an Edge Rule of a Pull Zone.
// The GUID field in the EdgeRule struct must not be set when creating a
// pull-zone.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_addedgerule
func (s *PullZoneService) AddOrUpdateEdgeRule(ctx context.Context, pullZoneID int64, opts *EdgeRule) error {
	req, err := s.client.newPostRequest(fmt.Sprintf("pullzone/%d/edgerules/addOrUpdate", pullZoneID), opts)
	if err != nil {
		return err
	}

	return s.client.sendRequest(ctx, req, nil)
}
