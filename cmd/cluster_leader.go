package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/autopilothq/consensus/client"
	"github.com/autopilothq/banks/types"
	term "github.com/autopilothq/goterminus"
	"github.com/autopilothq/lg"
)

// # Display the cluster topology
// # Where addresses is one or more addresses that acceptors are listening on
// raglan cluster topology -a <addresses> <clusterID>
var clusterLeaderCmd = &cobra.Command{
	Use:   "leader [clusterID] [nodeAddr]",
	Short: "Returns the node id of the cluster leader",
	Long: `Returns the node id of the cluster leader

You must provide address and port of the raft endpoint of any node in the
cluster.

The leader info can be displayed as either plain text (the default) or as
json. In plain text mode the columns that are output are:
	id, address, current state, previous state
`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		lg.RemoveOutput(os.Stdout)
		lg.AddOutput(os.Stderr, lg.MinLevel(lg.LevelError))

		if len(args) <= 1 {
			return ErrNoTargetNodeAddr
		}

		var (
			nodeURL = url.URL{
				Scheme: "http", // TODO support https
				Host:   fmt.Sprintf(args[1]),
			}
			cnf = viper.GetViper()

			jsonOutput, humanOutput bool
		)
		clusterID, err := types.DecodeClusterID(cnf.GetString("Banks.Cluster.ID"))
		if err != nil {
			// we should never reach here before it should have been already handled
			// in config validation
			return err
		}

		leader, err := client.LeaderRequest(nodeURL, clusterID)
		if err != nil {
			return err
		}

		jsonOutput, humanOutput, err = getClusterFlags(cmd)
		if err != nil {
			return err
		}

		if jsonOutput {
			return dumpJSON(leader, humanOutput)
		}

		if humanOutput {
			term.Table(
				[]string{"id", "address", "state", "prev state"},
				[]string{
					leader.Id.String(),
					leader.Addr,
					leader.State.String(),
					leader.PrevState.String(),
				})

			return nil
		}

		fmt.Println(leader.Id.String(), leader.Addr, leader.State, leader.PrevState)

		return nil
	},
}

func init() {
	clusterCmd.AddCommand(clusterLeaderCmd)
}
