package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var trimCmd = &cobra.Command{
	Use:     "trim [-dw] [-r rate] [-c|-s] pool [device...]",
	Short:   "initiate or manage TRIM on devices",
	GroupID: "maintenance",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trimCmd).Standalone()

	trimCmd.Flags().BoolS("c", "c", false, "cancel TRIM")
	trimCmd.Flags().BoolS("d", "d", false, "secure TRIM")
	trimCmd.Flags().StringS("r", "r", "", "TRIM rate limit")
	trimCmd.Flags().BoolS("s", "s", false, "suspend TRIM")
	trimCmd.Flags().BoolS("w", "w", false, "wait until TRIM completes")

	rootCmd.AddCommand(trimCmd)

	carapace.Gen(trimCmd).PositionalCompletion(
		zfs.ActionPools(),
	)

	carapace.Gen(trimCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return zfs.ActionPoolDevices(c.Args[0])
			}
			return carapace.ActionValues()
		}),
	)
}
