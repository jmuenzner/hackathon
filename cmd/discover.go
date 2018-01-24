package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

var discoverCmd = &cobra.Command{
	Use:   "discover",
	Short: "Run various discovery operations on the cluster",
}

func init() {
	RootCmd.AddCommand(discoverCmd)
	pflags := discoverCmd.PersistentFlags()
	pflags.Duration("timeout", 5*time.Second, "how long to attempt to discovery "+
		"the topology before giving up")
	pflags.Int("minSize", 0, "pool until the cluster reaches a specific "+
		"size. Defaults to 0, which means any size")
	pflags.Bool("json", false, "True to use json output instead of plain text")
	pflags.BoolP("human", "r", false, "True to output in a human friendly format")
}

func getDiscoveryFlags(cmd *cobra.Command) (
	minSize int, timeout time.Duration, jsonOutput bool, human bool, err error,
) {
	flags := cmd.Flags()
	minSize, err = flags.GetInt("minSize")

	if err == nil {
		timeout, err = flags.GetDuration("timeout")
	}

	if err == nil {
		jsonOutput, err = flags.GetBool("json")
	}

	if err == nil {
		human, err = flags.GetBool("human")
	}

	return minSize, timeout, jsonOutput, human, err
}
