package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var checkpointCmd = &cobra.Command{
	Use:     "checkpoint [-d|-w] pool",
	Short:   "checkpoint a pool",
	GroupID: "maintenance",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkpointCmd).Standalone()

	checkpointCmd.Flags().BoolS("d", "d", false, "discard checkpoint")
	checkpointCmd.Flags().BoolS("w", "w", false, "wait until checkpoint completes")

	rootCmd.AddCommand(checkpointCmd)

	carapace.Gen(checkpointCmd).PositionalCompletion(
		zfs.ActionPools(),
	)
}
