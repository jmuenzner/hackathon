package cmd

import (
	"context"
	"fmt"
	"os"

	term "github.com/autopilothq/goterminus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/autopilothq/banks/discovery/cluster"
	"github.com/autopilothq/banks/natshelper"
	"github.com/autopilothq/lg"
)

// # Display the cluster topology
// # Where addresses is one or more addresses that acceptors are listening on
// raglan cluster topology -a <addresses> <clusterID>
var discoverClusterCmd = &cobra.Command{
	Use:   "cluster [clusterID]",
	Short: "Discovers the cluster topology",
	Long: `Discovers the cluster topology

Discovers and displays the cluster toplogy, and a short summary of each
known node. The results of the discovery may be limited during network
partitions.

You must specify one or more nats addresses to request the information from.
`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		lg.RemoveOutput(os.Stdout)
		lg.AddOutput(os.Stderr, lg.MinLevel(lg.LevelError))

		var (
			log   = lg.Extend()
			conf  = viper.GetViper()
			cd    *cluster.Discovery
			nodes cluster.Nodes
		)

		minSize, timeout, jsonOutput, humanOutput, err := getDiscoveryFlags(cmd)
		if err != nil {
			return err
		}

		conn, err := natshelper.FromConfig(conf, log)
		if err != nil {
			return err
		}

		cd, err = cluster.Make(conn, conf, log)
		if err != nil {
			return err
		}

		ctx, cancelFn := context.WithTimeout(context.Background(), timeout)
		defer cancelFn()
		nodes, err = cd.Discover(ctx, minSize)
		if err != nil {
			return err
		}

		if jsonOutput {
			return dumpJSON(nodes, humanOutput)
		}

		if humanOutput {
			rows := make([][]string, 0, len(nodes))
			for _, node := range nodes {
				rows = append(rows, []string{node.ID.String(), node.Addr})
			}

			term.Table([]string{"id", "address"}, rows...)
			return nil
		}

		for _, node := range nodes {
			fmt.Println(node.ID.String(), node.Addr)
		}

		return nil
	},
}

func init() {
	discoverCmd.AddCommand(discoverClusterCmd)
}
