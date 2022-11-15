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

func TestStorageZoneCRUD(t *testing.T) {
	clt := newClient(t)

	szName := randomResourceName("storagezone")
	szOrigin := "http://bunny.net"
	szRegion := "NY"
	szAddopts := bunny.StorageZoneAddOptions{
		Name: &szName,
		OriginURL: &szOrigin,
		Region: &szRegion,
		ReplicationRegions: []string{"DE"},
	}

	listSzBefore, err := clt.StorageZone.List(context.Background(), nil)
	require.NoError(t, err, "storage zone list failed before add")

	sz := createStorageZone(t, clt, &szAddopts)

	// get the newly created storage zone
	getSz, err := clt.StorageZone.Get(context.Background(), *sz.ID)
	require.NoError(t, err, "storage zone get failed after adding")
	assert.NotNil(t, getSz.ID)
	assert.Equal(
		t,
		getSz.ReplicationRegions[0],
		"DE",
		"storage zone replication region should be set correctly",
	)

	// update the storage zone
	szUpdateOrigin := szOrigin + "/updated"
	szUpdateRewrite404To200 := true
	updateOpts := bunny.StorageZoneUpdateOptions{
		OriginURL: &szUpdateOrigin,
		Rewrite404To200: &szUpdateRewrite404To200,
		ReplicationRegions: []string{"LA"},
	}
	updateErr := clt.StorageZone.Update(context.Background(), *sz.ID, &updateOpts)
	assert.Nil(t, updateErr)

	// get the updated storage zone and validate updated properties
	getUpdatedSz, err := clt.StorageZone.Get(context.Background(), *sz.ID)
	assert.NotNil(t, getUpdatedSz.ID)
	assert.Equal(
		t,
		"LA",
		getUpdatedSz.ReplicationRegions[len(getUpdatedSz.ReplicationRegions) - 1],
		"storage zone replication region should be updated correctly",
	)

	// check the total number of storage zones is the expected amount
	listSzAfter, err := clt.StorageZone.List(context.Background(), nil)
	require.NoError(t, err, "storage zone list failed after add")
	assert.Equal(
		t,
		*listSzBefore.TotalItems + 1,
		*listSzAfter.TotalItems,
		"storage zones total items should increase by exactly 1",
	)
}
