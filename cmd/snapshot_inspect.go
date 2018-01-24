package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/autopilothq/consensus/transport/snapshotter/fs"
	"github.com/autopilothq/banks/types"
	"github.com/autopilothq/lg"
)

/*
# Outputs a summary of the stats from a snapshot: cluster id, id node, terms, etc
#
# This can only be run from the node that has direct access to the data
# directory, this does mean that you can dump data from a Banks node even
# when that node has crashed.
#
# If the snapshot is invalid: output an error report.
journeys snapshot inspect <clusterID> <localOutputLocation>
*/
var snapshotInspectCmd = &cobra.Command{
	Use: "inspect [clusterID] [snapshot path]",
	Long: `Outputs a snapshot summary

Outputs a summary of the stats from a snapshot: cluster id, id node, terms, etc
`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		lg.RemoveOutput(os.Stdout)
		lg.AddOutput(os.Stderr, lg.MinLevel(lg.LevelError))

		var (
			conf = viper.GetViper()
		)
		if len(args) < 2 {
			return fmt.Errorf("snapshot path is required")
		}
		clusterID, err := types.DecodeClusterID(args[0])
		if err != nil {
			return err
		}
		snapshotPath := args[1]
		state, size, err := fs.InspectSnapshot(snapshotPath)
		if err != nil {
			return err
		}
		if types.ClusterID(state.ClusterID) != clusterID {
			return fmt.Errorf("The given snapshot from"+
				" a different cluster, expected %s, got %s",
				clusterID, state.ClusterID)
		}
		showSnapshotState(snapshotPath, size, state, conf.GetBool("json"))
		return nil
	},
}

func init() {
	snapshotCmd.AddCommand(snapshotInspectCmd)
}
