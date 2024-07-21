package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateAccount{}

func NewMsgCreateAccount(
	creator string,
	address string,
	balance sdk.Coin,

) *MsgCreateAccount {
	return &MsgCreateAccount{
		Creator: creator,
		Address: address,
		Balance: balance,
	}
}

func (msg *MsgCreateAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateAccount{}

func NewMsgUpdateAccount(
	creator string,
	address string,
	balance sdk.Coin,

) *MsgUpdateAccount {
	return &MsgUpdateAccount{
		Creator: creator,
		Address: address,
		Balance: balance,
	}
}

func (msg *MsgUpdateAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteAccount{}

func NewMsgDeleteAccount(
	creator string,
	address string,

) *MsgDeleteAccount {
	return &MsgDeleteAccount{
		Creator: creator,
		Address: address,
	}
}

func (msg *MsgDeleteAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
