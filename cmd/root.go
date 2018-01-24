package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	bankscfg "github.com/autopilothq/banks/config"

	"github.com/hackathon/journeys/config"
)

var (
	configFile string
	consulURL  string
	vaultURL   string
	prefix     string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:               "journeys",
	Short:             "journeys loves your data",
	Long:              "journeys loves your data",
	PersistentPreRunE: setupAndValidateConfig,
}

// ConfigSource is a human readable indication of the source of
// config (eg file name, consul url, etc)
var ConfigSource string

func init() {

	pflags := RootCmd.PersistentFlags()

	// --config-file, c
	pflags.StringVarP(
		&configFile, "config-file", "c", "",
		"configuration file (eg. /etc/journeys/journeys.yml)",
	)

	// --consul-url
	pflags.StringVar(
		&consulURL, "consul-url", "",
		"consul url (eg. https://consul.somewhere)",
	)

	// --vault-url
	pflags.StringVar(
		&vaultURL, "vault-url", "",
		"vault url (eg. https://vault.somewhere)",
	)

	// --prefix
	pflags.StringVar(
		&prefix, "prefix", "journeys",
		"consul/vault key prefix",
	)

	// --nats
	pflags.StringSlice("nats", []string{}, "URLs of NATS hosts")
	viper.BindPFlag("Banks.NATS.Addrs", pflags.Lookup("nats"))

	// --nats-tls
	pflags.Bool("nats-tls", false, "Enable TLS for connection to NATS")
	viper.BindPFlag("Banks.NATS.TLSEnabled", pflags.Lookup("nats-tls"))

	// --nats-tls-cert
	pflags.String("nats-tls-cert", "",
		"File containing TLS certificate to use for NATS")
	viper.BindPFlag("Banks.NATS.TLSCert", pflags.Lookup("nats-tls-cert"))

	// --nats-tls-key
	pflags.String("nats-tls-key", "",
		"File containing TLS key to use for NATS")
	viper.BindPFlag("Banks.NATS.TLSKey", pflags.Lookup("nats-tls-key"))

	// --nats-tls-ca-cert
	pflags.StringSlice("nats-tls-ca-cert", []string{},
		"File containing TLS CA certificate to use for NATS")
	viper.BindPFlag("Banks.NATS.TLSCaCert", pflags.Lookup("nats-tls-ca-cert"))

}

func setupAndValidateConfig(cmd *cobra.Command, args []string) error {
	var warnings []string
	err := setupConfig()
	if err != nil {
		return err
	}

	if len(args) > 0 {
		// The cluster id is always the first arg
		viper.GetViper().Set("Banks.Cluster.ID", args[0])
	}
	
	warnings, err = bankscfg.Validate(viper.GetViper())
	if err != nil {
		return err
	}
	err = config.SetupLog()
	if err != nil {
		return err
	}
	for _, w := range warnings {
		config.Log().Warn(w)
	}

	return nil
}

func setupConfig() error {
	config.ApplyDefaults()
	// Providing none of configFile, consulURL or vaultURL is valid,
	// and enables localDevMode
	if configFile == "" && consulURL == "" && vaultURL == "" {
		ConfigSource = "localdev"
		return nil
	}

	// Providing a configFile is valid, but only if there is no
	// consulURL or vaultURL
	if configFile != "" {
		if consulURL != "" || vaultURL != "" {
			return errors.New(
				"--config-file may not be combined with --consul-url/--vault-url",
			)
		}

		return readConfigFile()
	}

	// Providing no configFile is valid but only if both
	// consulURL and vaultURL are provided
	if consulURL == "" || vaultURL == "" {
		return errors.New(
			"--consul-url and --vault-url must be provided as a pair",
		)
	}

	return setupRemoteConfig()
}

func readConfigFile() error {
	f, err := os.Open(configFile)
	if err != nil {
		return err
	}

	// TODO do not swallow errors
	defer f.Close()

	viper.SetConfigType("yaml")

	if err := viper.ReadConfig(f); err != nil {
		return err
	}

	ConfigSource = fmt.Sprintf("file: %s", configFile)

	return nil
}

func setupRemoteConfig() error {
	return errors.New("Consul/Vault integration is not yet available")
	// viper.AddRemoteProvider("consul", consulURL, prefix)
}
