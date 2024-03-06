package kernel

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInfo(t *testing.T) {
	require := require.New(t)

	mn := NewMixinNetwork("https://kernel.mixin.dev")
	info, err := mn.GetInfo()
	require.Nil(err)
	require.NotNil(info)
	log.Println(info.Mint.Pledge)
	log.Println(info.Mint.Pool)
	require.True(info.Mint.Batch > 1000)
	require.True(len(info.Graph.Consensus) > 10)
}
