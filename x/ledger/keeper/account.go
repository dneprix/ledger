package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/dneprix/ledger/x/ledger/types"
)

// SetAccount set a specific account in the store from its index
func (k Keeper) SetAccount(ctx context.Context, account types.Account) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AccountKeyPrefix))
	b := k.cdc.MustMarshal(&account)
	store.Set(types.AccountKey(
		account.Address,
	), b)
}

// GetAccount returns a account from its index
func (k Keeper) GetAccount(
	ctx context.Context,
	address string,

) (val types.Account, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AccountKeyPrefix))

	b := store.Get(types.AccountKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAccount removes a account from the store
func (k Keeper) RemoveAccount(
	ctx context.Context,
	address string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AccountKeyPrefix))
	store.Delete(types.AccountKey(
		address,
	))
}

// GetAllAccount returns all account
func (k Keeper) GetAllAccount(ctx context.Context) (list []types.Account) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AccountKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Account
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
