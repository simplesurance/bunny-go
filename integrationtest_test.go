//go:build integrationtest
// +build integrationtest

package bunny_test

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	bunny "github.com/simplesurance/bunny-go"
	"github.com/stretchr/testify/require"
)

const envVarApiKeyName = "BUNNY_API_KEY"

// pullzoneNamePrefix is the prefix for all pullzones created by the integrationtests.
const pullzoneNamePrefix = "bunny-go-test-"

func newClient(t *testing.T) *bunny.Client {
	t.Helper()

	apiKey := os.Getenv(envVarApiKeyName)
	if apiKey == "" {
		t.Fatalf("the environment variable %q is unset or empty, it must be set to a valid API key that is used for running integration tests",
			envVarApiKeyName)
	}

	return bunny.NewClient(apiKey, bunny.WithHTTPRequestLogger(t.Logf))
}

func randomPullZoneName() string {
	return pullzoneNamePrefix + uuid.New().String()
}

// createPullZone creates a Pull Zone via the bunny client and registers a
// testing cleanup function to remove it when the test terminates.
// If creating the Pull Zone fails, t.Fatal is called.
func createPullZone(t *testing.T, clt *bunny.Client, opts *bunny.PullZoneAddOptions) *bunny.PullZone {
	t.Helper()

	pz, err := clt.PullZone.Add(context.Background(), opts)
	require.NoError(t, err, "creating pull zone failed")
	require.NotNil(t, pz.ID, "add returned pull zone with nil id")
	require.NotNil(t, pz.Name, "add returned pull zone with nil name")

	t.Logf("created pull zone: %q, id: %d", *pz.Name, *pz.ID)

	t.Cleanup(func() {
		err := clt.PullZone.Delete(context.Background(), *pz.ID)
		if err != nil {
			t.Errorf("coult not delete pull zone (id: %d, name: %q) on test cleanup: %s", *pz.ID, *pz.Name, err)
			return

		}
		t.Logf("cleanup: deleted pull zone: %q, id: %d", *pz.Name, *pz.ID)
	})

	return pz
}
