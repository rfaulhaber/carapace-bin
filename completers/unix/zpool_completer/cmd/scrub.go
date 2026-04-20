package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var scrubCmd = &cobra.Command{
	Use:     "scrub [-s|-p|-w] [-e] pool...",
	Short:   "start or manage a scrub",
	GroupID: "maintenance",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scrubCmd).Standalone()

	scrubCmd.Flags().BoolS("e", "e", false, "only scrub error blocks")
	scrubCmd.Flags().BoolS("p", "p", false, "pause scrub")
	scrubCmd.Flags().BoolS("s", "s", false, "stop scrub")
	scrubCmd.Flags().BoolS("w", "w", false, "wait until scrub completes")

	rootCmd.AddCommand(scrubCmd)

	carapace.Gen(scrubCmd).PositionalAnyCompletion(
		zfs.ActionPools(),
	)
}
