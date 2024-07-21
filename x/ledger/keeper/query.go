package keeper

import (
	"github.com/dneprix/ledger/x/ledger/types"
)

var _ types.QueryServer = Keeper{}
