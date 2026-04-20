package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var initializeCmd = &cobra.Command{
	Use:     "initialize [-c|-s|-u] [-w] pool [device...]",
	Short:   "initialize pool devices",
	GroupID: "maintenance",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initializeCmd).Standalone()

	initializeCmd.Flags().BoolS("c", "c", false, "cancel initialization")
	initializeCmd.Flags().BoolS("s", "s", false, "suspend initialization")
	initializeCmd.Flags().BoolS("u", "u", false, "uninitialize")
	initializeCmd.Flags().BoolS("w", "w", false, "wait until initialization completes")

	rootCmd.AddCommand(initializeCmd)

	carapace.Gen(initializeCmd).PositionalCompletion(
		zfs.ActionPools(),
	)

	carapace.Gen(initializeCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return zfs.ActionPoolDevices(c.Args[0])
			}
			return carapace.ActionValues()
		}),
	)
}
