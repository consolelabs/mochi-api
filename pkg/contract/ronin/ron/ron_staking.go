// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ron

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RonMetaData contains all meta data concerning the Ron contract.
var RonMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"admin\",\"type\":\"address\",\"indexed\":false}],\"name\":\"ErrAdminOfAnyActivePoolForbidden\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrCallerMustBeValidatorContract\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"addr\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"extraInfo\",\"type\":\"string\",\"indexed\":false}],\"name\":\"ErrCannotInitTransferRON\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrCannotTransferRON\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"poolAddr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"ErrInactivePool\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrInsufficientBalance\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrInsufficientDelegatingAmount\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrInsufficientStakingAmount\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrInvalidArrays\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrInvalidCommissionRate\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrInvalidPoolShare\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrOnlyPoolAdminAllowed\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrPoolAdminForbidden\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrRecipientRevert\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrStakingAmountLeft\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrThreeInteractionAddrsNotEqual\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrThreeOperationAddrsNotDistinct\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrUndelegateTooEarly\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrUndelegateZeroAmount\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrUnstakeTooEarly\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrUnstakeZeroAmount\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrZeroCodeContract\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrZeroValue\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"minSecs\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"CooldownSecsToUndelegateUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"consensuAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"Delegated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false}],\"name\":\"Initialized\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"maxRate\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"MaxCommissionRateUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"threshold\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"MinValidatorStakingAmountUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"validator\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"admin\",\"type\":\"address\",\"indexed\":true}],\"name\":\"PoolApproved\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"period\",\"type\":\"uint256\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"poolAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"PoolSharesUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"validator\",\"type\":\"address[]\",\"indexed\":false}],\"name\":\"PoolsDeprecated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"period\",\"type\":\"uint256\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"poolAddrs\",\"type\":\"address[]\",\"indexed\":false}],\"name\":\"PoolsUpdateConflicted\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"period\",\"type\":\"uint256\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"poolAddrs\",\"type\":\"address[]\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"rewards\",\"type\":\"uint256[]\",\"indexed\":false}],\"name\":\"PoolsUpdateFailed\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"period\",\"type\":\"uint256\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"poolAddrs\",\"type\":\"address[]\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"aRps\",\"type\":\"uint256[]\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"shares\",\"type\":\"uint256[]\",\"indexed\":false}],\"name\":\"PoolsUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"poolAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"user\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"RewardClaimed\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"consensuAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"Staked\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"validator\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"contractBalance\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"StakingAmountDeductFailed\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"validator\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"admin\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"contractBalance\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"StakingAmountTransferFailed\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"consensuAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"Undelegated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"consensuAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"Unstaked\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"poolAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"user\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"debited\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"UserRewardUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"name\":\"ValidatorContractUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"secs\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"WaitingSecsToRevokeUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":null,\"name\":\"\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"payable\",\"type\":\"fallback\",\"anonymous\":false},{\"inputs\":[],\"name\":\"DEFAULT_ADDITION_GAS\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"PERIOD_DURATION\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_candidateAdmin\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_consensusAddr\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_treasuryAddr\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_bridgeOperatorAddr\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_commissionRate\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"applyValidatorCandidate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"payable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_consensusAddrs\",\"type\":\"address[]\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_amounts\",\"type\":\"uint256[]\",\"indexed\":false}],\"name\":\"bulkUndelegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_consensusAddrList\",\"type\":\"address[]\",\"indexed\":false}],\"name\":\"claimRewards\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"cooldownSecsToUndelegate\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_consensusAddr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"delegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"payable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_consensusAddrList\",\"type\":\"address[]\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_consensusAddrDst\",\"type\":\"address\",\"indexed\":false}],\"name\":\"delegateRewards\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_consensusAddr\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"execDeductStakingAmount\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_actualDeductingAmount\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_pools\",\"type\":\"address[]\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_newPeriod\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"execDeprecatePools\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_consensusAddrs\",\"type\":\"address[]\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_rewards\",\"type\":\"uint256[]\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_period\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"execRecordRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"payable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_pools\",\"type\":\"address[]\",\"indexed\":false}],\"name\":\"getManySelfStakings\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_selfStakings\",\"type\":\"uint256[]\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_poolAddrs\",\"type\":\"address[]\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_userList\",\"type\":\"address[]\",\"indexed\":false}],\"name\":\"getManyStakingAmounts\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_stakingAmounts\",\"type\":\"uint256[]\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_poolList\",\"type\":\"address[]\",\"indexed\":false}],\"name\":\"getManyStakingTotals\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_stakingAmounts\",\"type\":\"uint256[]\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_poolAdminAddr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"getPoolAddressOf\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_poolAddr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"getPoolDetail\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_admin\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_stakingAmount\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_stakingTotal\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_poolAddr\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_user\",\"type\":\"address\",\"indexed\":false}],\"name\":\"getReward\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_user\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_poolAddrList\",\"type\":\"address[]\",\"indexed\":false}],\"name\":\"getRewards\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_rewards\",\"type\":\"uint256[]\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_poolAddr\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_user\",\"type\":\"address\",\"indexed\":false}],\"name\":\"getStakingAmount\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_poolAddr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"getStakingTotal\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"__validatorContract\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"__minValidatorStakingAmount\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"__maxCommissionRate\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"__cooldownSecsToUndelegate\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"__waitingSecsToRevoke\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_poolAdminAddr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"isAdminOfActivePool\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bool\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"maxCommissionRate\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"minValidatorStakingAmount\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_consensusAddrSrc\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_consensusAddrDst\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"redelegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_consensusAddr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"requestEmergencyExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_consensusAddr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"requestRenounce\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_consensusAddr\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_effectiveDaysOnwards\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_commissionRate\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"requestUpdateCommissionRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_cooldownSecs\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"setCooldownSecsToUndelegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_maxRate\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"setMaxCommissionRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_threshold\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"setMinValidatorStakingAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_addr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"setValidatorContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_secs\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"setWaitingSecsToRevoke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_consensusAddr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"stake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"payable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_consensusAddr\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"undelegate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_consensusAddr\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_amount\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"unstake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"validatorContract\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"waitingSecsToRevoke\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":null,\"name\":\"\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"payable\",\"type\":\"receive\",\"anonymous\":false}]",
}

// RonABI is the input ABI used to generate the binding from.
// Deprecated: Use RonMetaData.ABI instead.
var RonABI = RonMetaData.ABI

// Ron is an auto generated Go binding around an Ethereum contract.
type Ron struct {
	RonCaller     // Read-only binding to the contract
	RonTransactor // Write-only binding to the contract
	RonFilterer   // Log filterer for contract events
}

// RonCaller is an auto generated read-only Go binding around an Ethereum contract.
type RonCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RonTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RonTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RonFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RonFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RonSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RonSession struct {
	Contract     *Ron              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RonCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RonCallerSession struct {
	Contract *RonCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RonTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RonTransactorSession struct {
	Contract     *RonTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RonRaw is an auto generated low-level Go binding around an Ethereum contract.
type RonRaw struct {
	Contract *Ron // Generic contract binding to access the raw methods on
}

// RonCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RonCallerRaw struct {
	Contract *RonCaller // Generic read-only contract binding to access the raw methods on
}

// RonTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RonTransactorRaw struct {
	Contract *RonTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRon creates a new instance of Ron, bound to a specific deployed contract.
func NewRon(address common.Address, backend bind.ContractBackend) (*Ron, error) {
	contract, err := bindRon(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ron{RonCaller: RonCaller{contract: contract}, RonTransactor: RonTransactor{contract: contract}, RonFilterer: RonFilterer{contract: contract}}, nil
}

// NewRonCaller creates a new read-only instance of Ron, bound to a specific deployed contract.
func NewRonCaller(address common.Address, caller bind.ContractCaller) (*RonCaller, error) {
	contract, err := bindRon(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RonCaller{contract: contract}, nil
}

// NewRonTransactor creates a new write-only instance of Ron, bound to a specific deployed contract.
func NewRonTransactor(address common.Address, transactor bind.ContractTransactor) (*RonTransactor, error) {
	contract, err := bindRon(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RonTransactor{contract: contract}, nil
}

// NewRonFilterer creates a new log filterer instance of Ron, bound to a specific deployed contract.
func NewRonFilterer(address common.Address, filterer bind.ContractFilterer) (*RonFilterer, error) {
	contract, err := bindRon(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RonFilterer{contract: contract}, nil
}

// bindRon binds a generic wrapper to an already deployed contract.
func bindRon(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RonABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ron *RonRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ron.Contract.RonCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ron *RonRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ron.Contract.RonTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ron *RonRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ron.Contract.RonTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ron *RonCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ron.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ron *RonTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ron.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ron *RonTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ron.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADDITIONGAS is a free data retrieval call binding the contract method 0x03827884.
//
// Solidity: function DEFAULT_ADDITION_GAS() view returns(uint256)
func (_Ron *RonCaller) DEFAULTADDITIONGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ron.contract.Call(opts, &out, "DEFAULT_ADDITION_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTADDITIONGAS is a free data retrieval call binding the contract method 0x03827884.
//
// Solidity: function DEFAULT_ADDITION_GAS() view returns(uint256)
func (_Ron *RonSession) DEFAULTADDITIONGAS() (*big.Int, error) {
	return _Ron.Contract.DEFAULTADDITIONGAS(&_Ron.CallOpts)
}

// DEFAULTADDITIONGAS is a free data retrieval call binding the contract method 0x03827884.
//
// Solidity: function DEFAULT_ADDITION_GAS() view returns(uint256)
func (_Ron *RonCallerSession) DEFAULTADDITIONGAS() (*big.Int, error) {
	return _Ron.Contract.DEFAULTADDITIONGAS(&_Ron.CallOpts)
}

// PERIODDURATION is a free data retrieval call binding the contract method 0x6558954f.
//
// Solidity: function PERIOD_DURATION() view returns(uint256)
func (_Ron *RonCaller) PERIODDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ron.contract.Call(opts, &out, "PERIOD_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERIODDURATION is a free data retrieval call binding the contract method 0x6558954f.
//
// Solidity: function PERIOD_DURATION() view returns(uint256)
func (_Ron *RonSession) PERIODDURATION() (*big.Int, error) {
	return _Ron.Contract.PERIODDURATION(&_Ron.CallOpts)
}

// PERIODDURATION is a free data retrieval call binding the contract method 0x6558954f.
//
// Solidity: function PERIOD_DURATION() view returns(uint256)
func (_Ron *RonCallerSession) PERIODDURATION() (*big.Int, error) {
	return _Ron.Contract.PERIODDURATION(&_Ron.CallOpts)
}

// CooldownSecsToUndelegate is a free data retrieval call binding the contract method 0x0682e8fa.
//
// Solidity: function cooldownSecsToUndelegate() view returns(uint256)
func (_Ron *RonCaller) CooldownSecsToUndelegate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ron.contract.Call(opts, &out, "cooldownSecsToUndelegate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CooldownSecsToUndelegate is a free data retrieval call binding the contract method 0x0682e8fa.
//
// Solidity: function cooldownSecsToUndelegate() view returns(uint256)
func (_Ron *RonSession) CooldownSecsToUndelegate() (*big.Int, error) {
	return _Ron.Contract.CooldownSecsToUndelegate(&_Ron.CallOpts)
}

// CooldownSecsToUndelegate is a free data retrieval call binding the contract method 0x0682e8fa.
//
// Solidity: function cooldownSecsToUndelegate() view returns(uint256)
func (_Ron *RonCallerSession) CooldownSecsToUndelegate() (*big.Int, error) {
	return _Ron.Contract.CooldownSecsToUndelegate(&_Ron.CallOpts)
}

// GetManySelfStakings is a free data retrieval call binding the contract method 0x42ef3c34.
//
// Solidity: function getManySelfStakings(address[] _pools) view returns(uint256[] _selfStakings)
func (_Ron *RonCaller) GetManySelfStakings(opts *bind.CallOpts, _pools []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Ron.contract.Call(opts, &out, "getManySelfStakings", _pools)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetManySelfStakings is a free data retrieval call binding the contract method 0x42ef3c34.
//
// Solidity: function getManySelfStakings(address[] _pools) view returns(uint256[] _selfStakings)
func (_Ron *RonSession) GetManySelfStakings(_pools []common.Address) ([]*big.Int, error) {
	return _Ron.Contract.GetManySelfStakings(&_Ron.CallOpts, _pools)
}

// GetManySelfStakings is a free data retrieval call binding the contract method 0x42ef3c34.
//
// Solidity: function getManySelfStakings(address[] _pools) view returns(uint256[] _selfStakings)
func (_Ron *RonCallerSession) GetManySelfStakings(_pools []common.Address) ([]*big.Int, error) {
	return _Ron.Contract.GetManySelfStakings(&_Ron.CallOpts, _pools)
}

// GetManyStakingAmounts is a free data retrieval call binding the contract method 0x095f6475.
//
// Solidity: function getManyStakingAmounts(address[] _poolAddrs, address[] _userList) view returns(uint256[] _stakingAmounts)
func (_Ron *RonCaller) GetManyStakingAmounts(opts *bind.CallOpts, _poolAddrs []common.Address, _userList []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Ron.contract.Call(opts, &out, "getManyStakingAmounts", _poolAddrs, _userList)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetManyStakingAmounts is a free data retrieval call binding the contract method 0x095f6475.
//
// Solidity: function getManyStakingAmounts(address[] _poolAddrs, address[] _userList) view returns(uint256[] _stakingAmounts)
func (_Ron *RonSession) GetManyStakingAmounts(_poolAddrs []common.Address, _userList []common.Address) ([]*big.Int, error) {
	return _Ron.Contract.GetManyStakingAmounts(&_Ron.CallOpts, _poolAddrs, _userList)
}

// GetManyStakingAmounts is a free data retrieval call binding the contract method 0x095f6475.
//
// Solidity: function getManyStakingAmounts(address[] _poolAddrs, address[] _userList) view returns(uint256[] _stakingAmounts)
func (_Ron *RonCallerSession) GetManyStakingAmounts(_poolAddrs []common.Address, _userList []common.Address) ([]*big.Int, error) {
	return _Ron.Contract.GetManyStakingAmounts(&_Ron.CallOpts, _poolAddrs, _userList)
}

// GetManyStakingTotals is a free data retrieval call binding the contract method 0x91f8723f.
//
// Solidity: function getManyStakingTotals(address[] _poolList) view returns(uint256[] _stakingAmounts)
func (_Ron *RonCaller) GetManyStakingTotals(opts *bind.CallOpts, _poolList []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Ron.contract.Call(opts, &out, "getManyStakingTotals", _poolList)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetManyStakingTotals is a free data retrieval call binding the contract method 0x91f8723f.
//
// Solidity: function getManyStakingTotals(address[] _poolList) view returns(uint256[] _stakingAmounts)
func (_Ron *RonSession) GetManyStakingTotals(_poolList []common.Address) ([]*big.Int, error) {
	return _Ron.Contract.GetManyStakingTotals(&_Ron.CallOpts, _poolList)
}

// GetManyStakingTotals is a free data retrieval call binding the contract method 0x91f8723f.
//
// Solidity: function getManyStakingTotals(address[] _poolList) view returns(uint256[] _stakingAmounts)
func (_Ron *RonCallerSession) GetManyStakingTotals(_poolList []common.Address) ([]*big.Int, error) {
	return _Ron.Contract.GetManyStakingTotals(&_Ron.CallOpts, _poolList)
}

// GetPoolAddressOf is a free data retrieval call binding the contract method 0xc5087003.
//
// Solidity: function getPoolAddressOf(address _poolAdminAddr) view returns(address)
func (_Ron *RonCaller) GetPoolAddressOf(opts *bind.CallOpts, _poolAdminAddr common.Address) (common.Address, error) {
	var out []interface{}
	err := _Ron.contract.Call(opts, &out, "getPoolAddressOf", _poolAdminAddr)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolAddressOf is a free data retrieval call binding the contract method 0xc5087003.
//
// Solidity: function getPoolAddressOf(address _poolAdminAddr) view returns(address)
func (_Ron *RonSession) GetPoolAddressOf(_poolAdminAddr common.Address) (common.Address, error) {
	return _Ron.Contract.GetPoolAddressOf(&_Ron.CallOpts, _poolAdminAddr)
}

// GetPoolAddressOf is a free data retrieval call binding the contract method 0xc5087003.
//
// Solidity: function getPoolAddressOf(address _poolAdminAddr) view returns(address)
func (_Ron *RonCallerSession) GetPoolAddressOf(_poolAdminAddr common.Address) (common.Address, error) {
	return _Ron.Contract.GetPoolAddressOf(&_Ron.CallOpts, _poolAdminAddr)
}

// GetPoolDetail is a free data retrieval call binding the contract method 0xd01b8eed.
//
// Solidity: function getPoolDetail(address _poolAddr) view returns(address _admin, uint256 _stakingAmount, uint256 _stakingTotal)
func (_Ron *RonCaller) GetPoolDetail(opts *bind.CallOpts, _poolAddr common.Address) (struct {
	Admin         common.Address
	StakingAmount *big.Int
	StakingTotal  *big.Int
}, error) {
	var out []interface{}
	err := _Ron.contract.Call(opts, &out, "getPoolDetail", _poolAddr)

	outstruct := new(struct {
		Admin         common.Address
		StakingAmount *big.Int
		StakingTotal  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Admin = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StakingAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StakingTotal = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPoolDetail is a free data retrieval call binding the contract method 0xd01b8eed.
//
// Solidity: function getPoolDetail(address _poolAddr) view returns(address _admin, uint256 _stakingAmount, uint256 _stakingTotal)
func (_Ron *RonSession) GetPoolDetail(_poolAddr common.Address) (struct {
	Admin         common.Address
	StakingAmount *big.Int
	StakingTotal  *big.Int
}, error) {
	return _Ron.Contract.GetPoolDetail(&_Ron.CallOpts, _poolAddr)
}

// GetPoolDetail is a free data retrieval call binding the contract method 0xd01b8eed.
//
// Solidity: function getPoolDetail(address _poolAddr) view returns(address _admin, uint256 _stakingAmount, uint256 _stakingTotal)
func (_Ron *RonCallerSession) GetPoolDetail(_poolAddr common.Address) (struct {
	Admin         common.Address
	StakingAmount *big.Int
	StakingTotal  *big.Int
}, error) {
	return _Ron.Contract.GetPoolDetail(&_Ron.CallOpts, _poolAddr)
}

// GetReward is a free data retrieval call binding the contract method 0x6b091695.
//
// Solidity: function getReward(address _poolAddr, address _user) view returns(uint256)
func (_Ron *RonCaller) GetReward(opts *bind.CallOpts, _poolAddr common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ron.contract.Call(opts, &out, "getReward", _poolAddr, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReward is a free data retrieval call binding the contract method 0x6b091695.
//
// Solidity: function getReward(address _poolAddr, address _user) view returns(uint256)
func (_Ron *RonSession) GetReward(_poolAddr common.Address, _user common.Address) (*big.Int, error) {
	return _Ron.Contract.GetReward(&_Ron.CallOpts, _poolAddr, _user)
}

// GetReward is a free data retrieval call binding the contract method 0x6b091695.
//
// Solidity: function getReward(address _poolAddr, address _user) view returns(uint256)
func (_Ron *RonCallerSession) GetReward(_poolAddr common.Address, _user common.Address) (*big.Int, error) {
	return _Ron.Contract.GetReward(&_Ron.CallOpts, _poolAddr, _user)
}

// GetRewards is a free data retrieval call binding the contract method 0x3d8e846e.
//
// Solidity: function getRewards(address _user, address[] _poolAddrList) view returns(uint256[] _rewards)
func (_Ron *RonCaller) GetRewards(opts *bind.CallOpts, _user common.Address, _poolAddrList []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Ron.contract.Call(opts, &out, "getRewards", _user, _poolAddrList)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetRewards is a free data retrieval call binding the contract method 0x3d8e846e.
//
// Solidity: function getRewards(address _user, address[] _poolAddrList) view returns(uint256[] _rewards)
func (_Ron *RonSession) GetRewards(_user common.Address, _poolAddrList []common.Address) ([]*big.Int, error) {
	return _Ron.Contract.GetRewards(&_Ron.CallOpts, _user, _poolAddrList)
}

// GetRewards is a free data retrieval call binding the contract method 0x3d8e846e.
//
// Solidity: function getRewards(address _user, address[] _poolAddrList) view returns(uint256[] _rewards)
func (_Ron *RonCallerSession) GetRewards(_user common.Address, _poolAddrList []common.Address) ([]*big.Int, error) {
	return _Ron.Contract.GetRewards(&_Ron.CallOpts, _user, _poolAddrList)
}

// GetStakingAmount is a free data retrieval call binding the contract method 0x76664b65.
//
// Solidity: function getStakingAmount(address _poolAddr, address _user) view returns(uint256)
func (_Ron *RonCaller) GetStakingAmount(opts *bind.CallOpts, _poolAddr common.Address, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ron.contract.Call(opts, &out, "getStakingAmount", _poolAddr, _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingAmount is a free data retrieval call binding the contract method 0x76664b65.
//
// Solidity: function getStakingAmount(address _poolAddr, address _user) view returns(uint256)
func (_Ron *RonSession) GetStakingAmount(_poolAddr common.Address, _user common.Address) (*big.Int, error) {
	return _Ron.Contract.GetStakingAmount(&_Ron.CallOpts, _poolAddr, _user)
}

// GetStakingAmount is a free data retrieval call binding the contract method 0x76664b65.
//
// Solidity: function getStakingAmount(address _poolAddr, address _user) view returns(uint256)
func (_Ron *RonCallerSession) GetStakingAmount(_poolAddr common.Address, _user common.Address) (*big.Int, error) {
	return _Ron.Contract.GetStakingAmount(&_Ron.CallOpts, _poolAddr, _user)
}

// GetStakingTotal is a free data retrieval call binding the contract method 0x895ab742.
//
// Solidity: function getStakingTotal(address _poolAddr) view returns(uint256)
func (_Ron *RonCaller) GetStakingTotal(opts *bind.CallOpts, _poolAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ron.contract.Call(opts, &out, "getStakingTotal", _poolAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingTotal is a free data retrieval call binding the contract method 0x895ab742.
//
// Solidity: function getStakingTotal(address _poolAddr) view returns(uint256)
func (_Ron *RonSession) GetStakingTotal(_poolAddr common.Address) (*big.Int, error) {
	return _Ron.Contract.GetStakingTotal(&_Ron.CallOpts, _poolAddr)
}

// GetStakingTotal is a free data retrieval call binding the contract method 0x895ab742.
//
// Solidity: function getStakingTotal(address _poolAddr) view returns(uint256)
func (_Ron *RonCallerSession) GetStakingTotal(_poolAddr common.Address) (*big.Int, error) {
	return _Ron.Contract.GetStakingTotal(&_Ron.CallOpts, _poolAddr)
}

// IsAdminOfActivePool is a free data retrieval call binding the contract method 0x42e0c408.
//
// Solidity: function isAdminOfActivePool(address _poolAdminAddr) view returns(bool)
func (_Ron *RonCaller) IsAdminOfActivePool(opts *bind.CallOpts, _poolAdminAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Ron.contract.Call(opts, &out, "isAdminOfActivePool", _poolAdminAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdminOfActivePool is a free data retrieval call binding the contract method 0x42e0c408.
//
// Solidity: function isAdminOfActivePool(address _poolAdminAddr) view returns(bool)
func (_Ron *RonSession) IsAdminOfActivePool(_poolAdminAddr common.Address) (bool, error) {
	return _Ron.Contract.IsAdminOfActivePool(&_Ron.CallOpts, _poolAdminAddr)
}

// IsAdminOfActivePool is a free data retrieval call binding the contract method 0x42e0c408.
//
// Solidity: function isAdminOfActivePool(address _poolAdminAddr) view returns(bool)
func (_Ron *RonCallerSession) IsAdminOfActivePool(_poolAdminAddr common.Address) (bool, error) {
	return _Ron.Contract.IsAdminOfActivePool(&_Ron.CallOpts, _poolAdminAddr)
}

// MaxCommissionRate is a free data retrieval call binding the contract method 0xc673316c.
//
// Solidity: function maxCommissionRate() view returns(uint256)
func (_Ron *RonCaller) MaxCommissionRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ron.contract.Call(opts, &out, "maxCommissionRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxCommissionRate is a free data retrieval call binding the contract method 0xc673316c.
//
// Solidity: function maxCommissionRate() view returns(uint256)
func (_Ron *RonSession) MaxCommissionRate() (*big.Int, error) {
	return _Ron.Contract.MaxCommissionRate(&_Ron.CallOpts)
}

// MaxCommissionRate is a free data retrieval call binding the contract method 0xc673316c.
//
// Solidity: function maxCommissionRate() view returns(uint256)
func (_Ron *RonCallerSession) MaxCommissionRate() (*big.Int, error) {
	return _Ron.Contract.MaxCommissionRate(&_Ron.CallOpts)
}

// MinValidatorStakingAmount is a free data retrieval call binding the contract method 0x909791dd.
//
// Solidity: function minValidatorStakingAmount() view returns(uint256)
func (_Ron *RonCaller) MinValidatorStakingAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ron.contract.Call(opts, &out, "minValidatorStakingAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinValidatorStakingAmount is a free data retrieval call binding the contract method 0x909791dd.
//
// Solidity: function minValidatorStakingAmount() view returns(uint256)
func (_Ron *RonSession) MinValidatorStakingAmount() (*big.Int, error) {
	return _Ron.Contract.MinValidatorStakingAmount(&_Ron.CallOpts)
}

// MinValidatorStakingAmount is a free data retrieval call binding the contract method 0x909791dd.
//
// Solidity: function minValidatorStakingAmount() view returns(uint256)
func (_Ron *RonCallerSession) MinValidatorStakingAmount() (*big.Int, error) {
	return _Ron.Contract.MinValidatorStakingAmount(&_Ron.CallOpts)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_Ron *RonCaller) ValidatorContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ron.contract.Call(opts, &out, "validatorContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_Ron *RonSession) ValidatorContract() (common.Address, error) {
	return _Ron.Contract.ValidatorContract(&_Ron.CallOpts)
}

// ValidatorContract is a free data retrieval call binding the contract method 0x99439089.
//
// Solidity: function validatorContract() view returns(address)
func (_Ron *RonCallerSession) ValidatorContract() (common.Address, error) {
	return _Ron.Contract.ValidatorContract(&_Ron.CallOpts)
}

// WaitingSecsToRevoke is a free data retrieval call binding the contract method 0xaf245429.
//
// Solidity: function waitingSecsToRevoke() view returns(uint256)
func (_Ron *RonCaller) WaitingSecsToRevoke(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ron.contract.Call(opts, &out, "waitingSecsToRevoke")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WaitingSecsToRevoke is a free data retrieval call binding the contract method 0xaf245429.
//
// Solidity: function waitingSecsToRevoke() view returns(uint256)
func (_Ron *RonSession) WaitingSecsToRevoke() (*big.Int, error) {
	return _Ron.Contract.WaitingSecsToRevoke(&_Ron.CallOpts)
}

// WaitingSecsToRevoke is a free data retrieval call binding the contract method 0xaf245429.
//
// Solidity: function waitingSecsToRevoke() view returns(uint256)
func (_Ron *RonCallerSession) WaitingSecsToRevoke() (*big.Int, error) {
	return _Ron.Contract.WaitingSecsToRevoke(&_Ron.CallOpts)
}

// ApplyValidatorCandidate is a paid mutator transaction binding the contract method 0xe5376f54.
//
// Solidity: function applyValidatorCandidate(address _candidateAdmin, address _consensusAddr, address _treasuryAddr, address _bridgeOperatorAddr, uint256 _commissionRate) payable returns()
func (_Ron *RonTransactor) ApplyValidatorCandidate(opts *bind.TransactOpts, _candidateAdmin common.Address, _consensusAddr common.Address, _treasuryAddr common.Address, _bridgeOperatorAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Ron.contract.Transact(opts, "applyValidatorCandidate", _candidateAdmin, _consensusAddr, _treasuryAddr, _bridgeOperatorAddr, _commissionRate)
}

// ApplyValidatorCandidate is a paid mutator transaction binding the contract method 0xe5376f54.
//
// Solidity: function applyValidatorCandidate(address _candidateAdmin, address _consensusAddr, address _treasuryAddr, address _bridgeOperatorAddr, uint256 _commissionRate) payable returns()
func (_Ron *RonSession) ApplyValidatorCandidate(_candidateAdmin common.Address, _consensusAddr common.Address, _treasuryAddr common.Address, _bridgeOperatorAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.ApplyValidatorCandidate(&_Ron.TransactOpts, _candidateAdmin, _consensusAddr, _treasuryAddr, _bridgeOperatorAddr, _commissionRate)
}

// ApplyValidatorCandidate is a paid mutator transaction binding the contract method 0xe5376f54.
//
// Solidity: function applyValidatorCandidate(address _candidateAdmin, address _consensusAddr, address _treasuryAddr, address _bridgeOperatorAddr, uint256 _commissionRate) payable returns()
func (_Ron *RonTransactorSession) ApplyValidatorCandidate(_candidateAdmin common.Address, _consensusAddr common.Address, _treasuryAddr common.Address, _bridgeOperatorAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.ApplyValidatorCandidate(&_Ron.TransactOpts, _candidateAdmin, _consensusAddr, _treasuryAddr, _bridgeOperatorAddr, _commissionRate)
}

// BulkUndelegate is a paid mutator transaction binding the contract method 0x9488e4e9.
//
// Solidity: function bulkUndelegate(address[] _consensusAddrs, uint256[] _amounts) returns()
func (_Ron *RonTransactor) BulkUndelegate(opts *bind.TransactOpts, _consensusAddrs []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Ron.contract.Transact(opts, "bulkUndelegate", _consensusAddrs, _amounts)
}

// BulkUndelegate is a paid mutator transaction binding the contract method 0x9488e4e9.
//
// Solidity: function bulkUndelegate(address[] _consensusAddrs, uint256[] _amounts) returns()
func (_Ron *RonSession) BulkUndelegate(_consensusAddrs []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Ron.Contract.BulkUndelegate(&_Ron.TransactOpts, _consensusAddrs, _amounts)
}

// BulkUndelegate is a paid mutator transaction binding the contract method 0x9488e4e9.
//
// Solidity: function bulkUndelegate(address[] _consensusAddrs, uint256[] _amounts) returns()
func (_Ron *RonTransactorSession) BulkUndelegate(_consensusAddrs []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Ron.Contract.BulkUndelegate(&_Ron.TransactOpts, _consensusAddrs, _amounts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xf9f031df.
//
// Solidity: function claimRewards(address[] _consensusAddrList) returns(uint256 _amount)
func (_Ron *RonTransactor) ClaimRewards(opts *bind.TransactOpts, _consensusAddrList []common.Address) (*types.Transaction, error) {
	return _Ron.contract.Transact(opts, "claimRewards", _consensusAddrList)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xf9f031df.
//
// Solidity: function claimRewards(address[] _consensusAddrList) returns(uint256 _amount)
func (_Ron *RonSession) ClaimRewards(_consensusAddrList []common.Address) (*types.Transaction, error) {
	return _Ron.Contract.ClaimRewards(&_Ron.TransactOpts, _consensusAddrList)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xf9f031df.
//
// Solidity: function claimRewards(address[] _consensusAddrList) returns(uint256 _amount)
func (_Ron *RonTransactorSession) ClaimRewards(_consensusAddrList []common.Address) (*types.Transaction, error) {
	return _Ron.Contract.ClaimRewards(&_Ron.TransactOpts, _consensusAddrList)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address _consensusAddr) payable returns()
func (_Ron *RonTransactor) Delegate(opts *bind.TransactOpts, _consensusAddr common.Address) (*types.Transaction, error) {
	return _Ron.contract.Transact(opts, "delegate", _consensusAddr)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address _consensusAddr) payable returns()
func (_Ron *RonSession) Delegate(_consensusAddr common.Address) (*types.Transaction, error) {
	return _Ron.Contract.Delegate(&_Ron.TransactOpts, _consensusAddr)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address _consensusAddr) payable returns()
func (_Ron *RonTransactorSession) Delegate(_consensusAddr common.Address) (*types.Transaction, error) {
	return _Ron.Contract.Delegate(&_Ron.TransactOpts, _consensusAddr)
}

// DelegateRewards is a paid mutator transaction binding the contract method 0x097e4a9d.
//
// Solidity: function delegateRewards(address[] _consensusAddrList, address _consensusAddrDst) returns(uint256 _amount)
func (_Ron *RonTransactor) DelegateRewards(opts *bind.TransactOpts, _consensusAddrList []common.Address, _consensusAddrDst common.Address) (*types.Transaction, error) {
	return _Ron.contract.Transact(opts, "delegateRewards", _consensusAddrList, _consensusAddrDst)
}

// DelegateRewards is a paid mutator transaction binding the contract method 0x097e4a9d.
//
// Solidity: function delegateRewards(address[] _consensusAddrList, address _consensusAddrDst) returns(uint256 _amount)
func (_Ron *RonSession) DelegateRewards(_consensusAddrList []common.Address, _consensusAddrDst common.Address) (*types.Transaction, error) {
	return _Ron.Contract.DelegateRewards(&_Ron.TransactOpts, _consensusAddrList, _consensusAddrDst)
}

// DelegateRewards is a paid mutator transaction binding the contract method 0x097e4a9d.
//
// Solidity: function delegateRewards(address[] _consensusAddrList, address _consensusAddrDst) returns(uint256 _amount)
func (_Ron *RonTransactorSession) DelegateRewards(_consensusAddrList []common.Address, _consensusAddrDst common.Address) (*types.Transaction, error) {
	return _Ron.Contract.DelegateRewards(&_Ron.TransactOpts, _consensusAddrList, _consensusAddrDst)
}

// ExecDeductStakingAmount is a paid mutator transaction binding the contract method 0x2715805e.
//
// Solidity: function execDeductStakingAmount(address _consensusAddr, uint256 _amount) returns(uint256 _actualDeductingAmount)
func (_Ron *RonTransactor) ExecDeductStakingAmount(opts *bind.TransactOpts, _consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ron.contract.Transact(opts, "execDeductStakingAmount", _consensusAddr, _amount)
}

// ExecDeductStakingAmount is a paid mutator transaction binding the contract method 0x2715805e.
//
// Solidity: function execDeductStakingAmount(address _consensusAddr, uint256 _amount) returns(uint256 _actualDeductingAmount)
func (_Ron *RonSession) ExecDeductStakingAmount(_consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.ExecDeductStakingAmount(&_Ron.TransactOpts, _consensusAddr, _amount)
}

// ExecDeductStakingAmount is a paid mutator transaction binding the contract method 0x2715805e.
//
// Solidity: function execDeductStakingAmount(address _consensusAddr, uint256 _amount) returns(uint256 _actualDeductingAmount)
func (_Ron *RonTransactorSession) ExecDeductStakingAmount(_consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.ExecDeductStakingAmount(&_Ron.TransactOpts, _consensusAddr, _amount)
}

// ExecDeprecatePools is a paid mutator transaction binding the contract method 0xe22d1c9d.
//
// Solidity: function execDeprecatePools(address[] _pools, uint256 _newPeriod) returns()
func (_Ron *RonTransactor) ExecDeprecatePools(opts *bind.TransactOpts, _pools []common.Address, _newPeriod *big.Int) (*types.Transaction, error) {
	return _Ron.contract.Transact(opts, "execDeprecatePools", _pools, _newPeriod)
}

// ExecDeprecatePools is a paid mutator transaction binding the contract method 0xe22d1c9d.
//
// Solidity: function execDeprecatePools(address[] _pools, uint256 _newPeriod) returns()
func (_Ron *RonSession) ExecDeprecatePools(_pools []common.Address, _newPeriod *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.ExecDeprecatePools(&_Ron.TransactOpts, _pools, _newPeriod)
}

// ExecDeprecatePools is a paid mutator transaction binding the contract method 0xe22d1c9d.
//
// Solidity: function execDeprecatePools(address[] _pools, uint256 _newPeriod) returns()
func (_Ron *RonTransactorSession) ExecDeprecatePools(_pools []common.Address, _newPeriod *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.ExecDeprecatePools(&_Ron.TransactOpts, _pools, _newPeriod)
}

// ExecRecordRewards is a paid mutator transaction binding the contract method 0xacd79c46.
//
// Solidity: function execRecordRewards(address[] _consensusAddrs, uint256[] _rewards, uint256 _period) payable returns()
func (_Ron *RonTransactor) ExecRecordRewards(opts *bind.TransactOpts, _consensusAddrs []common.Address, _rewards []*big.Int, _period *big.Int) (*types.Transaction, error) {
	return _Ron.contract.Transact(opts, "execRecordRewards", _consensusAddrs, _rewards, _period)
}

// ExecRecordRewards is a paid mutator transaction binding the contract method 0xacd79c46.
//
// Solidity: function execRecordRewards(address[] _consensusAddrs, uint256[] _rewards, uint256 _period) payable returns()
func (_Ron *RonSession) ExecRecordRewards(_consensusAddrs []common.Address, _rewards []*big.Int, _period *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.ExecRecordRewards(&_Ron.TransactOpts, _consensusAddrs, _rewards, _period)
}

// ExecRecordRewards is a paid mutator transaction binding the contract method 0xacd79c46.
//
// Solidity: function execRecordRewards(address[] _consensusAddrs, uint256[] _rewards, uint256 _period) payable returns()
func (_Ron *RonTransactorSession) ExecRecordRewards(_consensusAddrs []common.Address, _rewards []*big.Int, _period *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.ExecRecordRewards(&_Ron.TransactOpts, _consensusAddrs, _rewards, _period)
}

// Initialize is a paid mutator transaction binding the contract method 0xf92ad219.
//
// Solidity: function initialize(address __validatorContract, uint256 __minValidatorStakingAmount, uint256 __maxCommissionRate, uint256 __cooldownSecsToUndelegate, uint256 __waitingSecsToRevoke) returns()
func (_Ron *RonTransactor) Initialize(opts *bind.TransactOpts, __validatorContract common.Address, __minValidatorStakingAmount *big.Int, __maxCommissionRate *big.Int, __cooldownSecsToUndelegate *big.Int, __waitingSecsToRevoke *big.Int) (*types.Transaction, error) {
	return _Ron.contract.Transact(opts, "initialize", __validatorContract, __minValidatorStakingAmount, __maxCommissionRate, __cooldownSecsToUndelegate, __waitingSecsToRevoke)
}

// Initialize is a paid mutator transaction binding the contract method 0xf92ad219.
//
// Solidity: function initialize(address __validatorContract, uint256 __minValidatorStakingAmount, uint256 __maxCommissionRate, uint256 __cooldownSecsToUndelegate, uint256 __waitingSecsToRevoke) returns()
func (_Ron *RonSession) Initialize(__validatorContract common.Address, __minValidatorStakingAmount *big.Int, __maxCommissionRate *big.Int, __cooldownSecsToUndelegate *big.Int, __waitingSecsToRevoke *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.Initialize(&_Ron.TransactOpts, __validatorContract, __minValidatorStakingAmount, __maxCommissionRate, __cooldownSecsToUndelegate, __waitingSecsToRevoke)
}

// Initialize is a paid mutator transaction binding the contract method 0xf92ad219.
//
// Solidity: function initialize(address __validatorContract, uint256 __minValidatorStakingAmount, uint256 __maxCommissionRate, uint256 __cooldownSecsToUndelegate, uint256 __waitingSecsToRevoke) returns()
func (_Ron *RonTransactorSession) Initialize(__validatorContract common.Address, __minValidatorStakingAmount *big.Int, __maxCommissionRate *big.Int, __cooldownSecsToUndelegate *big.Int, __waitingSecsToRevoke *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.Initialize(&_Ron.TransactOpts, __validatorContract, __minValidatorStakingAmount, __maxCommissionRate, __cooldownSecsToUndelegate, __waitingSecsToRevoke)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address _consensusAddrSrc, address _consensusAddrDst, uint256 _amount) returns()
func (_Ron *RonTransactor) Redelegate(opts *bind.TransactOpts, _consensusAddrSrc common.Address, _consensusAddrDst common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ron.contract.Transact(opts, "redelegate", _consensusAddrSrc, _consensusAddrDst, _amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address _consensusAddrSrc, address _consensusAddrDst, uint256 _amount) returns()
func (_Ron *RonSession) Redelegate(_consensusAddrSrc common.Address, _consensusAddrDst common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.Redelegate(&_Ron.TransactOpts, _consensusAddrSrc, _consensusAddrDst, _amount)
}

// Redelegate is a paid mutator transaction binding the contract method 0x6bd8f804.
//
// Solidity: function redelegate(address _consensusAddrSrc, address _consensusAddrDst, uint256 _amount) returns()
func (_Ron *RonTransactorSession) Redelegate(_consensusAddrSrc common.Address, _consensusAddrDst common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.Redelegate(&_Ron.TransactOpts, _consensusAddrSrc, _consensusAddrDst, _amount)
}

// RequestEmergencyExit is a paid mutator transaction binding the contract method 0xaa15a6fd.
//
// Solidity: function requestEmergencyExit(address _consensusAddr) returns()
func (_Ron *RonTransactor) RequestEmergencyExit(opts *bind.TransactOpts, _consensusAddr common.Address) (*types.Transaction, error) {
	return _Ron.contract.Transact(opts, "requestEmergencyExit", _consensusAddr)
}

// RequestEmergencyExit is a paid mutator transaction binding the contract method 0xaa15a6fd.
//
// Solidity: function requestEmergencyExit(address _consensusAddr) returns()
func (_Ron *RonSession) RequestEmergencyExit(_consensusAddr common.Address) (*types.Transaction, error) {
	return _Ron.Contract.RequestEmergencyExit(&_Ron.TransactOpts, _consensusAddr)
}

// RequestEmergencyExit is a paid mutator transaction binding the contract method 0xaa15a6fd.
//
// Solidity: function requestEmergencyExit(address _consensusAddr) returns()
func (_Ron *RonTransactorSession) RequestEmergencyExit(_consensusAddr common.Address) (*types.Transaction, error) {
	return _Ron.Contract.RequestEmergencyExit(&_Ron.TransactOpts, _consensusAddr)
}

// RequestRenounce is a paid mutator transaction binding the contract method 0x1658c86e.
//
// Solidity: function requestRenounce(address _consensusAddr) returns()
func (_Ron *RonTransactor) RequestRenounce(opts *bind.TransactOpts, _consensusAddr common.Address) (*types.Transaction, error) {
	return _Ron.contract.Transact(opts, "requestRenounce", _consensusAddr)
}

// RequestRenounce is a paid mutator transaction binding the contract method 0x1658c86e.
//
// Solidity: function requestRenounce(address _consensusAddr) returns()
func (_Ron *RonSession) RequestRenounce(_consensusAddr common.Address) (*types.Transaction, error) {
	return _Ron.Contract.RequestRenounce(&_Ron.TransactOpts, _consensusAddr)
}

// RequestRenounce is a paid mutator transaction binding the contract method 0x1658c86e.
//
// Solidity: function requestRenounce(address _consensusAddr) returns()
func (_Ron *RonTransactorSession) RequestRenounce(_consensusAddr common.Address) (*types.Transaction, error) {
	return _Ron.Contract.RequestRenounce(&_Ron.TransactOpts, _consensusAddr)
}

// RequestUpdateCommissionRate is a paid mutator transaction binding the contract method 0x924f081e.
//
// Solidity: function requestUpdateCommissionRate(address _consensusAddr, uint256 _effectiveDaysOnwards, uint256 _commissionRate) returns()
func (_Ron *RonTransactor) RequestUpdateCommissionRate(opts *bind.TransactOpts, _consensusAddr common.Address, _effectiveDaysOnwards *big.Int, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Ron.contract.Transact(opts, "requestUpdateCommissionRate", _consensusAddr, _effectiveDaysOnwards, _commissionRate)
}

// RequestUpdateCommissionRate is a paid mutator transaction binding the contract method 0x924f081e.
//
// Solidity: function requestUpdateCommissionRate(address _consensusAddr, uint256 _effectiveDaysOnwards, uint256 _commissionRate) returns()
func (_Ron *RonSession) RequestUpdateCommissionRate(_consensusAddr common.Address, _effectiveDaysOnwards *big.Int, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.RequestUpdateCommissionRate(&_Ron.TransactOpts, _consensusAddr, _effectiveDaysOnwards, _commissionRate)
}

// RequestUpdateCommissionRate is a paid mutator transaction binding the contract method 0x924f081e.
//
// Solidity: function requestUpdateCommissionRate(address _consensusAddr, uint256 _effectiveDaysOnwards, uint256 _commissionRate) returns()
func (_Ron *RonTransactorSession) RequestUpdateCommissionRate(_consensusAddr common.Address, _effectiveDaysOnwards *big.Int, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.RequestUpdateCommissionRate(&_Ron.TransactOpts, _consensusAddr, _effectiveDaysOnwards, _commissionRate)
}

// SetCooldownSecsToUndelegate is a paid mutator transaction binding the contract method 0x888b9ae9.
//
// Solidity: function setCooldownSecsToUndelegate(uint256 _cooldownSecs) returns()
func (_Ron *RonTransactor) SetCooldownSecsToUndelegate(opts *bind.TransactOpts, _cooldownSecs *big.Int) (*types.Transaction, error) {
	return _Ron.contract.Transact(opts, "setCooldownSecsToUndelegate", _cooldownSecs)
}

// SetCooldownSecsToUndelegate is a paid mutator transaction binding the contract method 0x888b9ae9.
//
// Solidity: function setCooldownSecsToUndelegate(uint256 _cooldownSecs) returns()
func (_Ron *RonSession) SetCooldownSecsToUndelegate(_cooldownSecs *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.SetCooldownSecsToUndelegate(&_Ron.TransactOpts, _cooldownSecs)
}

// SetCooldownSecsToUndelegate is a paid mutator transaction binding the contract method 0x888b9ae9.
//
// Solidity: function setCooldownSecsToUndelegate(uint256 _cooldownSecs) returns()
func (_Ron *RonTransactorSession) SetCooldownSecsToUndelegate(_cooldownSecs *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.SetCooldownSecsToUndelegate(&_Ron.TransactOpts, _cooldownSecs)
}

// SetMaxCommissionRate is a paid mutator transaction binding the contract method 0xb78b5e41.
//
// Solidity: function setMaxCommissionRate(uint256 _maxRate) returns()
func (_Ron *RonTransactor) SetMaxCommissionRate(opts *bind.TransactOpts, _maxRate *big.Int) (*types.Transaction, error) {
	return _Ron.contract.Transact(opts, "setMaxCommissionRate", _maxRate)
}

// SetMaxCommissionRate is a paid mutator transaction binding the contract method 0xb78b5e41.
//
// Solidity: function setMaxCommissionRate(uint256 _maxRate) returns()
func (_Ron *RonSession) SetMaxCommissionRate(_maxRate *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.SetMaxCommissionRate(&_Ron.TransactOpts, _maxRate)
}

// SetMaxCommissionRate is a paid mutator transaction binding the contract method 0xb78b5e41.
//
// Solidity: function setMaxCommissionRate(uint256 _maxRate) returns()
func (_Ron *RonTransactorSession) SetMaxCommissionRate(_maxRate *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.SetMaxCommissionRate(&_Ron.TransactOpts, _maxRate)
}

// SetMinValidatorStakingAmount is a paid mutator transaction binding the contract method 0x679a6e43.
//
// Solidity: function setMinValidatorStakingAmount(uint256 _threshold) returns()
func (_Ron *RonTransactor) SetMinValidatorStakingAmount(opts *bind.TransactOpts, _threshold *big.Int) (*types.Transaction, error) {
	return _Ron.contract.Transact(opts, "setMinValidatorStakingAmount", _threshold)
}

// SetMinValidatorStakingAmount is a paid mutator transaction binding the contract method 0x679a6e43.
//
// Solidity: function setMinValidatorStakingAmount(uint256 _threshold) returns()
func (_Ron *RonSession) SetMinValidatorStakingAmount(_threshold *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.SetMinValidatorStakingAmount(&_Ron.TransactOpts, _threshold)
}

// SetMinValidatorStakingAmount is a paid mutator transaction binding the contract method 0x679a6e43.
//
// Solidity: function setMinValidatorStakingAmount(uint256 _threshold) returns()
func (_Ron *RonTransactorSession) SetMinValidatorStakingAmount(_threshold *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.SetMinValidatorStakingAmount(&_Ron.TransactOpts, _threshold)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_Ron *RonTransactor) SetValidatorContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Ron.contract.Transact(opts, "setValidatorContract", _addr)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_Ron *RonSession) SetValidatorContract(_addr common.Address) (*types.Transaction, error) {
	return _Ron.Contract.SetValidatorContract(&_Ron.TransactOpts, _addr)
}

// SetValidatorContract is a paid mutator transaction binding the contract method 0xcdf64a76.
//
// Solidity: function setValidatorContract(address _addr) returns()
func (_Ron *RonTransactorSession) SetValidatorContract(_addr common.Address) (*types.Transaction, error) {
	return _Ron.Contract.SetValidatorContract(&_Ron.TransactOpts, _addr)
}

// SetWaitingSecsToRevoke is a paid mutator transaction binding the contract method 0x969ffc14.
//
// Solidity: function setWaitingSecsToRevoke(uint256 _secs) returns()
func (_Ron *RonTransactor) SetWaitingSecsToRevoke(opts *bind.TransactOpts, _secs *big.Int) (*types.Transaction, error) {
	return _Ron.contract.Transact(opts, "setWaitingSecsToRevoke", _secs)
}

// SetWaitingSecsToRevoke is a paid mutator transaction binding the contract method 0x969ffc14.
//
// Solidity: function setWaitingSecsToRevoke(uint256 _secs) returns()
func (_Ron *RonSession) SetWaitingSecsToRevoke(_secs *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.SetWaitingSecsToRevoke(&_Ron.TransactOpts, _secs)
}

// SetWaitingSecsToRevoke is a paid mutator transaction binding the contract method 0x969ffc14.
//
// Solidity: function setWaitingSecsToRevoke(uint256 _secs) returns()
func (_Ron *RonTransactorSession) SetWaitingSecsToRevoke(_secs *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.SetWaitingSecsToRevoke(&_Ron.TransactOpts, _secs)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address _consensusAddr) payable returns()
func (_Ron *RonTransactor) Stake(opts *bind.TransactOpts, _consensusAddr common.Address) (*types.Transaction, error) {
	return _Ron.contract.Transact(opts, "stake", _consensusAddr)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address _consensusAddr) payable returns()
func (_Ron *RonSession) Stake(_consensusAddr common.Address) (*types.Transaction, error) {
	return _Ron.Contract.Stake(&_Ron.TransactOpts, _consensusAddr)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address _consensusAddr) payable returns()
func (_Ron *RonTransactorSession) Stake(_consensusAddr common.Address) (*types.Transaction, error) {
	return _Ron.Contract.Stake(&_Ron.TransactOpts, _consensusAddr)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _consensusAddr, uint256 _amount) returns()
func (_Ron *RonTransactor) Undelegate(opts *bind.TransactOpts, _consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ron.contract.Transact(opts, "undelegate", _consensusAddr, _amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _consensusAddr, uint256 _amount) returns()
func (_Ron *RonSession) Undelegate(_consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.Undelegate(&_Ron.TransactOpts, _consensusAddr, _amount)
}

// Undelegate is a paid mutator transaction binding the contract method 0x4d99dd16.
//
// Solidity: function undelegate(address _consensusAddr, uint256 _amount) returns()
func (_Ron *RonTransactorSession) Undelegate(_consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.Undelegate(&_Ron.TransactOpts, _consensusAddr, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _consensusAddr, uint256 _amount) returns()
func (_Ron *RonTransactor) Unstake(opts *bind.TransactOpts, _consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ron.contract.Transact(opts, "unstake", _consensusAddr, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _consensusAddr, uint256 _amount) returns()
func (_Ron *RonSession) Unstake(_consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.Unstake(&_Ron.TransactOpts, _consensusAddr, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address _consensusAddr, uint256 _amount) returns()
func (_Ron *RonTransactorSession) Unstake(_consensusAddr common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Ron.Contract.Unstake(&_Ron.TransactOpts, _consensusAddr, _amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Ron *RonTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Ron.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Ron *RonSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Ron.Contract.Fallback(&_Ron.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Ron *RonTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Ron.Contract.Fallback(&_Ron.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ron *RonTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ron.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ron *RonSession) Receive() (*types.Transaction, error) {
	return _Ron.Contract.Receive(&_Ron.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Ron *RonTransactorSession) Receive() (*types.Transaction, error) {
	return _Ron.Contract.Receive(&_Ron.TransactOpts)
}

// RonCooldownSecsToUndelegateUpdatedIterator is returned from FilterCooldownSecsToUndelegateUpdated and is used to iterate over the raw logs and unpacked data for CooldownSecsToUndelegateUpdated events raised by the Ron contract.
type RonCooldownSecsToUndelegateUpdatedIterator struct {
	Event *RonCooldownSecsToUndelegateUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RonCooldownSecsToUndelegateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonCooldownSecsToUndelegateUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RonCooldownSecsToUndelegateUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RonCooldownSecsToUndelegateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonCooldownSecsToUndelegateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonCooldownSecsToUndelegateUpdated represents a CooldownSecsToUndelegateUpdated event raised by the Ron contract.
type RonCooldownSecsToUndelegateUpdated struct {
	MinSecs *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCooldownSecsToUndelegateUpdated is a free log retrieval operation binding the contract event 0x4956b65267b8f1e642284bcb5037116c69a9c78d9ca576beeae0974737a4872a.
//
// Solidity: event CooldownSecsToUndelegateUpdated(uint256 minSecs)
func (_Ron *RonFilterer) FilterCooldownSecsToUndelegateUpdated(opts *bind.FilterOpts) (*RonCooldownSecsToUndelegateUpdatedIterator, error) {

	logs, sub, err := _Ron.contract.FilterLogs(opts, "CooldownSecsToUndelegateUpdated")
	if err != nil {
		return nil, err
	}
	return &RonCooldownSecsToUndelegateUpdatedIterator{contract: _Ron.contract, event: "CooldownSecsToUndelegateUpdated", logs: logs, sub: sub}, nil
}

// WatchCooldownSecsToUndelegateUpdated is a free log subscription operation binding the contract event 0x4956b65267b8f1e642284bcb5037116c69a9c78d9ca576beeae0974737a4872a.
//
// Solidity: event CooldownSecsToUndelegateUpdated(uint256 minSecs)
func (_Ron *RonFilterer) WatchCooldownSecsToUndelegateUpdated(opts *bind.WatchOpts, sink chan<- *RonCooldownSecsToUndelegateUpdated) (event.Subscription, error) {

	logs, sub, err := _Ron.contract.WatchLogs(opts, "CooldownSecsToUndelegateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonCooldownSecsToUndelegateUpdated)
				if err := _Ron.contract.UnpackLog(event, "CooldownSecsToUndelegateUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCooldownSecsToUndelegateUpdated is a log parse operation binding the contract event 0x4956b65267b8f1e642284bcb5037116c69a9c78d9ca576beeae0974737a4872a.
//
// Solidity: event CooldownSecsToUndelegateUpdated(uint256 minSecs)
func (_Ron *RonFilterer) ParseCooldownSecsToUndelegateUpdated(log types.Log) (*RonCooldownSecsToUndelegateUpdated, error) {
	event := new(RonCooldownSecsToUndelegateUpdated)
	if err := _Ron.contract.UnpackLog(event, "CooldownSecsToUndelegateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonDelegatedIterator is returned from FilterDelegated and is used to iterate over the raw logs and unpacked data for Delegated events raised by the Ron contract.
type RonDelegatedIterator struct {
	Event *RonDelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RonDelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonDelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RonDelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RonDelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonDelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonDelegated represents a Delegated event raised by the Ron contract.
type RonDelegated struct {
	Delegator    common.Address
	ConsensuAddr common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegated is a free log retrieval operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed delegator, address indexed consensuAddr, uint256 amount)
func (_Ron *RonFilterer) FilterDelegated(opts *bind.FilterOpts, delegator []common.Address, consensuAddr []common.Address) (*RonDelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var consensuAddrRule []interface{}
	for _, consensuAddrItem := range consensuAddr {
		consensuAddrRule = append(consensuAddrRule, consensuAddrItem)
	}

	logs, sub, err := _Ron.contract.FilterLogs(opts, "Delegated", delegatorRule, consensuAddrRule)
	if err != nil {
		return nil, err
	}
	return &RonDelegatedIterator{contract: _Ron.contract, event: "Delegated", logs: logs, sub: sub}, nil
}

// WatchDelegated is a free log subscription operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed delegator, address indexed consensuAddr, uint256 amount)
func (_Ron *RonFilterer) WatchDelegated(opts *bind.WatchOpts, sink chan<- *RonDelegated, delegator []common.Address, consensuAddr []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var consensuAddrRule []interface{}
	for _, consensuAddrItem := range consensuAddr {
		consensuAddrRule = append(consensuAddrRule, consensuAddrItem)
	}

	logs, sub, err := _Ron.contract.WatchLogs(opts, "Delegated", delegatorRule, consensuAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonDelegated)
				if err := _Ron.contract.UnpackLog(event, "Delegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelegated is a log parse operation binding the contract event 0xe5541a6b6103d4fa7e021ed54fad39c66f27a76bd13d374cf6240ae6bd0bb72b.
//
// Solidity: event Delegated(address indexed delegator, address indexed consensuAddr, uint256 amount)
func (_Ron *RonFilterer) ParseDelegated(log types.Log) (*RonDelegated, error) {
	event := new(RonDelegated)
	if err := _Ron.contract.UnpackLog(event, "Delegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Ron contract.
type RonInitializedIterator struct {
	Event *RonInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RonInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RonInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RonInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonInitialized represents a Initialized event raised by the Ron contract.
type RonInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ron *RonFilterer) FilterInitialized(opts *bind.FilterOpts) (*RonInitializedIterator, error) {

	logs, sub, err := _Ron.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RonInitializedIterator{contract: _Ron.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ron *RonFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RonInitialized) (event.Subscription, error) {

	logs, sub, err := _Ron.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonInitialized)
				if err := _Ron.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Ron *RonFilterer) ParseInitialized(log types.Log) (*RonInitialized, error) {
	event := new(RonInitialized)
	if err := _Ron.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonMaxCommissionRateUpdatedIterator is returned from FilterMaxCommissionRateUpdated and is used to iterate over the raw logs and unpacked data for MaxCommissionRateUpdated events raised by the Ron contract.
type RonMaxCommissionRateUpdatedIterator struct {
	Event *RonMaxCommissionRateUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RonMaxCommissionRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonMaxCommissionRateUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RonMaxCommissionRateUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RonMaxCommissionRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonMaxCommissionRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonMaxCommissionRateUpdated represents a MaxCommissionRateUpdated event raised by the Ron contract.
type RonMaxCommissionRateUpdated struct {
	MaxRate *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMaxCommissionRateUpdated is a free log retrieval operation binding the contract event 0x774069781371d65424b3b0b101c1d40014532cac040f979595b99a3fcf8ce08c.
//
// Solidity: event MaxCommissionRateUpdated(uint256 maxRate)
func (_Ron *RonFilterer) FilterMaxCommissionRateUpdated(opts *bind.FilterOpts) (*RonMaxCommissionRateUpdatedIterator, error) {

	logs, sub, err := _Ron.contract.FilterLogs(opts, "MaxCommissionRateUpdated")
	if err != nil {
		return nil, err
	}
	return &RonMaxCommissionRateUpdatedIterator{contract: _Ron.contract, event: "MaxCommissionRateUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxCommissionRateUpdated is a free log subscription operation binding the contract event 0x774069781371d65424b3b0b101c1d40014532cac040f979595b99a3fcf8ce08c.
//
// Solidity: event MaxCommissionRateUpdated(uint256 maxRate)
func (_Ron *RonFilterer) WatchMaxCommissionRateUpdated(opts *bind.WatchOpts, sink chan<- *RonMaxCommissionRateUpdated) (event.Subscription, error) {

	logs, sub, err := _Ron.contract.WatchLogs(opts, "MaxCommissionRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonMaxCommissionRateUpdated)
				if err := _Ron.contract.UnpackLog(event, "MaxCommissionRateUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMaxCommissionRateUpdated is a log parse operation binding the contract event 0x774069781371d65424b3b0b101c1d40014532cac040f979595b99a3fcf8ce08c.
//
// Solidity: event MaxCommissionRateUpdated(uint256 maxRate)
func (_Ron *RonFilterer) ParseMaxCommissionRateUpdated(log types.Log) (*RonMaxCommissionRateUpdated, error) {
	event := new(RonMaxCommissionRateUpdated)
	if err := _Ron.contract.UnpackLog(event, "MaxCommissionRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonMinValidatorStakingAmountUpdatedIterator is returned from FilterMinValidatorStakingAmountUpdated and is used to iterate over the raw logs and unpacked data for MinValidatorStakingAmountUpdated events raised by the Ron contract.
type RonMinValidatorStakingAmountUpdatedIterator struct {
	Event *RonMinValidatorStakingAmountUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RonMinValidatorStakingAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonMinValidatorStakingAmountUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RonMinValidatorStakingAmountUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RonMinValidatorStakingAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonMinValidatorStakingAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonMinValidatorStakingAmountUpdated represents a MinValidatorStakingAmountUpdated event raised by the Ron contract.
type RonMinValidatorStakingAmountUpdated struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinValidatorStakingAmountUpdated is a free log retrieval operation binding the contract event 0x372bbdb8d72373b0012f84ee5a11671e5fb72b8bea902ebca93a19cb45d32be2.
//
// Solidity: event MinValidatorStakingAmountUpdated(uint256 threshold)
func (_Ron *RonFilterer) FilterMinValidatorStakingAmountUpdated(opts *bind.FilterOpts) (*RonMinValidatorStakingAmountUpdatedIterator, error) {

	logs, sub, err := _Ron.contract.FilterLogs(opts, "MinValidatorStakingAmountUpdated")
	if err != nil {
		return nil, err
	}
	return &RonMinValidatorStakingAmountUpdatedIterator{contract: _Ron.contract, event: "MinValidatorStakingAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchMinValidatorStakingAmountUpdated is a free log subscription operation binding the contract event 0x372bbdb8d72373b0012f84ee5a11671e5fb72b8bea902ebca93a19cb45d32be2.
//
// Solidity: event MinValidatorStakingAmountUpdated(uint256 threshold)
func (_Ron *RonFilterer) WatchMinValidatorStakingAmountUpdated(opts *bind.WatchOpts, sink chan<- *RonMinValidatorStakingAmountUpdated) (event.Subscription, error) {

	logs, sub, err := _Ron.contract.WatchLogs(opts, "MinValidatorStakingAmountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonMinValidatorStakingAmountUpdated)
				if err := _Ron.contract.UnpackLog(event, "MinValidatorStakingAmountUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMinValidatorStakingAmountUpdated is a log parse operation binding the contract event 0x372bbdb8d72373b0012f84ee5a11671e5fb72b8bea902ebca93a19cb45d32be2.
//
// Solidity: event MinValidatorStakingAmountUpdated(uint256 threshold)
func (_Ron *RonFilterer) ParseMinValidatorStakingAmountUpdated(log types.Log) (*RonMinValidatorStakingAmountUpdated, error) {
	event := new(RonMinValidatorStakingAmountUpdated)
	if err := _Ron.contract.UnpackLog(event, "MinValidatorStakingAmountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonPoolApprovedIterator is returned from FilterPoolApproved and is used to iterate over the raw logs and unpacked data for PoolApproved events raised by the Ron contract.
type RonPoolApprovedIterator struct {
	Event *RonPoolApproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RonPoolApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonPoolApproved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RonPoolApproved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RonPoolApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonPoolApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonPoolApproved represents a PoolApproved event raised by the Ron contract.
type RonPoolApproved struct {
	Validator common.Address
	Admin     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPoolApproved is a free log retrieval operation binding the contract event 0xfc1f1e73948cbc47c5b7f90e5601b7daccd9ad7173218486ccc74bdd051d05e8.
//
// Solidity: event PoolApproved(address indexed validator, address indexed admin)
func (_Ron *RonFilterer) FilterPoolApproved(opts *bind.FilterOpts, validator []common.Address, admin []common.Address) (*RonPoolApprovedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Ron.contract.FilterLogs(opts, "PoolApproved", validatorRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &RonPoolApprovedIterator{contract: _Ron.contract, event: "PoolApproved", logs: logs, sub: sub}, nil
}

// WatchPoolApproved is a free log subscription operation binding the contract event 0xfc1f1e73948cbc47c5b7f90e5601b7daccd9ad7173218486ccc74bdd051d05e8.
//
// Solidity: event PoolApproved(address indexed validator, address indexed admin)
func (_Ron *RonFilterer) WatchPoolApproved(opts *bind.WatchOpts, sink chan<- *RonPoolApproved, validator []common.Address, admin []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Ron.contract.WatchLogs(opts, "PoolApproved", validatorRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonPoolApproved)
				if err := _Ron.contract.UnpackLog(event, "PoolApproved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolApproved is a log parse operation binding the contract event 0xfc1f1e73948cbc47c5b7f90e5601b7daccd9ad7173218486ccc74bdd051d05e8.
//
// Solidity: event PoolApproved(address indexed validator, address indexed admin)
func (_Ron *RonFilterer) ParsePoolApproved(log types.Log) (*RonPoolApproved, error) {
	event := new(RonPoolApproved)
	if err := _Ron.contract.UnpackLog(event, "PoolApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonPoolSharesUpdatedIterator is returned from FilterPoolSharesUpdated and is used to iterate over the raw logs and unpacked data for PoolSharesUpdated events raised by the Ron contract.
type RonPoolSharesUpdatedIterator struct {
	Event *RonPoolSharesUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RonPoolSharesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonPoolSharesUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RonPoolSharesUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RonPoolSharesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonPoolSharesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonPoolSharesUpdated represents a PoolSharesUpdated event raised by the Ron contract.
type RonPoolSharesUpdated struct {
	Period   *big.Int
	PoolAddr common.Address
	Shares   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPoolSharesUpdated is a free log retrieval operation binding the contract event 0x81faf50e2aaf52eaba2ab841071efb9f6f0850a3e7d008b1336e6001d3d4963c.
//
// Solidity: event PoolSharesUpdated(uint256 indexed period, address indexed poolAddr, uint256 shares)
func (_Ron *RonFilterer) FilterPoolSharesUpdated(opts *bind.FilterOpts, period []*big.Int, poolAddr []common.Address) (*RonPoolSharesUpdatedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var poolAddrRule []interface{}
	for _, poolAddrItem := range poolAddr {
		poolAddrRule = append(poolAddrRule, poolAddrItem)
	}

	logs, sub, err := _Ron.contract.FilterLogs(opts, "PoolSharesUpdated", periodRule, poolAddrRule)
	if err != nil {
		return nil, err
	}
	return &RonPoolSharesUpdatedIterator{contract: _Ron.contract, event: "PoolSharesUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolSharesUpdated is a free log subscription operation binding the contract event 0x81faf50e2aaf52eaba2ab841071efb9f6f0850a3e7d008b1336e6001d3d4963c.
//
// Solidity: event PoolSharesUpdated(uint256 indexed period, address indexed poolAddr, uint256 shares)
func (_Ron *RonFilterer) WatchPoolSharesUpdated(opts *bind.WatchOpts, sink chan<- *RonPoolSharesUpdated, period []*big.Int, poolAddr []common.Address) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var poolAddrRule []interface{}
	for _, poolAddrItem := range poolAddr {
		poolAddrRule = append(poolAddrRule, poolAddrItem)
	}

	logs, sub, err := _Ron.contract.WatchLogs(opts, "PoolSharesUpdated", periodRule, poolAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonPoolSharesUpdated)
				if err := _Ron.contract.UnpackLog(event, "PoolSharesUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolSharesUpdated is a log parse operation binding the contract event 0x81faf50e2aaf52eaba2ab841071efb9f6f0850a3e7d008b1336e6001d3d4963c.
//
// Solidity: event PoolSharesUpdated(uint256 indexed period, address indexed poolAddr, uint256 shares)
func (_Ron *RonFilterer) ParsePoolSharesUpdated(log types.Log) (*RonPoolSharesUpdated, error) {
	event := new(RonPoolSharesUpdated)
	if err := _Ron.contract.UnpackLog(event, "PoolSharesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonPoolsDeprecatedIterator is returned from FilterPoolsDeprecated and is used to iterate over the raw logs and unpacked data for PoolsDeprecated events raised by the Ron contract.
type RonPoolsDeprecatedIterator struct {
	Event *RonPoolsDeprecated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RonPoolsDeprecatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonPoolsDeprecated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RonPoolsDeprecated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RonPoolsDeprecatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonPoolsDeprecatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonPoolsDeprecated represents a PoolsDeprecated event raised by the Ron contract.
type RonPoolsDeprecated struct {
	Validator []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPoolsDeprecated is a free log retrieval operation binding the contract event 0x4f257d3ba23679d338f1d94296086bba5724af341b7fa31aa0ff297bfcdc62d8.
//
// Solidity: event PoolsDeprecated(address[] validator)
func (_Ron *RonFilterer) FilterPoolsDeprecated(opts *bind.FilterOpts) (*RonPoolsDeprecatedIterator, error) {

	logs, sub, err := _Ron.contract.FilterLogs(opts, "PoolsDeprecated")
	if err != nil {
		return nil, err
	}
	return &RonPoolsDeprecatedIterator{contract: _Ron.contract, event: "PoolsDeprecated", logs: logs, sub: sub}, nil
}

// WatchPoolsDeprecated is a free log subscription operation binding the contract event 0x4f257d3ba23679d338f1d94296086bba5724af341b7fa31aa0ff297bfcdc62d8.
//
// Solidity: event PoolsDeprecated(address[] validator)
func (_Ron *RonFilterer) WatchPoolsDeprecated(opts *bind.WatchOpts, sink chan<- *RonPoolsDeprecated) (event.Subscription, error) {

	logs, sub, err := _Ron.contract.WatchLogs(opts, "PoolsDeprecated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonPoolsDeprecated)
				if err := _Ron.contract.UnpackLog(event, "PoolsDeprecated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolsDeprecated is a log parse operation binding the contract event 0x4f257d3ba23679d338f1d94296086bba5724af341b7fa31aa0ff297bfcdc62d8.
//
// Solidity: event PoolsDeprecated(address[] validator)
func (_Ron *RonFilterer) ParsePoolsDeprecated(log types.Log) (*RonPoolsDeprecated, error) {
	event := new(RonPoolsDeprecated)
	if err := _Ron.contract.UnpackLog(event, "PoolsDeprecated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonPoolsUpdateConflictedIterator is returned from FilterPoolsUpdateConflicted and is used to iterate over the raw logs and unpacked data for PoolsUpdateConflicted events raised by the Ron contract.
type RonPoolsUpdateConflictedIterator struct {
	Event *RonPoolsUpdateConflicted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RonPoolsUpdateConflictedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonPoolsUpdateConflicted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RonPoolsUpdateConflicted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RonPoolsUpdateConflictedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonPoolsUpdateConflictedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonPoolsUpdateConflicted represents a PoolsUpdateConflicted event raised by the Ron contract.
type RonPoolsUpdateConflicted struct {
	Period    *big.Int
	PoolAddrs []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPoolsUpdateConflicted is a free log retrieval operation binding the contract event 0xee74f10cc50bf4b7e57fd36be7d46288795f3a9151dae97505b718b392ba14a3.
//
// Solidity: event PoolsUpdateConflicted(uint256 indexed period, address[] poolAddrs)
func (_Ron *RonFilterer) FilterPoolsUpdateConflicted(opts *bind.FilterOpts, period []*big.Int) (*RonPoolsUpdateConflictedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _Ron.contract.FilterLogs(opts, "PoolsUpdateConflicted", periodRule)
	if err != nil {
		return nil, err
	}
	return &RonPoolsUpdateConflictedIterator{contract: _Ron.contract, event: "PoolsUpdateConflicted", logs: logs, sub: sub}, nil
}

// WatchPoolsUpdateConflicted is a free log subscription operation binding the contract event 0xee74f10cc50bf4b7e57fd36be7d46288795f3a9151dae97505b718b392ba14a3.
//
// Solidity: event PoolsUpdateConflicted(uint256 indexed period, address[] poolAddrs)
func (_Ron *RonFilterer) WatchPoolsUpdateConflicted(opts *bind.WatchOpts, sink chan<- *RonPoolsUpdateConflicted, period []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _Ron.contract.WatchLogs(opts, "PoolsUpdateConflicted", periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonPoolsUpdateConflicted)
				if err := _Ron.contract.UnpackLog(event, "PoolsUpdateConflicted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolsUpdateConflicted is a log parse operation binding the contract event 0xee74f10cc50bf4b7e57fd36be7d46288795f3a9151dae97505b718b392ba14a3.
//
// Solidity: event PoolsUpdateConflicted(uint256 indexed period, address[] poolAddrs)
func (_Ron *RonFilterer) ParsePoolsUpdateConflicted(log types.Log) (*RonPoolsUpdateConflicted, error) {
	event := new(RonPoolsUpdateConflicted)
	if err := _Ron.contract.UnpackLog(event, "PoolsUpdateConflicted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonPoolsUpdateFailedIterator is returned from FilterPoolsUpdateFailed and is used to iterate over the raw logs and unpacked data for PoolsUpdateFailed events raised by the Ron contract.
type RonPoolsUpdateFailedIterator struct {
	Event *RonPoolsUpdateFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RonPoolsUpdateFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonPoolsUpdateFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RonPoolsUpdateFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RonPoolsUpdateFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonPoolsUpdateFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonPoolsUpdateFailed represents a PoolsUpdateFailed event raised by the Ron contract.
type RonPoolsUpdateFailed struct {
	Period    *big.Int
	PoolAddrs []common.Address
	Rewards   []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPoolsUpdateFailed is a free log retrieval operation binding the contract event 0xae52c603227f64e4c6101dde593aa9790a16b3ac77546bd746d758511e9560a5.
//
// Solidity: event PoolsUpdateFailed(uint256 indexed period, address[] poolAddrs, uint256[] rewards)
func (_Ron *RonFilterer) FilterPoolsUpdateFailed(opts *bind.FilterOpts, period []*big.Int) (*RonPoolsUpdateFailedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _Ron.contract.FilterLogs(opts, "PoolsUpdateFailed", periodRule)
	if err != nil {
		return nil, err
	}
	return &RonPoolsUpdateFailedIterator{contract: _Ron.contract, event: "PoolsUpdateFailed", logs: logs, sub: sub}, nil
}

// WatchPoolsUpdateFailed is a free log subscription operation binding the contract event 0xae52c603227f64e4c6101dde593aa9790a16b3ac77546bd746d758511e9560a5.
//
// Solidity: event PoolsUpdateFailed(uint256 indexed period, address[] poolAddrs, uint256[] rewards)
func (_Ron *RonFilterer) WatchPoolsUpdateFailed(opts *bind.WatchOpts, sink chan<- *RonPoolsUpdateFailed, period []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _Ron.contract.WatchLogs(opts, "PoolsUpdateFailed", periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonPoolsUpdateFailed)
				if err := _Ron.contract.UnpackLog(event, "PoolsUpdateFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolsUpdateFailed is a log parse operation binding the contract event 0xae52c603227f64e4c6101dde593aa9790a16b3ac77546bd746d758511e9560a5.
//
// Solidity: event PoolsUpdateFailed(uint256 indexed period, address[] poolAddrs, uint256[] rewards)
func (_Ron *RonFilterer) ParsePoolsUpdateFailed(log types.Log) (*RonPoolsUpdateFailed, error) {
	event := new(RonPoolsUpdateFailed)
	if err := _Ron.contract.UnpackLog(event, "PoolsUpdateFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonPoolsUpdatedIterator is returned from FilterPoolsUpdated and is used to iterate over the raw logs and unpacked data for PoolsUpdated events raised by the Ron contract.
type RonPoolsUpdatedIterator struct {
	Event *RonPoolsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RonPoolsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonPoolsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RonPoolsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RonPoolsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonPoolsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonPoolsUpdated represents a PoolsUpdated event raised by the Ron contract.
type RonPoolsUpdated struct {
	Period    *big.Int
	PoolAddrs []common.Address
	ARps      []*big.Int
	Shares    []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPoolsUpdated is a free log retrieval operation binding the contract event 0x0e54e0485f70f0f63bc25889ddbf01ce1269ad6f07fdb2df573a0fbdb4d66f88.
//
// Solidity: event PoolsUpdated(uint256 indexed period, address[] poolAddrs, uint256[] aRps, uint256[] shares)
func (_Ron *RonFilterer) FilterPoolsUpdated(opts *bind.FilterOpts, period []*big.Int) (*RonPoolsUpdatedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _Ron.contract.FilterLogs(opts, "PoolsUpdated", periodRule)
	if err != nil {
		return nil, err
	}
	return &RonPoolsUpdatedIterator{contract: _Ron.contract, event: "PoolsUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolsUpdated is a free log subscription operation binding the contract event 0x0e54e0485f70f0f63bc25889ddbf01ce1269ad6f07fdb2df573a0fbdb4d66f88.
//
// Solidity: event PoolsUpdated(uint256 indexed period, address[] poolAddrs, uint256[] aRps, uint256[] shares)
func (_Ron *RonFilterer) WatchPoolsUpdated(opts *bind.WatchOpts, sink chan<- *RonPoolsUpdated, period []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _Ron.contract.WatchLogs(opts, "PoolsUpdated", periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonPoolsUpdated)
				if err := _Ron.contract.UnpackLog(event, "PoolsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePoolsUpdated is a log parse operation binding the contract event 0x0e54e0485f70f0f63bc25889ddbf01ce1269ad6f07fdb2df573a0fbdb4d66f88.
//
// Solidity: event PoolsUpdated(uint256 indexed period, address[] poolAddrs, uint256[] aRps, uint256[] shares)
func (_Ron *RonFilterer) ParsePoolsUpdated(log types.Log) (*RonPoolsUpdated, error) {
	event := new(RonPoolsUpdated)
	if err := _Ron.contract.UnpackLog(event, "PoolsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonRewardClaimedIterator is returned from FilterRewardClaimed and is used to iterate over the raw logs and unpacked data for RewardClaimed events raised by the Ron contract.
type RonRewardClaimedIterator struct {
	Event *RonRewardClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RonRewardClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonRewardClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RonRewardClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RonRewardClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonRewardClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonRewardClaimed represents a RewardClaimed event raised by the Ron contract.
type RonRewardClaimed struct {
	PoolAddr common.Address
	User     common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRewardClaimed is a free log retrieval operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed poolAddr, address indexed user, uint256 amount)
func (_Ron *RonFilterer) FilterRewardClaimed(opts *bind.FilterOpts, poolAddr []common.Address, user []common.Address) (*RonRewardClaimedIterator, error) {

	var poolAddrRule []interface{}
	for _, poolAddrItem := range poolAddr {
		poolAddrRule = append(poolAddrRule, poolAddrItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Ron.contract.FilterLogs(opts, "RewardClaimed", poolAddrRule, userRule)
	if err != nil {
		return nil, err
	}
	return &RonRewardClaimedIterator{contract: _Ron.contract, event: "RewardClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardClaimed is a free log subscription operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed poolAddr, address indexed user, uint256 amount)
func (_Ron *RonFilterer) WatchRewardClaimed(opts *bind.WatchOpts, sink chan<- *RonRewardClaimed, poolAddr []common.Address, user []common.Address) (event.Subscription, error) {

	var poolAddrRule []interface{}
	for _, poolAddrItem := range poolAddr {
		poolAddrRule = append(poolAddrRule, poolAddrItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Ron.contract.WatchLogs(opts, "RewardClaimed", poolAddrRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonRewardClaimed)
				if err := _Ron.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardClaimed is a log parse operation binding the contract event 0x0aa4d283470c904c551d18bb894d37e17674920f3261a7f854be501e25f421b7.
//
// Solidity: event RewardClaimed(address indexed poolAddr, address indexed user, uint256 amount)
func (_Ron *RonFilterer) ParseRewardClaimed(log types.Log) (*RonRewardClaimed, error) {
	event := new(RonRewardClaimed)
	if err := _Ron.contract.UnpackLog(event, "RewardClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Ron contract.
type RonStakedIterator struct {
	Event *RonStaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RonStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonStaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RonStaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RonStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonStaked represents a Staked event raised by the Ron contract.
type RonStaked struct {
	ConsensuAddr common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed consensuAddr, uint256 amount)
func (_Ron *RonFilterer) FilterStaked(opts *bind.FilterOpts, consensuAddr []common.Address) (*RonStakedIterator, error) {

	var consensuAddrRule []interface{}
	for _, consensuAddrItem := range consensuAddr {
		consensuAddrRule = append(consensuAddrRule, consensuAddrItem)
	}

	logs, sub, err := _Ron.contract.FilterLogs(opts, "Staked", consensuAddrRule)
	if err != nil {
		return nil, err
	}
	return &RonStakedIterator{contract: _Ron.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed consensuAddr, uint256 amount)
func (_Ron *RonFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *RonStaked, consensuAddr []common.Address) (event.Subscription, error) {

	var consensuAddrRule []interface{}
	for _, consensuAddrItem := range consensuAddr {
		consensuAddrRule = append(consensuAddrRule, consensuAddrItem)
	}

	logs, sub, err := _Ron.contract.WatchLogs(opts, "Staked", consensuAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonStaked)
				if err := _Ron.contract.UnpackLog(event, "Staked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed consensuAddr, uint256 amount)
func (_Ron *RonFilterer) ParseStaked(log types.Log) (*RonStaked, error) {
	event := new(RonStaked)
	if err := _Ron.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonStakingAmountDeductFailedIterator is returned from FilterStakingAmountDeductFailed and is used to iterate over the raw logs and unpacked data for StakingAmountDeductFailed events raised by the Ron contract.
type RonStakingAmountDeductFailedIterator struct {
	Event *RonStakingAmountDeductFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RonStakingAmountDeductFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonStakingAmountDeductFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RonStakingAmountDeductFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RonStakingAmountDeductFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonStakingAmountDeductFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonStakingAmountDeductFailed represents a StakingAmountDeductFailed event raised by the Ron contract.
type RonStakingAmountDeductFailed struct {
	Validator       common.Address
	Recipient       common.Address
	Amount          *big.Int
	ContractBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakingAmountDeductFailed is a free log retrieval operation binding the contract event 0x63701cd972aa3c7f87898aab145c972e52185beab07d6e39380a998d334cf6c8.
//
// Solidity: event StakingAmountDeductFailed(address indexed validator, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_Ron *RonFilterer) FilterStakingAmountDeductFailed(opts *bind.FilterOpts, validator []common.Address, recipient []common.Address) (*RonStakingAmountDeductFailedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Ron.contract.FilterLogs(opts, "StakingAmountDeductFailed", validatorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &RonStakingAmountDeductFailedIterator{contract: _Ron.contract, event: "StakingAmountDeductFailed", logs: logs, sub: sub}, nil
}

// WatchStakingAmountDeductFailed is a free log subscription operation binding the contract event 0x63701cd972aa3c7f87898aab145c972e52185beab07d6e39380a998d334cf6c8.
//
// Solidity: event StakingAmountDeductFailed(address indexed validator, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_Ron *RonFilterer) WatchStakingAmountDeductFailed(opts *bind.WatchOpts, sink chan<- *RonStakingAmountDeductFailed, validator []common.Address, recipient []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Ron.contract.WatchLogs(opts, "StakingAmountDeductFailed", validatorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonStakingAmountDeductFailed)
				if err := _Ron.contract.UnpackLog(event, "StakingAmountDeductFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakingAmountDeductFailed is a log parse operation binding the contract event 0x63701cd972aa3c7f87898aab145c972e52185beab07d6e39380a998d334cf6c8.
//
// Solidity: event StakingAmountDeductFailed(address indexed validator, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_Ron *RonFilterer) ParseStakingAmountDeductFailed(log types.Log) (*RonStakingAmountDeductFailed, error) {
	event := new(RonStakingAmountDeductFailed)
	if err := _Ron.contract.UnpackLog(event, "StakingAmountDeductFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonStakingAmountTransferFailedIterator is returned from FilterStakingAmountTransferFailed and is used to iterate over the raw logs and unpacked data for StakingAmountTransferFailed events raised by the Ron contract.
type RonStakingAmountTransferFailedIterator struct {
	Event *RonStakingAmountTransferFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RonStakingAmountTransferFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonStakingAmountTransferFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RonStakingAmountTransferFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RonStakingAmountTransferFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonStakingAmountTransferFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonStakingAmountTransferFailed represents a StakingAmountTransferFailed event raised by the Ron contract.
type RonStakingAmountTransferFailed struct {
	Validator       common.Address
	Admin           common.Address
	Amount          *big.Int
	ContractBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakingAmountTransferFailed is a free log retrieval operation binding the contract event 0x7dc5115a5aba081f5a174f56a3d02eea582824783322a4ac03f7bd388f444194.
//
// Solidity: event StakingAmountTransferFailed(address indexed validator, address indexed admin, uint256 amount, uint256 contractBalance)
func (_Ron *RonFilterer) FilterStakingAmountTransferFailed(opts *bind.FilterOpts, validator []common.Address, admin []common.Address) (*RonStakingAmountTransferFailedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Ron.contract.FilterLogs(opts, "StakingAmountTransferFailed", validatorRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &RonStakingAmountTransferFailedIterator{contract: _Ron.contract, event: "StakingAmountTransferFailed", logs: logs, sub: sub}, nil
}

// WatchStakingAmountTransferFailed is a free log subscription operation binding the contract event 0x7dc5115a5aba081f5a174f56a3d02eea582824783322a4ac03f7bd388f444194.
//
// Solidity: event StakingAmountTransferFailed(address indexed validator, address indexed admin, uint256 amount, uint256 contractBalance)
func (_Ron *RonFilterer) WatchStakingAmountTransferFailed(opts *bind.WatchOpts, sink chan<- *RonStakingAmountTransferFailed, validator []common.Address, admin []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Ron.contract.WatchLogs(opts, "StakingAmountTransferFailed", validatorRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonStakingAmountTransferFailed)
				if err := _Ron.contract.UnpackLog(event, "StakingAmountTransferFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakingAmountTransferFailed is a log parse operation binding the contract event 0x7dc5115a5aba081f5a174f56a3d02eea582824783322a4ac03f7bd388f444194.
//
// Solidity: event StakingAmountTransferFailed(address indexed validator, address indexed admin, uint256 amount, uint256 contractBalance)
func (_Ron *RonFilterer) ParseStakingAmountTransferFailed(log types.Log) (*RonStakingAmountTransferFailed, error) {
	event := new(RonStakingAmountTransferFailed)
	if err := _Ron.contract.UnpackLog(event, "StakingAmountTransferFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonUndelegatedIterator is returned from FilterUndelegated and is used to iterate over the raw logs and unpacked data for Undelegated events raised by the Ron contract.
type RonUndelegatedIterator struct {
	Event *RonUndelegated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RonUndelegatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonUndelegated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RonUndelegated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RonUndelegatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonUndelegatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonUndelegated represents a Undelegated event raised by the Ron contract.
type RonUndelegated struct {
	Delegator    common.Address
	ConsensuAddr common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUndelegated is a free log retrieval operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed delegator, address indexed consensuAddr, uint256 amount)
func (_Ron *RonFilterer) FilterUndelegated(opts *bind.FilterOpts, delegator []common.Address, consensuAddr []common.Address) (*RonUndelegatedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var consensuAddrRule []interface{}
	for _, consensuAddrItem := range consensuAddr {
		consensuAddrRule = append(consensuAddrRule, consensuAddrItem)
	}

	logs, sub, err := _Ron.contract.FilterLogs(opts, "Undelegated", delegatorRule, consensuAddrRule)
	if err != nil {
		return nil, err
	}
	return &RonUndelegatedIterator{contract: _Ron.contract, event: "Undelegated", logs: logs, sub: sub}, nil
}

// WatchUndelegated is a free log subscription operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed delegator, address indexed consensuAddr, uint256 amount)
func (_Ron *RonFilterer) WatchUndelegated(opts *bind.WatchOpts, sink chan<- *RonUndelegated, delegator []common.Address, consensuAddr []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var consensuAddrRule []interface{}
	for _, consensuAddrItem := range consensuAddr {
		consensuAddrRule = append(consensuAddrRule, consensuAddrItem)
	}

	logs, sub, err := _Ron.contract.WatchLogs(opts, "Undelegated", delegatorRule, consensuAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonUndelegated)
				if err := _Ron.contract.UnpackLog(event, "Undelegated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUndelegated is a log parse operation binding the contract event 0x4d10bd049775c77bd7f255195afba5088028ecb3c7c277d393ccff7934f2f92c.
//
// Solidity: event Undelegated(address indexed delegator, address indexed consensuAddr, uint256 amount)
func (_Ron *RonFilterer) ParseUndelegated(log types.Log) (*RonUndelegated, error) {
	event := new(RonUndelegated)
	if err := _Ron.contract.UnpackLog(event, "Undelegated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the Ron contract.
type RonUnstakedIterator struct {
	Event *RonUnstaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RonUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonUnstaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RonUnstaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RonUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonUnstaked represents a Unstaked event raised by the Ron contract.
type RonUnstaked struct {
	ConsensuAddr common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed consensuAddr, uint256 amount)
func (_Ron *RonFilterer) FilterUnstaked(opts *bind.FilterOpts, consensuAddr []common.Address) (*RonUnstakedIterator, error) {

	var consensuAddrRule []interface{}
	for _, consensuAddrItem := range consensuAddr {
		consensuAddrRule = append(consensuAddrRule, consensuAddrItem)
	}

	logs, sub, err := _Ron.contract.FilterLogs(opts, "Unstaked", consensuAddrRule)
	if err != nil {
		return nil, err
	}
	return &RonUnstakedIterator{contract: _Ron.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed consensuAddr, uint256 amount)
func (_Ron *RonFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *RonUnstaked, consensuAddr []common.Address) (event.Subscription, error) {

	var consensuAddrRule []interface{}
	for _, consensuAddrItem := range consensuAddr {
		consensuAddrRule = append(consensuAddrRule, consensuAddrItem)
	}

	logs, sub, err := _Ron.contract.WatchLogs(opts, "Unstaked", consensuAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonUnstaked)
				if err := _Ron.contract.UnpackLog(event, "Unstaked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnstaked is a log parse operation binding the contract event 0x0f5bb82176feb1b5e747e28471aa92156a04d9f3ab9f45f28e2d704232b93f75.
//
// Solidity: event Unstaked(address indexed consensuAddr, uint256 amount)
func (_Ron *RonFilterer) ParseUnstaked(log types.Log) (*RonUnstaked, error) {
	event := new(RonUnstaked)
	if err := _Ron.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonUserRewardUpdatedIterator is returned from FilterUserRewardUpdated and is used to iterate over the raw logs and unpacked data for UserRewardUpdated events raised by the Ron contract.
type RonUserRewardUpdatedIterator struct {
	Event *RonUserRewardUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RonUserRewardUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonUserRewardUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RonUserRewardUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RonUserRewardUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonUserRewardUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonUserRewardUpdated represents a UserRewardUpdated event raised by the Ron contract.
type RonUserRewardUpdated struct {
	PoolAddr common.Address
	User     common.Address
	Debited  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUserRewardUpdated is a free log retrieval operation binding the contract event 0xaa7c29611027fd4be148712bb54960253b7a7d5998c17769bfc424c2f5f185ad.
//
// Solidity: event UserRewardUpdated(address indexed poolAddr, address indexed user, uint256 debited)
func (_Ron *RonFilterer) FilterUserRewardUpdated(opts *bind.FilterOpts, poolAddr []common.Address, user []common.Address) (*RonUserRewardUpdatedIterator, error) {

	var poolAddrRule []interface{}
	for _, poolAddrItem := range poolAddr {
		poolAddrRule = append(poolAddrRule, poolAddrItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Ron.contract.FilterLogs(opts, "UserRewardUpdated", poolAddrRule, userRule)
	if err != nil {
		return nil, err
	}
	return &RonUserRewardUpdatedIterator{contract: _Ron.contract, event: "UserRewardUpdated", logs: logs, sub: sub}, nil
}

// WatchUserRewardUpdated is a free log subscription operation binding the contract event 0xaa7c29611027fd4be148712bb54960253b7a7d5998c17769bfc424c2f5f185ad.
//
// Solidity: event UserRewardUpdated(address indexed poolAddr, address indexed user, uint256 debited)
func (_Ron *RonFilterer) WatchUserRewardUpdated(opts *bind.WatchOpts, sink chan<- *RonUserRewardUpdated, poolAddr []common.Address, user []common.Address) (event.Subscription, error) {

	var poolAddrRule []interface{}
	for _, poolAddrItem := range poolAddr {
		poolAddrRule = append(poolAddrRule, poolAddrItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Ron.contract.WatchLogs(opts, "UserRewardUpdated", poolAddrRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonUserRewardUpdated)
				if err := _Ron.contract.UnpackLog(event, "UserRewardUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserRewardUpdated is a log parse operation binding the contract event 0xaa7c29611027fd4be148712bb54960253b7a7d5998c17769bfc424c2f5f185ad.
//
// Solidity: event UserRewardUpdated(address indexed poolAddr, address indexed user, uint256 debited)
func (_Ron *RonFilterer) ParseUserRewardUpdated(log types.Log) (*RonUserRewardUpdated, error) {
	event := new(RonUserRewardUpdated)
	if err := _Ron.contract.UnpackLog(event, "UserRewardUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonValidatorContractUpdatedIterator is returned from FilterValidatorContractUpdated and is used to iterate over the raw logs and unpacked data for ValidatorContractUpdated events raised by the Ron contract.
type RonValidatorContractUpdatedIterator struct {
	Event *RonValidatorContractUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RonValidatorContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonValidatorContractUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RonValidatorContractUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RonValidatorContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonValidatorContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonValidatorContractUpdated represents a ValidatorContractUpdated event raised by the Ron contract.
type RonValidatorContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterValidatorContractUpdated is a free log retrieval operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_Ron *RonFilterer) FilterValidatorContractUpdated(opts *bind.FilterOpts) (*RonValidatorContractUpdatedIterator, error) {

	logs, sub, err := _Ron.contract.FilterLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return &RonValidatorContractUpdatedIterator{contract: _Ron.contract, event: "ValidatorContractUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorContractUpdated is a free log subscription operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_Ron *RonFilterer) WatchValidatorContractUpdated(opts *bind.WatchOpts, sink chan<- *RonValidatorContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Ron.contract.WatchLogs(opts, "ValidatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonValidatorContractUpdated)
				if err := _Ron.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorContractUpdated is a log parse operation binding the contract event 0xef40dc07567635f84f5edbd2f8dbc16b40d9d282dd8e7e6f4ff58236b6836169.
//
// Solidity: event ValidatorContractUpdated(address arg0)
func (_Ron *RonFilterer) ParseValidatorContractUpdated(log types.Log) (*RonValidatorContractUpdated, error) {
	event := new(RonValidatorContractUpdated)
	if err := _Ron.contract.UnpackLog(event, "ValidatorContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RonWaitingSecsToRevokeUpdatedIterator is returned from FilterWaitingSecsToRevokeUpdated and is used to iterate over the raw logs and unpacked data for WaitingSecsToRevokeUpdated events raised by the Ron contract.
type RonWaitingSecsToRevokeUpdatedIterator struct {
	Event *RonWaitingSecsToRevokeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RonWaitingSecsToRevokeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RonWaitingSecsToRevokeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RonWaitingSecsToRevokeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RonWaitingSecsToRevokeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RonWaitingSecsToRevokeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RonWaitingSecsToRevokeUpdated represents a WaitingSecsToRevokeUpdated event raised by the Ron contract.
type RonWaitingSecsToRevokeUpdated struct {
	Secs *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWaitingSecsToRevokeUpdated is a free log retrieval operation binding the contract event 0x02be0b73b597f2c0f138aebee162b3b0e25d5b5a26854c15dcf79176e9a1c678.
//
// Solidity: event WaitingSecsToRevokeUpdated(uint256 secs)
func (_Ron *RonFilterer) FilterWaitingSecsToRevokeUpdated(opts *bind.FilterOpts) (*RonWaitingSecsToRevokeUpdatedIterator, error) {

	logs, sub, err := _Ron.contract.FilterLogs(opts, "WaitingSecsToRevokeUpdated")
	if err != nil {
		return nil, err
	}
	return &RonWaitingSecsToRevokeUpdatedIterator{contract: _Ron.contract, event: "WaitingSecsToRevokeUpdated", logs: logs, sub: sub}, nil
}

// WatchWaitingSecsToRevokeUpdated is a free log subscription operation binding the contract event 0x02be0b73b597f2c0f138aebee162b3b0e25d5b5a26854c15dcf79176e9a1c678.
//
// Solidity: event WaitingSecsToRevokeUpdated(uint256 secs)
func (_Ron *RonFilterer) WatchWaitingSecsToRevokeUpdated(opts *bind.WatchOpts, sink chan<- *RonWaitingSecsToRevokeUpdated) (event.Subscription, error) {

	logs, sub, err := _Ron.contract.WatchLogs(opts, "WaitingSecsToRevokeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RonWaitingSecsToRevokeUpdated)
				if err := _Ron.contract.UnpackLog(event, "WaitingSecsToRevokeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWaitingSecsToRevokeUpdated is a log parse operation binding the contract event 0x02be0b73b597f2c0f138aebee162b3b0e25d5b5a26854c15dcf79176e9a1c678.
//
// Solidity: event WaitingSecsToRevokeUpdated(uint256 secs)
func (_Ron *RonFilterer) ParseWaitingSecsToRevokeUpdated(log types.Log) (*RonWaitingSecsToRevokeUpdated, error) {
	event := new(RonWaitingSecsToRevokeUpdated)
	if err := _Ron.contract.UnpackLog(event, "WaitingSecsToRevokeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
