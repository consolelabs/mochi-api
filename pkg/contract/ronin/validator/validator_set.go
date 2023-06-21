// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validator

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

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Admin              common.Address
	ConsensusAddr      common.Address
	TreasuryAddr       common.Address
	BridgeOperatorAddr common.Address
	CommissionRate     *big.Int
	RevokingTimestamp  *big.Int
	TopupDeadline      *big.Int
}

// Struct1 is an auto generated low-level Go binding around an user-defined struct.
type Struct1 struct {
	EffectiveTimestamp *big.Int
	CommissionRate     *big.Int
}

// ValidatorMetaData contains all meta data concerning the Validator contract.
var ValidatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrAlreadyRequestedEmergencyExit\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrAlreadyRequestedRevokingCandidate\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrAlreadyRequestedUpdatingCommissionRate\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrAlreadyWrappedEpoch\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrAtEndOfEpochOnly\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrCallPrecompiled\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrCallerMustBeBridgeTrackingContract\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrCallerMustBeCoinbase\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrCallerMustBeMaintenanceContract\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrCallerMustBeRoninTrustedOrgContract\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrCallerMustBeSlashIndicatorContract\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrCallerMustBeStakingContract\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrCallerMustBeStakingVestingContract\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"validator\",\"type\":\"address\",\"indexed\":false}],\"name\":\"ErrCannotBailout\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrExceedsMaxNumberOfCandidate\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_bridgeOperatorAddr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"ErrExistentBridgeOperator\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrExistentCandidate\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_candidateAdminAddr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"ErrExistentCandidateAdmin\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_treasuryAddr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"ErrExistentTreasury\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrInsufficientBalance\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrInvalidCommissionRate\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrInvalidEffectiveDaysOnwards\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrInvalidMaxPrioritizedValidatorNumber\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrInvalidMinEffectiveDaysOnwards\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrNonExistentCandidate\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrRecipientRevert\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrTrustedOrgCannotRenounce\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrUnauthorizedReceiveRON\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"ErrZeroCodeContract\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[],\"name\":\"NonExistentRecyclingInfo\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"error\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"period\",\"type\":\"uint256\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"epoch\",\"type\":\"uint256\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"consensusAddrs\",\"type\":\"address[]\",\"indexed\":false}],\"name\":\"BlockProducerSetUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"coinbaseAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"rewardAmount\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"deprecatedType\",\"type\":\"uint8\",\"indexed\":false}],\"name\":\"BlockRewardDeprecated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"coinbaseAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"submittedAmount\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"bonusAmount\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"BlockRewardSubmitted\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"consensusAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"bridgeOperator\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"recipientAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"BridgeOperatorRewardDistributed\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"consensusAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"bridgeOperator\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"contractBalance\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"BridgeOperatorRewardDistributionFailed\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"period\",\"type\":\"uint256\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"epoch\",\"type\":\"uint256\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"bridgeOperators\",\"type\":\"address[]\",\"indexed\":false}],\"name\":\"BridgeOperatorSetUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"name\":\"BridgeTrackingContractUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[],\"name\":\"BridgeTrackingIncorrectlyResponded\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"consensusAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"treasuryAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"admin\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"bridgeOperator\",\"type\":\"address\",\"indexed\":false}],\"name\":\"CandidateGranted\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"consensusAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"revokingTimestamp\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"CandidateRevokingTimestampUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"consensusAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"topupDeadline\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"CandidateTopupDeadlineUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"consensusAddrs\",\"type\":\"address[]\",\"indexed\":false}],\"name\":\"CandidatesRevoked\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"consensusAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"effectiveTimestamp\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"rate\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"CommissionRateUpdateScheduled\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"consensusAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"rate\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"CommissionRateUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"recipientAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"balance\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"DeprecatedRewardRecycleFailed\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"recipientAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"DeprecatedRewardRecycled\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"EmergencyExitLockedAmountUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"consensusAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"unlockedAmount\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"EmergencyExitLockedFundReleased\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"consensusAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"unlockedAmount\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"contractBalance\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"EmergencyExitLockedFundReleasingFailed\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"consensusAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"lockedAmount\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"EmergencyExitRequested\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"EmergencyExpiryDurationUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false}],\"name\":\"Initialized\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"name\":\"MaintenanceContractUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"MaxPrioritizedValidatorNumberUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"threshold\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"MaxValidatorCandidateUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"MaxValidatorNumberUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"numOfDays\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"MinEffectiveDaysOnwardsUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"consensusAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"MiningRewardDistributed\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"consensusAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"contractBalance\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"MiningRewardDistributionFailed\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"name\":\"RoninTrustedOrganizationContractUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"name\":\"SlashIndicatorContractUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"name\":\"StakingContractUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"totalAmount\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"consensusAddrs\",\"type\":\"address[]\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"amounts\",\"type\":\"uint256[]\",\"indexed\":false}],\"name\":\"StakingRewardDistributed\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"totalAmount\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"consensusAddrs\",\"type\":\"address[]\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"amounts\",\"type\":\"uint256[]\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"contractBalance\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"StakingRewardDistributionFailed\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"name\":\"StakingVestingContractUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"consensusAddr\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"period\",\"type\":\"uint256\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"jailedUntil\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"deductedStakingAmount\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"blockProducerRewardDeprecated\",\"type\":\"bool\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"bridgeOperatorRewardDeprecated\",\"type\":\"bool\",\"indexed\":false}],\"name\":\"ValidatorPunished\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"period\",\"type\":\"uint256\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"consensusAddrs\",\"type\":\"address[]\",\"indexed\":false}],\"name\":\"ValidatorSetUpdated\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"validator\",\"type\":\"address\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"period\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"ValidatorUnjailed\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"periodNumber\",\"type\":\"uint256\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"epochNumber\",\"type\":\"uint256\",\"indexed\":true},{\"internal_type\":\"\",\"name\":\"periodEnding\",\"type\":\"bool\",\"indexed\":false}],\"name\":\"WrappedUpEpoch\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"\",\"type\":\"event\",\"anonymous\":false},{\"inputs\":null,\"name\":\"\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"payable\",\"type\":\"fallback\",\"anonymous\":false},{\"inputs\":[],\"name\":\"DEFAULT_ADDITION_GAS\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"PERIOD_DURATION\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"bridgeTrackingContract\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_consensusAddr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"checkBridgeRewardDeprecatedAtLatestPeriod\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_result\",\"type\":\"bool\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_consensusAddr\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_period\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"checkBridgeRewardDeprecatedAtPeriod\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_result\",\"type\":\"bool\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_addr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"checkJailed\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bool\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_addr\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_blockNum\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"checkJailedAtBlock\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bool\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_addrList\",\"type\":\"address[]\",\"indexed\":false}],\"name\":\"checkManyJailed\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_result\",\"type\":\"bool[]\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_blockProducer\",\"type\":\"address\",\"indexed\":false}],\"name\":\"checkMiningRewardDeprecated\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_result\",\"type\":\"bool\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_blockProducer\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_period\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"checkMiningRewardDeprecatedAtPeriod\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_result\",\"type\":\"bool\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"currentPeriod\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"currentPeriodStartAtBlock\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"emergencyExitLockedAmount\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"emergencyExpiryDuration\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_block\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"epochEndingAt\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bool\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_block\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"epochOf\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_candidateAdmin\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_consensusAddr\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_treasuryAddr\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_bridgeOperatorAddr\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_commissionRate\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"execApplyValidatorCandidate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_validatorAddr\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_period\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"execBailOut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_consensusAddr\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_secLeftToRevoke\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"execEmergencyExit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_consensusAddr\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_recipient\",\"type\":\"address\",\"indexed\":false}],\"name\":\"execReleaseLockedFundForEmergencyExitRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_consensusAddr\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_secsLeft\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"execRequestRenounceCandidate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_consensusAddr\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_effectiveDaysOnwards\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_commissionRate\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"execRequestUpdateCommissionRate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_validatorAddr\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_newJailedUntil\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_slashAmount\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_cannotBailout\",\"type\":\"bool\",\"indexed\":false}],\"name\":\"execSlash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"getBlockProducers\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_result\",\"type\":\"address[]\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"getBridgeOperators\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_result\",\"type\":\"address[]\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_validatorAddrs\",\"type\":\"address[]\",\"indexed\":false}],\"name\":\"getBridgeOperatorsOf\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_result\",\"type\":\"address[]\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_candidate\",\"type\":\"address\",\"indexed\":false}],\"name\":\"getCandidateInfo\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"internal_type\":\"\",\"name\":\"admin\",\"type\":\"address\"},{\"internal_type\":\"\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internal_type\":\"\",\"name\":\"treasuryAddr\",\"type\":\"address\"},{\"internal_type\":\"\",\"name\":\"bridgeOperatorAddr\",\"type\":\"address\"},{\"internal_type\":\"\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internal_type\":\"\",\"name\":\"revokingTimestamp\",\"type\":\"uint256\"},{\"internal_type\":\"\",\"name\":\"topupDeadline\",\"type\":\"uint256\"}],\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"getCandidateInfos\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_list\",\"type\":\"tuple[]\",\"components\":[{\"internal_type\":\"\",\"name\":\"admin\",\"type\":\"address\"},{\"internal_type\":\"\",\"name\":\"consensusAddr\",\"type\":\"address\"},{\"internal_type\":\"\",\"name\":\"treasuryAddr\",\"type\":\"address\"},{\"internal_type\":\"\",\"name\":\"bridgeOperatorAddr\",\"type\":\"address\"},{\"internal_type\":\"\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internal_type\":\"\",\"name\":\"revokingTimestamp\",\"type\":\"uint256\"},{\"internal_type\":\"\",\"name\":\"topupDeadline\",\"type\":\"uint256\"}],\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_candidate\",\"type\":\"address\",\"indexed\":false}],\"name\":\"getCommissionChangeSchedule\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"tuple\",\"components\":[{\"internal_type\":\"\",\"name\":\"effectiveTimestamp\",\"type\":\"uint256\"},{\"internal_type\":\"\",\"name\":\"commissionRate\",\"type\":\"uint256\"}],\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_consensusAddr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"getEmergencyExitInfo\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_info\",\"type\":\"tuple\",\"components\":[{\"internal_type\":\"\",\"name\":\"lockedAmount\",\"type\":\"uint256\"},{\"internal_type\":\"\",\"name\":\"recyclingAt\",\"type\":\"uint256\"}],\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_addr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"getJailedTimeLeft\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"isJailed_\",\"type\":\"bool\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"blockLeft_\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"epochLeft_\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_addr\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_blockNum\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"getJailedTimeLeftAtBlock\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"isJailed_\",\"type\":\"bool\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"blockLeft_\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"epochLeft_\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"getLastUpdatedBlock\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"getValidatorCandidates\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address[]\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_validatorList\",\"type\":\"address[]\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_bridgeOperators\",\"type\":\"address[]\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_flags\",\"type\":\"uint8[]\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"__slashIndicatorContract\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"__stakingContract\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"__stakingVestingContract\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"__maintenanceContract\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"__roninTrustedOrganizationContract\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"__bridgeTrackingContract\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"__maxValidatorNumber\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"__maxValidatorCandidate\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"__maxPrioritizedValidatorNumber\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"__minEffectiveDaysOnwards\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"__numberOfBlocksInEpoch\",\"type\":\"uint256\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"__emergencyExitConfigs\",\"type\":\"uint256[2]\",\"indexed\":false}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_addr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"isBlockProducer\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bool\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_bridgeOperatorAddr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"isBridgeOperator\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_isOperator\",\"type\":\"bool\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_candidate\",\"type\":\"address\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_admin\",\"type\":\"address\",\"indexed\":false}],\"name\":\"isCandidateAdmin\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bool\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_consensusAddr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"isOperatingBridge\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bool\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"isPeriodEnding\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bool\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_addr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"isValidator\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bool\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_addr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"isValidatorCandidate\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"bool\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"maintenanceContract\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"maxPrioritizedValidatorNumber\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_maximumPrioritizedValidatorNumber\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"maxValidatorCandidate\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"maxValidatorNumber\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_maximumValidatorNumber\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"minEffectiveDaysOnwards\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"numberOfBlocksInEpoch\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_numberOfBlocks\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"precompilePickValidatorSetAddress\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"precompileSortValidatorsAddress\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"roninTrustedOrganizationContract\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_addr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"setBridgeTrackingContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_emergencyExitLockedAmount\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"setEmergencyExitLockedAmount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_emergencyExpiryDuration\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"setEmergencyExpiryDuration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_addr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"setMaintenanceContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_number\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"setMaxPrioritizedValidatorNumber\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_number\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"setMaxValidatorCandidate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_max\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"setMaxValidatorNumber\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_numOfDays\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"setMinEffectiveDaysOnwards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_addr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"setRoninTrustedOrganizationContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_addr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"setSlashIndicatorContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_addr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"setStakingContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_addr\",\"type\":\"address\",\"indexed\":false}],\"name\":\"setStakingVestingContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"slashIndicatorContract\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"stakingContract\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"stakingVestingContract\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"address\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"submitBlockReward\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"payable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"totalBlockProducers\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_total\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"totalBridgeOperators\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_total\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"totalDeprecatedReward\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[{\"internal_type\":\"\",\"name\":\"_epoch\",\"type\":\"uint256\",\"indexed\":false}],\"name\":\"tryGetPeriodOfEpoch\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"_filled\",\"type\":\"bool\",\"indexed\":false},{\"internal_type\":\"\",\"name\":\"_periodNumber\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"validatorCount\",\"outputs\":[{\"internal_type\":\"\",\"name\":\"\",\"type\":\"uint256\",\"indexed\":false}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":[],\"name\":\"wrapUpEpoch\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"payable\",\"type\":\"function\",\"anonymous\":false},{\"inputs\":null,\"name\":\"\",\"outputs\":null,\"payable\":false,\"stateMutability\":\"payable\",\"type\":\"receive\",\"anonymous\":false}]",
}

// ValidatorABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorMetaData.ABI instead.
var ValidatorABI = ValidatorMetaData.ABI

// Validator is an auto generated Go binding around an Ethereum contract.
type Validator struct {
	ValidatorCaller     // Read-only binding to the contract
	ValidatorTransactor // Write-only binding to the contract
	ValidatorFilterer   // Log filterer for contract events
}

// ValidatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorSession struct {
	Contract     *Validator        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorCallerSession struct {
	Contract *ValidatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ValidatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorTransactorSession struct {
	Contract     *ValidatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ValidatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorRaw struct {
	Contract *Validator // Generic contract binding to access the raw methods on
}

// ValidatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorCallerRaw struct {
	Contract *ValidatorCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorTransactorRaw struct {
	Contract *ValidatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidator creates a new instance of Validator, bound to a specific deployed contract.
func NewValidator(address common.Address, backend bind.ContractBackend) (*Validator, error) {
	contract, err := bindValidator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validator{ValidatorCaller: ValidatorCaller{contract: contract}, ValidatorTransactor: ValidatorTransactor{contract: contract}, ValidatorFilterer: ValidatorFilterer{contract: contract}}, nil
}

// NewValidatorCaller creates a new read-only instance of Validator, bound to a specific deployed contract.
func NewValidatorCaller(address common.Address, caller bind.ContractCaller) (*ValidatorCaller, error) {
	contract, err := bindValidator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorCaller{contract: contract}, nil
}

// NewValidatorTransactor creates a new write-only instance of Validator, bound to a specific deployed contract.
func NewValidatorTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorTransactor, error) {
	contract, err := bindValidator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorTransactor{contract: contract}, nil
}

// NewValidatorFilterer creates a new log filterer instance of Validator, bound to a specific deployed contract.
func NewValidatorFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorFilterer, error) {
	contract, err := bindValidator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorFilterer{contract: contract}, nil
}

// bindValidator binds a generic wrapper to an already deployed contract.
func bindValidator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.ValidatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.ValidatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validator *ValidatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validator *ValidatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validator *ValidatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validator.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADDITIONGAS is a free data retrieval call binding the contract method 0x03827884.
//
// Solidity: function DEFAULT_ADDITION_GAS() view returns(uint256)
func (_Validator *ValidatorCaller) DEFAULTADDITIONGAS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "DEFAULT_ADDITION_GAS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTADDITIONGAS is a free data retrieval call binding the contract method 0x03827884.
//
// Solidity: function DEFAULT_ADDITION_GAS() view returns(uint256)
func (_Validator *ValidatorSession) DEFAULTADDITIONGAS() (*big.Int, error) {
	return _Validator.Contract.DEFAULTADDITIONGAS(&_Validator.CallOpts)
}

// DEFAULTADDITIONGAS is a free data retrieval call binding the contract method 0x03827884.
//
// Solidity: function DEFAULT_ADDITION_GAS() view returns(uint256)
func (_Validator *ValidatorCallerSession) DEFAULTADDITIONGAS() (*big.Int, error) {
	return _Validator.Contract.DEFAULTADDITIONGAS(&_Validator.CallOpts)
}

// PERIODDURATION is a free data retrieval call binding the contract method 0x6558954f.
//
// Solidity: function PERIOD_DURATION() view returns(uint256)
func (_Validator *ValidatorCaller) PERIODDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "PERIOD_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PERIODDURATION is a free data retrieval call binding the contract method 0x6558954f.
//
// Solidity: function PERIOD_DURATION() view returns(uint256)
func (_Validator *ValidatorSession) PERIODDURATION() (*big.Int, error) {
	return _Validator.Contract.PERIODDURATION(&_Validator.CallOpts)
}

// PERIODDURATION is a free data retrieval call binding the contract method 0x6558954f.
//
// Solidity: function PERIOD_DURATION() view returns(uint256)
func (_Validator *ValidatorCallerSession) PERIODDURATION() (*big.Int, error) {
	return _Validator.Contract.PERIODDURATION(&_Validator.CallOpts)
}

// BridgeTrackingContract is a free data retrieval call binding the contract method 0x4493421e.
//
// Solidity: function bridgeTrackingContract() view returns(address)
func (_Validator *ValidatorCaller) BridgeTrackingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "bridgeTrackingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeTrackingContract is a free data retrieval call binding the contract method 0x4493421e.
//
// Solidity: function bridgeTrackingContract() view returns(address)
func (_Validator *ValidatorSession) BridgeTrackingContract() (common.Address, error) {
	return _Validator.Contract.BridgeTrackingContract(&_Validator.CallOpts)
}

// BridgeTrackingContract is a free data retrieval call binding the contract method 0x4493421e.
//
// Solidity: function bridgeTrackingContract() view returns(address)
func (_Validator *ValidatorCallerSession) BridgeTrackingContract() (common.Address, error) {
	return _Validator.Contract.BridgeTrackingContract(&_Validator.CallOpts)
}

// CheckBridgeRewardDeprecatedAtLatestPeriod is a free data retrieval call binding the contract method 0x1ab4a34c.
//
// Solidity: function checkBridgeRewardDeprecatedAtLatestPeriod(address _consensusAddr) view returns(bool _result)
func (_Validator *ValidatorCaller) CheckBridgeRewardDeprecatedAtLatestPeriod(opts *bind.CallOpts, _consensusAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "checkBridgeRewardDeprecatedAtLatestPeriod", _consensusAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckBridgeRewardDeprecatedAtLatestPeriod is a free data retrieval call binding the contract method 0x1ab4a34c.
//
// Solidity: function checkBridgeRewardDeprecatedAtLatestPeriod(address _consensusAddr) view returns(bool _result)
func (_Validator *ValidatorSession) CheckBridgeRewardDeprecatedAtLatestPeriod(_consensusAddr common.Address) (bool, error) {
	return _Validator.Contract.CheckBridgeRewardDeprecatedAtLatestPeriod(&_Validator.CallOpts, _consensusAddr)
}

// CheckBridgeRewardDeprecatedAtLatestPeriod is a free data retrieval call binding the contract method 0x1ab4a34c.
//
// Solidity: function checkBridgeRewardDeprecatedAtLatestPeriod(address _consensusAddr) view returns(bool _result)
func (_Validator *ValidatorCallerSession) CheckBridgeRewardDeprecatedAtLatestPeriod(_consensusAddr common.Address) (bool, error) {
	return _Validator.Contract.CheckBridgeRewardDeprecatedAtLatestPeriod(&_Validator.CallOpts, _consensusAddr)
}

// CheckBridgeRewardDeprecatedAtPeriod is a free data retrieval call binding the contract method 0xd5a0744f.
//
// Solidity: function checkBridgeRewardDeprecatedAtPeriod(address _consensusAddr, uint256 _period) view returns(bool _result)
func (_Validator *ValidatorCaller) CheckBridgeRewardDeprecatedAtPeriod(opts *bind.CallOpts, _consensusAddr common.Address, _period *big.Int) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "checkBridgeRewardDeprecatedAtPeriod", _consensusAddr, _period)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckBridgeRewardDeprecatedAtPeriod is a free data retrieval call binding the contract method 0xd5a0744f.
//
// Solidity: function checkBridgeRewardDeprecatedAtPeriod(address _consensusAddr, uint256 _period) view returns(bool _result)
func (_Validator *ValidatorSession) CheckBridgeRewardDeprecatedAtPeriod(_consensusAddr common.Address, _period *big.Int) (bool, error) {
	return _Validator.Contract.CheckBridgeRewardDeprecatedAtPeriod(&_Validator.CallOpts, _consensusAddr, _period)
}

// CheckBridgeRewardDeprecatedAtPeriod is a free data retrieval call binding the contract method 0xd5a0744f.
//
// Solidity: function checkBridgeRewardDeprecatedAtPeriod(address _consensusAddr, uint256 _period) view returns(bool _result)
func (_Validator *ValidatorCallerSession) CheckBridgeRewardDeprecatedAtPeriod(_consensusAddr common.Address, _period *big.Int) (bool, error) {
	return _Validator.Contract.CheckBridgeRewardDeprecatedAtPeriod(&_Validator.CallOpts, _consensusAddr, _period)
}

// CheckJailed is a free data retrieval call binding the contract method 0x2924de71.
//
// Solidity: function checkJailed(address _addr) view returns(bool)
func (_Validator *ValidatorCaller) CheckJailed(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "checkJailed", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckJailed is a free data retrieval call binding the contract method 0x2924de71.
//
// Solidity: function checkJailed(address _addr) view returns(bool)
func (_Validator *ValidatorSession) CheckJailed(_addr common.Address) (bool, error) {
	return _Validator.Contract.CheckJailed(&_Validator.CallOpts, _addr)
}

// CheckJailed is a free data retrieval call binding the contract method 0x2924de71.
//
// Solidity: function checkJailed(address _addr) view returns(bool)
func (_Validator *ValidatorCallerSession) CheckJailed(_addr common.Address) (bool, error) {
	return _Validator.Contract.CheckJailed(&_Validator.CallOpts, _addr)
}

// CheckJailedAtBlock is a free data retrieval call binding the contract method 0x23c65eb0.
//
// Solidity: function checkJailedAtBlock(address _addr, uint256 _blockNum) view returns(bool)
func (_Validator *ValidatorCaller) CheckJailedAtBlock(opts *bind.CallOpts, _addr common.Address, _blockNum *big.Int) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "checkJailedAtBlock", _addr, _blockNum)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckJailedAtBlock is a free data retrieval call binding the contract method 0x23c65eb0.
//
// Solidity: function checkJailedAtBlock(address _addr, uint256 _blockNum) view returns(bool)
func (_Validator *ValidatorSession) CheckJailedAtBlock(_addr common.Address, _blockNum *big.Int) (bool, error) {
	return _Validator.Contract.CheckJailedAtBlock(&_Validator.CallOpts, _addr, _blockNum)
}

// CheckJailedAtBlock is a free data retrieval call binding the contract method 0x23c65eb0.
//
// Solidity: function checkJailedAtBlock(address _addr, uint256 _blockNum) view returns(bool)
func (_Validator *ValidatorCallerSession) CheckJailedAtBlock(_addr common.Address, _blockNum *big.Int) (bool, error) {
	return _Validator.Contract.CheckJailedAtBlock(&_Validator.CallOpts, _addr, _blockNum)
}

// CheckManyJailed is a free data retrieval call binding the contract method 0x4de2b735.
//
// Solidity: function checkManyJailed(address[] _addrList) view returns(bool[] _result)
func (_Validator *ValidatorCaller) CheckManyJailed(opts *bind.CallOpts, _addrList []common.Address) ([]bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "checkManyJailed", _addrList)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// CheckManyJailed is a free data retrieval call binding the contract method 0x4de2b735.
//
// Solidity: function checkManyJailed(address[] _addrList) view returns(bool[] _result)
func (_Validator *ValidatorSession) CheckManyJailed(_addrList []common.Address) ([]bool, error) {
	return _Validator.Contract.CheckManyJailed(&_Validator.CallOpts, _addrList)
}

// CheckManyJailed is a free data retrieval call binding the contract method 0x4de2b735.
//
// Solidity: function checkManyJailed(address[] _addrList) view returns(bool[] _result)
func (_Validator *ValidatorCallerSession) CheckManyJailed(_addrList []common.Address) ([]bool, error) {
	return _Validator.Contract.CheckManyJailed(&_Validator.CallOpts, _addrList)
}

// CheckMiningRewardDeprecated is a free data retrieval call binding the contract method 0x873a5a70.
//
// Solidity: function checkMiningRewardDeprecated(address _blockProducer) view returns(bool _result)
func (_Validator *ValidatorCaller) CheckMiningRewardDeprecated(opts *bind.CallOpts, _blockProducer common.Address) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "checkMiningRewardDeprecated", _blockProducer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckMiningRewardDeprecated is a free data retrieval call binding the contract method 0x873a5a70.
//
// Solidity: function checkMiningRewardDeprecated(address _blockProducer) view returns(bool _result)
func (_Validator *ValidatorSession) CheckMiningRewardDeprecated(_blockProducer common.Address) (bool, error) {
	return _Validator.Contract.CheckMiningRewardDeprecated(&_Validator.CallOpts, _blockProducer)
}

// CheckMiningRewardDeprecated is a free data retrieval call binding the contract method 0x873a5a70.
//
// Solidity: function checkMiningRewardDeprecated(address _blockProducer) view returns(bool _result)
func (_Validator *ValidatorCallerSession) CheckMiningRewardDeprecated(_blockProducer common.Address) (bool, error) {
	return _Validator.Contract.CheckMiningRewardDeprecated(&_Validator.CallOpts, _blockProducer)
}

// CheckMiningRewardDeprecatedAtPeriod is a free data retrieval call binding the contract method 0x31a8aef5.
//
// Solidity: function checkMiningRewardDeprecatedAtPeriod(address _blockProducer, uint256 _period) view returns(bool _result)
func (_Validator *ValidatorCaller) CheckMiningRewardDeprecatedAtPeriod(opts *bind.CallOpts, _blockProducer common.Address, _period *big.Int) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "checkMiningRewardDeprecatedAtPeriod", _blockProducer, _period)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckMiningRewardDeprecatedAtPeriod is a free data retrieval call binding the contract method 0x31a8aef5.
//
// Solidity: function checkMiningRewardDeprecatedAtPeriod(address _blockProducer, uint256 _period) view returns(bool _result)
func (_Validator *ValidatorSession) CheckMiningRewardDeprecatedAtPeriod(_blockProducer common.Address, _period *big.Int) (bool, error) {
	return _Validator.Contract.CheckMiningRewardDeprecatedAtPeriod(&_Validator.CallOpts, _blockProducer, _period)
}

// CheckMiningRewardDeprecatedAtPeriod is a free data retrieval call binding the contract method 0x31a8aef5.
//
// Solidity: function checkMiningRewardDeprecatedAtPeriod(address _blockProducer, uint256 _period) view returns(bool _result)
func (_Validator *ValidatorCallerSession) CheckMiningRewardDeprecatedAtPeriod(_blockProducer common.Address, _period *big.Int) (bool, error) {
	return _Validator.Contract.CheckMiningRewardDeprecatedAtPeriod(&_Validator.CallOpts, _blockProducer, _period)
}

// CurrentPeriod is a free data retrieval call binding the contract method 0x06040618.
//
// Solidity: function currentPeriod() view returns(uint256)
func (_Validator *ValidatorCaller) CurrentPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "currentPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPeriod is a free data retrieval call binding the contract method 0x06040618.
//
// Solidity: function currentPeriod() view returns(uint256)
func (_Validator *ValidatorSession) CurrentPeriod() (*big.Int, error) {
	return _Validator.Contract.CurrentPeriod(&_Validator.CallOpts)
}

// CurrentPeriod is a free data retrieval call binding the contract method 0x06040618.
//
// Solidity: function currentPeriod() view returns(uint256)
func (_Validator *ValidatorCallerSession) CurrentPeriod() (*big.Int, error) {
	return _Validator.Contract.CurrentPeriod(&_Validator.CallOpts)
}

// CurrentPeriodStartAtBlock is a free data retrieval call binding the contract method 0x297a8fca.
//
// Solidity: function currentPeriodStartAtBlock() view returns(uint256)
func (_Validator *ValidatorCaller) CurrentPeriodStartAtBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "currentPeriodStartAtBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPeriodStartAtBlock is a free data retrieval call binding the contract method 0x297a8fca.
//
// Solidity: function currentPeriodStartAtBlock() view returns(uint256)
func (_Validator *ValidatorSession) CurrentPeriodStartAtBlock() (*big.Int, error) {
	return _Validator.Contract.CurrentPeriodStartAtBlock(&_Validator.CallOpts)
}

// CurrentPeriodStartAtBlock is a free data retrieval call binding the contract method 0x297a8fca.
//
// Solidity: function currentPeriodStartAtBlock() view returns(uint256)
func (_Validator *ValidatorCallerSession) CurrentPeriodStartAtBlock() (*big.Int, error) {
	return _Validator.Contract.CurrentPeriodStartAtBlock(&_Validator.CallOpts)
}

// EmergencyExitLockedAmount is a free data retrieval call binding the contract method 0x690b7536.
//
// Solidity: function emergencyExitLockedAmount() view returns(uint256)
func (_Validator *ValidatorCaller) EmergencyExitLockedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "emergencyExitLockedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmergencyExitLockedAmount is a free data retrieval call binding the contract method 0x690b7536.
//
// Solidity: function emergencyExitLockedAmount() view returns(uint256)
func (_Validator *ValidatorSession) EmergencyExitLockedAmount() (*big.Int, error) {
	return _Validator.Contract.EmergencyExitLockedAmount(&_Validator.CallOpts)
}

// EmergencyExitLockedAmount is a free data retrieval call binding the contract method 0x690b7536.
//
// Solidity: function emergencyExitLockedAmount() view returns(uint256)
func (_Validator *ValidatorCallerSession) EmergencyExitLockedAmount() (*big.Int, error) {
	return _Validator.Contract.EmergencyExitLockedAmount(&_Validator.CallOpts)
}

// EmergencyExpiryDuration is a free data retrieval call binding the contract method 0xa66c0f77.
//
// Solidity: function emergencyExpiryDuration() view returns(uint256)
func (_Validator *ValidatorCaller) EmergencyExpiryDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "emergencyExpiryDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmergencyExpiryDuration is a free data retrieval call binding the contract method 0xa66c0f77.
//
// Solidity: function emergencyExpiryDuration() view returns(uint256)
func (_Validator *ValidatorSession) EmergencyExpiryDuration() (*big.Int, error) {
	return _Validator.Contract.EmergencyExpiryDuration(&_Validator.CallOpts)
}

// EmergencyExpiryDuration is a free data retrieval call binding the contract method 0xa66c0f77.
//
// Solidity: function emergencyExpiryDuration() view returns(uint256)
func (_Validator *ValidatorCallerSession) EmergencyExpiryDuration() (*big.Int, error) {
	return _Validator.Contract.EmergencyExpiryDuration(&_Validator.CallOpts)
}

// EpochEndingAt is a free data retrieval call binding the contract method 0x7593ff71.
//
// Solidity: function epochEndingAt(uint256 _block) view returns(bool)
func (_Validator *ValidatorCaller) EpochEndingAt(opts *bind.CallOpts, _block *big.Int) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "epochEndingAt", _block)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EpochEndingAt is a free data retrieval call binding the contract method 0x7593ff71.
//
// Solidity: function epochEndingAt(uint256 _block) view returns(bool)
func (_Validator *ValidatorSession) EpochEndingAt(_block *big.Int) (bool, error) {
	return _Validator.Contract.EpochEndingAt(&_Validator.CallOpts, _block)
}

// EpochEndingAt is a free data retrieval call binding the contract method 0x7593ff71.
//
// Solidity: function epochEndingAt(uint256 _block) view returns(bool)
func (_Validator *ValidatorCallerSession) EpochEndingAt(_block *big.Int) (bool, error) {
	return _Validator.Contract.EpochEndingAt(&_Validator.CallOpts, _block)
}

// EpochOf is a free data retrieval call binding the contract method 0xa3d545f5.
//
// Solidity: function epochOf(uint256 _block) view returns(uint256)
func (_Validator *ValidatorCaller) EpochOf(opts *bind.CallOpts, _block *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "epochOf", _block)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochOf is a free data retrieval call binding the contract method 0xa3d545f5.
//
// Solidity: function epochOf(uint256 _block) view returns(uint256)
func (_Validator *ValidatorSession) EpochOf(_block *big.Int) (*big.Int, error) {
	return _Validator.Contract.EpochOf(&_Validator.CallOpts, _block)
}

// EpochOf is a free data retrieval call binding the contract method 0xa3d545f5.
//
// Solidity: function epochOf(uint256 _block) view returns(uint256)
func (_Validator *ValidatorCallerSession) EpochOf(_block *big.Int) (*big.Int, error) {
	return _Validator.Contract.EpochOf(&_Validator.CallOpts, _block)
}

// GetBlockProducers is a free data retrieval call binding the contract method 0x49096d26.
//
// Solidity: function getBlockProducers() view returns(address[] _result)
func (_Validator *ValidatorCaller) GetBlockProducers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getBlockProducers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBlockProducers is a free data retrieval call binding the contract method 0x49096d26.
//
// Solidity: function getBlockProducers() view returns(address[] _result)
func (_Validator *ValidatorSession) GetBlockProducers() ([]common.Address, error) {
	return _Validator.Contract.GetBlockProducers(&_Validator.CallOpts)
}

// GetBlockProducers is a free data retrieval call binding the contract method 0x49096d26.
//
// Solidity: function getBlockProducers() view returns(address[] _result)
func (_Validator *ValidatorCallerSession) GetBlockProducers() ([]common.Address, error) {
	return _Validator.Contract.GetBlockProducers(&_Validator.CallOpts)
}

// GetBridgeOperators is a free data retrieval call binding the contract method 0x9b19dbfd.
//
// Solidity: function getBridgeOperators() view returns(address[] _result)
func (_Validator *ValidatorCaller) GetBridgeOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getBridgeOperators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBridgeOperators is a free data retrieval call binding the contract method 0x9b19dbfd.
//
// Solidity: function getBridgeOperators() view returns(address[] _result)
func (_Validator *ValidatorSession) GetBridgeOperators() ([]common.Address, error) {
	return _Validator.Contract.GetBridgeOperators(&_Validator.CallOpts)
}

// GetBridgeOperators is a free data retrieval call binding the contract method 0x9b19dbfd.
//
// Solidity: function getBridgeOperators() view returns(address[] _result)
func (_Validator *ValidatorCallerSession) GetBridgeOperators() ([]common.Address, error) {
	return _Validator.Contract.GetBridgeOperators(&_Validator.CallOpts)
}

// GetBridgeOperatorsOf is a free data retrieval call binding the contract method 0x4244d4c9.
//
// Solidity: function getBridgeOperatorsOf(address[] _validatorAddrs) view returns(address[] _result)
func (_Validator *ValidatorCaller) GetBridgeOperatorsOf(opts *bind.CallOpts, _validatorAddrs []common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getBridgeOperatorsOf", _validatorAddrs)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetBridgeOperatorsOf is a free data retrieval call binding the contract method 0x4244d4c9.
//
// Solidity: function getBridgeOperatorsOf(address[] _validatorAddrs) view returns(address[] _result)
func (_Validator *ValidatorSession) GetBridgeOperatorsOf(_validatorAddrs []common.Address) ([]common.Address, error) {
	return _Validator.Contract.GetBridgeOperatorsOf(&_Validator.CallOpts, _validatorAddrs)
}

// GetBridgeOperatorsOf is a free data retrieval call binding the contract method 0x4244d4c9.
//
// Solidity: function getBridgeOperatorsOf(address[] _validatorAddrs) view returns(address[] _result)
func (_Validator *ValidatorCallerSession) GetBridgeOperatorsOf(_validatorAddrs []common.Address) ([]common.Address, error) {
	return _Validator.Contract.GetBridgeOperatorsOf(&_Validator.CallOpts, _validatorAddrs)
}

// GetCandidateInfo is a free data retrieval call binding the contract method 0x28bde1e1.
//
// Solidity: function getCandidateInfo(address _candidate) view returns((address,address,address,address,uint256,uint256,uint256))
func (_Validator *ValidatorCaller) GetCandidateInfo(opts *bind.CallOpts, _candidate common.Address) (Struct0, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getCandidateInfo", _candidate)

	if err != nil {
		return *new(Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct0)).(*Struct0)

	return out0, err

}

// GetCandidateInfo is a free data retrieval call binding the contract method 0x28bde1e1.
//
// Solidity: function getCandidateInfo(address _candidate) view returns((address,address,address,address,uint256,uint256,uint256))
func (_Validator *ValidatorSession) GetCandidateInfo(_candidate common.Address) (Struct0, error) {
	return _Validator.Contract.GetCandidateInfo(&_Validator.CallOpts, _candidate)
}

// GetCandidateInfo is a free data retrieval call binding the contract method 0x28bde1e1.
//
// Solidity: function getCandidateInfo(address _candidate) view returns((address,address,address,address,uint256,uint256,uint256))
func (_Validator *ValidatorCallerSession) GetCandidateInfo(_candidate common.Address) (Struct0, error) {
	return _Validator.Contract.GetCandidateInfo(&_Validator.CallOpts, _candidate)
}

// GetCandidateInfos is a free data retrieval call binding the contract method 0x5248184a.
//
// Solidity: function getCandidateInfos() view returns((address,address,address,address,uint256,uint256,uint256)[] _list)
func (_Validator *ValidatorCaller) GetCandidateInfos(opts *bind.CallOpts) ([]Struct0, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getCandidateInfos")

	if err != nil {
		return *new([]Struct0), err
	}

	out0 := *abi.ConvertType(out[0], new([]Struct0)).(*[]Struct0)

	return out0, err

}

// GetCandidateInfos is a free data retrieval call binding the contract method 0x5248184a.
//
// Solidity: function getCandidateInfos() view returns((address,address,address,address,uint256,uint256,uint256)[] _list)
func (_Validator *ValidatorSession) GetCandidateInfos() ([]Struct0, error) {
	return _Validator.Contract.GetCandidateInfos(&_Validator.CallOpts)
}

// GetCandidateInfos is a free data retrieval call binding the contract method 0x5248184a.
//
// Solidity: function getCandidateInfos() view returns((address,address,address,address,uint256,uint256,uint256)[] _list)
func (_Validator *ValidatorCallerSession) GetCandidateInfos() ([]Struct0, error) {
	return _Validator.Contract.GetCandidateInfos(&_Validator.CallOpts)
}

// GetCommissionChangeSchedule is a free data retrieval call binding the contract method 0xedb194bb.
//
// Solidity: function getCommissionChangeSchedule(address _candidate) view returns((uint256,uint256))
func (_Validator *ValidatorCaller) GetCommissionChangeSchedule(opts *bind.CallOpts, _candidate common.Address) (Struct1, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getCommissionChangeSchedule", _candidate)

	if err != nil {
		return *new(Struct1), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct1)).(*Struct1)

	return out0, err

}

// GetCommissionChangeSchedule is a free data retrieval call binding the contract method 0xedb194bb.
//
// Solidity: function getCommissionChangeSchedule(address _candidate) view returns((uint256,uint256))
func (_Validator *ValidatorSession) GetCommissionChangeSchedule(_candidate common.Address) (Struct1, error) {
	return _Validator.Contract.GetCommissionChangeSchedule(&_Validator.CallOpts, _candidate)
}

// GetCommissionChangeSchedule is a free data retrieval call binding the contract method 0xedb194bb.
//
// Solidity: function getCommissionChangeSchedule(address _candidate) view returns((uint256,uint256))
func (_Validator *ValidatorCallerSession) GetCommissionChangeSchedule(_candidate common.Address) (Struct1, error) {
	return _Validator.Contract.GetCommissionChangeSchedule(&_Validator.CallOpts, _candidate)
}

// GetEmergencyExitInfo is a free data retrieval call binding the contract method 0x2d784a98.
//
// Solidity: function getEmergencyExitInfo(address _consensusAddr) view returns((uint256,uint256) _info)
func (_Validator *ValidatorCaller) GetEmergencyExitInfo(opts *bind.CallOpts, _consensusAddr common.Address) (Struct1, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getEmergencyExitInfo", _consensusAddr)

	if err != nil {
		return *new(Struct1), err
	}

	out0 := *abi.ConvertType(out[0], new(Struct1)).(*Struct1)

	return out0, err

}

// GetEmergencyExitInfo is a free data retrieval call binding the contract method 0x2d784a98.
//
// Solidity: function getEmergencyExitInfo(address _consensusAddr) view returns((uint256,uint256) _info)
func (_Validator *ValidatorSession) GetEmergencyExitInfo(_consensusAddr common.Address) (Struct1, error) {
	return _Validator.Contract.GetEmergencyExitInfo(&_Validator.CallOpts, _consensusAddr)
}

// GetEmergencyExitInfo is a free data retrieval call binding the contract method 0x2d784a98.
//
// Solidity: function getEmergencyExitInfo(address _consensusAddr) view returns((uint256,uint256) _info)
func (_Validator *ValidatorCallerSession) GetEmergencyExitInfo(_consensusAddr common.Address) (Struct1, error) {
	return _Validator.Contract.GetEmergencyExitInfo(&_Validator.CallOpts, _consensusAddr)
}

// GetJailedTimeLeft is a free data retrieval call binding the contract method 0x96585fc2.
//
// Solidity: function getJailedTimeLeft(address _addr) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_Validator *ValidatorCaller) GetJailedTimeLeft(opts *bind.CallOpts, _addr common.Address) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getJailedTimeLeft", _addr)

	outstruct := new(struct {
		IsJailed  bool
		BlockLeft *big.Int
		EpochLeft *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsJailed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.BlockLeft = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EpochLeft = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetJailedTimeLeft is a free data retrieval call binding the contract method 0x96585fc2.
//
// Solidity: function getJailedTimeLeft(address _addr) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_Validator *ValidatorSession) GetJailedTimeLeft(_addr common.Address) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	return _Validator.Contract.GetJailedTimeLeft(&_Validator.CallOpts, _addr)
}

// GetJailedTimeLeft is a free data retrieval call binding the contract method 0x96585fc2.
//
// Solidity: function getJailedTimeLeft(address _addr) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_Validator *ValidatorCallerSession) GetJailedTimeLeft(_addr common.Address) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	return _Validator.Contract.GetJailedTimeLeft(&_Validator.CallOpts, _addr)
}

// GetJailedTimeLeftAtBlock is a free data retrieval call binding the contract method 0x11662dc2.
//
// Solidity: function getJailedTimeLeftAtBlock(address _addr, uint256 _blockNum) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_Validator *ValidatorCaller) GetJailedTimeLeftAtBlock(opts *bind.CallOpts, _addr common.Address, _blockNum *big.Int) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getJailedTimeLeftAtBlock", _addr, _blockNum)

	outstruct := new(struct {
		IsJailed  bool
		BlockLeft *big.Int
		EpochLeft *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsJailed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.BlockLeft = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EpochLeft = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetJailedTimeLeftAtBlock is a free data retrieval call binding the contract method 0x11662dc2.
//
// Solidity: function getJailedTimeLeftAtBlock(address _addr, uint256 _blockNum) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_Validator *ValidatorSession) GetJailedTimeLeftAtBlock(_addr common.Address, _blockNum *big.Int) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	return _Validator.Contract.GetJailedTimeLeftAtBlock(&_Validator.CallOpts, _addr, _blockNum)
}

// GetJailedTimeLeftAtBlock is a free data retrieval call binding the contract method 0x11662dc2.
//
// Solidity: function getJailedTimeLeftAtBlock(address _addr, uint256 _blockNum) view returns(bool isJailed_, uint256 blockLeft_, uint256 epochLeft_)
func (_Validator *ValidatorCallerSession) GetJailedTimeLeftAtBlock(_addr common.Address, _blockNum *big.Int) (struct {
	IsJailed  bool
	BlockLeft *big.Int
	EpochLeft *big.Int
}, error) {
	return _Validator.Contract.GetJailedTimeLeftAtBlock(&_Validator.CallOpts, _addr, _blockNum)
}

// GetLastUpdatedBlock is a free data retrieval call binding the contract method 0x87c891bd.
//
// Solidity: function getLastUpdatedBlock() view returns(uint256)
func (_Validator *ValidatorCaller) GetLastUpdatedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getLastUpdatedBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastUpdatedBlock is a free data retrieval call binding the contract method 0x87c891bd.
//
// Solidity: function getLastUpdatedBlock() view returns(uint256)
func (_Validator *ValidatorSession) GetLastUpdatedBlock() (*big.Int, error) {
	return _Validator.Contract.GetLastUpdatedBlock(&_Validator.CallOpts)
}

// GetLastUpdatedBlock is a free data retrieval call binding the contract method 0x87c891bd.
//
// Solidity: function getLastUpdatedBlock() view returns(uint256)
func (_Validator *ValidatorCallerSession) GetLastUpdatedBlock() (*big.Int, error) {
	return _Validator.Contract.GetLastUpdatedBlock(&_Validator.CallOpts)
}

// GetValidatorCandidates is a free data retrieval call binding the contract method 0xba77b06c.
//
// Solidity: function getValidatorCandidates() view returns(address[])
func (_Validator *ValidatorCaller) GetValidatorCandidates(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getValidatorCandidates")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidatorCandidates is a free data retrieval call binding the contract method 0xba77b06c.
//
// Solidity: function getValidatorCandidates() view returns(address[])
func (_Validator *ValidatorSession) GetValidatorCandidates() ([]common.Address, error) {
	return _Validator.Contract.GetValidatorCandidates(&_Validator.CallOpts)
}

// GetValidatorCandidates is a free data retrieval call binding the contract method 0xba77b06c.
//
// Solidity: function getValidatorCandidates() view returns(address[])
func (_Validator *ValidatorCallerSession) GetValidatorCandidates() ([]common.Address, error) {
	return _Validator.Contract.GetValidatorCandidates(&_Validator.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[] _validatorList, address[] _bridgeOperators, uint8[] _flags)
func (_Validator *ValidatorCaller) GetValidators(opts *bind.CallOpts) (struct {
	ValidatorList   []common.Address
	BridgeOperators []common.Address
	Flags           []uint8
}, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "getValidators")

	outstruct := new(struct {
		ValidatorList   []common.Address
		BridgeOperators []common.Address
		Flags           []uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ValidatorList = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.BridgeOperators = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Flags = *abi.ConvertType(out[2], new([]uint8)).(*[]uint8)

	return *outstruct, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[] _validatorList, address[] _bridgeOperators, uint8[] _flags)
func (_Validator *ValidatorSession) GetValidators() (struct {
	ValidatorList   []common.Address
	BridgeOperators []common.Address
	Flags           []uint8
}, error) {
	return _Validator.Contract.GetValidators(&_Validator.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[] _validatorList, address[] _bridgeOperators, uint8[] _flags)
func (_Validator *ValidatorCallerSession) GetValidators() (struct {
	ValidatorList   []common.Address
	BridgeOperators []common.Address
	Flags           []uint8
}, error) {
	return _Validator.Contract.GetValidators(&_Validator.CallOpts)
}

// IsBlockProducer is a free data retrieval call binding the contract method 0x65244ece.
//
// Solidity: function isBlockProducer(address _addr) view returns(bool)
func (_Validator *ValidatorCaller) IsBlockProducer(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "isBlockProducer", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlockProducer is a free data retrieval call binding the contract method 0x65244ece.
//
// Solidity: function isBlockProducer(address _addr) view returns(bool)
func (_Validator *ValidatorSession) IsBlockProducer(_addr common.Address) (bool, error) {
	return _Validator.Contract.IsBlockProducer(&_Validator.CallOpts, _addr)
}

// IsBlockProducer is a free data retrieval call binding the contract method 0x65244ece.
//
// Solidity: function isBlockProducer(address _addr) view returns(bool)
func (_Validator *ValidatorCallerSession) IsBlockProducer(_addr common.Address) (bool, error) {
	return _Validator.Contract.IsBlockProducer(&_Validator.CallOpts, _addr)
}

// IsBridgeOperator is a free data retrieval call binding the contract method 0xb405aaf2.
//
// Solidity: function isBridgeOperator(address _bridgeOperatorAddr) view returns(bool _isOperator)
func (_Validator *ValidatorCaller) IsBridgeOperator(opts *bind.CallOpts, _bridgeOperatorAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "isBridgeOperator", _bridgeOperatorAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBridgeOperator is a free data retrieval call binding the contract method 0xb405aaf2.
//
// Solidity: function isBridgeOperator(address _bridgeOperatorAddr) view returns(bool _isOperator)
func (_Validator *ValidatorSession) IsBridgeOperator(_bridgeOperatorAddr common.Address) (bool, error) {
	return _Validator.Contract.IsBridgeOperator(&_Validator.CallOpts, _bridgeOperatorAddr)
}

// IsBridgeOperator is a free data retrieval call binding the contract method 0xb405aaf2.
//
// Solidity: function isBridgeOperator(address _bridgeOperatorAddr) view returns(bool _isOperator)
func (_Validator *ValidatorCallerSession) IsBridgeOperator(_bridgeOperatorAddr common.Address) (bool, error) {
	return _Validator.Contract.IsBridgeOperator(&_Validator.CallOpts, _bridgeOperatorAddr)
}

// IsCandidateAdmin is a free data retrieval call binding the contract method 0x04d971ab.
//
// Solidity: function isCandidateAdmin(address _candidate, address _admin) view returns(bool)
func (_Validator *ValidatorCaller) IsCandidateAdmin(opts *bind.CallOpts, _candidate common.Address, _admin common.Address) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "isCandidateAdmin", _candidate, _admin)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCandidateAdmin is a free data retrieval call binding the contract method 0x04d971ab.
//
// Solidity: function isCandidateAdmin(address _candidate, address _admin) view returns(bool)
func (_Validator *ValidatorSession) IsCandidateAdmin(_candidate common.Address, _admin common.Address) (bool, error) {
	return _Validator.Contract.IsCandidateAdmin(&_Validator.CallOpts, _candidate, _admin)
}

// IsCandidateAdmin is a free data retrieval call binding the contract method 0x04d971ab.
//
// Solidity: function isCandidateAdmin(address _candidate, address _admin) view returns(bool)
func (_Validator *ValidatorCallerSession) IsCandidateAdmin(_candidate common.Address, _admin common.Address) (bool, error) {
	return _Validator.Contract.IsCandidateAdmin(&_Validator.CallOpts, _candidate, _admin)
}

// IsOperatingBridge is a free data retrieval call binding the contract method 0x1f628801.
//
// Solidity: function isOperatingBridge(address _consensusAddr) view returns(bool)
func (_Validator *ValidatorCaller) IsOperatingBridge(opts *bind.CallOpts, _consensusAddr common.Address) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "isOperatingBridge", _consensusAddr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperatingBridge is a free data retrieval call binding the contract method 0x1f628801.
//
// Solidity: function isOperatingBridge(address _consensusAddr) view returns(bool)
func (_Validator *ValidatorSession) IsOperatingBridge(_consensusAddr common.Address) (bool, error) {
	return _Validator.Contract.IsOperatingBridge(&_Validator.CallOpts, _consensusAddr)
}

// IsOperatingBridge is a free data retrieval call binding the contract method 0x1f628801.
//
// Solidity: function isOperatingBridge(address _consensusAddr) view returns(bool)
func (_Validator *ValidatorCallerSession) IsOperatingBridge(_consensusAddr common.Address) (bool, error) {
	return _Validator.Contract.IsOperatingBridge(&_Validator.CallOpts, _consensusAddr)
}

// IsPeriodEnding is a free data retrieval call binding the contract method 0x217f35c2.
//
// Solidity: function isPeriodEnding() view returns(bool)
func (_Validator *ValidatorCaller) IsPeriodEnding(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "isPeriodEnding")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPeriodEnding is a free data retrieval call binding the contract method 0x217f35c2.
//
// Solidity: function isPeriodEnding() view returns(bool)
func (_Validator *ValidatorSession) IsPeriodEnding() (bool, error) {
	return _Validator.Contract.IsPeriodEnding(&_Validator.CallOpts)
}

// IsPeriodEnding is a free data retrieval call binding the contract method 0x217f35c2.
//
// Solidity: function isPeriodEnding() view returns(bool)
func (_Validator *ValidatorCallerSession) IsPeriodEnding() (bool, error) {
	return _Validator.Contract.IsPeriodEnding(&_Validator.CallOpts)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _addr) view returns(bool)
func (_Validator *ValidatorCaller) IsValidator(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "isValidator", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _addr) view returns(bool)
func (_Validator *ValidatorSession) IsValidator(_addr common.Address) (bool, error) {
	return _Validator.Contract.IsValidator(&_Validator.CallOpts, _addr)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address _addr) view returns(bool)
func (_Validator *ValidatorCallerSession) IsValidator(_addr common.Address) (bool, error) {
	return _Validator.Contract.IsValidator(&_Validator.CallOpts, _addr)
}

// IsValidatorCandidate is a free data retrieval call binding the contract method 0xa0c3f2d2.
//
// Solidity: function isValidatorCandidate(address _addr) view returns(bool)
func (_Validator *ValidatorCaller) IsValidatorCandidate(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "isValidatorCandidate", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorCandidate is a free data retrieval call binding the contract method 0xa0c3f2d2.
//
// Solidity: function isValidatorCandidate(address _addr) view returns(bool)
func (_Validator *ValidatorSession) IsValidatorCandidate(_addr common.Address) (bool, error) {
	return _Validator.Contract.IsValidatorCandidate(&_Validator.CallOpts, _addr)
}

// IsValidatorCandidate is a free data retrieval call binding the contract method 0xa0c3f2d2.
//
// Solidity: function isValidatorCandidate(address _addr) view returns(bool)
func (_Validator *ValidatorCallerSession) IsValidatorCandidate(_addr common.Address) (bool, error) {
	return _Validator.Contract.IsValidatorCandidate(&_Validator.CallOpts, _addr)
}

// MaintenanceContract is a free data retrieval call binding the contract method 0xd2cb215e.
//
// Solidity: function maintenanceContract() view returns(address)
func (_Validator *ValidatorCaller) MaintenanceContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "maintenanceContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MaintenanceContract is a free data retrieval call binding the contract method 0xd2cb215e.
//
// Solidity: function maintenanceContract() view returns(address)
func (_Validator *ValidatorSession) MaintenanceContract() (common.Address, error) {
	return _Validator.Contract.MaintenanceContract(&_Validator.CallOpts)
}

// MaintenanceContract is a free data retrieval call binding the contract method 0xd2cb215e.
//
// Solidity: function maintenanceContract() view returns(address)
func (_Validator *ValidatorCallerSession) MaintenanceContract() (common.Address, error) {
	return _Validator.Contract.MaintenanceContract(&_Validator.CallOpts)
}

// MaxPrioritizedValidatorNumber is a free data retrieval call binding the contract method 0xeeb629a8.
//
// Solidity: function maxPrioritizedValidatorNumber() view returns(uint256 _maximumPrioritizedValidatorNumber)
func (_Validator *ValidatorCaller) MaxPrioritizedValidatorNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "maxPrioritizedValidatorNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxPrioritizedValidatorNumber is a free data retrieval call binding the contract method 0xeeb629a8.
//
// Solidity: function maxPrioritizedValidatorNumber() view returns(uint256 _maximumPrioritizedValidatorNumber)
func (_Validator *ValidatorSession) MaxPrioritizedValidatorNumber() (*big.Int, error) {
	return _Validator.Contract.MaxPrioritizedValidatorNumber(&_Validator.CallOpts)
}

// MaxPrioritizedValidatorNumber is a free data retrieval call binding the contract method 0xeeb629a8.
//
// Solidity: function maxPrioritizedValidatorNumber() view returns(uint256 _maximumPrioritizedValidatorNumber)
func (_Validator *ValidatorCallerSession) MaxPrioritizedValidatorNumber() (*big.Int, error) {
	return _Validator.Contract.MaxPrioritizedValidatorNumber(&_Validator.CallOpts)
}

// MaxValidatorCandidate is a free data retrieval call binding the contract method 0x605239a1.
//
// Solidity: function maxValidatorCandidate() view returns(uint256)
func (_Validator *ValidatorCaller) MaxValidatorCandidate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "maxValidatorCandidate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxValidatorCandidate is a free data retrieval call binding the contract method 0x605239a1.
//
// Solidity: function maxValidatorCandidate() view returns(uint256)
func (_Validator *ValidatorSession) MaxValidatorCandidate() (*big.Int, error) {
	return _Validator.Contract.MaxValidatorCandidate(&_Validator.CallOpts)
}

// MaxValidatorCandidate is a free data retrieval call binding the contract method 0x605239a1.
//
// Solidity: function maxValidatorCandidate() view returns(uint256)
func (_Validator *ValidatorCallerSession) MaxValidatorCandidate() (*big.Int, error) {
	return _Validator.Contract.MaxValidatorCandidate(&_Validator.CallOpts)
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() view returns(uint256 _maximumValidatorNumber)
func (_Validator *ValidatorCaller) MaxValidatorNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "maxValidatorNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() view returns(uint256 _maximumValidatorNumber)
func (_Validator *ValidatorSession) MaxValidatorNumber() (*big.Int, error) {
	return _Validator.Contract.MaxValidatorNumber(&_Validator.CallOpts)
}

// MaxValidatorNumber is a free data retrieval call binding the contract method 0xd09f1ab4.
//
// Solidity: function maxValidatorNumber() view returns(uint256 _maximumValidatorNumber)
func (_Validator *ValidatorCallerSession) MaxValidatorNumber() (*big.Int, error) {
	return _Validator.Contract.MaxValidatorNumber(&_Validator.CallOpts)
}

// MinEffectiveDaysOnwards is a free data retrieval call binding the contract method 0xcba44de9.
//
// Solidity: function minEffectiveDaysOnwards() view returns(uint256)
func (_Validator *ValidatorCaller) MinEffectiveDaysOnwards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "minEffectiveDaysOnwards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinEffectiveDaysOnwards is a free data retrieval call binding the contract method 0xcba44de9.
//
// Solidity: function minEffectiveDaysOnwards() view returns(uint256)
func (_Validator *ValidatorSession) MinEffectiveDaysOnwards() (*big.Int, error) {
	return _Validator.Contract.MinEffectiveDaysOnwards(&_Validator.CallOpts)
}

// MinEffectiveDaysOnwards is a free data retrieval call binding the contract method 0xcba44de9.
//
// Solidity: function minEffectiveDaysOnwards() view returns(uint256)
func (_Validator *ValidatorCallerSession) MinEffectiveDaysOnwards() (*big.Int, error) {
	return _Validator.Contract.MinEffectiveDaysOnwards(&_Validator.CallOpts)
}

// NumberOfBlocksInEpoch is a free data retrieval call binding the contract method 0x6aa1c2ef.
//
// Solidity: function numberOfBlocksInEpoch() view returns(uint256 _numberOfBlocks)
func (_Validator *ValidatorCaller) NumberOfBlocksInEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "numberOfBlocksInEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfBlocksInEpoch is a free data retrieval call binding the contract method 0x6aa1c2ef.
//
// Solidity: function numberOfBlocksInEpoch() view returns(uint256 _numberOfBlocks)
func (_Validator *ValidatorSession) NumberOfBlocksInEpoch() (*big.Int, error) {
	return _Validator.Contract.NumberOfBlocksInEpoch(&_Validator.CallOpts)
}

// NumberOfBlocksInEpoch is a free data retrieval call binding the contract method 0x6aa1c2ef.
//
// Solidity: function numberOfBlocksInEpoch() view returns(uint256 _numberOfBlocks)
func (_Validator *ValidatorCallerSession) NumberOfBlocksInEpoch() (*big.Int, error) {
	return _Validator.Contract.NumberOfBlocksInEpoch(&_Validator.CallOpts)
}

// PrecompilePickValidatorSetAddress is a free data retrieval call binding the contract method 0x3b3159b6.
//
// Solidity: function precompilePickValidatorSetAddress() view returns(address)
func (_Validator *ValidatorCaller) PrecompilePickValidatorSetAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "precompilePickValidatorSetAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrecompilePickValidatorSetAddress is a free data retrieval call binding the contract method 0x3b3159b6.
//
// Solidity: function precompilePickValidatorSetAddress() view returns(address)
func (_Validator *ValidatorSession) PrecompilePickValidatorSetAddress() (common.Address, error) {
	return _Validator.Contract.PrecompilePickValidatorSetAddress(&_Validator.CallOpts)
}

// PrecompilePickValidatorSetAddress is a free data retrieval call binding the contract method 0x3b3159b6.
//
// Solidity: function precompilePickValidatorSetAddress() view returns(address)
func (_Validator *ValidatorCallerSession) PrecompilePickValidatorSetAddress() (common.Address, error) {
	return _Validator.Contract.PrecompilePickValidatorSetAddress(&_Validator.CallOpts)
}

// PrecompileSortValidatorsAddress is a free data retrieval call binding the contract method 0x8d559c38.
//
// Solidity: function precompileSortValidatorsAddress() view returns(address)
func (_Validator *ValidatorCaller) PrecompileSortValidatorsAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "precompileSortValidatorsAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrecompileSortValidatorsAddress is a free data retrieval call binding the contract method 0x8d559c38.
//
// Solidity: function precompileSortValidatorsAddress() view returns(address)
func (_Validator *ValidatorSession) PrecompileSortValidatorsAddress() (common.Address, error) {
	return _Validator.Contract.PrecompileSortValidatorsAddress(&_Validator.CallOpts)
}

// PrecompileSortValidatorsAddress is a free data retrieval call binding the contract method 0x8d559c38.
//
// Solidity: function precompileSortValidatorsAddress() view returns(address)
func (_Validator *ValidatorCallerSession) PrecompileSortValidatorsAddress() (common.Address, error) {
	return _Validator.Contract.PrecompileSortValidatorsAddress(&_Validator.CallOpts)
}

// RoninTrustedOrganizationContract is a free data retrieval call binding the contract method 0x5511cde1.
//
// Solidity: function roninTrustedOrganizationContract() view returns(address)
func (_Validator *ValidatorCaller) RoninTrustedOrganizationContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "roninTrustedOrganizationContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RoninTrustedOrganizationContract is a free data retrieval call binding the contract method 0x5511cde1.
//
// Solidity: function roninTrustedOrganizationContract() view returns(address)
func (_Validator *ValidatorSession) RoninTrustedOrganizationContract() (common.Address, error) {
	return _Validator.Contract.RoninTrustedOrganizationContract(&_Validator.CallOpts)
}

// RoninTrustedOrganizationContract is a free data retrieval call binding the contract method 0x5511cde1.
//
// Solidity: function roninTrustedOrganizationContract() view returns(address)
func (_Validator *ValidatorCallerSession) RoninTrustedOrganizationContract() (common.Address, error) {
	return _Validator.Contract.RoninTrustedOrganizationContract(&_Validator.CallOpts)
}

// SlashIndicatorContract is a free data retrieval call binding the contract method 0x5a08482d.
//
// Solidity: function slashIndicatorContract() view returns(address)
func (_Validator *ValidatorCaller) SlashIndicatorContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "slashIndicatorContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SlashIndicatorContract is a free data retrieval call binding the contract method 0x5a08482d.
//
// Solidity: function slashIndicatorContract() view returns(address)
func (_Validator *ValidatorSession) SlashIndicatorContract() (common.Address, error) {
	return _Validator.Contract.SlashIndicatorContract(&_Validator.CallOpts)
}

// SlashIndicatorContract is a free data retrieval call binding the contract method 0x5a08482d.
//
// Solidity: function slashIndicatorContract() view returns(address)
func (_Validator *ValidatorCallerSession) SlashIndicatorContract() (common.Address, error) {
	return _Validator.Contract.SlashIndicatorContract(&_Validator.CallOpts)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_Validator *ValidatorCaller) StakingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "stakingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_Validator *ValidatorSession) StakingContract() (common.Address, error) {
	return _Validator.Contract.StakingContract(&_Validator.CallOpts)
}

// StakingContract is a free data retrieval call binding the contract method 0xee99205c.
//
// Solidity: function stakingContract() view returns(address)
func (_Validator *ValidatorCallerSession) StakingContract() (common.Address, error) {
	return _Validator.Contract.StakingContract(&_Validator.CallOpts)
}

// StakingVestingContract is a free data retrieval call binding the contract method 0x3529214b.
//
// Solidity: function stakingVestingContract() view returns(address)
func (_Validator *ValidatorCaller) StakingVestingContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "stakingVestingContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingVestingContract is a free data retrieval call binding the contract method 0x3529214b.
//
// Solidity: function stakingVestingContract() view returns(address)
func (_Validator *ValidatorSession) StakingVestingContract() (common.Address, error) {
	return _Validator.Contract.StakingVestingContract(&_Validator.CallOpts)
}

// StakingVestingContract is a free data retrieval call binding the contract method 0x3529214b.
//
// Solidity: function stakingVestingContract() view returns(address)
func (_Validator *ValidatorCallerSession) StakingVestingContract() (common.Address, error) {
	return _Validator.Contract.StakingVestingContract(&_Validator.CallOpts)
}

// TotalBlockProducers is a free data retrieval call binding the contract method 0x9e94b9ec.
//
// Solidity: function totalBlockProducers() view returns(uint256 _total)
func (_Validator *ValidatorCaller) TotalBlockProducers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "totalBlockProducers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBlockProducers is a free data retrieval call binding the contract method 0x9e94b9ec.
//
// Solidity: function totalBlockProducers() view returns(uint256 _total)
func (_Validator *ValidatorSession) TotalBlockProducers() (*big.Int, error) {
	return _Validator.Contract.TotalBlockProducers(&_Validator.CallOpts)
}

// TotalBlockProducers is a free data retrieval call binding the contract method 0x9e94b9ec.
//
// Solidity: function totalBlockProducers() view returns(uint256 _total)
func (_Validator *ValidatorCallerSession) TotalBlockProducers() (*big.Int, error) {
	return _Validator.Contract.TotalBlockProducers(&_Validator.CallOpts)
}

// TotalBridgeOperators is a free data retrieval call binding the contract method 0x562d5304.
//
// Solidity: function totalBridgeOperators() view returns(uint256 _total)
func (_Validator *ValidatorCaller) TotalBridgeOperators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "totalBridgeOperators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBridgeOperators is a free data retrieval call binding the contract method 0x562d5304.
//
// Solidity: function totalBridgeOperators() view returns(uint256 _total)
func (_Validator *ValidatorSession) TotalBridgeOperators() (*big.Int, error) {
	return _Validator.Contract.TotalBridgeOperators(&_Validator.CallOpts)
}

// TotalBridgeOperators is a free data retrieval call binding the contract method 0x562d5304.
//
// Solidity: function totalBridgeOperators() view returns(uint256 _total)
func (_Validator *ValidatorCallerSession) TotalBridgeOperators() (*big.Int, error) {
	return _Validator.Contract.TotalBridgeOperators(&_Validator.CallOpts)
}

// TotalDeprecatedReward is a free data retrieval call binding the contract method 0x4ee4d72b.
//
// Solidity: function totalDeprecatedReward() view returns(uint256)
func (_Validator *ValidatorCaller) TotalDeprecatedReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "totalDeprecatedReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDeprecatedReward is a free data retrieval call binding the contract method 0x4ee4d72b.
//
// Solidity: function totalDeprecatedReward() view returns(uint256)
func (_Validator *ValidatorSession) TotalDeprecatedReward() (*big.Int, error) {
	return _Validator.Contract.TotalDeprecatedReward(&_Validator.CallOpts)
}

// TotalDeprecatedReward is a free data retrieval call binding the contract method 0x4ee4d72b.
//
// Solidity: function totalDeprecatedReward() view returns(uint256)
func (_Validator *ValidatorCallerSession) TotalDeprecatedReward() (*big.Int, error) {
	return _Validator.Contract.TotalDeprecatedReward(&_Validator.CallOpts)
}

// TryGetPeriodOfEpoch is a free data retrieval call binding the contract method 0x468c96ae.
//
// Solidity: function tryGetPeriodOfEpoch(uint256 _epoch) view returns(bool _filled, uint256 _periodNumber)
func (_Validator *ValidatorCaller) TryGetPeriodOfEpoch(opts *bind.CallOpts, _epoch *big.Int) (struct {
	Filled       bool
	PeriodNumber *big.Int
}, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "tryGetPeriodOfEpoch", _epoch)

	outstruct := new(struct {
		Filled       bool
		PeriodNumber *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Filled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PeriodNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TryGetPeriodOfEpoch is a free data retrieval call binding the contract method 0x468c96ae.
//
// Solidity: function tryGetPeriodOfEpoch(uint256 _epoch) view returns(bool _filled, uint256 _periodNumber)
func (_Validator *ValidatorSession) TryGetPeriodOfEpoch(_epoch *big.Int) (struct {
	Filled       bool
	PeriodNumber *big.Int
}, error) {
	return _Validator.Contract.TryGetPeriodOfEpoch(&_Validator.CallOpts, _epoch)
}

// TryGetPeriodOfEpoch is a free data retrieval call binding the contract method 0x468c96ae.
//
// Solidity: function tryGetPeriodOfEpoch(uint256 _epoch) view returns(bool _filled, uint256 _periodNumber)
func (_Validator *ValidatorCallerSession) TryGetPeriodOfEpoch(_epoch *big.Int) (struct {
	Filled       bool
	PeriodNumber *big.Int
}, error) {
	return _Validator.Contract.TryGetPeriodOfEpoch(&_Validator.CallOpts, _epoch)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_Validator *ValidatorCaller) ValidatorCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validator.contract.Call(opts, &out, "validatorCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_Validator *ValidatorSession) ValidatorCount() (*big.Int, error) {
	return _Validator.Contract.ValidatorCount(&_Validator.CallOpts)
}

// ValidatorCount is a free data retrieval call binding the contract method 0x0f43a677.
//
// Solidity: function validatorCount() view returns(uint256)
func (_Validator *ValidatorCallerSession) ValidatorCount() (*big.Int, error) {
	return _Validator.Contract.ValidatorCount(&_Validator.CallOpts)
}

// ExecApplyValidatorCandidate is a paid mutator transaction binding the contract method 0x1104e528.
//
// Solidity: function execApplyValidatorCandidate(address _candidateAdmin, address _consensusAddr, address _treasuryAddr, address _bridgeOperatorAddr, uint256 _commissionRate) returns()
func (_Validator *ValidatorTransactor) ExecApplyValidatorCandidate(opts *bind.TransactOpts, _candidateAdmin common.Address, _consensusAddr common.Address, _treasuryAddr common.Address, _bridgeOperatorAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "execApplyValidatorCandidate", _candidateAdmin, _consensusAddr, _treasuryAddr, _bridgeOperatorAddr, _commissionRate)
}

// ExecApplyValidatorCandidate is a paid mutator transaction binding the contract method 0x1104e528.
//
// Solidity: function execApplyValidatorCandidate(address _candidateAdmin, address _consensusAddr, address _treasuryAddr, address _bridgeOperatorAddr, uint256 _commissionRate) returns()
func (_Validator *ValidatorSession) ExecApplyValidatorCandidate(_candidateAdmin common.Address, _consensusAddr common.Address, _treasuryAddr common.Address, _bridgeOperatorAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecApplyValidatorCandidate(&_Validator.TransactOpts, _candidateAdmin, _consensusAddr, _treasuryAddr, _bridgeOperatorAddr, _commissionRate)
}

// ExecApplyValidatorCandidate is a paid mutator transaction binding the contract method 0x1104e528.
//
// Solidity: function execApplyValidatorCandidate(address _candidateAdmin, address _consensusAddr, address _treasuryAddr, address _bridgeOperatorAddr, uint256 _commissionRate) returns()
func (_Validator *ValidatorTransactorSession) ExecApplyValidatorCandidate(_candidateAdmin common.Address, _consensusAddr common.Address, _treasuryAddr common.Address, _bridgeOperatorAddr common.Address, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecApplyValidatorCandidate(&_Validator.TransactOpts, _candidateAdmin, _consensusAddr, _treasuryAddr, _bridgeOperatorAddr, _commissionRate)
}

// ExecBailOut is a paid mutator transaction binding the contract method 0x15b5ebde.
//
// Solidity: function execBailOut(address _validatorAddr, uint256 _period) returns()
func (_Validator *ValidatorTransactor) ExecBailOut(opts *bind.TransactOpts, _validatorAddr common.Address, _period *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "execBailOut", _validatorAddr, _period)
}

// ExecBailOut is a paid mutator transaction binding the contract method 0x15b5ebde.
//
// Solidity: function execBailOut(address _validatorAddr, uint256 _period) returns()
func (_Validator *ValidatorSession) ExecBailOut(_validatorAddr common.Address, _period *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecBailOut(&_Validator.TransactOpts, _validatorAddr, _period)
}

// ExecBailOut is a paid mutator transaction binding the contract method 0x15b5ebde.
//
// Solidity: function execBailOut(address _validatorAddr, uint256 _period) returns()
func (_Validator *ValidatorTransactorSession) ExecBailOut(_validatorAddr common.Address, _period *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecBailOut(&_Validator.TransactOpts, _validatorAddr, _period)
}

// ExecEmergencyExit is a paid mutator transaction binding the contract method 0xa7c2f119.
//
// Solidity: function execEmergencyExit(address _consensusAddr, uint256 _secLeftToRevoke) returns()
func (_Validator *ValidatorTransactor) ExecEmergencyExit(opts *bind.TransactOpts, _consensusAddr common.Address, _secLeftToRevoke *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "execEmergencyExit", _consensusAddr, _secLeftToRevoke)
}

// ExecEmergencyExit is a paid mutator transaction binding the contract method 0xa7c2f119.
//
// Solidity: function execEmergencyExit(address _consensusAddr, uint256 _secLeftToRevoke) returns()
func (_Validator *ValidatorSession) ExecEmergencyExit(_consensusAddr common.Address, _secLeftToRevoke *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecEmergencyExit(&_Validator.TransactOpts, _consensusAddr, _secLeftToRevoke)
}

// ExecEmergencyExit is a paid mutator transaction binding the contract method 0xa7c2f119.
//
// Solidity: function execEmergencyExit(address _consensusAddr, uint256 _secLeftToRevoke) returns()
func (_Validator *ValidatorTransactorSession) ExecEmergencyExit(_consensusAddr common.Address, _secLeftToRevoke *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecEmergencyExit(&_Validator.TransactOpts, _consensusAddr, _secLeftToRevoke)
}

// ExecReleaseLockedFundForEmergencyExitRequest is a paid mutator transaction binding the contract method 0xc3c8b5d6.
//
// Solidity: function execReleaseLockedFundForEmergencyExitRequest(address _consensusAddr, address _recipient) returns()
func (_Validator *ValidatorTransactor) ExecReleaseLockedFundForEmergencyExitRequest(opts *bind.TransactOpts, _consensusAddr common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "execReleaseLockedFundForEmergencyExitRequest", _consensusAddr, _recipient)
}

// ExecReleaseLockedFundForEmergencyExitRequest is a paid mutator transaction binding the contract method 0xc3c8b5d6.
//
// Solidity: function execReleaseLockedFundForEmergencyExitRequest(address _consensusAddr, address _recipient) returns()
func (_Validator *ValidatorSession) ExecReleaseLockedFundForEmergencyExitRequest(_consensusAddr common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _Validator.Contract.ExecReleaseLockedFundForEmergencyExitRequest(&_Validator.TransactOpts, _consensusAddr, _recipient)
}

// ExecReleaseLockedFundForEmergencyExitRequest is a paid mutator transaction binding the contract method 0xc3c8b5d6.
//
// Solidity: function execReleaseLockedFundForEmergencyExitRequest(address _consensusAddr, address _recipient) returns()
func (_Validator *ValidatorTransactorSession) ExecReleaseLockedFundForEmergencyExitRequest(_consensusAddr common.Address, _recipient common.Address) (*types.Transaction, error) {
	return _Validator.Contract.ExecReleaseLockedFundForEmergencyExitRequest(&_Validator.TransactOpts, _consensusAddr, _recipient)
}

// ExecRequestRenounceCandidate is a paid mutator transaction binding the contract method 0xdd716ad3.
//
// Solidity: function execRequestRenounceCandidate(address _consensusAddr, uint256 _secsLeft) returns()
func (_Validator *ValidatorTransactor) ExecRequestRenounceCandidate(opts *bind.TransactOpts, _consensusAddr common.Address, _secsLeft *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "execRequestRenounceCandidate", _consensusAddr, _secsLeft)
}

// ExecRequestRenounceCandidate is a paid mutator transaction binding the contract method 0xdd716ad3.
//
// Solidity: function execRequestRenounceCandidate(address _consensusAddr, uint256 _secsLeft) returns()
func (_Validator *ValidatorSession) ExecRequestRenounceCandidate(_consensusAddr common.Address, _secsLeft *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecRequestRenounceCandidate(&_Validator.TransactOpts, _consensusAddr, _secsLeft)
}

// ExecRequestRenounceCandidate is a paid mutator transaction binding the contract method 0xdd716ad3.
//
// Solidity: function execRequestRenounceCandidate(address _consensusAddr, uint256 _secsLeft) returns()
func (_Validator *ValidatorTransactorSession) ExecRequestRenounceCandidate(_consensusAddr common.Address, _secsLeft *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecRequestRenounceCandidate(&_Validator.TransactOpts, _consensusAddr, _secsLeft)
}

// ExecRequestUpdateCommissionRate is a paid mutator transaction binding the contract method 0xe5125a1d.
//
// Solidity: function execRequestUpdateCommissionRate(address _consensusAddr, uint256 _effectiveDaysOnwards, uint256 _commissionRate) returns()
func (_Validator *ValidatorTransactor) ExecRequestUpdateCommissionRate(opts *bind.TransactOpts, _consensusAddr common.Address, _effectiveDaysOnwards *big.Int, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "execRequestUpdateCommissionRate", _consensusAddr, _effectiveDaysOnwards, _commissionRate)
}

// ExecRequestUpdateCommissionRate is a paid mutator transaction binding the contract method 0xe5125a1d.
//
// Solidity: function execRequestUpdateCommissionRate(address _consensusAddr, uint256 _effectiveDaysOnwards, uint256 _commissionRate) returns()
func (_Validator *ValidatorSession) ExecRequestUpdateCommissionRate(_consensusAddr common.Address, _effectiveDaysOnwards *big.Int, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecRequestUpdateCommissionRate(&_Validator.TransactOpts, _consensusAddr, _effectiveDaysOnwards, _commissionRate)
}

// ExecRequestUpdateCommissionRate is a paid mutator transaction binding the contract method 0xe5125a1d.
//
// Solidity: function execRequestUpdateCommissionRate(address _consensusAddr, uint256 _effectiveDaysOnwards, uint256 _commissionRate) returns()
func (_Validator *ValidatorTransactorSession) ExecRequestUpdateCommissionRate(_consensusAddr common.Address, _effectiveDaysOnwards *big.Int, _commissionRate *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.ExecRequestUpdateCommissionRate(&_Validator.TransactOpts, _consensusAddr, _effectiveDaysOnwards, _commissionRate)
}

// ExecSlash is a paid mutator transaction binding the contract method 0x2f78204c.
//
// Solidity: function execSlash(address _validatorAddr, uint256 _newJailedUntil, uint256 _slashAmount, bool _cannotBailout) returns()
func (_Validator *ValidatorTransactor) ExecSlash(opts *bind.TransactOpts, _validatorAddr common.Address, _newJailedUntil *big.Int, _slashAmount *big.Int, _cannotBailout bool) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "execSlash", _validatorAddr, _newJailedUntil, _slashAmount, _cannotBailout)
}

// ExecSlash is a paid mutator transaction binding the contract method 0x2f78204c.
//
// Solidity: function execSlash(address _validatorAddr, uint256 _newJailedUntil, uint256 _slashAmount, bool _cannotBailout) returns()
func (_Validator *ValidatorSession) ExecSlash(_validatorAddr common.Address, _newJailedUntil *big.Int, _slashAmount *big.Int, _cannotBailout bool) (*types.Transaction, error) {
	return _Validator.Contract.ExecSlash(&_Validator.TransactOpts, _validatorAddr, _newJailedUntil, _slashAmount, _cannotBailout)
}

// ExecSlash is a paid mutator transaction binding the contract method 0x2f78204c.
//
// Solidity: function execSlash(address _validatorAddr, uint256 _newJailedUntil, uint256 _slashAmount, bool _cannotBailout) returns()
func (_Validator *ValidatorTransactorSession) ExecSlash(_validatorAddr common.Address, _newJailedUntil *big.Int, _slashAmount *big.Int, _cannotBailout bool) (*types.Transaction, error) {
	return _Validator.Contract.ExecSlash(&_Validator.TransactOpts, _validatorAddr, _newJailedUntil, _slashAmount, _cannotBailout)
}

// Initialize is a paid mutator transaction binding the contract method 0x367ec12b.
//
// Solidity: function initialize(address __slashIndicatorContract, address __stakingContract, address __stakingVestingContract, address __maintenanceContract, address __roninTrustedOrganizationContract, address __bridgeTrackingContract, uint256 __maxValidatorNumber, uint256 __maxValidatorCandidate, uint256 __maxPrioritizedValidatorNumber, uint256 __minEffectiveDaysOnwards, uint256 __numberOfBlocksInEpoch, uint256[2] __emergencyExitConfigs) returns()
func (_Validator *ValidatorTransactor) Initialize(opts *bind.TransactOpts, __slashIndicatorContract common.Address, __stakingContract common.Address, __stakingVestingContract common.Address, __maintenanceContract common.Address, __roninTrustedOrganizationContract common.Address, __bridgeTrackingContract common.Address, __maxValidatorNumber *big.Int, __maxValidatorCandidate *big.Int, __maxPrioritizedValidatorNumber *big.Int, __minEffectiveDaysOnwards *big.Int, __numberOfBlocksInEpoch *big.Int, __emergencyExitConfigs [2]*big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "initialize", __slashIndicatorContract, __stakingContract, __stakingVestingContract, __maintenanceContract, __roninTrustedOrganizationContract, __bridgeTrackingContract, __maxValidatorNumber, __maxValidatorCandidate, __maxPrioritizedValidatorNumber, __minEffectiveDaysOnwards, __numberOfBlocksInEpoch, __emergencyExitConfigs)
}

// Initialize is a paid mutator transaction binding the contract method 0x367ec12b.
//
// Solidity: function initialize(address __slashIndicatorContract, address __stakingContract, address __stakingVestingContract, address __maintenanceContract, address __roninTrustedOrganizationContract, address __bridgeTrackingContract, uint256 __maxValidatorNumber, uint256 __maxValidatorCandidate, uint256 __maxPrioritizedValidatorNumber, uint256 __minEffectiveDaysOnwards, uint256 __numberOfBlocksInEpoch, uint256[2] __emergencyExitConfigs) returns()
func (_Validator *ValidatorSession) Initialize(__slashIndicatorContract common.Address, __stakingContract common.Address, __stakingVestingContract common.Address, __maintenanceContract common.Address, __roninTrustedOrganizationContract common.Address, __bridgeTrackingContract common.Address, __maxValidatorNumber *big.Int, __maxValidatorCandidate *big.Int, __maxPrioritizedValidatorNumber *big.Int, __minEffectiveDaysOnwards *big.Int, __numberOfBlocksInEpoch *big.Int, __emergencyExitConfigs [2]*big.Int) (*types.Transaction, error) {
	return _Validator.Contract.Initialize(&_Validator.TransactOpts, __slashIndicatorContract, __stakingContract, __stakingVestingContract, __maintenanceContract, __roninTrustedOrganizationContract, __bridgeTrackingContract, __maxValidatorNumber, __maxValidatorCandidate, __maxPrioritizedValidatorNumber, __minEffectiveDaysOnwards, __numberOfBlocksInEpoch, __emergencyExitConfigs)
}

// Initialize is a paid mutator transaction binding the contract method 0x367ec12b.
//
// Solidity: function initialize(address __slashIndicatorContract, address __stakingContract, address __stakingVestingContract, address __maintenanceContract, address __roninTrustedOrganizationContract, address __bridgeTrackingContract, uint256 __maxValidatorNumber, uint256 __maxValidatorCandidate, uint256 __maxPrioritizedValidatorNumber, uint256 __minEffectiveDaysOnwards, uint256 __numberOfBlocksInEpoch, uint256[2] __emergencyExitConfigs) returns()
func (_Validator *ValidatorTransactorSession) Initialize(__slashIndicatorContract common.Address, __stakingContract common.Address, __stakingVestingContract common.Address, __maintenanceContract common.Address, __roninTrustedOrganizationContract common.Address, __bridgeTrackingContract common.Address, __maxValidatorNumber *big.Int, __maxValidatorCandidate *big.Int, __maxPrioritizedValidatorNumber *big.Int, __minEffectiveDaysOnwards *big.Int, __numberOfBlocksInEpoch *big.Int, __emergencyExitConfigs [2]*big.Int) (*types.Transaction, error) {
	return _Validator.Contract.Initialize(&_Validator.TransactOpts, __slashIndicatorContract, __stakingContract, __stakingVestingContract, __maintenanceContract, __roninTrustedOrganizationContract, __bridgeTrackingContract, __maxValidatorNumber, __maxValidatorCandidate, __maxPrioritizedValidatorNumber, __minEffectiveDaysOnwards, __numberOfBlocksInEpoch, __emergencyExitConfigs)
}

// SetBridgeTrackingContract is a paid mutator transaction binding the contract method 0x9c8d98da.
//
// Solidity: function setBridgeTrackingContract(address _addr) returns()
func (_Validator *ValidatorTransactor) SetBridgeTrackingContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setBridgeTrackingContract", _addr)
}

// SetBridgeTrackingContract is a paid mutator transaction binding the contract method 0x9c8d98da.
//
// Solidity: function setBridgeTrackingContract(address _addr) returns()
func (_Validator *ValidatorSession) SetBridgeTrackingContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetBridgeTrackingContract(&_Validator.TransactOpts, _addr)
}

// SetBridgeTrackingContract is a paid mutator transaction binding the contract method 0x9c8d98da.
//
// Solidity: function setBridgeTrackingContract(address _addr) returns()
func (_Validator *ValidatorTransactorSession) SetBridgeTrackingContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetBridgeTrackingContract(&_Validator.TransactOpts, _addr)
}

// SetEmergencyExitLockedAmount is a paid mutator transaction binding the contract method 0x6611f843.
//
// Solidity: function setEmergencyExitLockedAmount(uint256 _emergencyExitLockedAmount) returns()
func (_Validator *ValidatorTransactor) SetEmergencyExitLockedAmount(opts *bind.TransactOpts, _emergencyExitLockedAmount *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setEmergencyExitLockedAmount", _emergencyExitLockedAmount)
}

// SetEmergencyExitLockedAmount is a paid mutator transaction binding the contract method 0x6611f843.
//
// Solidity: function setEmergencyExitLockedAmount(uint256 _emergencyExitLockedAmount) returns()
func (_Validator *ValidatorSession) SetEmergencyExitLockedAmount(_emergencyExitLockedAmount *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetEmergencyExitLockedAmount(&_Validator.TransactOpts, _emergencyExitLockedAmount)
}

// SetEmergencyExitLockedAmount is a paid mutator transaction binding the contract method 0x6611f843.
//
// Solidity: function setEmergencyExitLockedAmount(uint256 _emergencyExitLockedAmount) returns()
func (_Validator *ValidatorTransactorSession) SetEmergencyExitLockedAmount(_emergencyExitLockedAmount *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetEmergencyExitLockedAmount(&_Validator.TransactOpts, _emergencyExitLockedAmount)
}

// SetEmergencyExpiryDuration is a paid mutator transaction binding the contract method 0x4d8df063.
//
// Solidity: function setEmergencyExpiryDuration(uint256 _emergencyExpiryDuration) returns()
func (_Validator *ValidatorTransactor) SetEmergencyExpiryDuration(opts *bind.TransactOpts, _emergencyExpiryDuration *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setEmergencyExpiryDuration", _emergencyExpiryDuration)
}

// SetEmergencyExpiryDuration is a paid mutator transaction binding the contract method 0x4d8df063.
//
// Solidity: function setEmergencyExpiryDuration(uint256 _emergencyExpiryDuration) returns()
func (_Validator *ValidatorSession) SetEmergencyExpiryDuration(_emergencyExpiryDuration *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetEmergencyExpiryDuration(&_Validator.TransactOpts, _emergencyExpiryDuration)
}

// SetEmergencyExpiryDuration is a paid mutator transaction binding the contract method 0x4d8df063.
//
// Solidity: function setEmergencyExpiryDuration(uint256 _emergencyExpiryDuration) returns()
func (_Validator *ValidatorTransactorSession) SetEmergencyExpiryDuration(_emergencyExpiryDuration *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetEmergencyExpiryDuration(&_Validator.TransactOpts, _emergencyExpiryDuration)
}

// SetMaintenanceContract is a paid mutator transaction binding the contract method 0x46fe9311.
//
// Solidity: function setMaintenanceContract(address _addr) returns()
func (_Validator *ValidatorTransactor) SetMaintenanceContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setMaintenanceContract", _addr)
}

// SetMaintenanceContract is a paid mutator transaction binding the contract method 0x46fe9311.
//
// Solidity: function setMaintenanceContract(address _addr) returns()
func (_Validator *ValidatorSession) SetMaintenanceContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetMaintenanceContract(&_Validator.TransactOpts, _addr)
}

// SetMaintenanceContract is a paid mutator transaction binding the contract method 0x46fe9311.
//
// Solidity: function setMaintenanceContract(address _addr) returns()
func (_Validator *ValidatorTransactorSession) SetMaintenanceContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetMaintenanceContract(&_Validator.TransactOpts, _addr)
}

// SetMaxPrioritizedValidatorNumber is a paid mutator transaction binding the contract method 0xc94aaa02.
//
// Solidity: function setMaxPrioritizedValidatorNumber(uint256 _number) returns()
func (_Validator *ValidatorTransactor) SetMaxPrioritizedValidatorNumber(opts *bind.TransactOpts, _number *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setMaxPrioritizedValidatorNumber", _number)
}

// SetMaxPrioritizedValidatorNumber is a paid mutator transaction binding the contract method 0xc94aaa02.
//
// Solidity: function setMaxPrioritizedValidatorNumber(uint256 _number) returns()
func (_Validator *ValidatorSession) SetMaxPrioritizedValidatorNumber(_number *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetMaxPrioritizedValidatorNumber(&_Validator.TransactOpts, _number)
}

// SetMaxPrioritizedValidatorNumber is a paid mutator transaction binding the contract method 0xc94aaa02.
//
// Solidity: function setMaxPrioritizedValidatorNumber(uint256 _number) returns()
func (_Validator *ValidatorTransactorSession) SetMaxPrioritizedValidatorNumber(_number *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetMaxPrioritizedValidatorNumber(&_Validator.TransactOpts, _number)
}

// SetMaxValidatorCandidate is a paid mutator transaction binding the contract method 0x4f2a693f.
//
// Solidity: function setMaxValidatorCandidate(uint256 _number) returns()
func (_Validator *ValidatorTransactor) SetMaxValidatorCandidate(opts *bind.TransactOpts, _number *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setMaxValidatorCandidate", _number)
}

// SetMaxValidatorCandidate is a paid mutator transaction binding the contract method 0x4f2a693f.
//
// Solidity: function setMaxValidatorCandidate(uint256 _number) returns()
func (_Validator *ValidatorSession) SetMaxValidatorCandidate(_number *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetMaxValidatorCandidate(&_Validator.TransactOpts, _number)
}

// SetMaxValidatorCandidate is a paid mutator transaction binding the contract method 0x4f2a693f.
//
// Solidity: function setMaxValidatorCandidate(uint256 _number) returns()
func (_Validator *ValidatorTransactorSession) SetMaxValidatorCandidate(_number *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetMaxValidatorCandidate(&_Validator.TransactOpts, _number)
}

// SetMaxValidatorNumber is a paid mutator transaction binding the contract method 0x823a7b9c.
//
// Solidity: function setMaxValidatorNumber(uint256 _max) returns()
func (_Validator *ValidatorTransactor) SetMaxValidatorNumber(opts *bind.TransactOpts, _max *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setMaxValidatorNumber", _max)
}

// SetMaxValidatorNumber is a paid mutator transaction binding the contract method 0x823a7b9c.
//
// Solidity: function setMaxValidatorNumber(uint256 _max) returns()
func (_Validator *ValidatorSession) SetMaxValidatorNumber(_max *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetMaxValidatorNumber(&_Validator.TransactOpts, _max)
}

// SetMaxValidatorNumber is a paid mutator transaction binding the contract method 0x823a7b9c.
//
// Solidity: function setMaxValidatorNumber(uint256 _max) returns()
func (_Validator *ValidatorTransactorSession) SetMaxValidatorNumber(_max *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetMaxValidatorNumber(&_Validator.TransactOpts, _max)
}

// SetMinEffectiveDaysOnwards is a paid mutator transaction binding the contract method 0x1196ab66.
//
// Solidity: function setMinEffectiveDaysOnwards(uint256 _numOfDays) returns()
func (_Validator *ValidatorTransactor) SetMinEffectiveDaysOnwards(opts *bind.TransactOpts, _numOfDays *big.Int) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setMinEffectiveDaysOnwards", _numOfDays)
}

// SetMinEffectiveDaysOnwards is a paid mutator transaction binding the contract method 0x1196ab66.
//
// Solidity: function setMinEffectiveDaysOnwards(uint256 _numOfDays) returns()
func (_Validator *ValidatorSession) SetMinEffectiveDaysOnwards(_numOfDays *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetMinEffectiveDaysOnwards(&_Validator.TransactOpts, _numOfDays)
}

// SetMinEffectiveDaysOnwards is a paid mutator transaction binding the contract method 0x1196ab66.
//
// Solidity: function setMinEffectiveDaysOnwards(uint256 _numOfDays) returns()
func (_Validator *ValidatorTransactorSession) SetMinEffectiveDaysOnwards(_numOfDays *big.Int) (*types.Transaction, error) {
	return _Validator.Contract.SetMinEffectiveDaysOnwards(&_Validator.TransactOpts, _numOfDays)
}

// SetRoninTrustedOrganizationContract is a paid mutator transaction binding the contract method 0xb5e337de.
//
// Solidity: function setRoninTrustedOrganizationContract(address _addr) returns()
func (_Validator *ValidatorTransactor) SetRoninTrustedOrganizationContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setRoninTrustedOrganizationContract", _addr)
}

// SetRoninTrustedOrganizationContract is a paid mutator transaction binding the contract method 0xb5e337de.
//
// Solidity: function setRoninTrustedOrganizationContract(address _addr) returns()
func (_Validator *ValidatorSession) SetRoninTrustedOrganizationContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetRoninTrustedOrganizationContract(&_Validator.TransactOpts, _addr)
}

// SetRoninTrustedOrganizationContract is a paid mutator transaction binding the contract method 0xb5e337de.
//
// Solidity: function setRoninTrustedOrganizationContract(address _addr) returns()
func (_Validator *ValidatorTransactorSession) SetRoninTrustedOrganizationContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetRoninTrustedOrganizationContract(&_Validator.TransactOpts, _addr)
}

// SetSlashIndicatorContract is a paid mutator transaction binding the contract method 0x2bcf3d15.
//
// Solidity: function setSlashIndicatorContract(address _addr) returns()
func (_Validator *ValidatorTransactor) SetSlashIndicatorContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setSlashIndicatorContract", _addr)
}

// SetSlashIndicatorContract is a paid mutator transaction binding the contract method 0x2bcf3d15.
//
// Solidity: function setSlashIndicatorContract(address _addr) returns()
func (_Validator *ValidatorSession) SetSlashIndicatorContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetSlashIndicatorContract(&_Validator.TransactOpts, _addr)
}

// SetSlashIndicatorContract is a paid mutator transaction binding the contract method 0x2bcf3d15.
//
// Solidity: function setSlashIndicatorContract(address _addr) returns()
func (_Validator *ValidatorTransactorSession) SetSlashIndicatorContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetSlashIndicatorContract(&_Validator.TransactOpts, _addr)
}

// SetStakingContract is a paid mutator transaction binding the contract method 0x9dd373b9.
//
// Solidity: function setStakingContract(address _addr) returns()
func (_Validator *ValidatorTransactor) SetStakingContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setStakingContract", _addr)
}

// SetStakingContract is a paid mutator transaction binding the contract method 0x9dd373b9.
//
// Solidity: function setStakingContract(address _addr) returns()
func (_Validator *ValidatorSession) SetStakingContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetStakingContract(&_Validator.TransactOpts, _addr)
}

// SetStakingContract is a paid mutator transaction binding the contract method 0x9dd373b9.
//
// Solidity: function setStakingContract(address _addr) returns()
func (_Validator *ValidatorTransactorSession) SetStakingContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetStakingContract(&_Validator.TransactOpts, _addr)
}

// SetStakingVestingContract is a paid mutator transaction binding the contract method 0xad295783.
//
// Solidity: function setStakingVestingContract(address _addr) returns()
func (_Validator *ValidatorTransactor) SetStakingVestingContract(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "setStakingVestingContract", _addr)
}

// SetStakingVestingContract is a paid mutator transaction binding the contract method 0xad295783.
//
// Solidity: function setStakingVestingContract(address _addr) returns()
func (_Validator *ValidatorSession) SetStakingVestingContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetStakingVestingContract(&_Validator.TransactOpts, _addr)
}

// SetStakingVestingContract is a paid mutator transaction binding the contract method 0xad295783.
//
// Solidity: function setStakingVestingContract(address _addr) returns()
func (_Validator *ValidatorTransactorSession) SetStakingVestingContract(_addr common.Address) (*types.Transaction, error) {
	return _Validator.Contract.SetStakingVestingContract(&_Validator.TransactOpts, _addr)
}

// SubmitBlockReward is a paid mutator transaction binding the contract method 0x52091f17.
//
// Solidity: function submitBlockReward() payable returns()
func (_Validator *ValidatorTransactor) SubmitBlockReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "submitBlockReward")
}

// SubmitBlockReward is a paid mutator transaction binding the contract method 0x52091f17.
//
// Solidity: function submitBlockReward() payable returns()
func (_Validator *ValidatorSession) SubmitBlockReward() (*types.Transaction, error) {
	return _Validator.Contract.SubmitBlockReward(&_Validator.TransactOpts)
}

// SubmitBlockReward is a paid mutator transaction binding the contract method 0x52091f17.
//
// Solidity: function submitBlockReward() payable returns()
func (_Validator *ValidatorTransactorSession) SubmitBlockReward() (*types.Transaction, error) {
	return _Validator.Contract.SubmitBlockReward(&_Validator.TransactOpts)
}

// WrapUpEpoch is a paid mutator transaction binding the contract method 0x72e46810.
//
// Solidity: function wrapUpEpoch() payable returns()
func (_Validator *ValidatorTransactor) WrapUpEpoch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.contract.Transact(opts, "wrapUpEpoch")
}

// WrapUpEpoch is a paid mutator transaction binding the contract method 0x72e46810.
//
// Solidity: function wrapUpEpoch() payable returns()
func (_Validator *ValidatorSession) WrapUpEpoch() (*types.Transaction, error) {
	return _Validator.Contract.WrapUpEpoch(&_Validator.TransactOpts)
}

// WrapUpEpoch is a paid mutator transaction binding the contract method 0x72e46810.
//
// Solidity: function wrapUpEpoch() payable returns()
func (_Validator *ValidatorTransactorSession) WrapUpEpoch() (*types.Transaction, error) {
	return _Validator.Contract.WrapUpEpoch(&_Validator.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Validator *ValidatorTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Validator.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Validator *ValidatorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Validator.Contract.Fallback(&_Validator.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Validator *ValidatorTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Validator.Contract.Fallback(&_Validator.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Validator *ValidatorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validator.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Validator *ValidatorSession) Receive() (*types.Transaction, error) {
	return _Validator.Contract.Receive(&_Validator.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Validator *ValidatorTransactorSession) Receive() (*types.Transaction, error) {
	return _Validator.Contract.Receive(&_Validator.TransactOpts)
}

// ValidatorBlockProducerSetUpdatedIterator is returned from FilterBlockProducerSetUpdated and is used to iterate over the raw logs and unpacked data for BlockProducerSetUpdated events raised by the Validator contract.
type ValidatorBlockProducerSetUpdatedIterator struct {
	Event *ValidatorBlockProducerSetUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorBlockProducerSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorBlockProducerSetUpdated)
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
		it.Event = new(ValidatorBlockProducerSetUpdated)
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
func (it *ValidatorBlockProducerSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorBlockProducerSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorBlockProducerSetUpdated represents a BlockProducerSetUpdated event raised by the Validator contract.
type ValidatorBlockProducerSetUpdated struct {
	Period         *big.Int
	Epoch          *big.Int
	ConsensusAddrs []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlockProducerSetUpdated is a free log retrieval operation binding the contract event 0x283b50d76057d5f828df85bc87724c6af604e9b55c363a07c9faa2a2cd688b82.
//
// Solidity: event BlockProducerSetUpdated(uint256 indexed period, uint256 indexed epoch, address[] consensusAddrs)
func (_Validator *ValidatorFilterer) FilterBlockProducerSetUpdated(opts *bind.FilterOpts, period []*big.Int, epoch []*big.Int) (*ValidatorBlockProducerSetUpdatedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "BlockProducerSetUpdated", periodRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorBlockProducerSetUpdatedIterator{contract: _Validator.contract, event: "BlockProducerSetUpdated", logs: logs, sub: sub}, nil
}

// WatchBlockProducerSetUpdated is a free log subscription operation binding the contract event 0x283b50d76057d5f828df85bc87724c6af604e9b55c363a07c9faa2a2cd688b82.
//
// Solidity: event BlockProducerSetUpdated(uint256 indexed period, uint256 indexed epoch, address[] consensusAddrs)
func (_Validator *ValidatorFilterer) WatchBlockProducerSetUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorBlockProducerSetUpdated, period []*big.Int, epoch []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "BlockProducerSetUpdated", periodRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorBlockProducerSetUpdated)
				if err := _Validator.contract.UnpackLog(event, "BlockProducerSetUpdated", log); err != nil {
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

// ParseBlockProducerSetUpdated is a log parse operation binding the contract event 0x283b50d76057d5f828df85bc87724c6af604e9b55c363a07c9faa2a2cd688b82.
//
// Solidity: event BlockProducerSetUpdated(uint256 indexed period, uint256 indexed epoch, address[] consensusAddrs)
func (_Validator *ValidatorFilterer) ParseBlockProducerSetUpdated(log types.Log) (*ValidatorBlockProducerSetUpdated, error) {
	event := new(ValidatorBlockProducerSetUpdated)
	if err := _Validator.contract.UnpackLog(event, "BlockProducerSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorBlockRewardDeprecatedIterator is returned from FilterBlockRewardDeprecated and is used to iterate over the raw logs and unpacked data for BlockRewardDeprecated events raised by the Validator contract.
type ValidatorBlockRewardDeprecatedIterator struct {
	Event *ValidatorBlockRewardDeprecated // Event containing the contract specifics and raw log

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
func (it *ValidatorBlockRewardDeprecatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorBlockRewardDeprecated)
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
		it.Event = new(ValidatorBlockRewardDeprecated)
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
func (it *ValidatorBlockRewardDeprecatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorBlockRewardDeprecatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorBlockRewardDeprecated represents a BlockRewardDeprecated event raised by the Validator contract.
type ValidatorBlockRewardDeprecated struct {
	CoinbaseAddr   common.Address
	RewardAmount   *big.Int
	DeprecatedType uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlockRewardDeprecated is a free log retrieval operation binding the contract event 0x4042bb9a70998f80a86d9963f0d2132e9b11c8ad94d207c6141c8e34b05ce53e.
//
// Solidity: event BlockRewardDeprecated(address indexed coinbaseAddr, uint256 rewardAmount, uint8 deprecatedType)
func (_Validator *ValidatorFilterer) FilterBlockRewardDeprecated(opts *bind.FilterOpts, coinbaseAddr []common.Address) (*ValidatorBlockRewardDeprecatedIterator, error) {

	var coinbaseAddrRule []interface{}
	for _, coinbaseAddrItem := range coinbaseAddr {
		coinbaseAddrRule = append(coinbaseAddrRule, coinbaseAddrItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "BlockRewardDeprecated", coinbaseAddrRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorBlockRewardDeprecatedIterator{contract: _Validator.contract, event: "BlockRewardDeprecated", logs: logs, sub: sub}, nil
}

// WatchBlockRewardDeprecated is a free log subscription operation binding the contract event 0x4042bb9a70998f80a86d9963f0d2132e9b11c8ad94d207c6141c8e34b05ce53e.
//
// Solidity: event BlockRewardDeprecated(address indexed coinbaseAddr, uint256 rewardAmount, uint8 deprecatedType)
func (_Validator *ValidatorFilterer) WatchBlockRewardDeprecated(opts *bind.WatchOpts, sink chan<- *ValidatorBlockRewardDeprecated, coinbaseAddr []common.Address) (event.Subscription, error) {

	var coinbaseAddrRule []interface{}
	for _, coinbaseAddrItem := range coinbaseAddr {
		coinbaseAddrRule = append(coinbaseAddrRule, coinbaseAddrItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "BlockRewardDeprecated", coinbaseAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorBlockRewardDeprecated)
				if err := _Validator.contract.UnpackLog(event, "BlockRewardDeprecated", log); err != nil {
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

// ParseBlockRewardDeprecated is a log parse operation binding the contract event 0x4042bb9a70998f80a86d9963f0d2132e9b11c8ad94d207c6141c8e34b05ce53e.
//
// Solidity: event BlockRewardDeprecated(address indexed coinbaseAddr, uint256 rewardAmount, uint8 deprecatedType)
func (_Validator *ValidatorFilterer) ParseBlockRewardDeprecated(log types.Log) (*ValidatorBlockRewardDeprecated, error) {
	event := new(ValidatorBlockRewardDeprecated)
	if err := _Validator.contract.UnpackLog(event, "BlockRewardDeprecated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorBlockRewardSubmittedIterator is returned from FilterBlockRewardSubmitted and is used to iterate over the raw logs and unpacked data for BlockRewardSubmitted events raised by the Validator contract.
type ValidatorBlockRewardSubmittedIterator struct {
	Event *ValidatorBlockRewardSubmitted // Event containing the contract specifics and raw log

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
func (it *ValidatorBlockRewardSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorBlockRewardSubmitted)
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
		it.Event = new(ValidatorBlockRewardSubmitted)
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
func (it *ValidatorBlockRewardSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorBlockRewardSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorBlockRewardSubmitted represents a BlockRewardSubmitted event raised by the Validator contract.
type ValidatorBlockRewardSubmitted struct {
	CoinbaseAddr    common.Address
	SubmittedAmount *big.Int
	BonusAmount     *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBlockRewardSubmitted is a free log retrieval operation binding the contract event 0x0ede5c3be8625943fa64003cd4b91230089411249f3059bac6500873543ca9b1.
//
// Solidity: event BlockRewardSubmitted(address indexed coinbaseAddr, uint256 submittedAmount, uint256 bonusAmount)
func (_Validator *ValidatorFilterer) FilterBlockRewardSubmitted(opts *bind.FilterOpts, coinbaseAddr []common.Address) (*ValidatorBlockRewardSubmittedIterator, error) {

	var coinbaseAddrRule []interface{}
	for _, coinbaseAddrItem := range coinbaseAddr {
		coinbaseAddrRule = append(coinbaseAddrRule, coinbaseAddrItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "BlockRewardSubmitted", coinbaseAddrRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorBlockRewardSubmittedIterator{contract: _Validator.contract, event: "BlockRewardSubmitted", logs: logs, sub: sub}, nil
}

// WatchBlockRewardSubmitted is a free log subscription operation binding the contract event 0x0ede5c3be8625943fa64003cd4b91230089411249f3059bac6500873543ca9b1.
//
// Solidity: event BlockRewardSubmitted(address indexed coinbaseAddr, uint256 submittedAmount, uint256 bonusAmount)
func (_Validator *ValidatorFilterer) WatchBlockRewardSubmitted(opts *bind.WatchOpts, sink chan<- *ValidatorBlockRewardSubmitted, coinbaseAddr []common.Address) (event.Subscription, error) {

	var coinbaseAddrRule []interface{}
	for _, coinbaseAddrItem := range coinbaseAddr {
		coinbaseAddrRule = append(coinbaseAddrRule, coinbaseAddrItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "BlockRewardSubmitted", coinbaseAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorBlockRewardSubmitted)
				if err := _Validator.contract.UnpackLog(event, "BlockRewardSubmitted", log); err != nil {
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

// ParseBlockRewardSubmitted is a log parse operation binding the contract event 0x0ede5c3be8625943fa64003cd4b91230089411249f3059bac6500873543ca9b1.
//
// Solidity: event BlockRewardSubmitted(address indexed coinbaseAddr, uint256 submittedAmount, uint256 bonusAmount)
func (_Validator *ValidatorFilterer) ParseBlockRewardSubmitted(log types.Log) (*ValidatorBlockRewardSubmitted, error) {
	event := new(ValidatorBlockRewardSubmitted)
	if err := _Validator.contract.UnpackLog(event, "BlockRewardSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorBridgeOperatorRewardDistributedIterator is returned from FilterBridgeOperatorRewardDistributed and is used to iterate over the raw logs and unpacked data for BridgeOperatorRewardDistributed events raised by the Validator contract.
type ValidatorBridgeOperatorRewardDistributedIterator struct {
	Event *ValidatorBridgeOperatorRewardDistributed // Event containing the contract specifics and raw log

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
func (it *ValidatorBridgeOperatorRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorBridgeOperatorRewardDistributed)
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
		it.Event = new(ValidatorBridgeOperatorRewardDistributed)
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
func (it *ValidatorBridgeOperatorRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorBridgeOperatorRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorBridgeOperatorRewardDistributed represents a BridgeOperatorRewardDistributed event raised by the Validator contract.
type ValidatorBridgeOperatorRewardDistributed struct {
	ConsensusAddr  common.Address
	BridgeOperator common.Address
	RecipientAddr  common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeOperatorRewardDistributed is a free log retrieval operation binding the contract event 0x72a57dc38837a1cba7881b7b1a5594d9e6b65cec6a985b54e2cee3e89369691c.
//
// Solidity: event BridgeOperatorRewardDistributed(address indexed consensusAddr, address indexed bridgeOperator, address indexed recipientAddr, uint256 amount)
func (_Validator *ValidatorFilterer) FilterBridgeOperatorRewardDistributed(opts *bind.FilterOpts, consensusAddr []common.Address, bridgeOperator []common.Address, recipientAddr []common.Address) (*ValidatorBridgeOperatorRewardDistributedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var bridgeOperatorRule []interface{}
	for _, bridgeOperatorItem := range bridgeOperator {
		bridgeOperatorRule = append(bridgeOperatorRule, bridgeOperatorItem)
	}
	var recipientAddrRule []interface{}
	for _, recipientAddrItem := range recipientAddr {
		recipientAddrRule = append(recipientAddrRule, recipientAddrItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "BridgeOperatorRewardDistributed", consensusAddrRule, bridgeOperatorRule, recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorBridgeOperatorRewardDistributedIterator{contract: _Validator.contract, event: "BridgeOperatorRewardDistributed", logs: logs, sub: sub}, nil
}

// WatchBridgeOperatorRewardDistributed is a free log subscription operation binding the contract event 0x72a57dc38837a1cba7881b7b1a5594d9e6b65cec6a985b54e2cee3e89369691c.
//
// Solidity: event BridgeOperatorRewardDistributed(address indexed consensusAddr, address indexed bridgeOperator, address indexed recipientAddr, uint256 amount)
func (_Validator *ValidatorFilterer) WatchBridgeOperatorRewardDistributed(opts *bind.WatchOpts, sink chan<- *ValidatorBridgeOperatorRewardDistributed, consensusAddr []common.Address, bridgeOperator []common.Address, recipientAddr []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var bridgeOperatorRule []interface{}
	for _, bridgeOperatorItem := range bridgeOperator {
		bridgeOperatorRule = append(bridgeOperatorRule, bridgeOperatorItem)
	}
	var recipientAddrRule []interface{}
	for _, recipientAddrItem := range recipientAddr {
		recipientAddrRule = append(recipientAddrRule, recipientAddrItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "BridgeOperatorRewardDistributed", consensusAddrRule, bridgeOperatorRule, recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorBridgeOperatorRewardDistributed)
				if err := _Validator.contract.UnpackLog(event, "BridgeOperatorRewardDistributed", log); err != nil {
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

// ParseBridgeOperatorRewardDistributed is a log parse operation binding the contract event 0x72a57dc38837a1cba7881b7b1a5594d9e6b65cec6a985b54e2cee3e89369691c.
//
// Solidity: event BridgeOperatorRewardDistributed(address indexed consensusAddr, address indexed bridgeOperator, address indexed recipientAddr, uint256 amount)
func (_Validator *ValidatorFilterer) ParseBridgeOperatorRewardDistributed(log types.Log) (*ValidatorBridgeOperatorRewardDistributed, error) {
	event := new(ValidatorBridgeOperatorRewardDistributed)
	if err := _Validator.contract.UnpackLog(event, "BridgeOperatorRewardDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorBridgeOperatorRewardDistributionFailedIterator is returned from FilterBridgeOperatorRewardDistributionFailed and is used to iterate over the raw logs and unpacked data for BridgeOperatorRewardDistributionFailed events raised by the Validator contract.
type ValidatorBridgeOperatorRewardDistributionFailedIterator struct {
	Event *ValidatorBridgeOperatorRewardDistributionFailed // Event containing the contract specifics and raw log

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
func (it *ValidatorBridgeOperatorRewardDistributionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorBridgeOperatorRewardDistributionFailed)
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
		it.Event = new(ValidatorBridgeOperatorRewardDistributionFailed)
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
func (it *ValidatorBridgeOperatorRewardDistributionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorBridgeOperatorRewardDistributionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorBridgeOperatorRewardDistributionFailed represents a BridgeOperatorRewardDistributionFailed event raised by the Validator contract.
type ValidatorBridgeOperatorRewardDistributionFailed struct {
	ConsensusAddr   common.Address
	BridgeOperator  common.Address
	Recipient       common.Address
	Amount          *big.Int
	ContractBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgeOperatorRewardDistributionFailed is a free log retrieval operation binding the contract event 0xd35d76d87d51ed89407fc7ceaaccf32cf72784b94530892ce33546540e141b72.
//
// Solidity: event BridgeOperatorRewardDistributionFailed(address indexed consensusAddr, address indexed bridgeOperator, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_Validator *ValidatorFilterer) FilterBridgeOperatorRewardDistributionFailed(opts *bind.FilterOpts, consensusAddr []common.Address, bridgeOperator []common.Address, recipient []common.Address) (*ValidatorBridgeOperatorRewardDistributionFailedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var bridgeOperatorRule []interface{}
	for _, bridgeOperatorItem := range bridgeOperator {
		bridgeOperatorRule = append(bridgeOperatorRule, bridgeOperatorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "BridgeOperatorRewardDistributionFailed", consensusAddrRule, bridgeOperatorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorBridgeOperatorRewardDistributionFailedIterator{contract: _Validator.contract, event: "BridgeOperatorRewardDistributionFailed", logs: logs, sub: sub}, nil
}

// WatchBridgeOperatorRewardDistributionFailed is a free log subscription operation binding the contract event 0xd35d76d87d51ed89407fc7ceaaccf32cf72784b94530892ce33546540e141b72.
//
// Solidity: event BridgeOperatorRewardDistributionFailed(address indexed consensusAddr, address indexed bridgeOperator, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_Validator *ValidatorFilterer) WatchBridgeOperatorRewardDistributionFailed(opts *bind.WatchOpts, sink chan<- *ValidatorBridgeOperatorRewardDistributionFailed, consensusAddr []common.Address, bridgeOperator []common.Address, recipient []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var bridgeOperatorRule []interface{}
	for _, bridgeOperatorItem := range bridgeOperator {
		bridgeOperatorRule = append(bridgeOperatorRule, bridgeOperatorItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "BridgeOperatorRewardDistributionFailed", consensusAddrRule, bridgeOperatorRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorBridgeOperatorRewardDistributionFailed)
				if err := _Validator.contract.UnpackLog(event, "BridgeOperatorRewardDistributionFailed", log); err != nil {
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

// ParseBridgeOperatorRewardDistributionFailed is a log parse operation binding the contract event 0xd35d76d87d51ed89407fc7ceaaccf32cf72784b94530892ce33546540e141b72.
//
// Solidity: event BridgeOperatorRewardDistributionFailed(address indexed consensusAddr, address indexed bridgeOperator, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_Validator *ValidatorFilterer) ParseBridgeOperatorRewardDistributionFailed(log types.Log) (*ValidatorBridgeOperatorRewardDistributionFailed, error) {
	event := new(ValidatorBridgeOperatorRewardDistributionFailed)
	if err := _Validator.contract.UnpackLog(event, "BridgeOperatorRewardDistributionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorBridgeOperatorSetUpdatedIterator is returned from FilterBridgeOperatorSetUpdated and is used to iterate over the raw logs and unpacked data for BridgeOperatorSetUpdated events raised by the Validator contract.
type ValidatorBridgeOperatorSetUpdatedIterator struct {
	Event *ValidatorBridgeOperatorSetUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorBridgeOperatorSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorBridgeOperatorSetUpdated)
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
		it.Event = new(ValidatorBridgeOperatorSetUpdated)
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
func (it *ValidatorBridgeOperatorSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorBridgeOperatorSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorBridgeOperatorSetUpdated represents a BridgeOperatorSetUpdated event raised by the Validator contract.
type ValidatorBridgeOperatorSetUpdated struct {
	Period          *big.Int
	Epoch           *big.Int
	BridgeOperators []common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgeOperatorSetUpdated is a free log retrieval operation binding the contract event 0x773d1888df530d69716b183a92450d45f97fba49f2a4bb34fae3b23da0e2cc6f.
//
// Solidity: event BridgeOperatorSetUpdated(uint256 indexed period, uint256 indexed epoch, address[] bridgeOperators)
func (_Validator *ValidatorFilterer) FilterBridgeOperatorSetUpdated(opts *bind.FilterOpts, period []*big.Int, epoch []*big.Int) (*ValidatorBridgeOperatorSetUpdatedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "BridgeOperatorSetUpdated", periodRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorBridgeOperatorSetUpdatedIterator{contract: _Validator.contract, event: "BridgeOperatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeOperatorSetUpdated is a free log subscription operation binding the contract event 0x773d1888df530d69716b183a92450d45f97fba49f2a4bb34fae3b23da0e2cc6f.
//
// Solidity: event BridgeOperatorSetUpdated(uint256 indexed period, uint256 indexed epoch, address[] bridgeOperators)
func (_Validator *ValidatorFilterer) WatchBridgeOperatorSetUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorBridgeOperatorSetUpdated, period []*big.Int, epoch []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "BridgeOperatorSetUpdated", periodRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorBridgeOperatorSetUpdated)
				if err := _Validator.contract.UnpackLog(event, "BridgeOperatorSetUpdated", log); err != nil {
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

// ParseBridgeOperatorSetUpdated is a log parse operation binding the contract event 0x773d1888df530d69716b183a92450d45f97fba49f2a4bb34fae3b23da0e2cc6f.
//
// Solidity: event BridgeOperatorSetUpdated(uint256 indexed period, uint256 indexed epoch, address[] bridgeOperators)
func (_Validator *ValidatorFilterer) ParseBridgeOperatorSetUpdated(log types.Log) (*ValidatorBridgeOperatorSetUpdated, error) {
	event := new(ValidatorBridgeOperatorSetUpdated)
	if err := _Validator.contract.UnpackLog(event, "BridgeOperatorSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorBridgeTrackingContractUpdatedIterator is returned from FilterBridgeTrackingContractUpdated and is used to iterate over the raw logs and unpacked data for BridgeTrackingContractUpdated events raised by the Validator contract.
type ValidatorBridgeTrackingContractUpdatedIterator struct {
	Event *ValidatorBridgeTrackingContractUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorBridgeTrackingContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorBridgeTrackingContractUpdated)
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
		it.Event = new(ValidatorBridgeTrackingContractUpdated)
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
func (it *ValidatorBridgeTrackingContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorBridgeTrackingContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorBridgeTrackingContractUpdated represents a BridgeTrackingContractUpdated event raised by the Validator contract.
type ValidatorBridgeTrackingContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBridgeTrackingContractUpdated is a free log retrieval operation binding the contract event 0x034c8da497df28467c79ddadbba1cc3cdd41f510ea73faae271e6f16a6111621.
//
// Solidity: event BridgeTrackingContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) FilterBridgeTrackingContractUpdated(opts *bind.FilterOpts) (*ValidatorBridgeTrackingContractUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "BridgeTrackingContractUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorBridgeTrackingContractUpdatedIterator{contract: _Validator.contract, event: "BridgeTrackingContractUpdated", logs: logs, sub: sub}, nil
}

// WatchBridgeTrackingContractUpdated is a free log subscription operation binding the contract event 0x034c8da497df28467c79ddadbba1cc3cdd41f510ea73faae271e6f16a6111621.
//
// Solidity: event BridgeTrackingContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) WatchBridgeTrackingContractUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorBridgeTrackingContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "BridgeTrackingContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorBridgeTrackingContractUpdated)
				if err := _Validator.contract.UnpackLog(event, "BridgeTrackingContractUpdated", log); err != nil {
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

// ParseBridgeTrackingContractUpdated is a log parse operation binding the contract event 0x034c8da497df28467c79ddadbba1cc3cdd41f510ea73faae271e6f16a6111621.
//
// Solidity: event BridgeTrackingContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) ParseBridgeTrackingContractUpdated(log types.Log) (*ValidatorBridgeTrackingContractUpdated, error) {
	event := new(ValidatorBridgeTrackingContractUpdated)
	if err := _Validator.contract.UnpackLog(event, "BridgeTrackingContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorBridgeTrackingIncorrectlyRespondedIterator is returned from FilterBridgeTrackingIncorrectlyResponded and is used to iterate over the raw logs and unpacked data for BridgeTrackingIncorrectlyResponded events raised by the Validator contract.
type ValidatorBridgeTrackingIncorrectlyRespondedIterator struct {
	Event *ValidatorBridgeTrackingIncorrectlyResponded // Event containing the contract specifics and raw log

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
func (it *ValidatorBridgeTrackingIncorrectlyRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorBridgeTrackingIncorrectlyResponded)
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
		it.Event = new(ValidatorBridgeTrackingIncorrectlyResponded)
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
func (it *ValidatorBridgeTrackingIncorrectlyRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorBridgeTrackingIncorrectlyRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorBridgeTrackingIncorrectlyResponded represents a BridgeTrackingIncorrectlyResponded event raised by the Validator contract.
type ValidatorBridgeTrackingIncorrectlyResponded struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBridgeTrackingIncorrectlyResponded is a free log retrieval operation binding the contract event 0x64ba7143ea5a17abea37667aa9ae137e3afba5033c5f504770c02829c128189c.
//
// Solidity: event BridgeTrackingIncorrectlyResponded()
func (_Validator *ValidatorFilterer) FilterBridgeTrackingIncorrectlyResponded(opts *bind.FilterOpts) (*ValidatorBridgeTrackingIncorrectlyRespondedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "BridgeTrackingIncorrectlyResponded")
	if err != nil {
		return nil, err
	}
	return &ValidatorBridgeTrackingIncorrectlyRespondedIterator{contract: _Validator.contract, event: "BridgeTrackingIncorrectlyResponded", logs: logs, sub: sub}, nil
}

// WatchBridgeTrackingIncorrectlyResponded is a free log subscription operation binding the contract event 0x64ba7143ea5a17abea37667aa9ae137e3afba5033c5f504770c02829c128189c.
//
// Solidity: event BridgeTrackingIncorrectlyResponded()
func (_Validator *ValidatorFilterer) WatchBridgeTrackingIncorrectlyResponded(opts *bind.WatchOpts, sink chan<- *ValidatorBridgeTrackingIncorrectlyResponded) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "BridgeTrackingIncorrectlyResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorBridgeTrackingIncorrectlyResponded)
				if err := _Validator.contract.UnpackLog(event, "BridgeTrackingIncorrectlyResponded", log); err != nil {
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

// ParseBridgeTrackingIncorrectlyResponded is a log parse operation binding the contract event 0x64ba7143ea5a17abea37667aa9ae137e3afba5033c5f504770c02829c128189c.
//
// Solidity: event BridgeTrackingIncorrectlyResponded()
func (_Validator *ValidatorFilterer) ParseBridgeTrackingIncorrectlyResponded(log types.Log) (*ValidatorBridgeTrackingIncorrectlyResponded, error) {
	event := new(ValidatorBridgeTrackingIncorrectlyResponded)
	if err := _Validator.contract.UnpackLog(event, "BridgeTrackingIncorrectlyResponded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorCandidateGrantedIterator is returned from FilterCandidateGranted and is used to iterate over the raw logs and unpacked data for CandidateGranted events raised by the Validator contract.
type ValidatorCandidateGrantedIterator struct {
	Event *ValidatorCandidateGranted // Event containing the contract specifics and raw log

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
func (it *ValidatorCandidateGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorCandidateGranted)
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
		it.Event = new(ValidatorCandidateGranted)
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
func (it *ValidatorCandidateGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorCandidateGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorCandidateGranted represents a CandidateGranted event raised by the Validator contract.
type ValidatorCandidateGranted struct {
	ConsensusAddr  common.Address
	TreasuryAddr   common.Address
	Admin          common.Address
	BridgeOperator common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCandidateGranted is a free log retrieval operation binding the contract event 0xd690f592ed983cfbc05717fbcf06c4e10ae328432c309fe49246cf4a4be69fcd.
//
// Solidity: event CandidateGranted(address indexed consensusAddr, address indexed treasuryAddr, address indexed admin, address bridgeOperator)
func (_Validator *ValidatorFilterer) FilterCandidateGranted(opts *bind.FilterOpts, consensusAddr []common.Address, treasuryAddr []common.Address, admin []common.Address) (*ValidatorCandidateGrantedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var treasuryAddrRule []interface{}
	for _, treasuryAddrItem := range treasuryAddr {
		treasuryAddrRule = append(treasuryAddrRule, treasuryAddrItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "CandidateGranted", consensusAddrRule, treasuryAddrRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorCandidateGrantedIterator{contract: _Validator.contract, event: "CandidateGranted", logs: logs, sub: sub}, nil
}

// WatchCandidateGranted is a free log subscription operation binding the contract event 0xd690f592ed983cfbc05717fbcf06c4e10ae328432c309fe49246cf4a4be69fcd.
//
// Solidity: event CandidateGranted(address indexed consensusAddr, address indexed treasuryAddr, address indexed admin, address bridgeOperator)
func (_Validator *ValidatorFilterer) WatchCandidateGranted(opts *bind.WatchOpts, sink chan<- *ValidatorCandidateGranted, consensusAddr []common.Address, treasuryAddr []common.Address, admin []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var treasuryAddrRule []interface{}
	for _, treasuryAddrItem := range treasuryAddr {
		treasuryAddrRule = append(treasuryAddrRule, treasuryAddrItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "CandidateGranted", consensusAddrRule, treasuryAddrRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorCandidateGranted)
				if err := _Validator.contract.UnpackLog(event, "CandidateGranted", log); err != nil {
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

// ParseCandidateGranted is a log parse operation binding the contract event 0xd690f592ed983cfbc05717fbcf06c4e10ae328432c309fe49246cf4a4be69fcd.
//
// Solidity: event CandidateGranted(address indexed consensusAddr, address indexed treasuryAddr, address indexed admin, address bridgeOperator)
func (_Validator *ValidatorFilterer) ParseCandidateGranted(log types.Log) (*ValidatorCandidateGranted, error) {
	event := new(ValidatorCandidateGranted)
	if err := _Validator.contract.UnpackLog(event, "CandidateGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorCandidateRevokingTimestampUpdatedIterator is returned from FilterCandidateRevokingTimestampUpdated and is used to iterate over the raw logs and unpacked data for CandidateRevokingTimestampUpdated events raised by the Validator contract.
type ValidatorCandidateRevokingTimestampUpdatedIterator struct {
	Event *ValidatorCandidateRevokingTimestampUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorCandidateRevokingTimestampUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorCandidateRevokingTimestampUpdated)
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
		it.Event = new(ValidatorCandidateRevokingTimestampUpdated)
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
func (it *ValidatorCandidateRevokingTimestampUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorCandidateRevokingTimestampUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorCandidateRevokingTimestampUpdated represents a CandidateRevokingTimestampUpdated event raised by the Validator contract.
type ValidatorCandidateRevokingTimestampUpdated struct {
	ConsensusAddr     common.Address
	RevokingTimestamp *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCandidateRevokingTimestampUpdated is a free log retrieval operation binding the contract event 0xb9a1e33376bfbda9092f2d1e37662c1b435aab0d3fa8da3acc8f37ee222f99e7.
//
// Solidity: event CandidateRevokingTimestampUpdated(address indexed consensusAddr, uint256 revokingTimestamp)
func (_Validator *ValidatorFilterer) FilterCandidateRevokingTimestampUpdated(opts *bind.FilterOpts, consensusAddr []common.Address) (*ValidatorCandidateRevokingTimestampUpdatedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "CandidateRevokingTimestampUpdated", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorCandidateRevokingTimestampUpdatedIterator{contract: _Validator.contract, event: "CandidateRevokingTimestampUpdated", logs: logs, sub: sub}, nil
}

// WatchCandidateRevokingTimestampUpdated is a free log subscription operation binding the contract event 0xb9a1e33376bfbda9092f2d1e37662c1b435aab0d3fa8da3acc8f37ee222f99e7.
//
// Solidity: event CandidateRevokingTimestampUpdated(address indexed consensusAddr, uint256 revokingTimestamp)
func (_Validator *ValidatorFilterer) WatchCandidateRevokingTimestampUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorCandidateRevokingTimestampUpdated, consensusAddr []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "CandidateRevokingTimestampUpdated", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorCandidateRevokingTimestampUpdated)
				if err := _Validator.contract.UnpackLog(event, "CandidateRevokingTimestampUpdated", log); err != nil {
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

// ParseCandidateRevokingTimestampUpdated is a log parse operation binding the contract event 0xb9a1e33376bfbda9092f2d1e37662c1b435aab0d3fa8da3acc8f37ee222f99e7.
//
// Solidity: event CandidateRevokingTimestampUpdated(address indexed consensusAddr, uint256 revokingTimestamp)
func (_Validator *ValidatorFilterer) ParseCandidateRevokingTimestampUpdated(log types.Log) (*ValidatorCandidateRevokingTimestampUpdated, error) {
	event := new(ValidatorCandidateRevokingTimestampUpdated)
	if err := _Validator.contract.UnpackLog(event, "CandidateRevokingTimestampUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorCandidateTopupDeadlineUpdatedIterator is returned from FilterCandidateTopupDeadlineUpdated and is used to iterate over the raw logs and unpacked data for CandidateTopupDeadlineUpdated events raised by the Validator contract.
type ValidatorCandidateTopupDeadlineUpdatedIterator struct {
	Event *ValidatorCandidateTopupDeadlineUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorCandidateTopupDeadlineUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorCandidateTopupDeadlineUpdated)
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
		it.Event = new(ValidatorCandidateTopupDeadlineUpdated)
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
func (it *ValidatorCandidateTopupDeadlineUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorCandidateTopupDeadlineUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorCandidateTopupDeadlineUpdated represents a CandidateTopupDeadlineUpdated event raised by the Validator contract.
type ValidatorCandidateTopupDeadlineUpdated struct {
	ConsensusAddr common.Address
	TopupDeadline *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCandidateTopupDeadlineUpdated is a free log retrieval operation binding the contract event 0x88f854e137380c14d63f6ed99781bf13402167cf55bac49bcd44d4f2d6a34275.
//
// Solidity: event CandidateTopupDeadlineUpdated(address indexed consensusAddr, uint256 topupDeadline)
func (_Validator *ValidatorFilterer) FilterCandidateTopupDeadlineUpdated(opts *bind.FilterOpts, consensusAddr []common.Address) (*ValidatorCandidateTopupDeadlineUpdatedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "CandidateTopupDeadlineUpdated", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorCandidateTopupDeadlineUpdatedIterator{contract: _Validator.contract, event: "CandidateTopupDeadlineUpdated", logs: logs, sub: sub}, nil
}

// WatchCandidateTopupDeadlineUpdated is a free log subscription operation binding the contract event 0x88f854e137380c14d63f6ed99781bf13402167cf55bac49bcd44d4f2d6a34275.
//
// Solidity: event CandidateTopupDeadlineUpdated(address indexed consensusAddr, uint256 topupDeadline)
func (_Validator *ValidatorFilterer) WatchCandidateTopupDeadlineUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorCandidateTopupDeadlineUpdated, consensusAddr []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "CandidateTopupDeadlineUpdated", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorCandidateTopupDeadlineUpdated)
				if err := _Validator.contract.UnpackLog(event, "CandidateTopupDeadlineUpdated", log); err != nil {
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

// ParseCandidateTopupDeadlineUpdated is a log parse operation binding the contract event 0x88f854e137380c14d63f6ed99781bf13402167cf55bac49bcd44d4f2d6a34275.
//
// Solidity: event CandidateTopupDeadlineUpdated(address indexed consensusAddr, uint256 topupDeadline)
func (_Validator *ValidatorFilterer) ParseCandidateTopupDeadlineUpdated(log types.Log) (*ValidatorCandidateTopupDeadlineUpdated, error) {
	event := new(ValidatorCandidateTopupDeadlineUpdated)
	if err := _Validator.contract.UnpackLog(event, "CandidateTopupDeadlineUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorCandidatesRevokedIterator is returned from FilterCandidatesRevoked and is used to iterate over the raw logs and unpacked data for CandidatesRevoked events raised by the Validator contract.
type ValidatorCandidatesRevokedIterator struct {
	Event *ValidatorCandidatesRevoked // Event containing the contract specifics and raw log

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
func (it *ValidatorCandidatesRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorCandidatesRevoked)
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
		it.Event = new(ValidatorCandidatesRevoked)
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
func (it *ValidatorCandidatesRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorCandidatesRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorCandidatesRevoked represents a CandidatesRevoked event raised by the Validator contract.
type ValidatorCandidatesRevoked struct {
	ConsensusAddrs []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCandidatesRevoked is a free log retrieval operation binding the contract event 0x4eaf233b9dc25a5552c1927feee1412eea69add17c2485c831c2e60e234f3c91.
//
// Solidity: event CandidatesRevoked(address[] consensusAddrs)
func (_Validator *ValidatorFilterer) FilterCandidatesRevoked(opts *bind.FilterOpts) (*ValidatorCandidatesRevokedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "CandidatesRevoked")
	if err != nil {
		return nil, err
	}
	return &ValidatorCandidatesRevokedIterator{contract: _Validator.contract, event: "CandidatesRevoked", logs: logs, sub: sub}, nil
}

// WatchCandidatesRevoked is a free log subscription operation binding the contract event 0x4eaf233b9dc25a5552c1927feee1412eea69add17c2485c831c2e60e234f3c91.
//
// Solidity: event CandidatesRevoked(address[] consensusAddrs)
func (_Validator *ValidatorFilterer) WatchCandidatesRevoked(opts *bind.WatchOpts, sink chan<- *ValidatorCandidatesRevoked) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "CandidatesRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorCandidatesRevoked)
				if err := _Validator.contract.UnpackLog(event, "CandidatesRevoked", log); err != nil {
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

// ParseCandidatesRevoked is a log parse operation binding the contract event 0x4eaf233b9dc25a5552c1927feee1412eea69add17c2485c831c2e60e234f3c91.
//
// Solidity: event CandidatesRevoked(address[] consensusAddrs)
func (_Validator *ValidatorFilterer) ParseCandidatesRevoked(log types.Log) (*ValidatorCandidatesRevoked, error) {
	event := new(ValidatorCandidatesRevoked)
	if err := _Validator.contract.UnpackLog(event, "CandidatesRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorCommissionRateUpdateScheduledIterator is returned from FilterCommissionRateUpdateScheduled and is used to iterate over the raw logs and unpacked data for CommissionRateUpdateScheduled events raised by the Validator contract.
type ValidatorCommissionRateUpdateScheduledIterator struct {
	Event *ValidatorCommissionRateUpdateScheduled // Event containing the contract specifics and raw log

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
func (it *ValidatorCommissionRateUpdateScheduledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorCommissionRateUpdateScheduled)
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
		it.Event = new(ValidatorCommissionRateUpdateScheduled)
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
func (it *ValidatorCommissionRateUpdateScheduledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorCommissionRateUpdateScheduledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorCommissionRateUpdateScheduled represents a CommissionRateUpdateScheduled event raised by the Validator contract.
type ValidatorCommissionRateUpdateScheduled struct {
	ConsensusAddr      common.Address
	EffectiveTimestamp *big.Int
	Rate               *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCommissionRateUpdateScheduled is a free log retrieval operation binding the contract event 0x6ebafd1bb6316b2f63198a81b05cff2149c6eaae1784466a6d062b4391900f21.
//
// Solidity: event CommissionRateUpdateScheduled(address indexed consensusAddr, uint256 effectiveTimestamp, uint256 rate)
func (_Validator *ValidatorFilterer) FilterCommissionRateUpdateScheduled(opts *bind.FilterOpts, consensusAddr []common.Address) (*ValidatorCommissionRateUpdateScheduledIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "CommissionRateUpdateScheduled", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorCommissionRateUpdateScheduledIterator{contract: _Validator.contract, event: "CommissionRateUpdateScheduled", logs: logs, sub: sub}, nil
}

// WatchCommissionRateUpdateScheduled is a free log subscription operation binding the contract event 0x6ebafd1bb6316b2f63198a81b05cff2149c6eaae1784466a6d062b4391900f21.
//
// Solidity: event CommissionRateUpdateScheduled(address indexed consensusAddr, uint256 effectiveTimestamp, uint256 rate)
func (_Validator *ValidatorFilterer) WatchCommissionRateUpdateScheduled(opts *bind.WatchOpts, sink chan<- *ValidatorCommissionRateUpdateScheduled, consensusAddr []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "CommissionRateUpdateScheduled", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorCommissionRateUpdateScheduled)
				if err := _Validator.contract.UnpackLog(event, "CommissionRateUpdateScheduled", log); err != nil {
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

// ParseCommissionRateUpdateScheduled is a log parse operation binding the contract event 0x6ebafd1bb6316b2f63198a81b05cff2149c6eaae1784466a6d062b4391900f21.
//
// Solidity: event CommissionRateUpdateScheduled(address indexed consensusAddr, uint256 effectiveTimestamp, uint256 rate)
func (_Validator *ValidatorFilterer) ParseCommissionRateUpdateScheduled(log types.Log) (*ValidatorCommissionRateUpdateScheduled, error) {
	event := new(ValidatorCommissionRateUpdateScheduled)
	if err := _Validator.contract.UnpackLog(event, "CommissionRateUpdateScheduled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorCommissionRateUpdatedIterator is returned from FilterCommissionRateUpdated and is used to iterate over the raw logs and unpacked data for CommissionRateUpdated events raised by the Validator contract.
type ValidatorCommissionRateUpdatedIterator struct {
	Event *ValidatorCommissionRateUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorCommissionRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorCommissionRateUpdated)
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
		it.Event = new(ValidatorCommissionRateUpdated)
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
func (it *ValidatorCommissionRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorCommissionRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorCommissionRateUpdated represents a CommissionRateUpdated event raised by the Validator contract.
type ValidatorCommissionRateUpdated struct {
	ConsensusAddr common.Address
	Rate          *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCommissionRateUpdated is a free log retrieval operation binding the contract event 0x86d576c20e383fc2413ef692209cc48ddad5e52f25db5b32f8f7ec5076461ae9.
//
// Solidity: event CommissionRateUpdated(address indexed consensusAddr, uint256 rate)
func (_Validator *ValidatorFilterer) FilterCommissionRateUpdated(opts *bind.FilterOpts, consensusAddr []common.Address) (*ValidatorCommissionRateUpdatedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "CommissionRateUpdated", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorCommissionRateUpdatedIterator{contract: _Validator.contract, event: "CommissionRateUpdated", logs: logs, sub: sub}, nil
}

// WatchCommissionRateUpdated is a free log subscription operation binding the contract event 0x86d576c20e383fc2413ef692209cc48ddad5e52f25db5b32f8f7ec5076461ae9.
//
// Solidity: event CommissionRateUpdated(address indexed consensusAddr, uint256 rate)
func (_Validator *ValidatorFilterer) WatchCommissionRateUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorCommissionRateUpdated, consensusAddr []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "CommissionRateUpdated", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorCommissionRateUpdated)
				if err := _Validator.contract.UnpackLog(event, "CommissionRateUpdated", log); err != nil {
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

// ParseCommissionRateUpdated is a log parse operation binding the contract event 0x86d576c20e383fc2413ef692209cc48ddad5e52f25db5b32f8f7ec5076461ae9.
//
// Solidity: event CommissionRateUpdated(address indexed consensusAddr, uint256 rate)
func (_Validator *ValidatorFilterer) ParseCommissionRateUpdated(log types.Log) (*ValidatorCommissionRateUpdated, error) {
	event := new(ValidatorCommissionRateUpdated)
	if err := _Validator.contract.UnpackLog(event, "CommissionRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorDeprecatedRewardRecycleFailedIterator is returned from FilterDeprecatedRewardRecycleFailed and is used to iterate over the raw logs and unpacked data for DeprecatedRewardRecycleFailed events raised by the Validator contract.
type ValidatorDeprecatedRewardRecycleFailedIterator struct {
	Event *ValidatorDeprecatedRewardRecycleFailed // Event containing the contract specifics and raw log

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
func (it *ValidatorDeprecatedRewardRecycleFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorDeprecatedRewardRecycleFailed)
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
		it.Event = new(ValidatorDeprecatedRewardRecycleFailed)
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
func (it *ValidatorDeprecatedRewardRecycleFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorDeprecatedRewardRecycleFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorDeprecatedRewardRecycleFailed represents a DeprecatedRewardRecycleFailed event raised by the Validator contract.
type ValidatorDeprecatedRewardRecycleFailed struct {
	RecipientAddr common.Address
	Amount        *big.Int
	Balance       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeprecatedRewardRecycleFailed is a free log retrieval operation binding the contract event 0xa0561a59abed308fcd0556834574739d778cc6229018039a24ddda0f86aa0b73.
//
// Solidity: event DeprecatedRewardRecycleFailed(address indexed recipientAddr, uint256 amount, uint256 balance)
func (_Validator *ValidatorFilterer) FilterDeprecatedRewardRecycleFailed(opts *bind.FilterOpts, recipientAddr []common.Address) (*ValidatorDeprecatedRewardRecycleFailedIterator, error) {

	var recipientAddrRule []interface{}
	for _, recipientAddrItem := range recipientAddr {
		recipientAddrRule = append(recipientAddrRule, recipientAddrItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "DeprecatedRewardRecycleFailed", recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorDeprecatedRewardRecycleFailedIterator{contract: _Validator.contract, event: "DeprecatedRewardRecycleFailed", logs: logs, sub: sub}, nil
}

// WatchDeprecatedRewardRecycleFailed is a free log subscription operation binding the contract event 0xa0561a59abed308fcd0556834574739d778cc6229018039a24ddda0f86aa0b73.
//
// Solidity: event DeprecatedRewardRecycleFailed(address indexed recipientAddr, uint256 amount, uint256 balance)
func (_Validator *ValidatorFilterer) WatchDeprecatedRewardRecycleFailed(opts *bind.WatchOpts, sink chan<- *ValidatorDeprecatedRewardRecycleFailed, recipientAddr []common.Address) (event.Subscription, error) {

	var recipientAddrRule []interface{}
	for _, recipientAddrItem := range recipientAddr {
		recipientAddrRule = append(recipientAddrRule, recipientAddrItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "DeprecatedRewardRecycleFailed", recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorDeprecatedRewardRecycleFailed)
				if err := _Validator.contract.UnpackLog(event, "DeprecatedRewardRecycleFailed", log); err != nil {
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

// ParseDeprecatedRewardRecycleFailed is a log parse operation binding the contract event 0xa0561a59abed308fcd0556834574739d778cc6229018039a24ddda0f86aa0b73.
//
// Solidity: event DeprecatedRewardRecycleFailed(address indexed recipientAddr, uint256 amount, uint256 balance)
func (_Validator *ValidatorFilterer) ParseDeprecatedRewardRecycleFailed(log types.Log) (*ValidatorDeprecatedRewardRecycleFailed, error) {
	event := new(ValidatorDeprecatedRewardRecycleFailed)
	if err := _Validator.contract.UnpackLog(event, "DeprecatedRewardRecycleFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorDeprecatedRewardRecycledIterator is returned from FilterDeprecatedRewardRecycled and is used to iterate over the raw logs and unpacked data for DeprecatedRewardRecycled events raised by the Validator contract.
type ValidatorDeprecatedRewardRecycledIterator struct {
	Event *ValidatorDeprecatedRewardRecycled // Event containing the contract specifics and raw log

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
func (it *ValidatorDeprecatedRewardRecycledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorDeprecatedRewardRecycled)
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
		it.Event = new(ValidatorDeprecatedRewardRecycled)
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
func (it *ValidatorDeprecatedRewardRecycledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorDeprecatedRewardRecycledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorDeprecatedRewardRecycled represents a DeprecatedRewardRecycled event raised by the Validator contract.
type ValidatorDeprecatedRewardRecycled struct {
	RecipientAddr common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeprecatedRewardRecycled is a free log retrieval operation binding the contract event 0xc447c884574da5878be39c403db2245c22530c99b579ea7bcbb3720e1d110dc8.
//
// Solidity: event DeprecatedRewardRecycled(address indexed recipientAddr, uint256 amount)
func (_Validator *ValidatorFilterer) FilterDeprecatedRewardRecycled(opts *bind.FilterOpts, recipientAddr []common.Address) (*ValidatorDeprecatedRewardRecycledIterator, error) {

	var recipientAddrRule []interface{}
	for _, recipientAddrItem := range recipientAddr {
		recipientAddrRule = append(recipientAddrRule, recipientAddrItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "DeprecatedRewardRecycled", recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorDeprecatedRewardRecycledIterator{contract: _Validator.contract, event: "DeprecatedRewardRecycled", logs: logs, sub: sub}, nil
}

// WatchDeprecatedRewardRecycled is a free log subscription operation binding the contract event 0xc447c884574da5878be39c403db2245c22530c99b579ea7bcbb3720e1d110dc8.
//
// Solidity: event DeprecatedRewardRecycled(address indexed recipientAddr, uint256 amount)
func (_Validator *ValidatorFilterer) WatchDeprecatedRewardRecycled(opts *bind.WatchOpts, sink chan<- *ValidatorDeprecatedRewardRecycled, recipientAddr []common.Address) (event.Subscription, error) {

	var recipientAddrRule []interface{}
	for _, recipientAddrItem := range recipientAddr {
		recipientAddrRule = append(recipientAddrRule, recipientAddrItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "DeprecatedRewardRecycled", recipientAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorDeprecatedRewardRecycled)
				if err := _Validator.contract.UnpackLog(event, "DeprecatedRewardRecycled", log); err != nil {
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

// ParseDeprecatedRewardRecycled is a log parse operation binding the contract event 0xc447c884574da5878be39c403db2245c22530c99b579ea7bcbb3720e1d110dc8.
//
// Solidity: event DeprecatedRewardRecycled(address indexed recipientAddr, uint256 amount)
func (_Validator *ValidatorFilterer) ParseDeprecatedRewardRecycled(log types.Log) (*ValidatorDeprecatedRewardRecycled, error) {
	event := new(ValidatorDeprecatedRewardRecycled)
	if err := _Validator.contract.UnpackLog(event, "DeprecatedRewardRecycled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorEmergencyExitLockedAmountUpdatedIterator is returned from FilterEmergencyExitLockedAmountUpdated and is used to iterate over the raw logs and unpacked data for EmergencyExitLockedAmountUpdated events raised by the Validator contract.
type ValidatorEmergencyExitLockedAmountUpdatedIterator struct {
	Event *ValidatorEmergencyExitLockedAmountUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorEmergencyExitLockedAmountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorEmergencyExitLockedAmountUpdated)
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
		it.Event = new(ValidatorEmergencyExitLockedAmountUpdated)
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
func (it *ValidatorEmergencyExitLockedAmountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorEmergencyExitLockedAmountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorEmergencyExitLockedAmountUpdated represents a EmergencyExitLockedAmountUpdated event raised by the Validator contract.
type ValidatorEmergencyExitLockedAmountUpdated struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExitLockedAmountUpdated is a free log retrieval operation binding the contract event 0x17a6c3eb965cdd7439982da25abf85be88f0f772ca33198f548e2f99fee0289a.
//
// Solidity: event EmergencyExitLockedAmountUpdated(uint256 amount)
func (_Validator *ValidatorFilterer) FilterEmergencyExitLockedAmountUpdated(opts *bind.FilterOpts) (*ValidatorEmergencyExitLockedAmountUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "EmergencyExitLockedAmountUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorEmergencyExitLockedAmountUpdatedIterator{contract: _Validator.contract, event: "EmergencyExitLockedAmountUpdated", logs: logs, sub: sub}, nil
}

// WatchEmergencyExitLockedAmountUpdated is a free log subscription operation binding the contract event 0x17a6c3eb965cdd7439982da25abf85be88f0f772ca33198f548e2f99fee0289a.
//
// Solidity: event EmergencyExitLockedAmountUpdated(uint256 amount)
func (_Validator *ValidatorFilterer) WatchEmergencyExitLockedAmountUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorEmergencyExitLockedAmountUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "EmergencyExitLockedAmountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorEmergencyExitLockedAmountUpdated)
				if err := _Validator.contract.UnpackLog(event, "EmergencyExitLockedAmountUpdated", log); err != nil {
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

// ParseEmergencyExitLockedAmountUpdated is a log parse operation binding the contract event 0x17a6c3eb965cdd7439982da25abf85be88f0f772ca33198f548e2f99fee0289a.
//
// Solidity: event EmergencyExitLockedAmountUpdated(uint256 amount)
func (_Validator *ValidatorFilterer) ParseEmergencyExitLockedAmountUpdated(log types.Log) (*ValidatorEmergencyExitLockedAmountUpdated, error) {
	event := new(ValidatorEmergencyExitLockedAmountUpdated)
	if err := _Validator.contract.UnpackLog(event, "EmergencyExitLockedAmountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorEmergencyExitLockedFundReleasedIterator is returned from FilterEmergencyExitLockedFundReleased and is used to iterate over the raw logs and unpacked data for EmergencyExitLockedFundReleased events raised by the Validator contract.
type ValidatorEmergencyExitLockedFundReleasedIterator struct {
	Event *ValidatorEmergencyExitLockedFundReleased // Event containing the contract specifics and raw log

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
func (it *ValidatorEmergencyExitLockedFundReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorEmergencyExitLockedFundReleased)
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
		it.Event = new(ValidatorEmergencyExitLockedFundReleased)
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
func (it *ValidatorEmergencyExitLockedFundReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorEmergencyExitLockedFundReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorEmergencyExitLockedFundReleased represents a EmergencyExitLockedFundReleased event raised by the Validator contract.
type ValidatorEmergencyExitLockedFundReleased struct {
	ConsensusAddr  common.Address
	Recipient      common.Address
	UnlockedAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExitLockedFundReleased is a free log retrieval operation binding the contract event 0x7229136a18186c71a86246c012af3bb1df6460ef163934bbdccd6368abdd41e4.
//
// Solidity: event EmergencyExitLockedFundReleased(address indexed consensusAddr, address indexed recipient, uint256 unlockedAmount)
func (_Validator *ValidatorFilterer) FilterEmergencyExitLockedFundReleased(opts *bind.FilterOpts, consensusAddr []common.Address, recipient []common.Address) (*ValidatorEmergencyExitLockedFundReleasedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "EmergencyExitLockedFundReleased", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorEmergencyExitLockedFundReleasedIterator{contract: _Validator.contract, event: "EmergencyExitLockedFundReleased", logs: logs, sub: sub}, nil
}

// WatchEmergencyExitLockedFundReleased is a free log subscription operation binding the contract event 0x7229136a18186c71a86246c012af3bb1df6460ef163934bbdccd6368abdd41e4.
//
// Solidity: event EmergencyExitLockedFundReleased(address indexed consensusAddr, address indexed recipient, uint256 unlockedAmount)
func (_Validator *ValidatorFilterer) WatchEmergencyExitLockedFundReleased(opts *bind.WatchOpts, sink chan<- *ValidatorEmergencyExitLockedFundReleased, consensusAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "EmergencyExitLockedFundReleased", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorEmergencyExitLockedFundReleased)
				if err := _Validator.contract.UnpackLog(event, "EmergencyExitLockedFundReleased", log); err != nil {
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

// ParseEmergencyExitLockedFundReleased is a log parse operation binding the contract event 0x7229136a18186c71a86246c012af3bb1df6460ef163934bbdccd6368abdd41e4.
//
// Solidity: event EmergencyExitLockedFundReleased(address indexed consensusAddr, address indexed recipient, uint256 unlockedAmount)
func (_Validator *ValidatorFilterer) ParseEmergencyExitLockedFundReleased(log types.Log) (*ValidatorEmergencyExitLockedFundReleased, error) {
	event := new(ValidatorEmergencyExitLockedFundReleased)
	if err := _Validator.contract.UnpackLog(event, "EmergencyExitLockedFundReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorEmergencyExitLockedFundReleasingFailedIterator is returned from FilterEmergencyExitLockedFundReleasingFailed and is used to iterate over the raw logs and unpacked data for EmergencyExitLockedFundReleasingFailed events raised by the Validator contract.
type ValidatorEmergencyExitLockedFundReleasingFailedIterator struct {
	Event *ValidatorEmergencyExitLockedFundReleasingFailed // Event containing the contract specifics and raw log

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
func (it *ValidatorEmergencyExitLockedFundReleasingFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorEmergencyExitLockedFundReleasingFailed)
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
		it.Event = new(ValidatorEmergencyExitLockedFundReleasingFailed)
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
func (it *ValidatorEmergencyExitLockedFundReleasingFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorEmergencyExitLockedFundReleasingFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorEmergencyExitLockedFundReleasingFailed represents a EmergencyExitLockedFundReleasingFailed event raised by the Validator contract.
type ValidatorEmergencyExitLockedFundReleasingFailed struct {
	ConsensusAddr   common.Address
	Recipient       common.Address
	UnlockedAmount  *big.Int
	ContractBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExitLockedFundReleasingFailed is a free log retrieval operation binding the contract event 0x3747d14eb72ad3e35cba9c3e00dab3b8d15b40cac6bdbd08402356e4f69f30a1.
//
// Solidity: event EmergencyExitLockedFundReleasingFailed(address indexed consensusAddr, address indexed recipient, uint256 unlockedAmount, uint256 contractBalance)
func (_Validator *ValidatorFilterer) FilterEmergencyExitLockedFundReleasingFailed(opts *bind.FilterOpts, consensusAddr []common.Address, recipient []common.Address) (*ValidatorEmergencyExitLockedFundReleasingFailedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "EmergencyExitLockedFundReleasingFailed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorEmergencyExitLockedFundReleasingFailedIterator{contract: _Validator.contract, event: "EmergencyExitLockedFundReleasingFailed", logs: logs, sub: sub}, nil
}

// WatchEmergencyExitLockedFundReleasingFailed is a free log subscription operation binding the contract event 0x3747d14eb72ad3e35cba9c3e00dab3b8d15b40cac6bdbd08402356e4f69f30a1.
//
// Solidity: event EmergencyExitLockedFundReleasingFailed(address indexed consensusAddr, address indexed recipient, uint256 unlockedAmount, uint256 contractBalance)
func (_Validator *ValidatorFilterer) WatchEmergencyExitLockedFundReleasingFailed(opts *bind.WatchOpts, sink chan<- *ValidatorEmergencyExitLockedFundReleasingFailed, consensusAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "EmergencyExitLockedFundReleasingFailed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorEmergencyExitLockedFundReleasingFailed)
				if err := _Validator.contract.UnpackLog(event, "EmergencyExitLockedFundReleasingFailed", log); err != nil {
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

// ParseEmergencyExitLockedFundReleasingFailed is a log parse operation binding the contract event 0x3747d14eb72ad3e35cba9c3e00dab3b8d15b40cac6bdbd08402356e4f69f30a1.
//
// Solidity: event EmergencyExitLockedFundReleasingFailed(address indexed consensusAddr, address indexed recipient, uint256 unlockedAmount, uint256 contractBalance)
func (_Validator *ValidatorFilterer) ParseEmergencyExitLockedFundReleasingFailed(log types.Log) (*ValidatorEmergencyExitLockedFundReleasingFailed, error) {
	event := new(ValidatorEmergencyExitLockedFundReleasingFailed)
	if err := _Validator.contract.UnpackLog(event, "EmergencyExitLockedFundReleasingFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorEmergencyExitRequestedIterator is returned from FilterEmergencyExitRequested and is used to iterate over the raw logs and unpacked data for EmergencyExitRequested events raised by the Validator contract.
type ValidatorEmergencyExitRequestedIterator struct {
	Event *ValidatorEmergencyExitRequested // Event containing the contract specifics and raw log

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
func (it *ValidatorEmergencyExitRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorEmergencyExitRequested)
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
		it.Event = new(ValidatorEmergencyExitRequested)
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
func (it *ValidatorEmergencyExitRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorEmergencyExitRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorEmergencyExitRequested represents a EmergencyExitRequested event raised by the Validator contract.
type ValidatorEmergencyExitRequested struct {
	ConsensusAddr common.Address
	LockedAmount  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExitRequested is a free log retrieval operation binding the contract event 0x77a1a819870c0f4d04c3ca4cc2881a0393136abc28bd651af50aedade94a27c4.
//
// Solidity: event EmergencyExitRequested(address indexed consensusAddr, uint256 lockedAmount)
func (_Validator *ValidatorFilterer) FilterEmergencyExitRequested(opts *bind.FilterOpts, consensusAddr []common.Address) (*ValidatorEmergencyExitRequestedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "EmergencyExitRequested", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorEmergencyExitRequestedIterator{contract: _Validator.contract, event: "EmergencyExitRequested", logs: logs, sub: sub}, nil
}

// WatchEmergencyExitRequested is a free log subscription operation binding the contract event 0x77a1a819870c0f4d04c3ca4cc2881a0393136abc28bd651af50aedade94a27c4.
//
// Solidity: event EmergencyExitRequested(address indexed consensusAddr, uint256 lockedAmount)
func (_Validator *ValidatorFilterer) WatchEmergencyExitRequested(opts *bind.WatchOpts, sink chan<- *ValidatorEmergencyExitRequested, consensusAddr []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "EmergencyExitRequested", consensusAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorEmergencyExitRequested)
				if err := _Validator.contract.UnpackLog(event, "EmergencyExitRequested", log); err != nil {
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

// ParseEmergencyExitRequested is a log parse operation binding the contract event 0x77a1a819870c0f4d04c3ca4cc2881a0393136abc28bd651af50aedade94a27c4.
//
// Solidity: event EmergencyExitRequested(address indexed consensusAddr, uint256 lockedAmount)
func (_Validator *ValidatorFilterer) ParseEmergencyExitRequested(log types.Log) (*ValidatorEmergencyExitRequested, error) {
	event := new(ValidatorEmergencyExitRequested)
	if err := _Validator.contract.UnpackLog(event, "EmergencyExitRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorEmergencyExpiryDurationUpdatedIterator is returned from FilterEmergencyExpiryDurationUpdated and is used to iterate over the raw logs and unpacked data for EmergencyExpiryDurationUpdated events raised by the Validator contract.
type ValidatorEmergencyExpiryDurationUpdatedIterator struct {
	Event *ValidatorEmergencyExpiryDurationUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorEmergencyExpiryDurationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorEmergencyExpiryDurationUpdated)
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
		it.Event = new(ValidatorEmergencyExpiryDurationUpdated)
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
func (it *ValidatorEmergencyExpiryDurationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorEmergencyExpiryDurationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorEmergencyExpiryDurationUpdated represents a EmergencyExpiryDurationUpdated event raised by the Validator contract.
type ValidatorEmergencyExpiryDurationUpdated struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyExpiryDurationUpdated is a free log retrieval operation binding the contract event 0x0a50c66137118f386332efb949231ddd3946100dbf880003daca37ddd9e0662b.
//
// Solidity: event EmergencyExpiryDurationUpdated(uint256 amount)
func (_Validator *ValidatorFilterer) FilterEmergencyExpiryDurationUpdated(opts *bind.FilterOpts) (*ValidatorEmergencyExpiryDurationUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "EmergencyExpiryDurationUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorEmergencyExpiryDurationUpdatedIterator{contract: _Validator.contract, event: "EmergencyExpiryDurationUpdated", logs: logs, sub: sub}, nil
}

// WatchEmergencyExpiryDurationUpdated is a free log subscription operation binding the contract event 0x0a50c66137118f386332efb949231ddd3946100dbf880003daca37ddd9e0662b.
//
// Solidity: event EmergencyExpiryDurationUpdated(uint256 amount)
func (_Validator *ValidatorFilterer) WatchEmergencyExpiryDurationUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorEmergencyExpiryDurationUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "EmergencyExpiryDurationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorEmergencyExpiryDurationUpdated)
				if err := _Validator.contract.UnpackLog(event, "EmergencyExpiryDurationUpdated", log); err != nil {
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

// ParseEmergencyExpiryDurationUpdated is a log parse operation binding the contract event 0x0a50c66137118f386332efb949231ddd3946100dbf880003daca37ddd9e0662b.
//
// Solidity: event EmergencyExpiryDurationUpdated(uint256 amount)
func (_Validator *ValidatorFilterer) ParseEmergencyExpiryDurationUpdated(log types.Log) (*ValidatorEmergencyExpiryDurationUpdated, error) {
	event := new(ValidatorEmergencyExpiryDurationUpdated)
	if err := _Validator.contract.UnpackLog(event, "EmergencyExpiryDurationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Validator contract.
type ValidatorInitializedIterator struct {
	Event *ValidatorInitialized // Event containing the contract specifics and raw log

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
func (it *ValidatorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorInitialized)
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
		it.Event = new(ValidatorInitialized)
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
func (it *ValidatorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorInitialized represents a Initialized event raised by the Validator contract.
type ValidatorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Validator *ValidatorFilterer) FilterInitialized(opts *bind.FilterOpts) (*ValidatorInitializedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ValidatorInitializedIterator{contract: _Validator.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Validator *ValidatorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ValidatorInitialized) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorInitialized)
				if err := _Validator.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Validator *ValidatorFilterer) ParseInitialized(log types.Log) (*ValidatorInitialized, error) {
	event := new(ValidatorInitialized)
	if err := _Validator.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorMaintenanceContractUpdatedIterator is returned from FilterMaintenanceContractUpdated and is used to iterate over the raw logs and unpacked data for MaintenanceContractUpdated events raised by the Validator contract.
type ValidatorMaintenanceContractUpdatedIterator struct {
	Event *ValidatorMaintenanceContractUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorMaintenanceContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorMaintenanceContractUpdated)
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
		it.Event = new(ValidatorMaintenanceContractUpdated)
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
func (it *ValidatorMaintenanceContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorMaintenanceContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorMaintenanceContractUpdated represents a MaintenanceContractUpdated event raised by the Validator contract.
type ValidatorMaintenanceContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMaintenanceContractUpdated is a free log retrieval operation binding the contract event 0x31a33f126a5bae3c5bdf6cfc2cd6dcfffe2fe9634bdb09e21c44762993889e3b.
//
// Solidity: event MaintenanceContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) FilterMaintenanceContractUpdated(opts *bind.FilterOpts) (*ValidatorMaintenanceContractUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "MaintenanceContractUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorMaintenanceContractUpdatedIterator{contract: _Validator.contract, event: "MaintenanceContractUpdated", logs: logs, sub: sub}, nil
}

// WatchMaintenanceContractUpdated is a free log subscription operation binding the contract event 0x31a33f126a5bae3c5bdf6cfc2cd6dcfffe2fe9634bdb09e21c44762993889e3b.
//
// Solidity: event MaintenanceContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) WatchMaintenanceContractUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorMaintenanceContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "MaintenanceContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorMaintenanceContractUpdated)
				if err := _Validator.contract.UnpackLog(event, "MaintenanceContractUpdated", log); err != nil {
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

// ParseMaintenanceContractUpdated is a log parse operation binding the contract event 0x31a33f126a5bae3c5bdf6cfc2cd6dcfffe2fe9634bdb09e21c44762993889e3b.
//
// Solidity: event MaintenanceContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) ParseMaintenanceContractUpdated(log types.Log) (*ValidatorMaintenanceContractUpdated, error) {
	event := new(ValidatorMaintenanceContractUpdated)
	if err := _Validator.contract.UnpackLog(event, "MaintenanceContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorMaxPrioritizedValidatorNumberUpdatedIterator is returned from FilterMaxPrioritizedValidatorNumberUpdated and is used to iterate over the raw logs and unpacked data for MaxPrioritizedValidatorNumberUpdated events raised by the Validator contract.
type ValidatorMaxPrioritizedValidatorNumberUpdatedIterator struct {
	Event *ValidatorMaxPrioritizedValidatorNumberUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorMaxPrioritizedValidatorNumberUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorMaxPrioritizedValidatorNumberUpdated)
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
		it.Event = new(ValidatorMaxPrioritizedValidatorNumberUpdated)
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
func (it *ValidatorMaxPrioritizedValidatorNumberUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorMaxPrioritizedValidatorNumberUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorMaxPrioritizedValidatorNumberUpdated represents a MaxPrioritizedValidatorNumberUpdated event raised by the Validator contract.
type ValidatorMaxPrioritizedValidatorNumberUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMaxPrioritizedValidatorNumberUpdated is a free log retrieval operation binding the contract event 0xa9588dc77416849bd922605ce4fc806712281ad8a8f32d4238d6c8cca548e15e.
//
// Solidity: event MaxPrioritizedValidatorNumberUpdated(uint256 arg0)
func (_Validator *ValidatorFilterer) FilterMaxPrioritizedValidatorNumberUpdated(opts *bind.FilterOpts) (*ValidatorMaxPrioritizedValidatorNumberUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "MaxPrioritizedValidatorNumberUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorMaxPrioritizedValidatorNumberUpdatedIterator{contract: _Validator.contract, event: "MaxPrioritizedValidatorNumberUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxPrioritizedValidatorNumberUpdated is a free log subscription operation binding the contract event 0xa9588dc77416849bd922605ce4fc806712281ad8a8f32d4238d6c8cca548e15e.
//
// Solidity: event MaxPrioritizedValidatorNumberUpdated(uint256 arg0)
func (_Validator *ValidatorFilterer) WatchMaxPrioritizedValidatorNumberUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorMaxPrioritizedValidatorNumberUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "MaxPrioritizedValidatorNumberUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorMaxPrioritizedValidatorNumberUpdated)
				if err := _Validator.contract.UnpackLog(event, "MaxPrioritizedValidatorNumberUpdated", log); err != nil {
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

// ParseMaxPrioritizedValidatorNumberUpdated is a log parse operation binding the contract event 0xa9588dc77416849bd922605ce4fc806712281ad8a8f32d4238d6c8cca548e15e.
//
// Solidity: event MaxPrioritizedValidatorNumberUpdated(uint256 arg0)
func (_Validator *ValidatorFilterer) ParseMaxPrioritizedValidatorNumberUpdated(log types.Log) (*ValidatorMaxPrioritizedValidatorNumberUpdated, error) {
	event := new(ValidatorMaxPrioritizedValidatorNumberUpdated)
	if err := _Validator.contract.UnpackLog(event, "MaxPrioritizedValidatorNumberUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorMaxValidatorCandidateUpdatedIterator is returned from FilterMaxValidatorCandidateUpdated and is used to iterate over the raw logs and unpacked data for MaxValidatorCandidateUpdated events raised by the Validator contract.
type ValidatorMaxValidatorCandidateUpdatedIterator struct {
	Event *ValidatorMaxValidatorCandidateUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorMaxValidatorCandidateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorMaxValidatorCandidateUpdated)
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
		it.Event = new(ValidatorMaxValidatorCandidateUpdated)
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
func (it *ValidatorMaxValidatorCandidateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorMaxValidatorCandidateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorMaxValidatorCandidateUpdated represents a MaxValidatorCandidateUpdated event raised by the Validator contract.
type ValidatorMaxValidatorCandidateUpdated struct {
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMaxValidatorCandidateUpdated is a free log retrieval operation binding the contract event 0x82d5dc32d1b741512ad09c32404d7e7921e8934c6222343d95f55f7a2b9b2ab4.
//
// Solidity: event MaxValidatorCandidateUpdated(uint256 threshold)
func (_Validator *ValidatorFilterer) FilterMaxValidatorCandidateUpdated(opts *bind.FilterOpts) (*ValidatorMaxValidatorCandidateUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "MaxValidatorCandidateUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorMaxValidatorCandidateUpdatedIterator{contract: _Validator.contract, event: "MaxValidatorCandidateUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxValidatorCandidateUpdated is a free log subscription operation binding the contract event 0x82d5dc32d1b741512ad09c32404d7e7921e8934c6222343d95f55f7a2b9b2ab4.
//
// Solidity: event MaxValidatorCandidateUpdated(uint256 threshold)
func (_Validator *ValidatorFilterer) WatchMaxValidatorCandidateUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorMaxValidatorCandidateUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "MaxValidatorCandidateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorMaxValidatorCandidateUpdated)
				if err := _Validator.contract.UnpackLog(event, "MaxValidatorCandidateUpdated", log); err != nil {
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

// ParseMaxValidatorCandidateUpdated is a log parse operation binding the contract event 0x82d5dc32d1b741512ad09c32404d7e7921e8934c6222343d95f55f7a2b9b2ab4.
//
// Solidity: event MaxValidatorCandidateUpdated(uint256 threshold)
func (_Validator *ValidatorFilterer) ParseMaxValidatorCandidateUpdated(log types.Log) (*ValidatorMaxValidatorCandidateUpdated, error) {
	event := new(ValidatorMaxValidatorCandidateUpdated)
	if err := _Validator.contract.UnpackLog(event, "MaxValidatorCandidateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorMaxValidatorNumberUpdatedIterator is returned from FilterMaxValidatorNumberUpdated and is used to iterate over the raw logs and unpacked data for MaxValidatorNumberUpdated events raised by the Validator contract.
type ValidatorMaxValidatorNumberUpdatedIterator struct {
	Event *ValidatorMaxValidatorNumberUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorMaxValidatorNumberUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorMaxValidatorNumberUpdated)
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
		it.Event = new(ValidatorMaxValidatorNumberUpdated)
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
func (it *ValidatorMaxValidatorNumberUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorMaxValidatorNumberUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorMaxValidatorNumberUpdated represents a MaxValidatorNumberUpdated event raised by the Validator contract.
type ValidatorMaxValidatorNumberUpdated struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMaxValidatorNumberUpdated is a free log retrieval operation binding the contract event 0xb5464c05fd0e0f000c535850116cda2742ee1f7b34384cb920ad7b8e802138b5.
//
// Solidity: event MaxValidatorNumberUpdated(uint256 arg0)
func (_Validator *ValidatorFilterer) FilterMaxValidatorNumberUpdated(opts *bind.FilterOpts) (*ValidatorMaxValidatorNumberUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "MaxValidatorNumberUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorMaxValidatorNumberUpdatedIterator{contract: _Validator.contract, event: "MaxValidatorNumberUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxValidatorNumberUpdated is a free log subscription operation binding the contract event 0xb5464c05fd0e0f000c535850116cda2742ee1f7b34384cb920ad7b8e802138b5.
//
// Solidity: event MaxValidatorNumberUpdated(uint256 arg0)
func (_Validator *ValidatorFilterer) WatchMaxValidatorNumberUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorMaxValidatorNumberUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "MaxValidatorNumberUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorMaxValidatorNumberUpdated)
				if err := _Validator.contract.UnpackLog(event, "MaxValidatorNumberUpdated", log); err != nil {
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

// ParseMaxValidatorNumberUpdated is a log parse operation binding the contract event 0xb5464c05fd0e0f000c535850116cda2742ee1f7b34384cb920ad7b8e802138b5.
//
// Solidity: event MaxValidatorNumberUpdated(uint256 arg0)
func (_Validator *ValidatorFilterer) ParseMaxValidatorNumberUpdated(log types.Log) (*ValidatorMaxValidatorNumberUpdated, error) {
	event := new(ValidatorMaxValidatorNumberUpdated)
	if err := _Validator.contract.UnpackLog(event, "MaxValidatorNumberUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorMinEffectiveDaysOnwardsUpdatedIterator is returned from FilterMinEffectiveDaysOnwardsUpdated and is used to iterate over the raw logs and unpacked data for MinEffectiveDaysOnwardsUpdated events raised by the Validator contract.
type ValidatorMinEffectiveDaysOnwardsUpdatedIterator struct {
	Event *ValidatorMinEffectiveDaysOnwardsUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorMinEffectiveDaysOnwardsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorMinEffectiveDaysOnwardsUpdated)
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
		it.Event = new(ValidatorMinEffectiveDaysOnwardsUpdated)
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
func (it *ValidatorMinEffectiveDaysOnwardsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorMinEffectiveDaysOnwardsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorMinEffectiveDaysOnwardsUpdated represents a MinEffectiveDaysOnwardsUpdated event raised by the Validator contract.
type ValidatorMinEffectiveDaysOnwardsUpdated struct {
	NumOfDays *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinEffectiveDaysOnwardsUpdated is a free log retrieval operation binding the contract event 0x266d432ffe659e3565750d26ec685b822a58041eee724b67a5afec3168a25267.
//
// Solidity: event MinEffectiveDaysOnwardsUpdated(uint256 numOfDays)
func (_Validator *ValidatorFilterer) FilterMinEffectiveDaysOnwardsUpdated(opts *bind.FilterOpts) (*ValidatorMinEffectiveDaysOnwardsUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "MinEffectiveDaysOnwardsUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorMinEffectiveDaysOnwardsUpdatedIterator{contract: _Validator.contract, event: "MinEffectiveDaysOnwardsUpdated", logs: logs, sub: sub}, nil
}

// WatchMinEffectiveDaysOnwardsUpdated is a free log subscription operation binding the contract event 0x266d432ffe659e3565750d26ec685b822a58041eee724b67a5afec3168a25267.
//
// Solidity: event MinEffectiveDaysOnwardsUpdated(uint256 numOfDays)
func (_Validator *ValidatorFilterer) WatchMinEffectiveDaysOnwardsUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorMinEffectiveDaysOnwardsUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "MinEffectiveDaysOnwardsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorMinEffectiveDaysOnwardsUpdated)
				if err := _Validator.contract.UnpackLog(event, "MinEffectiveDaysOnwardsUpdated", log); err != nil {
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

// ParseMinEffectiveDaysOnwardsUpdated is a log parse operation binding the contract event 0x266d432ffe659e3565750d26ec685b822a58041eee724b67a5afec3168a25267.
//
// Solidity: event MinEffectiveDaysOnwardsUpdated(uint256 numOfDays)
func (_Validator *ValidatorFilterer) ParseMinEffectiveDaysOnwardsUpdated(log types.Log) (*ValidatorMinEffectiveDaysOnwardsUpdated, error) {
	event := new(ValidatorMinEffectiveDaysOnwardsUpdated)
	if err := _Validator.contract.UnpackLog(event, "MinEffectiveDaysOnwardsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorMiningRewardDistributedIterator is returned from FilterMiningRewardDistributed and is used to iterate over the raw logs and unpacked data for MiningRewardDistributed events raised by the Validator contract.
type ValidatorMiningRewardDistributedIterator struct {
	Event *ValidatorMiningRewardDistributed // Event containing the contract specifics and raw log

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
func (it *ValidatorMiningRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorMiningRewardDistributed)
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
		it.Event = new(ValidatorMiningRewardDistributed)
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
func (it *ValidatorMiningRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorMiningRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorMiningRewardDistributed represents a MiningRewardDistributed event raised by the Validator contract.
type ValidatorMiningRewardDistributed struct {
	ConsensusAddr common.Address
	Recipient     common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMiningRewardDistributed is a free log retrieval operation binding the contract event 0x1ce7a1c4702402cd393500acb1de5bd927727a54e144a587d328f1b679abe4ec.
//
// Solidity: event MiningRewardDistributed(address indexed consensusAddr, address indexed recipient, uint256 amount)
func (_Validator *ValidatorFilterer) FilterMiningRewardDistributed(opts *bind.FilterOpts, consensusAddr []common.Address, recipient []common.Address) (*ValidatorMiningRewardDistributedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "MiningRewardDistributed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorMiningRewardDistributedIterator{contract: _Validator.contract, event: "MiningRewardDistributed", logs: logs, sub: sub}, nil
}

// WatchMiningRewardDistributed is a free log subscription operation binding the contract event 0x1ce7a1c4702402cd393500acb1de5bd927727a54e144a587d328f1b679abe4ec.
//
// Solidity: event MiningRewardDistributed(address indexed consensusAddr, address indexed recipient, uint256 amount)
func (_Validator *ValidatorFilterer) WatchMiningRewardDistributed(opts *bind.WatchOpts, sink chan<- *ValidatorMiningRewardDistributed, consensusAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "MiningRewardDistributed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorMiningRewardDistributed)
				if err := _Validator.contract.UnpackLog(event, "MiningRewardDistributed", log); err != nil {
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

// ParseMiningRewardDistributed is a log parse operation binding the contract event 0x1ce7a1c4702402cd393500acb1de5bd927727a54e144a587d328f1b679abe4ec.
//
// Solidity: event MiningRewardDistributed(address indexed consensusAddr, address indexed recipient, uint256 amount)
func (_Validator *ValidatorFilterer) ParseMiningRewardDistributed(log types.Log) (*ValidatorMiningRewardDistributed, error) {
	event := new(ValidatorMiningRewardDistributed)
	if err := _Validator.contract.UnpackLog(event, "MiningRewardDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorMiningRewardDistributionFailedIterator is returned from FilterMiningRewardDistributionFailed and is used to iterate over the raw logs and unpacked data for MiningRewardDistributionFailed events raised by the Validator contract.
type ValidatorMiningRewardDistributionFailedIterator struct {
	Event *ValidatorMiningRewardDistributionFailed // Event containing the contract specifics and raw log

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
func (it *ValidatorMiningRewardDistributionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorMiningRewardDistributionFailed)
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
		it.Event = new(ValidatorMiningRewardDistributionFailed)
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
func (it *ValidatorMiningRewardDistributionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorMiningRewardDistributionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorMiningRewardDistributionFailed represents a MiningRewardDistributionFailed event raised by the Validator contract.
type ValidatorMiningRewardDistributionFailed struct {
	ConsensusAddr   common.Address
	Recipient       common.Address
	Amount          *big.Int
	ContractBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMiningRewardDistributionFailed is a free log retrieval operation binding the contract event 0x6c69e09ee5c5ac33c0cd57787261c5bade070a392ab34a4b5487c6868f723f6e.
//
// Solidity: event MiningRewardDistributionFailed(address indexed consensusAddr, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_Validator *ValidatorFilterer) FilterMiningRewardDistributionFailed(opts *bind.FilterOpts, consensusAddr []common.Address, recipient []common.Address) (*ValidatorMiningRewardDistributionFailedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "MiningRewardDistributionFailed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorMiningRewardDistributionFailedIterator{contract: _Validator.contract, event: "MiningRewardDistributionFailed", logs: logs, sub: sub}, nil
}

// WatchMiningRewardDistributionFailed is a free log subscription operation binding the contract event 0x6c69e09ee5c5ac33c0cd57787261c5bade070a392ab34a4b5487c6868f723f6e.
//
// Solidity: event MiningRewardDistributionFailed(address indexed consensusAddr, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_Validator *ValidatorFilterer) WatchMiningRewardDistributionFailed(opts *bind.WatchOpts, sink chan<- *ValidatorMiningRewardDistributionFailed, consensusAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "MiningRewardDistributionFailed", consensusAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorMiningRewardDistributionFailed)
				if err := _Validator.contract.UnpackLog(event, "MiningRewardDistributionFailed", log); err != nil {
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

// ParseMiningRewardDistributionFailed is a log parse operation binding the contract event 0x6c69e09ee5c5ac33c0cd57787261c5bade070a392ab34a4b5487c6868f723f6e.
//
// Solidity: event MiningRewardDistributionFailed(address indexed consensusAddr, address indexed recipient, uint256 amount, uint256 contractBalance)
func (_Validator *ValidatorFilterer) ParseMiningRewardDistributionFailed(log types.Log) (*ValidatorMiningRewardDistributionFailed, error) {
	event := new(ValidatorMiningRewardDistributionFailed)
	if err := _Validator.contract.UnpackLog(event, "MiningRewardDistributionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorRoninTrustedOrganizationContractUpdatedIterator is returned from FilterRoninTrustedOrganizationContractUpdated and is used to iterate over the raw logs and unpacked data for RoninTrustedOrganizationContractUpdated events raised by the Validator contract.
type ValidatorRoninTrustedOrganizationContractUpdatedIterator struct {
	Event *ValidatorRoninTrustedOrganizationContractUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorRoninTrustedOrganizationContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorRoninTrustedOrganizationContractUpdated)
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
		it.Event = new(ValidatorRoninTrustedOrganizationContractUpdated)
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
func (it *ValidatorRoninTrustedOrganizationContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorRoninTrustedOrganizationContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorRoninTrustedOrganizationContractUpdated represents a RoninTrustedOrganizationContractUpdated event raised by the Validator contract.
type ValidatorRoninTrustedOrganizationContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRoninTrustedOrganizationContractUpdated is a free log retrieval operation binding the contract event 0xfd6f5f93d69a07c593a09be0b208bff13ab4ffd6017df3b33433d63bdc59b4d7.
//
// Solidity: event RoninTrustedOrganizationContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) FilterRoninTrustedOrganizationContractUpdated(opts *bind.FilterOpts) (*ValidatorRoninTrustedOrganizationContractUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "RoninTrustedOrganizationContractUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorRoninTrustedOrganizationContractUpdatedIterator{contract: _Validator.contract, event: "RoninTrustedOrganizationContractUpdated", logs: logs, sub: sub}, nil
}

// WatchRoninTrustedOrganizationContractUpdated is a free log subscription operation binding the contract event 0xfd6f5f93d69a07c593a09be0b208bff13ab4ffd6017df3b33433d63bdc59b4d7.
//
// Solidity: event RoninTrustedOrganizationContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) WatchRoninTrustedOrganizationContractUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorRoninTrustedOrganizationContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "RoninTrustedOrganizationContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorRoninTrustedOrganizationContractUpdated)
				if err := _Validator.contract.UnpackLog(event, "RoninTrustedOrganizationContractUpdated", log); err != nil {
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

// ParseRoninTrustedOrganizationContractUpdated is a log parse operation binding the contract event 0xfd6f5f93d69a07c593a09be0b208bff13ab4ffd6017df3b33433d63bdc59b4d7.
//
// Solidity: event RoninTrustedOrganizationContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) ParseRoninTrustedOrganizationContractUpdated(log types.Log) (*ValidatorRoninTrustedOrganizationContractUpdated, error) {
	event := new(ValidatorRoninTrustedOrganizationContractUpdated)
	if err := _Validator.contract.UnpackLog(event, "RoninTrustedOrganizationContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorSlashIndicatorContractUpdatedIterator is returned from FilterSlashIndicatorContractUpdated and is used to iterate over the raw logs and unpacked data for SlashIndicatorContractUpdated events raised by the Validator contract.
type ValidatorSlashIndicatorContractUpdatedIterator struct {
	Event *ValidatorSlashIndicatorContractUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorSlashIndicatorContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorSlashIndicatorContractUpdated)
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
		it.Event = new(ValidatorSlashIndicatorContractUpdated)
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
func (it *ValidatorSlashIndicatorContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorSlashIndicatorContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorSlashIndicatorContractUpdated represents a SlashIndicatorContractUpdated event raised by the Validator contract.
type ValidatorSlashIndicatorContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSlashIndicatorContractUpdated is a free log retrieval operation binding the contract event 0xaa5b07dd43aa44c69b70a6a2b9c3fcfed12b6e5f6323596ba7ac91035ab80a4f.
//
// Solidity: event SlashIndicatorContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) FilterSlashIndicatorContractUpdated(opts *bind.FilterOpts) (*ValidatorSlashIndicatorContractUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "SlashIndicatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorSlashIndicatorContractUpdatedIterator{contract: _Validator.contract, event: "SlashIndicatorContractUpdated", logs: logs, sub: sub}, nil
}

// WatchSlashIndicatorContractUpdated is a free log subscription operation binding the contract event 0xaa5b07dd43aa44c69b70a6a2b9c3fcfed12b6e5f6323596ba7ac91035ab80a4f.
//
// Solidity: event SlashIndicatorContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) WatchSlashIndicatorContractUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorSlashIndicatorContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "SlashIndicatorContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorSlashIndicatorContractUpdated)
				if err := _Validator.contract.UnpackLog(event, "SlashIndicatorContractUpdated", log); err != nil {
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

// ParseSlashIndicatorContractUpdated is a log parse operation binding the contract event 0xaa5b07dd43aa44c69b70a6a2b9c3fcfed12b6e5f6323596ba7ac91035ab80a4f.
//
// Solidity: event SlashIndicatorContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) ParseSlashIndicatorContractUpdated(log types.Log) (*ValidatorSlashIndicatorContractUpdated, error) {
	event := new(ValidatorSlashIndicatorContractUpdated)
	if err := _Validator.contract.UnpackLog(event, "SlashIndicatorContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorStakingContractUpdatedIterator is returned from FilterStakingContractUpdated and is used to iterate over the raw logs and unpacked data for StakingContractUpdated events raised by the Validator contract.
type ValidatorStakingContractUpdatedIterator struct {
	Event *ValidatorStakingContractUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorStakingContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorStakingContractUpdated)
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
		it.Event = new(ValidatorStakingContractUpdated)
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
func (it *ValidatorStakingContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorStakingContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorStakingContractUpdated represents a StakingContractUpdated event raised by the Validator contract.
type ValidatorStakingContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterStakingContractUpdated is a free log retrieval operation binding the contract event 0x6397f5b135542bb3f477cb346cfab5abdec1251d08dc8f8d4efb4ffe122ea0bf.
//
// Solidity: event StakingContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) FilterStakingContractUpdated(opts *bind.FilterOpts) (*ValidatorStakingContractUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "StakingContractUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorStakingContractUpdatedIterator{contract: _Validator.contract, event: "StakingContractUpdated", logs: logs, sub: sub}, nil
}

// WatchStakingContractUpdated is a free log subscription operation binding the contract event 0x6397f5b135542bb3f477cb346cfab5abdec1251d08dc8f8d4efb4ffe122ea0bf.
//
// Solidity: event StakingContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) WatchStakingContractUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorStakingContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "StakingContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorStakingContractUpdated)
				if err := _Validator.contract.UnpackLog(event, "StakingContractUpdated", log); err != nil {
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

// ParseStakingContractUpdated is a log parse operation binding the contract event 0x6397f5b135542bb3f477cb346cfab5abdec1251d08dc8f8d4efb4ffe122ea0bf.
//
// Solidity: event StakingContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) ParseStakingContractUpdated(log types.Log) (*ValidatorStakingContractUpdated, error) {
	event := new(ValidatorStakingContractUpdated)
	if err := _Validator.contract.UnpackLog(event, "StakingContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorStakingRewardDistributedIterator is returned from FilterStakingRewardDistributed and is used to iterate over the raw logs and unpacked data for StakingRewardDistributed events raised by the Validator contract.
type ValidatorStakingRewardDistributedIterator struct {
	Event *ValidatorStakingRewardDistributed // Event containing the contract specifics and raw log

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
func (it *ValidatorStakingRewardDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorStakingRewardDistributed)
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
		it.Event = new(ValidatorStakingRewardDistributed)
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
func (it *ValidatorStakingRewardDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorStakingRewardDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorStakingRewardDistributed represents a StakingRewardDistributed event raised by the Validator contract.
type ValidatorStakingRewardDistributed struct {
	TotalAmount    *big.Int
	ConsensusAddrs []common.Address
	Amounts        []*big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStakingRewardDistributed is a free log retrieval operation binding the contract event 0x9e242ca1ef9dde96eb71ef8d19a3f0f6a619b63e4c0d3998771387103656d087.
//
// Solidity: event StakingRewardDistributed(uint256 totalAmount, address[] consensusAddrs, uint256[] amounts)
func (_Validator *ValidatorFilterer) FilterStakingRewardDistributed(opts *bind.FilterOpts) (*ValidatorStakingRewardDistributedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "StakingRewardDistributed")
	if err != nil {
		return nil, err
	}
	return &ValidatorStakingRewardDistributedIterator{contract: _Validator.contract, event: "StakingRewardDistributed", logs: logs, sub: sub}, nil
}

// WatchStakingRewardDistributed is a free log subscription operation binding the contract event 0x9e242ca1ef9dde96eb71ef8d19a3f0f6a619b63e4c0d3998771387103656d087.
//
// Solidity: event StakingRewardDistributed(uint256 totalAmount, address[] consensusAddrs, uint256[] amounts)
func (_Validator *ValidatorFilterer) WatchStakingRewardDistributed(opts *bind.WatchOpts, sink chan<- *ValidatorStakingRewardDistributed) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "StakingRewardDistributed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorStakingRewardDistributed)
				if err := _Validator.contract.UnpackLog(event, "StakingRewardDistributed", log); err != nil {
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

// ParseStakingRewardDistributed is a log parse operation binding the contract event 0x9e242ca1ef9dde96eb71ef8d19a3f0f6a619b63e4c0d3998771387103656d087.
//
// Solidity: event StakingRewardDistributed(uint256 totalAmount, address[] consensusAddrs, uint256[] amounts)
func (_Validator *ValidatorFilterer) ParseStakingRewardDistributed(log types.Log) (*ValidatorStakingRewardDistributed, error) {
	event := new(ValidatorStakingRewardDistributed)
	if err := _Validator.contract.UnpackLog(event, "StakingRewardDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorStakingRewardDistributionFailedIterator is returned from FilterStakingRewardDistributionFailed and is used to iterate over the raw logs and unpacked data for StakingRewardDistributionFailed events raised by the Validator contract.
type ValidatorStakingRewardDistributionFailedIterator struct {
	Event *ValidatorStakingRewardDistributionFailed // Event containing the contract specifics and raw log

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
func (it *ValidatorStakingRewardDistributionFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorStakingRewardDistributionFailed)
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
		it.Event = new(ValidatorStakingRewardDistributionFailed)
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
func (it *ValidatorStakingRewardDistributionFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorStakingRewardDistributionFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorStakingRewardDistributionFailed represents a StakingRewardDistributionFailed event raised by the Validator contract.
type ValidatorStakingRewardDistributionFailed struct {
	TotalAmount     *big.Int
	ConsensusAddrs  []common.Address
	Amounts         []*big.Int
	ContractBalance *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakingRewardDistributionFailed is a free log retrieval operation binding the contract event 0xe5668ec1bb2b6bb144a50f810e388da4b1d7d3fc05fcb9d588a1aac59d248f89.
//
// Solidity: event StakingRewardDistributionFailed(uint256 totalAmount, address[] consensusAddrs, uint256[] amounts, uint256 contractBalance)
func (_Validator *ValidatorFilterer) FilterStakingRewardDistributionFailed(opts *bind.FilterOpts) (*ValidatorStakingRewardDistributionFailedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "StakingRewardDistributionFailed")
	if err != nil {
		return nil, err
	}
	return &ValidatorStakingRewardDistributionFailedIterator{contract: _Validator.contract, event: "StakingRewardDistributionFailed", logs: logs, sub: sub}, nil
}

// WatchStakingRewardDistributionFailed is a free log subscription operation binding the contract event 0xe5668ec1bb2b6bb144a50f810e388da4b1d7d3fc05fcb9d588a1aac59d248f89.
//
// Solidity: event StakingRewardDistributionFailed(uint256 totalAmount, address[] consensusAddrs, uint256[] amounts, uint256 contractBalance)
func (_Validator *ValidatorFilterer) WatchStakingRewardDistributionFailed(opts *bind.WatchOpts, sink chan<- *ValidatorStakingRewardDistributionFailed) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "StakingRewardDistributionFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorStakingRewardDistributionFailed)
				if err := _Validator.contract.UnpackLog(event, "StakingRewardDistributionFailed", log); err != nil {
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

// ParseStakingRewardDistributionFailed is a log parse operation binding the contract event 0xe5668ec1bb2b6bb144a50f810e388da4b1d7d3fc05fcb9d588a1aac59d248f89.
//
// Solidity: event StakingRewardDistributionFailed(uint256 totalAmount, address[] consensusAddrs, uint256[] amounts, uint256 contractBalance)
func (_Validator *ValidatorFilterer) ParseStakingRewardDistributionFailed(log types.Log) (*ValidatorStakingRewardDistributionFailed, error) {
	event := new(ValidatorStakingRewardDistributionFailed)
	if err := _Validator.contract.UnpackLog(event, "StakingRewardDistributionFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorStakingVestingContractUpdatedIterator is returned from FilterStakingVestingContractUpdated and is used to iterate over the raw logs and unpacked data for StakingVestingContractUpdated events raised by the Validator contract.
type ValidatorStakingVestingContractUpdatedIterator struct {
	Event *ValidatorStakingVestingContractUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorStakingVestingContractUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorStakingVestingContractUpdated)
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
		it.Event = new(ValidatorStakingVestingContractUpdated)
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
func (it *ValidatorStakingVestingContractUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorStakingVestingContractUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorStakingVestingContractUpdated represents a StakingVestingContractUpdated event raised by the Validator contract.
type ValidatorStakingVestingContractUpdated struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterStakingVestingContractUpdated is a free log retrieval operation binding the contract event 0xc328090a37d855191ab58469296f98f87a851ca57d5cdfd1e9ac3c83e9e7096d.
//
// Solidity: event StakingVestingContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) FilterStakingVestingContractUpdated(opts *bind.FilterOpts) (*ValidatorStakingVestingContractUpdatedIterator, error) {

	logs, sub, err := _Validator.contract.FilterLogs(opts, "StakingVestingContractUpdated")
	if err != nil {
		return nil, err
	}
	return &ValidatorStakingVestingContractUpdatedIterator{contract: _Validator.contract, event: "StakingVestingContractUpdated", logs: logs, sub: sub}, nil
}

// WatchStakingVestingContractUpdated is a free log subscription operation binding the contract event 0xc328090a37d855191ab58469296f98f87a851ca57d5cdfd1e9ac3c83e9e7096d.
//
// Solidity: event StakingVestingContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) WatchStakingVestingContractUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorStakingVestingContractUpdated) (event.Subscription, error) {

	logs, sub, err := _Validator.contract.WatchLogs(opts, "StakingVestingContractUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorStakingVestingContractUpdated)
				if err := _Validator.contract.UnpackLog(event, "StakingVestingContractUpdated", log); err != nil {
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

// ParseStakingVestingContractUpdated is a log parse operation binding the contract event 0xc328090a37d855191ab58469296f98f87a851ca57d5cdfd1e9ac3c83e9e7096d.
//
// Solidity: event StakingVestingContractUpdated(address arg0)
func (_Validator *ValidatorFilterer) ParseStakingVestingContractUpdated(log types.Log) (*ValidatorStakingVestingContractUpdated, error) {
	event := new(ValidatorStakingVestingContractUpdated)
	if err := _Validator.contract.UnpackLog(event, "StakingVestingContractUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorValidatorPunishedIterator is returned from FilterValidatorPunished and is used to iterate over the raw logs and unpacked data for ValidatorPunished events raised by the Validator contract.
type ValidatorValidatorPunishedIterator struct {
	Event *ValidatorValidatorPunished // Event containing the contract specifics and raw log

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
func (it *ValidatorValidatorPunishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorValidatorPunished)
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
		it.Event = new(ValidatorValidatorPunished)
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
func (it *ValidatorValidatorPunishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorValidatorPunishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorValidatorPunished represents a ValidatorPunished event raised by the Validator contract.
type ValidatorValidatorPunished struct {
	ConsensusAddr                  common.Address
	Period                         *big.Int
	JailedUntil                    *big.Int
	DeductedStakingAmount          *big.Int
	BlockProducerRewardDeprecated  bool
	BridgeOperatorRewardDeprecated bool
	Raw                            types.Log // Blockchain specific contextual infos
}

// FilterValidatorPunished is a free log retrieval operation binding the contract event 0x54ce99c5ce1fc9f61656d4a0fb2697974d0c973ac32eecaefe06fcf18b8ef68a.
//
// Solidity: event ValidatorPunished(address indexed consensusAddr, uint256 indexed period, uint256 jailedUntil, uint256 deductedStakingAmount, bool blockProducerRewardDeprecated, bool bridgeOperatorRewardDeprecated)
func (_Validator *ValidatorFilterer) FilterValidatorPunished(opts *bind.FilterOpts, consensusAddr []common.Address, period []*big.Int) (*ValidatorValidatorPunishedIterator, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "ValidatorPunished", consensusAddrRule, periodRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorValidatorPunishedIterator{contract: _Validator.contract, event: "ValidatorPunished", logs: logs, sub: sub}, nil
}

// WatchValidatorPunished is a free log subscription operation binding the contract event 0x54ce99c5ce1fc9f61656d4a0fb2697974d0c973ac32eecaefe06fcf18b8ef68a.
//
// Solidity: event ValidatorPunished(address indexed consensusAddr, uint256 indexed period, uint256 jailedUntil, uint256 deductedStakingAmount, bool blockProducerRewardDeprecated, bool bridgeOperatorRewardDeprecated)
func (_Validator *ValidatorFilterer) WatchValidatorPunished(opts *bind.WatchOpts, sink chan<- *ValidatorValidatorPunished, consensusAddr []common.Address, period []*big.Int) (event.Subscription, error) {

	var consensusAddrRule []interface{}
	for _, consensusAddrItem := range consensusAddr {
		consensusAddrRule = append(consensusAddrRule, consensusAddrItem)
	}
	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "ValidatorPunished", consensusAddrRule, periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorValidatorPunished)
				if err := _Validator.contract.UnpackLog(event, "ValidatorPunished", log); err != nil {
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

// ParseValidatorPunished is a log parse operation binding the contract event 0x54ce99c5ce1fc9f61656d4a0fb2697974d0c973ac32eecaefe06fcf18b8ef68a.
//
// Solidity: event ValidatorPunished(address indexed consensusAddr, uint256 indexed period, uint256 jailedUntil, uint256 deductedStakingAmount, bool blockProducerRewardDeprecated, bool bridgeOperatorRewardDeprecated)
func (_Validator *ValidatorFilterer) ParseValidatorPunished(log types.Log) (*ValidatorValidatorPunished, error) {
	event := new(ValidatorValidatorPunished)
	if err := _Validator.contract.UnpackLog(event, "ValidatorPunished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorValidatorSetUpdatedIterator is returned from FilterValidatorSetUpdated and is used to iterate over the raw logs and unpacked data for ValidatorSetUpdated events raised by the Validator contract.
type ValidatorValidatorSetUpdatedIterator struct {
	Event *ValidatorValidatorSetUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorValidatorSetUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorValidatorSetUpdated)
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
		it.Event = new(ValidatorValidatorSetUpdated)
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
func (it *ValidatorValidatorSetUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorValidatorSetUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorValidatorSetUpdated represents a ValidatorSetUpdated event raised by the Validator contract.
type ValidatorValidatorSetUpdated struct {
	Period         *big.Int
	ConsensusAddrs []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetUpdated is a free log retrieval operation binding the contract event 0x3d0eea40644a206ec25781dd5bb3b60eb4fa1264b993c3bddf3c73b14f29ef5e.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed period, address[] consensusAddrs)
func (_Validator *ValidatorFilterer) FilterValidatorSetUpdated(opts *bind.FilterOpts, period []*big.Int) (*ValidatorValidatorSetUpdatedIterator, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "ValidatorSetUpdated", periodRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorValidatorSetUpdatedIterator{contract: _Validator.contract, event: "ValidatorSetUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorSetUpdated is a free log subscription operation binding the contract event 0x3d0eea40644a206ec25781dd5bb3b60eb4fa1264b993c3bddf3c73b14f29ef5e.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed period, address[] consensusAddrs)
func (_Validator *ValidatorFilterer) WatchValidatorSetUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorValidatorSetUpdated, period []*big.Int) (event.Subscription, error) {

	var periodRule []interface{}
	for _, periodItem := range period {
		periodRule = append(periodRule, periodItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "ValidatorSetUpdated", periodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorValidatorSetUpdated)
				if err := _Validator.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
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

// ParseValidatorSetUpdated is a log parse operation binding the contract event 0x3d0eea40644a206ec25781dd5bb3b60eb4fa1264b993c3bddf3c73b14f29ef5e.
//
// Solidity: event ValidatorSetUpdated(uint256 indexed period, address[] consensusAddrs)
func (_Validator *ValidatorFilterer) ParseValidatorSetUpdated(log types.Log) (*ValidatorValidatorSetUpdated, error) {
	event := new(ValidatorValidatorSetUpdated)
	if err := _Validator.contract.UnpackLog(event, "ValidatorSetUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorValidatorUnjailedIterator is returned from FilterValidatorUnjailed and is used to iterate over the raw logs and unpacked data for ValidatorUnjailed events raised by the Validator contract.
type ValidatorValidatorUnjailedIterator struct {
	Event *ValidatorValidatorUnjailed // Event containing the contract specifics and raw log

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
func (it *ValidatorValidatorUnjailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorValidatorUnjailed)
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
		it.Event = new(ValidatorValidatorUnjailed)
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
func (it *ValidatorValidatorUnjailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorValidatorUnjailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorValidatorUnjailed represents a ValidatorUnjailed event raised by the Validator contract.
type ValidatorValidatorUnjailed struct {
	Validator common.Address
	Period    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorUnjailed is a free log retrieval operation binding the contract event 0x6bb2436cb6b6eb65d5a52fac2ae0373a77ade6661e523ef3004ee2d5524e6c6e.
//
// Solidity: event ValidatorUnjailed(address indexed validator, uint256 period)
func (_Validator *ValidatorFilterer) FilterValidatorUnjailed(opts *bind.FilterOpts, validator []common.Address) (*ValidatorValidatorUnjailedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "ValidatorUnjailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorValidatorUnjailedIterator{contract: _Validator.contract, event: "ValidatorUnjailed", logs: logs, sub: sub}, nil
}

// WatchValidatorUnjailed is a free log subscription operation binding the contract event 0x6bb2436cb6b6eb65d5a52fac2ae0373a77ade6661e523ef3004ee2d5524e6c6e.
//
// Solidity: event ValidatorUnjailed(address indexed validator, uint256 period)
func (_Validator *ValidatorFilterer) WatchValidatorUnjailed(opts *bind.WatchOpts, sink chan<- *ValidatorValidatorUnjailed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "ValidatorUnjailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorValidatorUnjailed)
				if err := _Validator.contract.UnpackLog(event, "ValidatorUnjailed", log); err != nil {
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

// ParseValidatorUnjailed is a log parse operation binding the contract event 0x6bb2436cb6b6eb65d5a52fac2ae0373a77ade6661e523ef3004ee2d5524e6c6e.
//
// Solidity: event ValidatorUnjailed(address indexed validator, uint256 period)
func (_Validator *ValidatorFilterer) ParseValidatorUnjailed(log types.Log) (*ValidatorValidatorUnjailed, error) {
	event := new(ValidatorValidatorUnjailed)
	if err := _Validator.contract.UnpackLog(event, "ValidatorUnjailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorWrappedUpEpochIterator is returned from FilterWrappedUpEpoch and is used to iterate over the raw logs and unpacked data for WrappedUpEpoch events raised by the Validator contract.
type ValidatorWrappedUpEpochIterator struct {
	Event *ValidatorWrappedUpEpoch // Event containing the contract specifics and raw log

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
func (it *ValidatorWrappedUpEpochIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorWrappedUpEpoch)
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
		it.Event = new(ValidatorWrappedUpEpoch)
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
func (it *ValidatorWrappedUpEpochIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorWrappedUpEpochIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorWrappedUpEpoch represents a WrappedUpEpoch event raised by the Validator contract.
type ValidatorWrappedUpEpoch struct {
	PeriodNumber *big.Int
	EpochNumber  *big.Int
	PeriodEnding bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWrappedUpEpoch is a free log retrieval operation binding the contract event 0x0195462033384fec211477c56217da64a58bd405e0bed331ba4ded67e4ae4ce7.
//
// Solidity: event WrappedUpEpoch(uint256 indexed periodNumber, uint256 indexed epochNumber, bool periodEnding)
func (_Validator *ValidatorFilterer) FilterWrappedUpEpoch(opts *bind.FilterOpts, periodNumber []*big.Int, epochNumber []*big.Int) (*ValidatorWrappedUpEpochIterator, error) {

	var periodNumberRule []interface{}
	for _, periodNumberItem := range periodNumber {
		periodNumberRule = append(periodNumberRule, periodNumberItem)
	}
	var epochNumberRule []interface{}
	for _, epochNumberItem := range epochNumber {
		epochNumberRule = append(epochNumberRule, epochNumberItem)
	}

	logs, sub, err := _Validator.contract.FilterLogs(opts, "WrappedUpEpoch", periodNumberRule, epochNumberRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorWrappedUpEpochIterator{contract: _Validator.contract, event: "WrappedUpEpoch", logs: logs, sub: sub}, nil
}

// WatchWrappedUpEpoch is a free log subscription operation binding the contract event 0x0195462033384fec211477c56217da64a58bd405e0bed331ba4ded67e4ae4ce7.
//
// Solidity: event WrappedUpEpoch(uint256 indexed periodNumber, uint256 indexed epochNumber, bool periodEnding)
func (_Validator *ValidatorFilterer) WatchWrappedUpEpoch(opts *bind.WatchOpts, sink chan<- *ValidatorWrappedUpEpoch, periodNumber []*big.Int, epochNumber []*big.Int) (event.Subscription, error) {

	var periodNumberRule []interface{}
	for _, periodNumberItem := range periodNumber {
		periodNumberRule = append(periodNumberRule, periodNumberItem)
	}
	var epochNumberRule []interface{}
	for _, epochNumberItem := range epochNumber {
		epochNumberRule = append(epochNumberRule, epochNumberItem)
	}

	logs, sub, err := _Validator.contract.WatchLogs(opts, "WrappedUpEpoch", periodNumberRule, epochNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorWrappedUpEpoch)
				if err := _Validator.contract.UnpackLog(event, "WrappedUpEpoch", log); err != nil {
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

// ParseWrappedUpEpoch is a log parse operation binding the contract event 0x0195462033384fec211477c56217da64a58bd405e0bed331ba4ded67e4ae4ce7.
//
// Solidity: event WrappedUpEpoch(uint256 indexed periodNumber, uint256 indexed epochNumber, bool periodEnding)
func (_Validator *ValidatorFilterer) ParseWrappedUpEpoch(log types.Log) (*ValidatorWrappedUpEpoch, error) {
	event := new(ValidatorWrappedUpEpoch)
	if err := _Validator.contract.UnpackLog(event, "WrappedUpEpoch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
