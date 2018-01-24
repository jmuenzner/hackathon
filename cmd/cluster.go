package cmd

import (
	"github.com/spf13/cobra"
)

var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Various cluster commands",
}

func init() {
	RootCmd.AddCommand(clusterCmd)
	pflags := clusterCmd.PersistentFlags()
	pflags.Bool("json", false, "True to use json output instead of plain text")
	pflags.BoolP("human", "r", false, "True to output in a human friendly format")
}

func getClusterFlags(cmd *cobra.Command) (
	jsonOutput bool, humanOutput bool, err error,
) {
	flags := cmd.Flags()
	jsonOutput, err = flags.GetBool("json")

	if err == nil {
		humanOutput, err = flags.GetBool("human")
	}

	return jsonOutput, humanOutput, err
}
