package bunny

import (
	"context"
	"fmt"
)

// AddOrUpdate adds or updates an Edge Rule for a Pull Zone.
// The GUID field must not be set when creating a pull-zone.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_addedgerule
func (s *EdgeRuleService) AddOrUpdate(ctx context.Context, opts *EdgeRule) error {
	req, err := s.client.newPostRequest(fmt.Sprintf("pullzone/%d/edgerules/addOrUpdate", s.pullZoneID), opts)
	if err != nil {
		return err
	}

	return s.client.sendRequest(ctx, req, nil)
}
