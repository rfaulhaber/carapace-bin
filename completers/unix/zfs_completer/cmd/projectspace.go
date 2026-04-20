package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var projectspaceCmd = &cobra.Command{
	Use:     "projectspace [-Hp] [-o field,...] [-s field] [-S field] dataset",
	Short:   "display project space usage",
	GroupID: "misc",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projectspaceCmd).Standalone()

	projectspaceCmd.Flags().BoolS("H", "H", false, "scripting mode")
	projectspaceCmd.Flags().StringS("o", "o", "", "fields to display")
	projectspaceCmd.Flags().BoolS("p", "p", false, "display exact values")
	projectspaceCmd.Flags().StringS("s", "s", "", "sort ascending by field")
	projectspaceCmd.Flags().StringS("S", "S", "", "sort descending by field")

	rootCmd.AddCommand(projectspaceCmd)

	carapace.Gen(projectspaceCmd).FlagCompletion(carapace.ActionMap{
		"o": carapace.ActionValues("type", "name", "used", "quota", "objused", "objquota").UniqueList(","),
		"s": carapace.ActionValues("type", "name", "used", "quota", "objused", "objquota"),
		"S": carapace.ActionValues("type", "name", "used", "quota", "objused", "objquota"),
	})

	carapace.Gen(projectspaceCmd).PositionalCompletion(
		zfs.ActionFilesystems(),
	)
}
