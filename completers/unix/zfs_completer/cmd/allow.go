package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var allowCmd = &cobra.Command{
	Use:     "allow [-dluge] [-s @setname] permission[,...] dataset",
	Short:   "delegate permissions on a dataset",
	GroupID: "delegation",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(allowCmd).Standalone()

	allowCmd.Flags().BoolS("c", "c", false, "set create-time permissions")
	allowCmd.Flags().BoolS("d", "d", false, "allow for descendants only")
	allowCmd.Flags().BoolS("e", "e", false, "delegate to everyone")
	allowCmd.Flags().StringS("g", "g", "", "delegate to group")
	allowCmd.Flags().BoolS("l", "l", false, "allow locally only")
	allowCmd.Flags().StringS("s", "s", "", "define or add to permission set")
	allowCmd.Flags().StringS("u", "u", "", "delegate to user")

	rootCmd.AddCommand(allowCmd)

	carapace.Gen(allowCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return zfs.ActionPermissions().UniqueList(",")
			}
			return zfs.ActionFilesystems()
		}),
	)
}
