package cmd

import (
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/autopilothq/consensus/client"
	"github.com/autopilothq/consensus/nodes"
	"github.com/autopilothq/consensus/types"
	term "github.com/autopilothq/goterminus"
	"github.com/autopilothq/lg"
)

// # Display the cluster topology
// # Where addresses is one or more addresses that acceptors are listening on
// raglan cluster topology -a <addresses> <clusterID>
var clusterTopologyCmd = &cobra.Command{
	Use:   "topology [clusterID] [nodeAddr]",
	Short: "Displays the cluster topology",
	Long: `Displays the general cluster toplogy, and a short summary of each known node.

You must provide address and port of the raft endpoint of any node in the
cluster.
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

			ns                      []*nodes.Node
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

		ns, err = client.TopologyRequest(nodeURL, clusterID)
		if err != nil {
			return err
		}

		jsonOutput, humanOutput, err = getClusterFlags(cmd)
		if err != nil {
			return err
		}

		if jsonOutput {
			return dumpJSON(ns, humanOutput)
		}

		if humanOutput {
			rows := make([][]string, 0, len(ns))

			for _, node := range ns {
				leaderState := ""
				if node.Id == leader.Id {
					leaderState = "leader"
				}

				rows = append(rows, []string{
					leader.Id.String(),
					leader.Addr,
					leader.State.String(),
					leader.PrevState.String(),
					leaderState,
				})
			}

			term.Table([]string{"id", "address", "state", "prev state", "leader"},
				rows...)

			return nil
		}

		for _, node := range ns {
			leaderState := ""
			if node.Id == leader.Id {
				leaderState = "leader"
			}

			fmt.Println(node.Id.String(), node.Addr, node.State, node.PrevState,
				leaderState)
		}

		return nil
	},
}

func init() {
	clusterCmd.AddCommand(clusterTopologyCmd)
}
