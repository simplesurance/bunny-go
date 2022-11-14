//go:build integrationtest
// +build integrationtest

package bunny_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	bunny "github.com/simplesurance/bunny-go"
)

func TestPullZoneAddRemoveHostname(t *testing.T) {
	clt := newClient(t)

	pzAddopts := bunny.PullZoneAddOptions{
		Name:      randomResourceName("pullzone"),
		OriginURL: "http://bunny.net",
	}

	pz := createPullZone(t, clt, &pzAddopts)

	hostname := "testhostname-" + uuid.New().String() + ".bunny.net"
	err := clt.PullZone.AddCustomHostname(context.Background(), *pz.ID, &bunny.AddCustomHostnameOptions{Hostname: &hostname})
	require.NoError(t, err, "add hostname to pull zone failed")

	getPz, err := clt.PullZone.Get(context.Background(), *pz.ID)
	require.NoError(t, err, "pull zone get failed after adding hostname")
	require.True(t, containsHostname(getPz.Hostnames, hostname), "hostname not returned by get after adding it")

	err = clt.PullZone.RemoveCustomHostname(context.Background(), *pz.ID, &bunny.RemoveCustomHostnameOptions{Hostname: &hostname})
	require.NoError(t, err, "removing hostname from pull zone failed")

	getPz, err = clt.PullZone.Get(context.Background(), *pz.ID)
	require.NoError(t, err, "pull zone get failed after removing hostname")
	require.False(t, containsHostname(getPz.Hostnames, hostname), "pull zone hostnames list is not empty after removing hostname")
}

func containsHostname(hostnames []*bunny.Hostname, hostname string) bool {
	for _, elem := range hostnames {
		if elem.Value != nil && *elem.Value == hostname {
			return true
		}
	}

	return false
}
