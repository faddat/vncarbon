package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateVncarbon{}

func NewMsgCreateVncarbon(creator string, carbon int32, emitter string, date string) *MsgCreateVncarbon {
	return &MsgCreateVncarbon{
		Creator: creator,
		Carbon:  carbon,
		Emitter: emitter,
		Date:    date,
	}
}

func (msg *MsgCreateVncarbon) Route() string {
	return RouterKey
}

func (msg *MsgCreateVncarbon) Type() string {
	return "CreateVncarbon"
}

func (msg *MsgCreateVncarbon) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateVncarbon) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateVncarbon) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateVncarbon{}

func NewMsgUpdateVncarbon(creator string, id string, carbon int32, emitter string, date string) *MsgUpdateVncarbon {
	return &MsgUpdateVncarbon{
		Id:      id,
		Creator: creator,
		Carbon:  carbon,
		Emitter: emitter,
		Date:    date,
	}
}

func (msg *MsgUpdateVncarbon) Route() string {
	return RouterKey
}

func (msg *MsgUpdateVncarbon) Type() string {
	return "UpdateVncarbon"
}

func (msg *MsgUpdateVncarbon) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateVncarbon) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateVncarbon) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgCreateVncarbon{}

func NewMsgDeleteVncarbon(creator string, id string) *MsgDeleteVncarbon {
	return &MsgDeleteVncarbon{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteVncarbon) Route() string {
	return RouterKey
}

func (msg *MsgDeleteVncarbon) Type() string {
	return "DeleteVncarbon"
}

func (msg *MsgDeleteVncarbon) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteVncarbon) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteVncarbon) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
