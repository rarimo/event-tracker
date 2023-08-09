package config

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

// eventTrackerSettingsYamlKey is a key for event tracker runner settings in config file
const eventTrackerSettingsYamlKey = "event_tracker"

// EventTrackerSettings is a struct that contains all the settings for the events.EventTracker
type EventTrackerSettings struct {
	ContractAddress common.Address    `fig:"contract_address,required"`
	StartBlock      uint64            `fig:"start_block,required"`
	IterationSize   uint64            `fig:"iteration_size,required"`
	RPC             *ethclient.Client `fig:"rpc,required"`
	RunnerSettings  `fig:"runner,required"`
}

// EventTrackerSettings is a method returning an event tracker runner settings from a config file
func (c *config) EventTrackerSettings() EventTrackerSettings {
	return c.eventTrackerRunnerSettingsOnce.Do(func() any {
		var yamlCfg EventTrackerSettings

		if err := figure.
			Out(&yamlCfg).
			With(figure.BaseHooks, evmHooks).
			From(kv.MustGetStringMap(c.getter, eventTrackerSettingsYamlKey)).
			Please(); err != nil {
			panic(err)
		}

		return yamlCfg
	}).(EventTrackerSettings)
}
