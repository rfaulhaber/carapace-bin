package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get [-Hp] [-o field,...] all|property[,...] [pool...]",
	Short:   "get pool properties",
	GroupID: "property",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()

	getCmd.Flags().BoolS("H", "H", false, "scripting mode")
	getCmd.Flags().BoolS("j", "j", false, "JSON output")
	getCmd.Flags().StringS("o", "o", "", "columns to display")
	getCmd.Flags().BoolS("p", "p", false, "display exact values")

	rootCmd.AddCommand(getCmd)

	carapace.Gen(getCmd).FlagCompletion(carapace.ActionMap{
		"o": carapace.ActionValues("name", "property", "value", "source").UniqueList(","),
	})

	carapace.Gen(getCmd).PositionalCompletion(
		carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				carapace.ActionValues("all"),
				zfs.ActionPoolProperties(),
				zfs.ActionReadonlyPoolProperties(),
			).ToA().FilterArgs()
		}),
	)

	carapace.Gen(getCmd).PositionalAnyCompletion(
		zfs.ActionPools(),
	)
}
