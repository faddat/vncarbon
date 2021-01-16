package vncarbon

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/faddat/vncarbon/x/vncarbon/keeper"
	"github.com/faddat/vncarbon/x/vncarbon/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case *types.MsgCreateVncarbon:
			return handleMsgCreateVncarbon(ctx, k, msg)

		case *types.MsgUpdateVncarbon:
			return handleMsgUpdateVncarbon(ctx, k, msg)

		case *types.MsgDeleteVncarbon:
			return handleMsgDeleteVncarbon(ctx, k, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
