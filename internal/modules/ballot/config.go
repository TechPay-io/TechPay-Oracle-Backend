// Package ballot handles closed ballot finalization.
// It implements an oracle which uses TechPay API to collect
// ballot participants' account total and feeds it into
// the smart contract to allow ballot winner calculation.
package ballot

import (
	"encoding/json"
	"io/ioutil"
)

// FinalizingOracleConfig represents a configuration
// of the ballot finalizing oracle module.
type FinalizingOracleConfig struct {
	Name            string `json:"name"`
	ApiUrl          string `json:"api_url"`
	ScanClosedDelay int64  `json:"scan_closed_delay"`
	KeyStore        string `json:"keystore"`
	KeySecret       string `json:"key_secret"`
	ScanActiveDelay int64  `json:"scan_active_delay"`
	ResultsWebHook  string `json:"results_web_hook"`
}

// configuration loads and parses the ballot finalizing oracle module configuration
// from a config JSON file.
func configuration(path string) (*FinalizingOracleConfig, error) {
	// read the config file
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// parse the modules json structure
	cfg := FinalizingOracleConfig{}
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
