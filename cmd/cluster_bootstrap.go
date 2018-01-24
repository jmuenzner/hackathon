package cmd

import (
	"fmt"
	"io"
	"net/url"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/autopilothq/consensus/client"
	"github.com/autopilothq/banks/types"
	"github.com/autopilothq/lg"
)

// TODO: support certs and TLS
var clusterBootstrapCmd = &cobra.Command{
	Use:   "bootstrap [clusterID] [nodeAddr]",
	Short: "Bootstraps a cluster of running Banks nodes",
	Long: `Bootstraps a cluster of running Banks nodes

You must provide address and port of the raft endpoint of any node in the
cluster. That node will be asked to perform the bootstrap for the entire
cluster.

Asking an existing cluster to bootstrap will do nothing.
`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		lg.RemoveOutput(os.Stdout)
		lg.AddOutput(os.Stderr, lg.MinLevel(lg.LevelError))

		if len(args) <= 1 {
			return ErrNoTargetNodeAddr
		}

		nodeURL := url.URL{
			Scheme: "http", // TODO support https
			Host:   fmt.Sprintf(args[1]),
		}

		cnf := viper.GetViper()
		clusterID, err := types.DecodeClusterID(cnf.GetString("Banks.Cluster.ID"))
		if err != nil {
			// we should never reach here before it should have been already handled
			// in config validation
			return err
		}

		fromSnapshot, err := cmd.Flags().GetString("from-snapshot")
		if err != nil {
			return err
		}
		var snapshot io.ReadCloser
		if fromSnapshot != "" {
			snapshot, err = os.Open(fromSnapshot)
			if err != nil {
				return err
			}
			defer snapshot.Close()
		}

		return client.BootstrapRequest(nodeURL, clusterID, snapshot)
	},
}

func init() {
	clusterCmd.AddCommand(clusterBootstrapCmd)
	clusterBootstrapCmd.Flags().String("from-snapshot", "",
		"The snapshot to bootstrap the cluster after disaster")
}
