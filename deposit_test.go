package kernel

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeposit(t *testing.T) {
	require := require.New(t)

	mn := NewMixinNetwork("https://kernel.mixin.dev")
	tx, err := mn.GetDepositTransaction("46dbd75ed416c82652554a2fd257df3adb5d8c68726db6631bf1300e7aa36f41", "64c308a64db9f15935c0396bc575328b3d988f467aa5abe4cad27afab77d8b76", 1)
	require.Nil(err)
	require.NotNil(tx)
	require.Equal("550b4c4a4dd13c1eefed7595f631c27e5d53b49f3bb1f6b662843e275d234218", tx.Hash)
}
