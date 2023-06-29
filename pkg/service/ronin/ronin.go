package ronin

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/defipod/mochi/pkg/cache"
	"github.com/defipod/mochi/pkg/config"
	"github.com/defipod/mochi/pkg/contract/ronin/axs"
	"github.com/defipod/mochi/pkg/contract/ronin/lp"
	"github.com/defipod/mochi/pkg/contract/ronin/ron"
	"github.com/defipod/mochi/pkg/contract/ronin/validator"
	"github.com/defipod/mochi/pkg/logger"
)

type ronin struct {
	axsstaking *axs.Axs
	ronstaking *ron.Ron
	validator  *validator.Validator

	// key   = staking token address
	// value = lp staking pool instance
	lpStakingPools map[string]*lp.Lp
	cache          cache.Cache
}

func New(cfg *config.Config, cache cache.Cache) (Service, error) {
	client, err := ethclient.Dial(cfg.RpcUrl.Ronin)
	if err != nil {
		return nil, err
	}

	log := logger.NewLogrusLogger()

	axsStakingPool := common.HexToAddress("0x05b0bb3c1c320b280501b86706c3551995bc8571")
	axs, err := axs.NewAxs(axsStakingPool, client)
	if err != nil {
		log.Fatal(err, "failed to init axs staking")
	}

	ronStakingPool := common.HexToAddress("0x545edb750eb8769c868429be9586f5857a768758")
	ron, err := ron.NewRon(ronStakingPool, client)
	if err != nil {
		log.Fatal(err, "failed to init ron staking")
	}

	validatorSet := common.HexToAddress("0x617c5d73662282EA7FfD231E020eCa6D2B0D552f")
	validator, err := validator.NewValidator(validatorSet, client)
	if err != nil {
		log.Fatal(err, "failed to init validator")
	}

	lpStakingAddrs := []string{
		"0xb9072cec557528f81dd25dc474d4d69564956e1e",
		"0xba1c32baff8f23252259a641fd5ca0bd211d4f65",
		"0xd4640c26c1a31cd632d8ae1a96fe5ac135d1eb52",
		"0x14327fa6a4027d8f08c0a1b7feddd178156e9527",
		"0x4e2d6466a53444248272b913c105e9281ec266d8",
		"0x487671acdea3745b6dac3ae8d1757b44a04bfe8a",
	}

	lpStakingPools := make(map[string]*lp.Lp)
	for _, addr := range lpStakingAddrs {
		lp, err := lp.NewLp(common.HexToAddress(addr), client)
		if err != nil {
			log.Fatalf(err, "failed to init lp staking %s", addr)
		}

		stakingToken, err := lp.GetStakingToken(&bind.CallOpts{})
		if err != nil {
			log.Fatalf(err, "failed to get staking token of lp %s", addr)
		}

		lpStakingPools[stakingToken.Hex()] = lp
	}

	return &ronin{
		axsstaking:     axs,
		ronstaking:     ron,
		lpStakingPools: lpStakingPools,
		validator:      validator,
		cache:          cache,
	}, nil
}

var (
	axieStakingAmountKey = "axie-staking-amount"
	axiePendingRewardKey = "axie-pending-reward"
	ronStakingAmountKey  = "ron-staking-amount"
	ronPendingRewardKey  = "ron-pending-reward"
	lpPendingRewardKey   = "lp-pending-reward"
)

func (r *ronin) GetAxsStakingAmount(address string) (float64, error) {
	var amount float64
	cached, err := r.doCacheAxieStakingAmount(address)
	if err == nil && cached != "" {
		defer r.doNetworkAxieStakingAmount(address)
		return amount, json.Unmarshal([]byte(cached), &amount)
	}

	// call network
	return r.doNetworkAxieStakingAmount(address)
}

func (r *ronin) GetAxsPendingRewards(address string) (float64, error) {
	var amount float64
	cached, err := r.doCacheAxiePendingReward(address)
	if err == nil && cached != "" {
		defer r.doNetworkAxiePendingReward(address)
		return amount, json.Unmarshal([]byte(cached), &amount)
	}

	// call network
	return r.doNetworkAxiePendingReward(address)
}

func (r *ronin) GetRonStakingAmount(address string) (float64, error) {
	var amount float64
	cached, err := r.doCacheRonStakingAmount(address)
	if err == nil && cached != "" {
		defer r.doNetworkRonStakingAmount(address)
		return amount, json.Unmarshal([]byte(cached), &amount)
	}

	// call network
	return r.doNetworkRonStakingAmount(address)
}

func (r *ronin) GetRonPendingRewards(address string) (float64, error) {
	var amount float64
	cached, err := r.doCacheRonPendingReward(address)
	if err == nil && cached != "" {
		defer r.doNetworkRonPendingReward(address)
		return amount, json.Unmarshal([]byte(cached), &amount)
	}

	// call network
	return r.doNetworkRonPendingReward(address)
}

func (r *ronin) GetLpPendingRewards(address string) (map[string]LpRewardData, error) {
	var amount map[string]LpRewardData
	cached, err := r.doCacheLpPendingRewards(address)
	if err == nil && cached != "" {
		defer r.doNetworkLpPendingRewards(address)
		return amount, json.Unmarshal([]byte(cached), &amount)
	}

	// call network
	return r.doNetworkLpPendingRewards(address)
}
