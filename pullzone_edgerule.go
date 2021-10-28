package bunny

// EdgeRuleService communicates with the /pullzone/{pullZoneId}/edgerules API endpoint.
//
// Bunny.net API docs: https://docs.bunny.net/reference/pullzonepublic_addedgerule
type EdgeRuleService struct {
	client     *Client
	pullZoneID int64
}
