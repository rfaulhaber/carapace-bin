package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var loadKeyCmd = &cobra.Command{
	Use:     "load-key [-rnL] [-l|-a] [dataset]",
	Short:   "load encryption key for a dataset",
	GroupID: "encryption",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loadKeyCmd).Standalone()

	loadKeyCmd.Flags().BoolS("L", "L", false, "prompt for key if not available")
	loadKeyCmd.Flags().BoolS("a", "a", false, "load keys for all encrypted datasets")
	loadKeyCmd.Flags().BoolS("n", "n", false, "dry-run")
	loadKeyCmd.Flags().BoolS("r", "r", false, "load keys recursively")

	rootCmd.AddCommand(loadKeyCmd)

	carapace.Gen(loadKeyCmd).PositionalCompletion(
		zfs.ActionFilesystems(),
	)
}
