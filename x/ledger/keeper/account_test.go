package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "github.com/dneprix/ledger/testutil/keeper"
	"github.com/dneprix/ledger/testutil/nullify"
	"github.com/dneprix/ledger/x/ledger/keeper"
	"github.com/dneprix/ledger/x/ledger/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNAccount(keeper keeper.Keeper, ctx context.Context, n int) []types.Account {
	items := make([]types.Account, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetAccount(ctx, items[i])
	}
	return items
}

func TestAccountGet(t *testing.T) {
	keeper, ctx := keepertest.LedgerKeeper(t)
	items := createNAccount(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetAccount(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestAccountRemove(t *testing.T) {
	keeper, ctx := keepertest.LedgerKeeper(t)
	items := createNAccount(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAccount(ctx,
			item.Address,
		)
		_, found := keeper.GetAccount(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestAccountGetAll(t *testing.T) {
	keeper, ctx := keepertest.LedgerKeeper(t)
	items := createNAccount(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAccount(ctx)),
	)
}
