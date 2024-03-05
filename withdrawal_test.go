package kernel

import (
	"encoding/hex"
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

	tx, err = mn.GetWithdrawalClaim("2e04fba82c1f874602a7baf1e5d0d56b6dfdc28b8af1c43b368ba1f5654aada6")
	require.Nil(err)
	require.NotNil(tx)
	require.Equal("b6fbe6f7d34906725c7a54e89c95c07717dee7cab033b3e6d69db485bbcb3089", tx.Hash)
	extra, _ := hex.DecodeString(tx.Extra)
	require.Equal("658aace0c5e2291ca91f2dd734dc533db68c447f78f2428682349396084dc4cc", string(extra[64:]))
}
