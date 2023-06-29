package ronin

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/defipod/mochi/pkg/util"
)

func (r *ronin) doCacheAxieStakingAmount(address string) (string, error) {
	return r.cache.GetString(fmt.Sprintf("%s-%s", axieStakingAmountKey, strings.ToLower(address)))
}

func (r *ronin) doCacheAxiePendingReward(address string) (string, error) {
	return r.cache.GetString(fmt.Sprintf("%s-%s", axiePendingRewardKey, strings.ToLower(address)))
}

func (r *ronin) doCacheRonStakingAmount(address string) (string, error) {
	return r.cache.GetString(fmt.Sprintf("%s-%s", ronStakingAmountKey, strings.ToLower(address)))
}

func (r *ronin) doCacheRonPendingReward(address string) (string, error) {
	return r.cache.GetString(fmt.Sprintf("%s-%s", ronPendingRewardKey, strings.ToLower(address)))
}

func (r *ronin) doCacheLpPendingRewards(address string) (string, error) {
	return r.cache.GetString(fmt.Sprintf("%s-%s", lpPendingRewardKey, strings.ToLower(address)))
}

func (r *ronin) doNetworkAxieStakingAmount(address string) (float64, error) {
	amount, err := r.axsstaking.GetStakingAmount(&bind.CallOpts{}, common.HexToAddress(address))
	if err != nil {
		return 0, nil
	}

	// cache krystal-balance-token-data
	// if error occurs -> ignore
	formatAmount := util.BigIntToFloat(amount, 18)
	bytes, _ := json.Marshal(formatAmount)
	r.cache.Set(axieStakingAmountKey+"-"+strings.ToLower(address), string(bytes), 7*24*time.Hour)

	return util.BigIntToFloat(amount, 18), nil
}

func (r *ronin) doNetworkAxiePendingReward(address string) (float64, error) {
	amount, err := r.axsstaking.GetPendingRewards(&bind.CallOpts{}, common.HexToAddress(address))
	if err != nil {
		return 0, nil
	}

	// cache krystal-balance-token-data
	// if error occurs -> ignore
	formatAmount := util.BigIntToFloat(amount, 18)
	bytes, _ := json.Marshal(formatAmount)
	r.cache.Set(axiePendingRewardKey+"-"+strings.ToLower(address), string(bytes), 7*24*time.Hour)

	return util.BigIntToFloat(amount, 18), nil
}

func (r *ronin) doNetworkRonStakingAmount(address string) (float64, error) {
	validators, err := r.validator.GetValidatorCandidates(&bind.CallOpts{})
	if err != nil {
		return 0, nil
	}

	totalStaking := big.NewInt(0)
	userAddr := common.HexToAddress(address)

	// total RON staking amount = SUM of staking amount by each validator
	for _, v := range validators {
		amount, err := r.ronstaking.GetStakingAmount(&bind.CallOpts{}, v, userAddr)
		if err != nil {
			return 0, nil
		}

		totalStaking.Add(totalStaking, amount)
	}

	// cache krystal-balance-token-data
	// if error occurs -> ignore
	formatTotalStaking := util.BigIntToFloat(totalStaking, 18)
	bytes, _ := json.Marshal(formatTotalStaking)
	r.cache.Set(ronStakingAmountKey+"-"+strings.ToLower(address), string(bytes), 7*24*time.Hour)

	return formatTotalStaking, nil
}

func (r *ronin) doNetworkRonPendingReward(address string) (float64, error) {
	// get all validators
	validators, err := r.validator.GetValidatorCandidates(&bind.CallOpts{})
	if err != nil {
		return 0, nil
	}

	totalReward := big.NewInt(0)
	userAddr := common.HexToAddress(address)

	// total RON staking rewards = SUM of rewards by each validator
	for _, v := range validators {
		amount, err := r.ronstaking.GetReward(&bind.CallOpts{}, v, userAddr)
		if err != nil {
			return 0, nil
		}

		totalReward.Add(totalReward, amount)
	}

	// cache krystal-balance-token-data
	// if error occurs -> ignore
	formatTotalReward := util.BigIntToFloat(totalReward, 18)
	bytes, _ := json.Marshal(formatTotalReward)
	r.cache.Set(ronPendingRewardKey+"-"+strings.ToLower(address), string(bytes), 7*24*time.Hour)

	return formatTotalReward, nil
}

func (r *ronin) doNetworkLpPendingRewards(address string) (map[string]LpRewardData, error) {
	// result data is a map with:
	// - key   = staking token address
	// - value = farming reward data
	result := make(map[string]LpRewardData)
	userAddr := common.HexToAddress(address)

	for stakingToken, p := range r.lpStakingPools {
		reward, err := p.GetPendingRewards(&bind.CallOpts{}, userAddr)
		if err != nil {
			return nil, err
		}

		rewardToken, err := p.GetRewardToken(&bind.CallOpts{})
		if err != nil {
			return nil, err
		}

		result[strings.ToLower(stakingToken)] = LpRewardData{
			RewardToken: rewardToken.Hex(),
			Reward:      util.BigIntToFloat(reward, 18),
		}
	}

	// cache krystal-balance-token-data
	// if error occurs -> ignore
	bytes, _ := json.Marshal(&result)
	r.cache.Set(lpPendingRewardKey+"-"+strings.ToLower(address), string(bytes), 7*24*time.Hour)

	return result, nil
}
