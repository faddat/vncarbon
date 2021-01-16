package vncarbon

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/faddat/vncarbon/x/vncarbon/keeper"
	"github.com/faddat/vncarbon/x/vncarbon/types"
)

func handleMsgCreateVncarbon(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateVncarbon) (*sdk.Result, error) {
	k.CreateVncarbon(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateVncarbon(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateVncarbon) (*sdk.Result, error) {
	var vncarbon = types.Vncarbon{
		Creator: msg.Creator,
		Id:      msg.Id,
		Carbon:  msg.Carbon,
		Emitter: msg.Emitter,
		Date:    msg.Date,
	}

	if msg.Creator != k.GetVncarbonOwner(ctx, msg.Id) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner") // If not, throw an error
	}

	k.UpdateVncarbon(ctx, vncarbon)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgDeleteVncarbon(ctx sdk.Context, k keeper.Keeper, msg *types.MsgDeleteVncarbon) (*sdk.Result, error) {
	if !k.HasVncarbon(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, msg.Id)
	}
	if msg.Creator != k.GetVncarbonOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner")
	}

	k.DeleteVncarbon(ctx, msg.Id)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
