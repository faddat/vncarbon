package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateVncarbon{}, "vncarbon/CreateVncarbon", nil)
	cdc.RegisterConcrete(&MsgUpdateVncarbon{}, "vncarbon/UpdateVncarbon", nil)
	cdc.RegisterConcrete(&MsgDeleteVncarbon{}, "vncarbon/DeleteVncarbon", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateVncarbon{},
		&MsgUpdateVncarbon{},
		&MsgDeleteVncarbon{},
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
