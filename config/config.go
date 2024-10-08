package config

import (
	"github.com/spf13/viper"
	"os"
)

// env constants
const (
	PlasmaDaHttpHost = "HTTP_HOST"
	PlasmaDaHttpPort = "HTTP_PORT"
	PlasmaDaType     = "DA_TYPE"
	PlasmaDaId       = "DA_ID"
	PlasmaDaHomeDir  = "HOME_DIR"
	PlasmaEvmRpcUrl  = "EVM_RPC_URL"
	PlasmaKeyFile    = "KEY_FILE"
	PlasmaPassphrase = "PASSPHRASE"
	PlasmaChainId    = "CHAIN_ID"
	PlasmaHubAddr    = "ALT_DA_HUB_ADDR"
)

type App struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	DA           string `json:"da"`
	DaID         string `json:"da_id"`
	HomeDir      string `json:"home_dir"`
	EvmRpcUrl    string `json:"evm_rpc_url"`
	KeyFile      string `json:"key_file"`
	Passphrase   string `json:"passphrase"`
	ChainId      int64  `json:"chain_id"`
	AltDaHubAddr string `json:"alt_da_hub_addr"`
}

func DefaultConfig() App {
	// default app config will read from the environment variables
	// if set by flag, it will override the default values
	homeDir, _ := os.UserHomeDir()
	cfg := App{
		Host:         "localhost",
		Port:         8087,
		DA:           "file",
		DaID:         "0x000c", // it is celestia
		HomeDir:      homeDir,
		EvmRpcUrl:    "https://sepolia.optimism.io",
		KeyFile:      "",
		Passphrase:   "passphrase",
		ChainId:      11155420,
		AltDaHubAddr: "0x865Cb4Ccdc123d5Ac1D8CDC1f967da1Eb111FcAd",
	}

	if homeDir := viper.GetString(PlasmaDaHomeDir); homeDir != "" {
		cfg.HomeDir = homeDir
	}
	if host := viper.GetString(PlasmaDaHttpHost); host != "" {
		cfg.Host = host
	}
	if port := viper.GetInt(PlasmaDaHttpPort); port > 0 {
		cfg.Port = port
	}
	if da := viper.GetString(PlasmaDaType); da != "" {
		cfg.DA = da
	}
	if daId := viper.GetString(PlasmaDaId); daId != "" {
		cfg.DaID = daId
	}
	if evmRpcUrl := viper.GetString(PlasmaEvmRpcUrl); evmRpcUrl != "" {
		cfg.EvmRpcUrl = evmRpcUrl
	}
	if keyFile := viper.GetString(PlasmaKeyFile); keyFile != "" {
		cfg.KeyFile = keyFile
	}
	if passphrase := viper.GetString(PlasmaPassphrase); passphrase != "" {
		cfg.Passphrase = passphrase
	}
	if chainId := viper.GetInt64(PlasmaChainId); chainId > 0 {
		cfg.ChainId = chainId
	}
	if plasmaHubAddr := viper.GetString(PlasmaHubAddr); plasmaHubAddr != "" {
		cfg.AltDaHubAddr = plasmaHubAddr
	}

	return cfg
}

//func init() {
//	viper.AutomaticEnv()
//	viper.SetEnvPrefix(PrefixEnv)
//}
