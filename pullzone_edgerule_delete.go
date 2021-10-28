package bunny

import (
	"context"
	"fmt"
)

// Delete removes an Edge Rule of a Pull Zone.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_deleteedgerule
func (s *EdgeRuleService) Delete(ctx context.Context, edgeRuleGUID string) error {
	req, err := s.client.newDeleteRequest(fmt.Sprintf("pullzone/%d/edgerules/%s", s.pullZoneID, edgeRuleGUID), nil)
	if err != nil {
		return err
	}

	return s.client.sendRequest(ctx, req, nil)
}
