package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/faddat/vncarbon/x/vncarbon/types"
	"github.com/spf13/cobra"
)

func CmdListVncarbon() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-vncarbon",
		Short: "list all vncarbon",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllVncarbonRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.VncarbonAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowVncarbon() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-vncarbon [id]",
		Short: "shows a vncarbon",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetVncarbonRequest{
				Id: args[0],
			}

			res, err := queryClient.Vncarbon(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
