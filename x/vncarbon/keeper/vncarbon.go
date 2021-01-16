package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/faddat/vncarbon/x/vncarbon/types"
	"strconv"
)

// GetVncarbonCount get the total number of vncarbon
func (k Keeper) GetVncarbonCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VncarbonCountKey))
	byteKey := types.KeyPrefix(types.VncarbonCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetVncarbonCount set the total number of vncarbon
func (k Keeper) SetVncarbonCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VncarbonCountKey))
	byteKey := types.KeyPrefix(types.VncarbonCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) CreateVncarbon(ctx sdk.Context, msg types.MsgCreateVncarbon) {
	// Create the vncarbon
	count := k.GetVncarbonCount(ctx)
	var vncarbon = types.Vncarbon{
		Creator: msg.Creator,
		Id:      strconv.FormatInt(count, 10),
		Carbon:  msg.Carbon,
		Emitter: msg.Emitter,
		Date:    msg.Date,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VncarbonKey))
	key := types.KeyPrefix(types.VncarbonKey + vncarbon.Id)
	value := k.cdc.MustMarshalBinaryBare(&vncarbon)
	store.Set(key, value)

	// Update vncarbon count
	k.SetVncarbonCount(ctx, count+1)
}

func (k Keeper) UpdateVncarbon(ctx sdk.Context, vncarbon types.Vncarbon) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VncarbonKey))
	b := k.cdc.MustMarshalBinaryBare(&vncarbon)
	store.Set(types.KeyPrefix(types.VncarbonKey+vncarbon.Id), b)
}

func (k Keeper) GetVncarbon(ctx sdk.Context, key string) types.Vncarbon {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VncarbonKey))
	var vncarbon types.Vncarbon
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.VncarbonKey+key)), &vncarbon)
	return vncarbon
}

func (k Keeper) HasVncarbon(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VncarbonKey))
	return store.Has(types.KeyPrefix(types.VncarbonKey + id))
}

func (k Keeper) GetVncarbonOwner(ctx sdk.Context, key string) string {
	return k.GetVncarbon(ctx, key).Creator
}

// DeleteVncarbon deletes a vncarbon
func (k Keeper) DeleteVncarbon(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VncarbonKey))
	store.Delete(types.KeyPrefix(types.VncarbonKey + key))
}

func (k Keeper) GetAllVncarbon(ctx sdk.Context) (msgs []types.Vncarbon) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VncarbonKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.VncarbonKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Vncarbon
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
