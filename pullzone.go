package bunny

// PullZoneService communicates with the /pullzone API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pull-zone
type PullZoneService struct {
	client *Client
}

// EdgeRuleService returns a Service to communicate with the Edge Rule API
// endpoint for the given Pull Zone.
func (s *PullZoneService) EdgeRuleService(pullZoneID int64) *EdgeRuleService {
	return &EdgeRuleService{
		client:     s.client,
		pullZoneID: pullZoneID,
	}
}
