package ronin

type Service interface {
	GetAxsStakingAmount(address string) (float64, error)
	GetAxsPendingRewards(address string) (float64, error)
	GetRonStakingAmount(address string) (float64, error)
	GetRonPendingRewards(address string) (float64, error)
	GetLpPendingRewards(address string) (map[string]LpRewardData, error)
}
