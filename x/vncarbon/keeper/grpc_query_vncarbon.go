package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/faddat/vncarbon/x/vncarbon/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) VncarbonAll(c context.Context, req *types.QueryAllVncarbonRequest) (*types.QueryAllVncarbonResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var vncarbons []*types.Vncarbon
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	vncarbonStore := prefix.NewStore(store, types.KeyPrefix(types.VncarbonKey))

	pageRes, err := query.Paginate(vncarbonStore, req.Pagination, func(key []byte, value []byte) error {
		var vncarbon types.Vncarbon
		if err := k.cdc.UnmarshalBinaryBare(value, &vncarbon); err != nil {
			return err
		}

		vncarbons = append(vncarbons, &vncarbon)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVncarbonResponse{Vncarbon: vncarbons, Pagination: pageRes}, nil
}

func (k Keeper) Vncarbon(c context.Context, req *types.QueryGetVncarbonRequest) (*types.QueryGetVncarbonResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var vncarbon types.Vncarbon
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VncarbonKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.VncarbonKey+req.Id)), &vncarbon)

	return &types.QueryGetVncarbonResponse{Vncarbon: &vncarbon}, nil
}
