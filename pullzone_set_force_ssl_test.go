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

func TestSetForceSSL(t *testing.T) {
	ctx := context.Background()
	clt := newClient(t)

	pzAddopts := bunny.PullZoneAddOptions{
		Name:      randomPullZoneName(),
		OriginURL: "http://bunny.net",
	}

	pz := createPullZone(t, clt, &pzAddopts)

	hostname := "testhostname-" + uuid.New().String() + ".bunnytftest.de"
	err := clt.PullZone.AddCustomHostname(ctx, *pz.ID, &bunny.AddCustomHostnameOptions{Hostname: &hostname})
	require.NoError(t, err, "add hostname to pull zone failed")

	trueVal := true
	err = clt.PullZone.SetForceSSL(ctx, *pz.ID, &bunny.SetForceSSLOptions{
		Hostname: &hostname,
		ForceSSL: &trueVal,
	})
	require.NoError(t, err, "enabling force ssl failed")

	pz, err = clt.PullZone.Get(ctx, *pz.ID)
	require.NoError(t, err, "retrieving pull zone failed")
	assertHostnameForceSSLValue(t, pz.Hostnames, hostname, true)

	falseVal := false
	err = clt.PullZone.SetForceSSL(ctx, *pz.ID, &bunny.SetForceSSLOptions{
		Hostname: &hostname,
		ForceSSL: &falseVal,
	})
	require.NoError(t, err, "enabling force ssl failed")

	pz, err = clt.PullZone.Get(ctx, *pz.ID)
	require.NoError(t, err, "retrieving pull zone failed")
	assertHostnameForceSSLValue(t, pz.Hostnames, hostname, false)

}

func assertHostnameForceSSLValue(t *testing.T, hostnames []*bunny.Hostname, hostname string, expectedForceSSLVal bool) {
	t.Helper()

	for _, elem := range hostnames {
		if elem.Value == nil {
			t.Errorf("hostname entry has nil Value field")
			continue
		}

		if *elem.Value == hostname {
			if elem.ForceSSL == nil {
				t.Errorf("hostname entry has nil ForceSSL field")
				return
			}
			if *elem.ForceSSL != expectedForceSSLVal {
				t.Errorf("expected %v ForceSSL value, got %v, for hostname %q", expectedForceSSLVal, *elem.ForceSSL, hostname)
				return
			}

			return
		}
	}

	t.Errorf("hostname %q not found in hostnames slices", hostname)
}
