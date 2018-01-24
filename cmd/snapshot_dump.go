package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/autopilothq/consensus/client"
	"github.com/autopilothq/consensus/types"
	"github.com/autopilothq/lg"
)

/*
# Dump a snapshot to disk from a specific node. This may not be consistent
# with the entire cluster as the node may be partitioned. This is designed
# to be used to get data about a specific node when the cluster may not be
# in a good state.
#
# This can only be run from the node that has direct access to the data
# directory, this does mean that you can dump data from a Banks node even
# when that node has crashed.
#
# After creating the new snapshot it verifies that it's valid.
#
# Ops usage of this is to run it on every node and manage 3 or 5 individual
# backups per cluster.
#
journeys snapshot dump <clusterID> <nodeID> -d <dataDir> -o <localOutputLocation>
*/
var snapshotDumpCmd = &cobra.Command{
	Use: "dump [clusterID] [nodeID]",
	Long: `Dumps the latest snapshot

Dump a snapshot to disk from a specific node. This may not be consistent
with the entire cluster as the node may be partitioned. This is designed
to be used to get data about a specific node when the cluster may not be
in a good state.
`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		lg.RemoveOutput(os.Stdout)
		lg.AddOutput(os.Stderr, lg.MinLevel(lg.LevelError))

		var (
			log  = lg.Extend()
			conf = viper.GetViper()
		)
		if len(args) == 0 {
			return fmt.Errorf("ClusterID is required")
		}
		clusterID, err := types.DecodeClusterID(args[0])
		if err != nil {
			return err
		}
		var nodeID types.ID
		if len(args) > 1 {
			nodeID, err = types.DecodeID(args[1])
			if err != nil {
				return err
			}
		}

		flags := cmd.Flags()
		snapDir, err := flags.GetString("snap-dir")
		if err != nil {
			return err
		}
		if snapDir == "" {
			dataDir, derr := flags.GetString("data-dir")
			if derr != nil {
				return derr
			}
			if dataDir == "" {
				return fmt.Errorf("data-dir or snap-dir must be set")
			}
			snapDir, err = client.LookupSnapshotDirectory(dataDir, nodeID)
			if err != nil {
				return err
			}
		}

		if snapDir == "" {
			return fmt.Errorf("Couldn't find snapshot dir")
		}
		output, err := flags.GetString("output")
		if err != nil {
			return err
		}
		if output == "" {
			return fmt.Errorf("output must be set")
		}
		state, size, err := client.DumpLatestSnapshot(snapDir, output, clusterID,
			nodeID, log)
		if err != nil {
			return err
		}
		// output json
		showSnapshotState(output, size, state, conf.GetBool("json"))
		return nil
	},
}

func init() {
	snapshotCmd.AddCommand(snapshotDumpCmd)
	flags := snapshotDumpCmd.Flags()
	flags.StringP("data-dir", "d", "", "the baks data directory")
	flags.String("snap-dir", "", "the snapshots directory")
	flags.StringP("output", "o", "", "local output location")
}
