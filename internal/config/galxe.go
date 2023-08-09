package config

import (
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

// galxeSubmitterSettingsYamlKey is a key for galxe submitter settings in config file
const galxeSubmitterSettingsYamlKey = "galxe_submitter"

// GalxeSubmitterSettings is a struct that contains all the settings for the galxe.Submitter
type GalxeSubmitterSettings struct {
	CredentialId int64  `fig:"credential_id"`
	AccessToken  string `fig:"access_token"`
	SkipSubmit   bool   `fig:"skip_submit"`
}

// GalxeSubmitterSettings is a method returning a galxe.Submitter settings from a config file
func (c *config) GalxeSubmitterSettings() GalxeSubmitterSettings {
	return c.galxeSubmitterSettingsOnce.Do(func() any {
		var yamlCfg GalxeSubmitterSettings

		if err := figure.
			Out(&yamlCfg).
			From(kv.MustGetStringMap(c.getter, galxeSubmitterSettingsYamlKey)).
			Please(); err != nil {
			panic(err)
		}

		return yamlCfg
	}).(GalxeSubmitterSettings)
}
