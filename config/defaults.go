package config

import (
	"github.com/autopilothq/banks/config"
	"github.com/spf13/viper"
)

// ApplyDefaults applies default values from the environment if preset,
// otherwise uses defaults
func ApplyDefaults() {
	config.SetDefaultFromEnv("Banks.BindAddr", "127.0.0.1", viper.SetDefault)
	config.SetDefaultFromEnv("Banks.NATS.Addrs", []string{"nats://localhost:4222"},
		viper.SetDefault)
	config.SetDefaultFromEnv("Banks.Log.Level", "info", viper.SetDefault)
	config.SetDefaultFromEnv("Banks.Log.Format", "text", viper.SetDefault)
	config.SetDefaultFromEnv("Banks.SubjectPrefix", "Banks", viper.SetDefault)
	config.SetDefaultFromEnv("Banks.Health", map[string]interface{}{
		"Endpoint": "0.0.0.0:4567",
		"Enabled":  false,
	}, viper.SetDefault)
	config.SetDefaultFromEnv("Banks.Cluster", map[string]interface{}{
		"Port":              4747,
		"DataDir":           "./data",
		"LeaderWaitTimeout": 7000000000, // 7 seconds in nanoseconds
		// MinSnapshotEntries is the number of entries before a new snapshot kicks off
		"MinSnapshotEntries": 10000,
		// MinSize is minimum size of the cluster
		"MinSize": 3,
		// AckDelay is delay before any acknowledges are emitted
		"AckDelay": "5ms",
		// MinAckTimeout is minimum work duration to acknowledge it
		"MinAckTimeout": "500ms",
		// JoinTimeout is the length of time we're willing to wait to
		// join the cluster.
		"JoinTimeout":   "60s",
	}, viper.SetDefault)

	// Use the following DNS server for discovery. An empty string means that
	// discovery should defer to resolv.conf, which is what you want in most
	// situations
	config.SetDefaultFromEnv("Banks.Discovery.CustomDNSServer", "",
		viper.SetDefault)

	config.SetDefaultFromEnv("Banks.Discovery.Types.NATS", map[string]interface{}{
		// Query the following DNS name to discover NATs nodes.
		"DnsName": "nats.banks.",
		// TimeoutNatsDiscovery is the length of time (in ms) that Nats
		// discovery is allowed to take. 1000 ms by default.
		"Timeout": "1000ms",
		// DefaultPort is the default port to assume Nats is running on
		"DefaultPort": 4222,
	}, viper.SetDefault)

	config.SetDefaultFromEnv("Banks.Discovery.Types.Cluster", map[string]interface{}{
		// Timeout is allow cluster node discovery. 5s by default.
		"Timeout": "5s",
	}, viper.SetDefault)

	config.SetDefaultFromEnv("Banks.Acceptors.Websocket", map[string]interface{}{
		"Port": 7777,
	}, viper.SetDefault)
	config.SetDefaultFromEnv("Banks.Storage.Cache.Size", 50000, viper.SetDefault)
}
