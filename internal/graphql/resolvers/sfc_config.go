package resolvers

import (
	"fantom-api-graphql/internal/repository"
	"fantom-api-graphql/internal/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/sync/singleflight"
)

// SfcConfig represents the SFC contract configuration resolver.
type SfcConfig struct {
	cg *singleflight.Group
}

// SfcConfig resolves the current SFC configuration.
func (rs *rootResolver) SfcConfig() SfcConfig {
	return SfcConfig{cg: new(singleflight.Group)}
}

// getConfig load the configuration from repository.
func (sc SfcConfig) getConfig() (*types.SfcConfig, error) {
	// get the SFC configuration only once
	cfg, err, _ := sc.cg.Do("cfg", func() (interface{}, error) {
		return repository.R().SfcConfiguration()
	})

	// loader failed
	if err != nil {
		return nil, err
	}
	return cfg.(*types.SfcConfig), nil
}

// MinValidatorStake resolves the minimal validator stake in WEI unit.
func (sc SfcConfig) MinValidatorStake() (hexutil.Big, error) {
	c, err := sc.getConfig()
	if err != nil {
		return hexutil.Big{}, err
	}
	return c.MinValidatorStake, nil
}

// MaxDelegatedRatio resolves the ratio between self stake
// and all received stake in 18 digits number multiplier.
func (sc SfcConfig) MaxDelegatedRatio() (hexutil.Big, error) {
	c, err := sc.getConfig()
	if err != nil {
		return hexutil.Big{}, err
	}
	return c.MaxDelegatedRatio, nil
}

// MinLockupDuration resolves the lowest lockup duration allowed.
func (sc SfcConfig) MinLockupDuration() (hexutil.Big, error) {
	c, err := sc.getConfig()
	if err != nil {
		return hexutil.Big{}, err
	}
	return c.MinLockupDuration, nil
}

// MaxLockupDuration resolves the highest lockup duration allowed.
func (sc SfcConfig) MaxLockupDuration() (hexutil.Big, error) {
	c, err := sc.getConfig()
	if err != nil {
		return hexutil.Big{}, err
	}
	return c.MaxLockupDuration, nil
}

// WithdrawalPeriodEpochs resolves the minimal number of epochs allowed
// between un-delegate and withdraw requests.
func (sc SfcConfig) WithdrawalPeriodEpochs() (hexutil.Big, error) {
	c, err := sc.getConfig()
	if err != nil {
		return hexutil.Big{}, err
	}
	return c.WithdrawalPeriodEpochs, nil
}

// WithdrawalPeriodTime resolves the minimal number of seconds allowed
// between un-delegate and withdraw requests.
func (sc SfcConfig) WithdrawalPeriodTime() (hexutil.Big, error) {
	c, err := sc.getConfig()
	if err != nil {
		return hexutil.Big{}, err
	}
	return c.WithdrawalPeriodTime, nil
}
