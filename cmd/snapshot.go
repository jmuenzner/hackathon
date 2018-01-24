package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/autopilothq/consensus/raftpb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Run various snapshot operations on the local node",
}

func init() {
	RootCmd.AddCommand(snapshotCmd)
	pflags := snapshotCmd.PersistentFlags()
	pflags.Bool("json", false, "True to use json output instead of plain text")
	viper.BindPFlag("json", pflags.Lookup("json"))
}

func showSnapshotState(output string, size int64, state raftpb.State,
	showJSON bool) {
	if showJSON {
		out := struct {
			Path      string
			Size      int64
			ClusterID string
			NodeID    uint64
			Index     uint64
			Term      uint64
			Created   int64
			Nodes     []uint64
		}{
			Path:      output,
			Size:      size,
			ClusterID: state.ClusterID,
			NodeID:    state.NodeID,
			Index:     state.Index,
			Term:      state.Term,
			Created:   state.Created,
			Nodes:     state.ConfState.Nodes,
		}
		json.NewEncoder(os.Stdout).Encode(out)
	} else {
		fmt.Printf("Path: %s, Size: %d\n", output, size)
		fmt.Printf("ClusterID: %s, NodeID: %d\n", state.ClusterID, state.NodeID)
		fmt.Printf("Index: %d, Term: %d\n", state.Index, state.Term)
		fmt.Printf("Raft configuration: %v\n", state.ConfState)
		fmt.Printf("Created: %s\n", time.Unix(state.Created, 0))
	}
}
