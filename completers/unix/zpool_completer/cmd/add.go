package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add [-fgLnP] [-o property=value] pool vdev...",
	Short:   "add vdevs to a pool",
	GroupID: "vdev",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()

	addCmd.Flags().BoolS("f", "f", false, "force use of vdevs")
	addCmd.Flags().BoolS("g", "g", false, "display vdev GUIDs")
	addCmd.Flags().BoolS("L", "L", false, "resolve symbolic links")
	addCmd.Flags().BoolS("n", "n", false, "dry-run")
	addCmd.Flags().StringArrayS("o", "o", nil, "set property")
	addCmd.Flags().BoolS("P", "P", false, "display full paths")

	rootCmd.AddCommand(addCmd)

	carapace.Gen(addCmd).FlagCompletion(carapace.ActionMap{
		"o": zfs.ActionPoolPropertyAssignments(),
	})

	carapace.Gen(addCmd).PositionalCompletion(
		zfs.ActionPools(),
	)

	carapace.Gen(addCmd).PositionalAnyCompletion(
		carapace.Batch(
			zfs.ActionVdevTypes(),
			carapace.ActionFiles(),
		).ToA(),
	)
}
