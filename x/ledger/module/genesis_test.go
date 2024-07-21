package ledger_test

import (
	"testing"

	keepertest "github.com/dneprix/ledger/testutil/keeper"
	"github.com/dneprix/ledger/testutil/nullify"
	ledger "github.com/dneprix/ledger/x/ledger/module"
	"github.com/dneprix/ledger/x/ledger/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		AccountList: []types.Account{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LedgerKeeper(t)
	ledger.InitGenesis(ctx, k, genesisState)
	got := ledger.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AccountList, got.AccountList)
	// this line is used by starport scaffolding # genesis/test/assert
}
