package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var rewriteCmd = &cobra.Command{
	Use:     "rewrite [-b | -c algorithm | -d | -u] [-r] filesystem|volume",
	Short:   "rewrite data blocks without modifying logical content",
	GroupID: "misc",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rewriteCmd).Standalone()

	rewriteCmd.Flags().BoolS("b", "b", false, "rewrite only partially-filled blocks")
	rewriteCmd.Flags().StringS("c", "c", "", "recompress using algorithm")
	rewriteCmd.Flags().BoolS("d", "d", false, "rewrite dedup blocks as non-dedup")
	rewriteCmd.Flags().BoolS("r", "r", false, "apply recursively")
	rewriteCmd.Flags().BoolS("u", "u", false, "rewrite uncompressed")

	rootCmd.AddCommand(rewriteCmd)

	carapace.Gen(rewriteCmd).FlagCompletion(carapace.ActionMap{
		"c": zfs.ActionPropertyValues("compression"),
	})

	carapace.Gen(rewriteCmd).PositionalCompletion(
		zfs.ActionFilesystems(),
	)
}
