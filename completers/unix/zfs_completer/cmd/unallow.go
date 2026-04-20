package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var unallowCmd = &cobra.Command{
	Use:     "unallow [-rlduge] [-s @setname] permission[,...] dataset",
	Short:   "revoke delegated permissions",
	GroupID: "delegation",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unallowCmd).Standalone()

	unallowCmd.Flags().BoolS("c", "c", false, "remove create-time permissions")
	unallowCmd.Flags().BoolS("d", "d", false, "remove from descendants only")
	unallowCmd.Flags().BoolS("e", "e", false, "remove from everyone")
	unallowCmd.Flags().StringS("g", "g", "", "remove from group")
	unallowCmd.Flags().BoolS("l", "l", false, "remove locally only")
	unallowCmd.Flags().BoolS("r", "r", false, "recursively remove")
	unallowCmd.Flags().StringS("s", "s", "", "remove from permission set")
	unallowCmd.Flags().StringS("u", "u", "", "remove from user")

	rootCmd.AddCommand(unallowCmd)

	carapace.Gen(unallowCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return zfs.ActionPermissions().UniqueList(",")
			}
			return zfs.ActionFilesystems()
		}),
	)
}
