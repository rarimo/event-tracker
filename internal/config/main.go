package config

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"reflect"
	"strings"
)

type Config interface {
	comfig.Logger
	pgdb.Databaser

	EventTrackerSettings() EventTrackerSettings
	GalxeSubmitterSettings() GalxeSubmitterSettings
}

type config struct {
	comfig.Logger
	pgdb.Databaser
	getter kv.Getter

	eventTrackerRunnerSettingsOnce comfig.Once
	galxeSubmitterSettingsOnce     comfig.Once
}

func New(getter kv.Getter) Config {
	return &config{
		getter:    getter,
		Databaser: pgdb.NewDatabaser(getter),
		Logger:    comfig.NewLogger(getter, comfig.LoggerOpts{}),
	}
}

var evmHooks = figure.Hooks{
	"*ethclient.Client": func(value any) (reflect.Value, error) {
		switch v := value.(type) {
		case string:
			client, err := ethclient.Dial(v)
			if err != nil {
				return reflect.Value{}, errors.Wrap(err, "failed to convert value into ethclient")
			}
			return reflect.ValueOf(client), nil
		default:
			return reflect.Value{}, fmt.Errorf("unsupported conversion from %T", value)
		}
	},
	"common.Address": func(value any) (reflect.Value, error) {
		switch v := value.(type) {
		case string:
			address := common.HexToAddress(strings.ToLower(v))
			return reflect.ValueOf(address), nil
		default:
			return reflect.Value{}, fmt.Errorf("unsupported conversion from %T", value)
		}
	},
}
