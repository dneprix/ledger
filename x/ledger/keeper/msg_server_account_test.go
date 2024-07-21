package keeper_test

import (
	"strconv"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "github.com/dneprix/ledger/testutil/keeper"
	"github.com/dneprix/ledger/x/ledger/keeper"
	"github.com/dneprix/ledger/x/ledger/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestAccountMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.LedgerKeeper(t)
	srv := keeper.NewMsgServerImpl(k)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateAccount{Creator: creator,
			Address: strconv.Itoa(i),
		}
		_, err := srv.CreateAccount(ctx, expected)
		require.NoError(t, err)
		rst, found := k.GetAccount(ctx,
			expected.Address,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestAccountMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateAccount
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateAccount{Creator: creator,
				Address: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateAccount{Creator: "B",
				Address: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateAccount{Creator: creator,
				Address: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.LedgerKeeper(t)
			srv := keeper.NewMsgServerImpl(k)
			expected := &types.MsgCreateAccount{Creator: creator,
				Address: strconv.Itoa(0),
			}
			_, err := srv.CreateAccount(ctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateAccount(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetAccount(ctx,
					expected.Address,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestAccountMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteAccount
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteAccount{Creator: creator,
				Address: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteAccount{Creator: "B",
				Address: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteAccount{Creator: creator,
				Address: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.LedgerKeeper(t)
			srv := keeper.NewMsgServerImpl(k)

			_, err := srv.CreateAccount(ctx, &types.MsgCreateAccount{Creator: creator,
				Address: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteAccount(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetAccount(ctx,
					tc.request.Address,
				)
				require.False(t, found)
			}
		})
	}
}
