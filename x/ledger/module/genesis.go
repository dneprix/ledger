package ledger

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dneprix/ledger/x/ledger/keeper"
	"github.com/dneprix/ledger/x/ledger/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the account
	for _, elem := range genState.AccountList {
		k.SetAccount(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AccountList = k.GetAllAccount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
