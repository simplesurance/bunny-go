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

// storagezoneNamePrefix is the prefix for all storage zones created by the integrationtests.
const storagezoneNamePrefix = "bunny-go-test-storage-"

// videoLibraryNamePrefix is the prefix for all video libraries created by the integrationtests.
const videoLibraryNamePrefix = "bunny-go-test-videolibrary-"

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

func randomStorageZoneName() string {
	return storagezoneNamePrefix + uuid.New().String()
}

func randomVideoLibraryName()  string {
	return videoLibraryNamePrefix + uuid.New().String()
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
			t.Errorf("could not delete pull zone (id: %d, name: %q) on test cleanup: %s", *pz.ID, *pz.Name, err)
			return

		}
		t.Logf("cleanup: deleted pull zone: %q, id: %d", *pz.Name, *pz.ID)
	})

	return pz
}


// createStorageZone creates a Storage Zone via the bunny client and registers a
// testing cleanup function to remove it when the test terminates.
// If creating the Storage Zone fails, t.Fatal is called.
func createStorageZone(t *testing.T, clt *bunny.Client, opts *bunny.StorageZoneAddOptions) *bunny.StorageZone {
	t.Helper()

	sz, err := clt.StorageZone.Add(context.Background(), opts)
	require.NoError(t, err, "creating storage zone failed")
	require.NotNil(t, sz.ID, "add returned storage zone with nil id")
	require.NotNil(t, sz.Name, "add returned storage zone with nil name")

	t.Logf("created storage zone: %q, id: %d", *sz.Name, *sz.ID)

	t.Cleanup(func() {
		err := clt.StorageZone.Delete(context.Background(), *sz.ID)
		if err != nil {
			t.Errorf("could not delete storage zone (id: %d, name: %q) on test cleanup: %s", *sz.ID, *sz.Name, err)
			return

		}
		t.Logf("cleanup: deleted storage zone: %q, id: %d", *sz.Name, *sz.ID)
	})

	return sz
}


// createVideoLibrary creates a Video Library via the bunny client and registers a
// testing cleanup function to remove it when the test terminates.
// If creating the Video Library fails, t.Fatal is called.
func createVideoLibrary(t *testing.T, clt *bunny.Client, opts *bunny.VideoLibraryAddOptions) *bunny.VideoLibrary {
	t.Helper()

	vl, err := clt.VideoLibrary.Add(context.Background(), opts)
	require.NoError(t, err, "creating video library failed")
	require.NotNil(t, vl.ID, "add returned video library with nil id")
	require.NotNil(t, vl.Name, "add returned video library with nil name")

	t.Logf("created video library: %q, id: %d", *vl.Name, *vl.ID)

	t.Cleanup(func() {
		err := clt.VideoLibrary.Delete(context.Background(), *vl.ID)
		if err != nil {
			t.Errorf("could not delete video library (id: %d, name: %q) on test cleanup: %s", *vl.ID, *vl.Name, err)
			return

		}
		t.Logf("cleanup: deleted video library: %q, id: %d", *vl.Name, *vl.ID)
	})

	return vl
}
