package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/hackathon/journeys/config"
	"github.com/autopilothq/banks/command"
	ct "github.com/autopilothq/banks/contract/types"
	"github.com/autopilothq/banks/types"
)


// StartCmd is the defintion of the cobra start command
var StartCmd = &cobra.Command{
	Use:   "start [clusterID]",
	Short: "Starts this journeys node",
	Long:  "Starts this journeys node",
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			cnf = viper.GetViper()
			clusterOptions = ct.ClusterOptions{}
		)

		clusterID, err := types.DecodeClusterID(cnf.GetString("Banks.Cluster.ID"))
		if err != nil {
			// we should never reach here before it should have been already handled
			// in config validation
			return err
		}

		clusterOptions.ID = clusterID
		clusterOptions.Bootstrap = cnf.GetBool("Banks.Cluster.Bootstrap")
		clusterOptions.Port = cnf.GetInt("Banks.Cluster.Port")
		clusterOptions.BootstrapSnapshot = cnf.GetString(
			"Banks.Cluster.BootstrapSnapshot")
		if clusterOptions.BootstrapSnapshot != "" && !clusterOptions.Bootstrap {
			return fmt.Errorf("cluster-bootstrap must be set" +
				" to bootstrap from snapshot")
		}
		nodeID := uint64(cnf.GetInt64("Banks.Cluster.NodeID"))

		return command.RunStart(
			nodeID,
			ConfigSource,
			clusterOptions,
			config.Reader(),
			config.SecretReader(),
			config.Log(),
			config.Reader().GetString("Banks.Log.Level"),
		)
	},
}

func init() {
	RootCmd.AddCommand(StartCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// StartCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

	StartCmd.Flags().String("id", "", "The default id for this node")

	viper.BindPFlag("Banks.Cluster.NodeID",
		StartCmd.Flags().Lookup("id"))

	StartCmd.Flags().String("cluster-id", "",
		"The id of the cluster to join")

	viper.BindPFlag("Banks.Cluster.ID", StartCmd.Flags().
		Lookup("cluster-id"))

	StartCmd.Flags().Int("cluster-port", 4747,
		"The default port of this node in the cluster")

	viper.BindPFlag("Banks.Cluster.Port", StartCmd.Flags().
		Lookup("cluster-port"))

	StartCmd.Flags().Bool("cluster-bootstrap", false,
		"If true this instructs the node to bootstrap a new cluster rather than "+
			"trying to join an existing one")

	viper.BindPFlag("Banks.Cluster.Bootstrap", StartCmd.Flags().
		Lookup("cluster-bootstrap"))

	StartCmd.Flags().String("bootstrap-from-snapshot", "",
		"The snapshot to bootstrap the cluster after disaster")

	viper.BindPFlag("Banks.Cluster.BootstrapSnapshot", StartCmd.Flags().
		Lookup("bootstrap-from-snapshot"))

	StartCmd.Flags().String("nodeid-path", "",
		"The nodeID file path")

	viper.BindPFlag("Banks.NodeIDPath", StartCmd.Flags().
		Lookup("nodeid-path"))

	StartCmd.Flags().Bool("profile", false, "profiles main execution if true")
}
