package kernel

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSnapshot(t *testing.T) {
	require := require.New(t)

	mn := NewMixinNetwork("https://kernel.mixin.dev")
	snapshots, err := mn.ListSnapshotsSince(0, 3, 0, 1)
	require.Nil(err)
	require.Len(snapshots, 3)

	snapshot, err := mn.GetSnapshot("7860c8b2d9734f8facf51a7a1cfa525f9c4f0a3cb452a40616be5064bdfcb49f")
	require.Nil(err)
	require.Equal(2, snapshot.Version)
	require.Len(snapshot.Transactions, 1)
}
