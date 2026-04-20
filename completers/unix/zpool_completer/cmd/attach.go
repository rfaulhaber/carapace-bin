package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var attachCmd = &cobra.Command{
	Use:     "attach [-fsw] [-o property=value] pool device new_device",
	Short:   "attach a device to a vdev",
	GroupID: "vdev",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attachCmd).Standalone()

	attachCmd.Flags().BoolS("f", "f", false, "force attach")
	attachCmd.Flags().StringArrayS("o", "o", nil, "set property")
	attachCmd.Flags().BoolS("s", "s", false, "do not resilver")
	attachCmd.Flags().BoolS("w", "w", false, "wait until resilvering completes")

	rootCmd.AddCommand(attachCmd)

	carapace.Gen(attachCmd).FlagCompletion(carapace.ActionMap{
		"o": zfs.ActionPoolPropertyAssignments(),
	})

	carapace.Gen(attachCmd).PositionalCompletion(
		zfs.ActionPools(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return zfs.ActionPoolDevices(c.Args[0])
		}),
		carapace.ActionFiles(),
	)
}
