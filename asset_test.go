package kernel

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAsset(t *testing.T) {
	require := require.New(t)

	mn := NewMixinNetwork("https://kernel.mixin.dev")
	asset, err := mn.GetAsset("fe6b7788944d328778f98e3e81588215b5a07de4f9a4a7de4db4535b404e65db")
	require.Nil(err)
	require.NotNil(asset)
	require.Equal("c6d0c728-2624-429b-8e0d-d9d19b6592fa", asset.AssetKey)
	require.Equal("fe6b7788944d328778f98e3e81588215b5a07de4f9a4a7de4db4535b404e65db", asset.Chain)
	require.Equal("fe6b7788944d328778f98e3e81588215b5a07de4f9a4a7de4db4535b404e65db", asset.ID)

	asset, err = mn.GetAsset("fe6b7788944d328778f98e3e81588215b5a07de4f9a4a7de4db4535b404e65cb")
	require.Nil(err)
	require.Nil(asset)
}
