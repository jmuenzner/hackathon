package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/autopilothq/banks/discovery/nats"
	"github.com/autopilothq/lg"
)

// # Display the cluster topology
// # Where addresses is one or more addresses that acceptors are listening on
// raglan cluster topology -a <addresses> <clusterID>
var discoverNatsCmd = &cobra.Command{
	Use:   "nats [clusterID]",
	Short: "Discovers the nats nodes",
	Long: `Discovers the nats nodes

Discovers and displays the available nats nodes. The results of the discovery
may be limited during network partitions.
`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		lg.RemoveOutput(os.Stdout)
		lg.AddOutput(os.Stderr, lg.MinLevel(lg.LevelError))

		var (
			log  = lg.Extend()
			conf = viper.GetViper()
		)

		_, _, jsonOutput, humanOutput, err := getDiscoveryFlags(cmd)
		if err != nil {
			return err
		}

		natsAddrs, err := nats.FromConfig(conf, log)
		if err != nil {
			return err
		}

		if jsonOutput {
			return dumpJSON(natsAddrs, humanOutput)
		}

		fmt.Println(strings.Join(natsAddrs, " "))
		return nil
	},
}

func init() {
	discoverCmd.AddCommand(discoverNatsCmd)

}
