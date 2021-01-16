package cli

import (
	"github.com/spf13/cobra"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/faddat/vncarbon/x/vncarbon/types"
)

func CmdCreateVncarbon() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-vncarbon [carbon] [emitter] [date]",
		Short: "Creates a new vncarbon",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsCarbon, _ := strconv.ParseInt(args[0], 10, 64)
			argsEmitter := string(args[1])
			argsDate := string(args[2])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateVncarbon(clientCtx.GetFromAddress().String(), int32(argsCarbon), string(argsEmitter), string(argsDate))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateVncarbon() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-vncarbon [id] [carbon] [emitter] [date]",
		Short: "Update a vncarbon",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsCarbon, _ := strconv.ParseInt(args[1], 10, 64)
			argsEmitter := string(args[2])
			argsDate := string(args[3])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateVncarbon(clientCtx.GetFromAddress().String(), id, int32(argsCarbon), string(argsEmitter), string(argsDate))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteVncarbon() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-vncarbon [id] [carbon] [emitter] [date]",
		Short: "Delete a vncarbon by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteVncarbon(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
