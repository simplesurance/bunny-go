//go:build integrationtest
// +build integrationtest

package bunny_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"

	bunny "github.com/simplesurance/bunny-go"
)

func TestVideoLibraryCRUD(t *testing.T) {
	clt := newClient(t)

	vlName := randomResourceName("videolibrary")
	vlRegion := "NY"
	// vlOrigin := "http://bunny.net"
	vlAddopts := bunny.VideoLibraryAddOptions{
		Name: &vlName,
		ReplicationRegions: []string{vlRegion},
	}

	listVlBefore, err := clt.VideoLibrary.List(context.Background(), nil)
	require.NoError(t, err, "video library list failed before add")

	vl := createVideoLibrary(t, clt, &vlAddopts)

	// get the newly created video library
	getVl, err := clt.VideoLibrary.Get(context.Background(), *vl.ID, &bunny.VideoLibraryGetOpts{false})
	require.NoError(t, err, "video library get failed after adding")
	assert.NotNil(t, getVl.ID)
	assert.Nil(t, getVl.APIAccessKey)
	assert.Equal(
		t,
		vlRegion,
		getVl.ReplicationRegions[0],
		"video library replication region should be set correctly",
	)

	// update the video library
	newName := vlName + "-updated"
	setTrue := true
	setFalse := false
	updateOpts := bunny.VideoLibraryUpdateOptions{
		Name: &newName,
		PlayerTokenAuthenticationEnabled: &setTrue,
		AllowDirectPlay: &setFalse,
	}
	updateErr := clt.VideoLibrary.Update(context.Background(), *vl.ID, &updateOpts)
	assert.Nil(t, updateErr)

	// get the updated video library and validate updated properties
	getUpdatedVl, err := clt.VideoLibrary.Get(context.Background(), *vl.ID, &bunny.VideoLibraryGetOpts{true})
	assert.NotNil(t, getUpdatedVl.ID)
	assert.NotNil(t, getUpdatedVl.APIAccessKey)
	assert.Equal(
		t,
		newName,
		*getUpdatedVl.Name,
		"video library Name should be updated correctly",
	)
	assert.Equal(
		t,
		true,
		*getUpdatedVl.PlayerTokenAuthenticationEnabled,
		"video library PlayerTokenAuthenticationEnabled should be updated correctly",
	)
	assert.Equal(
		t,
		false,
		*getUpdatedVl.AllowDirectPlay,
		"video library AllowDirectPlay should be updated correctly",
	)

	// check the total number of video libraries is the expected amount
	listVlAfter, err := clt.VideoLibrary.List(context.Background(), nil)
	require.NoError(t, err, "video library list failed after add")
	assert.Equal(
		t,
		*listVlBefore.TotalItems + 1,
		*listVlAfter.TotalItems,
		"video libraries total items should increase by exactly 1",
	)
}
