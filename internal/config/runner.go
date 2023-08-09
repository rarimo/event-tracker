package config

import (
	"time"
)

// RunnerSettings is a struct that contains all the settings for the runner
type RunnerSettings struct {
	// Name is a name of the runner
	Name            string `fig:"name,required"`
	BackoffSettings `fig:"backoff,required"`
}

// BackoffSettings is a struct that contains all the settings for the backoff
type BackoffSettings struct {
	// NormalPeriod is a period of time between runner function calls
	NormalPeriod      time.Duration `fig:"normal_period,required"`
	MinAbnormalPeriod time.Duration `fig:"min_abnormal_period,required"`
	MaxAbnormalPeriod time.Duration `fig:"max_abnormal_period,required"`
}
