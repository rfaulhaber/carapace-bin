package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list [-HgLpPv] [-o property,...] [-T d|u] [pool...] [interval [count]]",
	Short:   "list pools and their properties",
	GroupID: "property",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().BoolS("H", "H", false, "scripting mode")
	listCmd.Flags().BoolS("L", "L", false, "resolve symbolic links")
	listCmd.Flags().BoolS("P", "P", false, "display full paths")
	listCmd.Flags().StringS("T", "T", "", "timestamp format")
	listCmd.Flags().BoolS("g", "g", false, "display vdev GUIDs")
	listCmd.Flags().BoolS("j", "j", false, "JSON output")
	listCmd.Flags().StringS("o", "o", "", "properties to display")
	listCmd.Flags().BoolS("p", "p", false, "display exact values")
	listCmd.Flags().BoolS("v", "v", false, "verbose vdev statistics")

	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"T": carapace.ActionValues("d", "u"),
		"o": zfs.ActionPoolListColumns().UniqueList(","),
	})

	carapace.Gen(listCmd).PositionalAnyCompletion(
		zfs.ActionPools(),
	)
}
