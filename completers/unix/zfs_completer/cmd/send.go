package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var sendCmd = &cobra.Command{
	Use:     "send [-DLPbcehnpsvw] [-R [-X dataset,...]] [[-I|-i] snapshot] snapshot",
	Short:   "generate a send stream",
	GroupID: "send",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sendCmd).Standalone()

	sendCmd.Flags().BoolS("D", "D", false, "deduplicate (deprecated)")
	sendCmd.Flags().StringS("I", "I", "", "send all intermediary snapshots")
	sendCmd.Flags().BoolS("L", "L", false, "allow large blocks")
	sendCmd.Flags().BoolS("P", "P", false, "machine-parsable output")
	sendCmd.Flags().BoolS("R", "R", false, "replicate")
	sendCmd.Flags().StringS("S", "S", "", "send saved dataset state")
	sendCmd.Flags().BoolS("V", "V", false, "show data rate in process title")
	sendCmd.Flags().StringS("X", "X", "", "exclude datasets from replication")
	sendCmd.Flags().BoolS("b", "b", false, "send only received property values")
	sendCmd.Flags().BoolS("c", "c", false, "use compressed WRITE records")
	sendCmd.Flags().BoolS("e", "e", false, "use WRITE_EMBEDDED records")
	sendCmd.Flags().BoolS("h", "h", false, "include snapshot holds")
	sendCmd.Flags().StringS("i", "i", "", "send incremental from snapshot")
	sendCmd.Flags().BoolS("n", "n", false, "dry-run")
	sendCmd.Flags().BoolS("p", "p", false, "include dataset properties")
	sendCmd.Flags().BoolS("s", "s", false, "skip missing snapshots")
	sendCmd.Flags().StringS("t", "t", "", "resume send with token")
	sendCmd.Flags().BoolS("v", "v", false, "verbose")
	sendCmd.Flags().BoolS("w", "w", false, "raw send for encrypted datasets")
	sendCmd.Flags().String("redact", "", "redaction bookmark")

	rootCmd.AddCommand(sendCmd)

	carapace.Gen(sendCmd).FlagCompletion(carapace.ActionMap{
		"I":      zfs.ActionSnapshots(),
		"S":      zfs.ActionFilesystems(),
		"X":      zfs.ActionFilesystems().UniqueList(","),
		"i":      zfs.ActionDatasets(zfs.DatasetOpts{Snapshot: true, Bookmark: true}),
		"redact": zfs.ActionBookmarks(),
	})

	carapace.Gen(sendCmd).PositionalCompletion(
		zfs.ActionDatasets(zfs.DatasetOpts{Snapshot: true, Filesystem: true, Volume: true}),
	)
}
