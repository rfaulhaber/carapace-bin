package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:     "set property=value pool",
	Short:   "set pool property",
	GroupID: "property",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setCmd).Standalone()

	rootCmd.AddCommand(setCmd)

	carapace.Gen(setCmd).PositionalCompletion(
		zfs.ActionPoolPropertyAssignments(),
		zfs.ActionPools(),
	)
}
