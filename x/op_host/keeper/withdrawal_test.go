package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ProvenWithdrawal(t *testing.T) {
	ctx, input := createDefaultTestInput(t)

	input.OPHostKeeper.RecordProvenWithdrawal(ctx, 1, [32]byte{1, 2, 3})
	input.OPHostKeeper.RecordProvenWithdrawal(ctx, 1, [32]byte{4, 5, 6})
	input.OPHostKeeper.RecordProvenWithdrawal(ctx, 2, [32]byte{7, 8, 9})

	require.True(t, input.OPHostKeeper.HasProvenWithdrawal(ctx, 1, [32]byte{1, 2, 3}))
	require.True(t, input.OPHostKeeper.HasProvenWithdrawal(ctx, 1, [32]byte{4, 5, 6}))

	input.OPHostKeeper.IterateProvenWithdrawals(ctx, 1, func(bridgeId uint64, withdrawalHash [32]byte) bool {
		require.Equal(t, uint64(1), bridgeId)
		if withdrawalHash != [32]byte{1, 2, 3} {
			require.Equal(t, [32]byte{4, 5, 6}, withdrawalHash)
		}

		return false
	})
}
