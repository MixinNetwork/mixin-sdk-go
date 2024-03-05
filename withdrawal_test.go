package kernel

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWithdrawal(t *testing.T) {
	require := require.New(t)

	mn := NewMixinNetwork("https://kernel.mixin.dev")
	tx, err := mn.GetWithdrawalClaim("04f9195917b43461135d54c3cf8e588cbc22299680ec90b80be15f0bc02a026d")
	require.Nil(err)
	require.NotNil(tx)
	require.Equal("a99c2e0e2b1da4d648755ef19bd95139acbbe6564cfb06dec7cd34931ca72cdc", tx.Asset)
	require.Len(tx.Inputs, 1)
	require.Len(tx.Outputs, 1)
	require.Equal("876c0df9fdcac099c98b7b848dc92e53cff23ea1f84b7cdf04c7a475e661efc0", tx.Snapshot)
	require.Equal(uint8(5), tx.Version)
}
