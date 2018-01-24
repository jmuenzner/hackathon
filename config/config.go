package config

import (
	"github.com/spf13/viper"

	bankscfg "github.com/autopilothq/banks/config"
)

// Reader returns a bankscfg.KeyedReader to read config values
func Reader() bankscfg.KeyedReader {
	return viper.GetViper()
}

// SecretReader returns a bankscfg.KeyedReader to read secret credentials
func SecretReader() bankscfg.KeyedReader {
	// TODO: use Vault instead of falling back to regular config
	return Reader()
}
