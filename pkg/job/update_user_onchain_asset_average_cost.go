package job

import (
	"encoding/json"
	"fmt"

	"github.com/defipod/mochi/pkg/entities"
	"github.com/defipod/mochi/pkg/logger"
	"github.com/defipod/mochi/pkg/model"
	"github.com/defipod/mochi/pkg/service"
	dadapter "github.com/defipod/mochi/pkg/service/dune/adapter"
	"github.com/defipod/mochi/pkg/service/sentrygo"
)

type updateUserOnchainAssetAvgCostJob struct {
	entity  *entities.Entity
	service *service.Service
	log     logger.Logger
}

func NewUpdateUserOnchainAssetAvgCostJob(e *entities.Entity) Job {
	return &updateUserOnchainAssetAvgCostJob{
		entity:  e,
		service: e.GetSvc(),
		log:     e.GetLogger(),
	}
}

func (j *updateUserOnchainAssetAvgCostJob) Run() error {
	j.log.Info("update_user_asset_average_cost job started")

	// 1.0 Get top 50 active evm addresses
	evmAddrs, err := j.getTop50ActiveEvmAddresses()
	if err != nil {
		j.log.Error(err, "failed to get top 50 active evm addresses")
		j.logSentry(err, map[string]interface{}{"task": "getTop50ActiveEvmAddresses"})
		return err
	}

	// 2.0 Call dune to get the average cost of the assets of each address
	avgCosts, err := j.getEvmAssetAvgCosts(evmAddrs)
	if err != nil {
		j.log.Error(err, "failed to get avg cost of evm assets")
		j.logSentry(err, map[string]interface{}{"task": "getEvmAssetAvgCosts"})
		return err
	}

	// 3.0 Update the average cost of the assets of each address in the database
	if err := j.updateEvmAssetAvgCosts(avgCosts); err != nil {
		j.log.Error(err, "failed to update avg cost of evm assets")
		j.logSentry(err, map[string]interface{}{"task": "updateEvmAssetAvgCosts"})
		return err
	}

	j.log.Info("update_user_asset_average_cost job finished")
	return nil
}

func (j *updateUserOnchainAssetAvgCostJob) getTop50ActiveEvmAddresses() ([]string, error) {
	profiles, err := j.service.MochiProfile.GetTopActiveUsers(50)
	if err != nil {
		return nil, err
	}
	evmAddrs := make([]string, 0)
	for _, p := range profiles {
		for _, a := range p.AssociatedAccounts {
			if a.Platform == "evm-chain" {
				evmAddrs = append(evmAddrs, a.PlatformIdentifier)
			}
		}
	}
	return evmAddrs, nil
}

func (j *updateUserOnchainAssetAvgCostJob) getEvmAssetAvgCosts(evmAddrs []string) ([]model.OnchainAssetAvgCost, error) {
	avgCosts := make([]model.OnchainAssetAvgCost, 0)

	// // Call dune to get the average cost of the assets of each address
	// j.log.Info("executing query to get average cost of assets of each address")
	// var queryId int64 = 3782999 // https://dune.com/queries/3782999
	// params := make(map[string]interface{})
	// params["wallet_addresses"] = strings.Join(evmAddrs, ",")
	// executeQueryResp, err := j.service.Dune.ExecuteQuery(queryId, map[string]interface{}{})
	// if err != nil {
	// 	return nil, err
	// }
	// executionId := executeQueryResp.ExecutionId
	executionId := "01HZ6JBH10MHN88JAK8H2Z9XF4"
	j.log.Infof("finished execute query, executeId: %s", executionId)

	// // Execute query, wait for the query execution to finish
	// j.log.Info("waiting for query execution to finish")
	// time.Sleep(30 * time.Second)

	// for {
	// 	execStatus, err := j.service.Dune.GetExecutionStatus(executionId)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to get execution status: %w", err)
	// 	}

	// 	// If the query execution is finished, break the loop
	// 	if execStatus.IsExecutionFinished {
	// 		break
	// 	}
	// 	// Otherwise wait for 30 seconds
	// 	time.Sleep(30 * time.Second)
	// }

	// j.log.Info("execution finished, start getting execution result")

	// Get the execution result

	jsonData := `
	{
		"execution_id": "01HZ6JBH10MHN88JAK8H2Z9XF4",
		"query_id": 3782999,
		"is_execution_finished": true,
		"state": "QUERY_STATE_COMPLETED",
		"submitted_at": "2024-05-31T05:35:55.424406Z",
		"expires_at": "2024-08-29T05:36:03.25455Z",
		"execution_started_at": "2024-05-31T05:35:55.885817Z",
		"execution_ended_at": "2024-05-31T05:36:03.254549Z",
		"result": {
			"rows": [
				{
					"average_cost": 0.21824163010383088,
					"blockchain": "fantom",
					"symbol": "2OMB",
					"token_address": "0x7a6e4e3cc2ac9924605dca4ba31d1831c84b44ae",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 820.4611928107397,
					"blockchain": "fantom",
					"symbol": "2SHARES",
					"token_address": "0xc54a1684fd1bef1f077a336e6be4bd9a3096a6ca",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 7758.496181495126,
					"blockchain": "fantom",
					"symbol": "3SHARES",
					"token_address": "0x6437adac543583c4b31bf0323a0870430f5cc2e7",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 976455363876369400,
					"blockchain": "base",
					"symbol": "APEIN",
					"token_address": "0x280df118566625cb6678c036d0f9027069b73f46",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.05422514982805414,
					"blockchain": "ethereum",
					"symbol": "BAKE",
					"token_address": "0x44face2e310e543f6d85867eb06fb251e3bfe1fc",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.3080540502626088,
					"blockchain": "bnb",
					"symbol": "BAMI",
					"token_address": "0x8249bc1dea00660d2d38df126b53c6c9a733e942",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 6.530611929298341,
					"blockchain": "fantom",
					"symbol": "BASED",
					"token_address": "0x8d7d3409881b51466b483b11ea1b8a03cded89ae",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 755.0939001279603,
					"blockchain": "fantom",
					"symbol": "BASED",
					"token_address": "0x23654048e2ee5e8414eea67a7d1a1c02505b7e4a",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0,
					"blockchain": "bnb",
					"symbol": "BBANK",
					"token_address": "0xf4b5470523ccd314c6b9da041076e7d79e0df267",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 1312.3897207616983,
					"blockchain": "fantom",
					"symbol": "BNB",
					"token_address": "0xd67de0e0a0fd7b15dc8348bb9be742f3c5850454",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 5.843186673414637,
					"blockchain": "fantom",
					"symbol": "BOO",
					"token_address": "0x6e0aa9718c56ef5d19ccf57955284c7cd95737be",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.0642170772,
					"blockchain": "arbitrum",
					"symbol": "BRC",
					"token_address": "0xb5de3f06af62d8428a8bf7b4400ea42ad2e0bc53",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 8995.721551559122,
					"blockchain": "fantom",
					"symbol": "BSHARE",
					"token_address": "0x49c290ff692149a4e16611c694fded42c954ab7a",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 55.8688778129335,
					"blockchain": "ethereum",
					"symbol": "BTC",
					"token_address": "0xbd6323a83b613f668687014e8a5852079494fb68",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 67567.38751929099,
					"blockchain": "bnb",
					"symbol": "BTCB",
					"token_address": "0x7130d2a12b9bcbfae4f2634d864a1ee1ce3ead9c",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 231.57928544251192,
					"blockchain": "ethereum",
					"symbol": "BTRFLY",
					"token_address": "0xc55126051b22ebb829d00368f4b12bde432de5da",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0,
					"blockchain": "fantom",
					"symbol": "BUSD",
					"token_address": "0xc931f61b1534eb21d8c11b24f3f5ab2471d4ab50",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.0002964401606414518,
					"blockchain": "fantom",
					"symbol": "BUTT",
					"token_address": "0xf42cc7284389fbf749590f26539002ca931323d0",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 1.0064166931581247,
					"blockchain": "fantom",
					"symbol": "DAI",
					"token_address": "0x8d11ec38a3eb5e956b052f67da8bdc9bef8abf3e",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0,
					"blockchain": "ethereum",
					"symbol": "DEFROGS",
					"token_address": "0xd555498a524612c67f286df0e0a9a64a73a7cdc7",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 1.2160268850194316,
					"blockchain": "bnb",
					"symbol": "DODO",
					"token_address": "0x67ee3cb086f8a16f34bee3ca72fad36f7db929e2",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.000018869991904493564,
					"blockchain": "ethereum",
					"symbol": "FADE",
					"token_address": "0xdb83becd16164b81b5eee04a10ae8537d1f327e5",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.00337718942294572,
					"blockchain": "fantom",
					"symbol": "FASTMOON",
					"token_address": "0x58fb8cbab7253d988bad2e7ca9079acc77ed430c",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 1.0054134868944322,
					"blockchain": "fantom",
					"symbol": "FRAX",
					"token_address": "0xdc301622e621166bd8e82f2ca0a26c13ad0be355",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.410139112800325,
					"blockchain": "base",
					"symbol": "FREN",
					"token_address": "0x12063cc18a7096d170e5fc410d8623ad97ee24b3",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 3.2696693955525644,
					"blockchain": "base",
					"symbol": "GB",
					"token_address": "0x2af864fb54b55900cd58d19c7102d9e4fa8d84a3",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 10.78939125170614,
					"blockchain": "fantom",
					"symbol": "GEIST",
					"token_address": "0xd8321aa83fb0a4ecd6348d4577431310a6e0814d",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.20746813143326165,
					"blockchain": "arbitrum",
					"symbol": "GETH",
					"token_address": "0xdd69db25f6d620a7bad3023c5d32761d353d3de9",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 40.648220488467864,
					"blockchain": "arbitrum",
					"symbol": "GMX",
					"token_address": "0xfc5a1a6eb076a2c7ad06ed22c90d7e710e35ad0a",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0,
					"blockchain": "bnb",
					"symbol": "HDODO",
					"token_address": "0xfed2e6a6105e48a781d0808e69460bd5ba32d3d3",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.017308345046701878,
					"blockchain": "fantom",
					"symbol": "HIGH",
					"token_address": "0x33fc7706f98d804a1aac4d16a322bebc98276b24",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0,
					"blockchain": "fantom",
					"symbol": "ICE",
					"token_address": "0xf16e81dce15b08f326220742020379b855b87df9",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0,
					"blockchain": "arbitrum",
					"symbol": "JOE",
					"token_address": "0x371c7ec6d8039ff7933a2aa28eb827ffe1f52f07",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 1.0625619262810195,
					"blockchain": "arbitrum",
					"symbol": "JONES",
					"token_address": "0x10393c20975cf177a3513071bc110f7962cd67da",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 7.807375218324466,
					"blockchain": "arbitrum",
					"symbol": "LINK",
					"token_address": "0xf97f4df75117a78c1a5a0dbb814af92458539fb4",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 48.69779364417139,
					"blockchain": "fantom",
					"symbol": "LINK",
					"token_address": "0xb3654dc3d10ea7645f8319668e8f54d2574fbdc8",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 9.036921147370382,
					"blockchain": "arbitrum",
					"symbol": "LPT",
					"token_address": "0x289ba1701c2f088cf0faf8b3705246331cb8a839",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.27884513603584815,
					"blockchain": "fantom",
					"symbol": "MCLB",
					"token_address": "0x5deb27e51dbeef691ba1175a2e563870499c2acb",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 1.000720954922238,
					"blockchain": "fantom",
					"symbol": "MIM",
					"token_address": "0x82f0b8b456c1a451378467398982d4834b6829c1",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.0860982835740171,
					"blockchain": "arbitrum",
					"symbol": "MYC",
					"token_address": "0xc74fe4c715510ec2f8c61d70d397b32043f55abe",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 1139.481539755199,
					"blockchain": "ethereum",
					"symbol": "NO",
					"token_address": "0xe99620545740a30dbf8b4601d2f7ddc99cc1dc42",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 24.52745246700746,
					"blockchain": "base",
					"symbol": "OBTC",
					"token_address": "0x54dc305fd3542d402bafdcde0574eb9a0a893408",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 4.837666941998499,
					"blockchain": "ethereum",
					"symbol": "POLS",
					"token_address": "0x83e6f1e41cdd28eaceb20cb649155049fac3d5aa",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0,
					"blockchain": "fantom",
					"symbol": "RAINSPIRIT",
					"token_address": "0xf9c6e3c123f0494a4447100bd7dbd536f43cc33a",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.0095986190345319,
					"blockchain": "arbitrum",
					"symbol": "SPA",
					"token_address": "0x5575552988a3a80504bbaeb1311674fcfd40ad4b",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.003714803238231209,
					"blockchain": "fantom",
					"symbol": "SPELL",
					"token_address": "0x468003b688943977e6130f4f68f23aad939a1040",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.614652165032766,
					"blockchain": "arbitrum",
					"symbol": "STG",
					"token_address": "0x6694340fc020c5e6b96567843da2df01b2ce1eb6",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 2.504164366711412,
					"blockchain": "ethereum",
					"symbol": "SUPER",
					"token_address": "0xe53ec727dbdeb9e2d5456c3be40cff031ab40a55",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 1.370677104484951,
					"blockchain": "arbitrum",
					"symbol": "SYN",
					"token_address": "0x080f6aed32fc474dd5717105dba5ea57268f46eb",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.04226967867502745,
					"blockchain": "fantom",
					"symbol": "TBOND",
					"token_address": "0x24248cd1747348bdc971a5395f4b3cd7fee94ea0",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0,
					"blockchain": "polygon",
					"symbol": "TITAN",
					"token_address": "0xaaa5b9e6c589642f98a1cda99b9d024b8407285a",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.023017518032397873,
					"blockchain": "polygon",
					"symbol": "TRADE",
					"token_address": "0x692ac1e363ae34b6b489148152b12e2785a3d8d6",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 3.9534659803566486e-11,
					"blockchain": "bnb",
					"symbol": "TWOGE",
					"token_address": "0xd5ffab1841b9137d5528ed09d6ebb66c3088aede",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 17.275045648466335,
					"blockchain": "arbitrum",
					"symbol": "UMAMI",
					"token_address": "0x1622bf67e6e5747b81866fe0b85178a93c7f86e3",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 1.0002744913887442,
					"blockchain": "arbitrum",
					"symbol": "USDC",
					"token_address": "0xaf88d065e77c8cc2239327c5edb3a432268e5831",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.9956680875711911,
					"blockchain": "avalanche_c",
					"symbol": "USDC",
					"token_address": "0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 1.000431773938665,
					"blockchain": "base",
					"symbol": "USDC",
					"token_address": "0x833589fcd6edb6e08f4c7c32d4f71b54bda02913",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 1.0010235434514274,
					"blockchain": "fantom",
					"symbol": "USDC",
					"token_address": "0x04068da6c83afcfa0e13ba15a6696662335d5b75",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 1.0031352363125876,
					"blockchain": "fantom",
					"symbol": "USDC",
					"token_address": "0x28a92dde19d9989f39a49905d7c9c2fac7799bdf",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.9508950543050174,
					"blockchain": "optimism",
					"symbol": "USDC",
					"token_address": "0x7f5c764cbc14f9669b88837ca1490cca17c31607",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 1.0021272569709632,
					"blockchain": "polygon",
					"symbol": "USDC",
					"token_address": "0x2791bca1f2de4661ed88a30c99a7a9449aa84174",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 1.0058689261073228,
					"blockchain": "avalanche_c",
					"symbol": "USDC.e",
					"token_address": "0xa7d7079b0fead91f3e65f86e8915cb59c1a4c664",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 1.0003275292976863,
					"blockchain": "polygon",
					"symbol": "USDT",
					"token_address": "0xc2132d05d31c914a87c6611c10748aeb04b58e8f",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 18.479051295736518,
					"blockchain": "avalanche_c",
					"symbol": "USDT.e",
					"token_address": "0xc7198437980c041c805a1edcba50c1ce5db95118",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 1.0019015276217471,
					"blockchain": "base",
					"symbol": "USDbC",
					"token_address": "0xd9aaec86b65d86f6a7b5b1b0c42ffa531710b6ca",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 9.762923082499752,
					"blockchain": "avalanche_c",
					"symbol": "USDt",
					"token_address": "0x9702230a8ea53601f5cd2dc00fdbc13d4df4a8c7",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 77165949342883,
					"blockchain": "bnb",
					"symbol": "UST",
					"token_address": "0x23396cf899ca06c4472205fc903bdb4de249d6fc",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.4355928658612501,
					"blockchain": "arbitrum",
					"symbol": "VSTA",
					"token_address": "0xa684cd057951541187f288294a1e1c2646aa2d24",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.3992321323887188,
					"blockchain": "bnb",
					"symbol": "WATCH",
					"token_address": "0x7a9f28eb62c791422aa23ceae1da9c847cbec9b0",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 242.7881078032517,
					"blockchain": "avalanche_c",
					"symbol": "WAVAX",
					"token_address": "0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 63965.35,
					"blockchain": "ethereum",
					"symbol": "WBTC",
					"token_address": "0x2260fac5e5542a773aa44fbcfedf7c193bc2c599",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 1690.1975107848525,
					"blockchain": "arbitrum",
					"symbol": "WETH",
					"token_address": "0x82af49447d8a07e3bd95bd0d56f35241523fbab1",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 3794.52,
					"blockchain": "zora",
					"symbol": "WETH",
					"token_address": "0x4200000000000000000000000000000000000006",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0,
					"blockchain": "avalanche_c",
					"symbol": "WETH.e",
					"token_address": "0x49d5c2bdffac6ce2bfdb6640f4f80f226bc10bab",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.4912039434212945,
					"blockchain": "fantom",
					"symbol": "WFTM",
					"token_address": "0x21be370d5312f44cb42ce377bc9b8a0cef1a4c83",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.008097590266459969,
					"blockchain": "fantom",
					"symbol": "fBOMB",
					"token_address": "0x74ccbe53f77b08632ce0cb91d3a545bf6b8e0979",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0,
					"blockchain": "fantom",
					"symbol": "fUSDT",
					"token_address": "0x049d68029688eabf473097a2fc38ef61633a3c7a",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 0.999250312278662,
					"blockchain": "fantom",
					"symbol": "miMATIC",
					"token_address": "0xfb98b335551a418cd0737375a2ea0ded62ea213b",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 16.616317592019083,
					"blockchain": "fantom",
					"symbol": "pFTM",
					"token_address": "0x112df7e3b4b7ab424f07319d4e92f41e6608c48b",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 23398.728220943845,
					"blockchain": "optimism",
					"symbol": "sETH",
					"token_address": "0xe405de8f52ba7559f9df3c368500b6e6ae6cee49",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 2.54661772235859,
					"blockchain": "optimism",
					"symbol": "sUSD",
					"token_address": "0x8c6f28f2f1a3c87f0f938b96d27520d9751ec8d9",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 28626.433862115406,
					"blockchain": "avalanche_c",
					"symbol": "wMEMO",
					"token_address": "0x0da67235dd5787d67955420c84ca1cecd4e5bb3b",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 19.898608784555574,
					"blockchain": "ethereum",
					"symbol": "xSUSHI",
					"token_address": "0x8798249c2e607446efb7ad49ec89dd1865ff4272",
					"wallet_address": "0x5417a03667abb6a059b3f174c1f67b1e83753046"
				},
				{
					"average_cost": 1.0266160801053446,
					"blockchain": "ethereum",
					"symbol": "0xBET",
					"token_address": "0x78993f9bee8b68f2531a92427595405f294161db",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 0.00654309215458559,
					"blockchain": "ethereum",
					"symbol": "AIL",
					"token_address": "0xd155fa55c40d010335aa152891aa687e2f3090bd",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 0,
					"blockchain": "ethereum",
					"symbol": "AITAX",
					"token_address": "0x9f04c2bd696a6191246144ba762456a24c457520",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 0.19013684546063597,
					"blockchain": "ethereum",
					"symbol": "ATH",
					"token_address": "0xbe7458bc543cf2df43ac109d2f713dffe6417aa4",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 1.531366202633671,
					"blockchain": "ethereum",
					"symbol": "BAI",
					"token_address": "0x36c79f0b8a2e8a3c0230c254c452973e7a3ba155",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 0.0000069279994315155175,
					"blockchain": "base",
					"symbol": "BIDEN",
					"token_address": "0xa82a1f2e7517ca8115ebe3f56a1f12e9b941613b",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 164.42784744821847,
					"blockchain": "base",
					"symbol": "BLARB",
					"token_address": "0x0d30be9d9c2cf90aeff4fef5b2e8c3d0b02596a0",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 67.55242775989778,
					"blockchain": "ethereum",
					"symbol": "BTC",
					"token_address": "0xbd6323a83b613f668687014e8a5852079494fb68",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 248.07747716204204,
					"blockchain": "ethereum",
					"symbol": "BTRFLY",
					"token_address": "0xc55126051b22ebb829d00368f4b12bde432de5da",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 0.000025167633607094852,
					"blockchain": "base",
					"symbol": "BUILD",
					"token_address": "0x3c281a39944a2319aa653d81cfd93ca10983d234",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 1.001768,
					"blockchain": "bnb",
					"symbol": "BUSD",
					"token_address": "0xe9e7cea3dedca5984780bafc599bd69add087d56",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 0.000005306576135318418,
					"blockchain": "ethereum",
					"symbol": "CA",
					"token_address": "0xa0c7e61ee4faa9fcefdc8e8fc5697d54bf8c8141",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 1.98650144073974,
					"blockchain": "ethereum",
					"symbol": "CHAT",
					"token_address": "0xbb3d7f42c58abd83616ad7c8c72473ee46df2678",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 0.050822742107023686,
					"blockchain": "base",
					"symbol": "DEFIDO",
					"token_address": "0xd064c53f043d5aee2ac9503b13ee012bf2def1d0",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 1.2284896516462591,
					"blockchain": "ethereum",
					"symbol": "DESK",
					"token_address": "0xe00924736426393035b22770d94a188f25fc3b16",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 0.02338694876958988,
					"blockchain": "ethereum",
					"symbol": "DEUS",
					"token_address": "0xba7d732929b5dce3a12ed6642a711f37630adf57",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 971.16698364762,
					"blockchain": "arbitrum",
					"symbol": "DPX",
					"token_address": "0x6c2c06790b3e3e3c38e12ee22f8183b37a13ee55",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 2357.86,
					"blockchain": "bnb",
					"symbol": "ETH",
					"token_address": "0x2170ed0880ac9a755fd29b2688956bd959f933f8",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 0.5401590104051568,
					"blockchain": "ethereum",
					"symbol": "FR33",
					"token_address": "0x0caa29cc8fc91f09c3d0510a941a9c11e31f180b",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 0,
					"blockchain": "ethereum",
					"symbol": "GPU",
					"token_address": "0x1258d60b224c0c5cd888d37bbf31aa5fcfb7e870",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 0.026466550390227994,
					"blockchain": "ethereum",
					"symbol": "IPT",
					"token_address": "0x9244b75890af6a37704a750d41227c504138b1fb",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 0.12866361086326492,
					"blockchain": "ethereum",
					"symbol": "JUDGE",
					"token_address": "0x01202c9a1adfc1475c960c23bdf7530698330fa0",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 0.5956744193245138,
					"blockchain": "ethereum",
					"symbol": "NGL",
					"token_address": "0x12652c6d93fdb6f4f37d48a8687783c782bb0d10",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 8.110246066730593,
					"blockchain": "base",
					"symbol": "NORMIE",
					"token_address": "0x7f12d13b34f5f4f0a9449c16bcd42f0da47af200",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 1.7721497021758854,
					"blockchain": "ethereum",
					"symbol": "QUEST",
					"token_address": "0x9b8f1cfc6ad3459c49ca61296627911e8c4c431d",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 0.00839709780664303,
					"blockchain": "ethereum",
					"symbol": "RSR",
					"token_address": "0x320623b8e4ff03373931769a31fc52a4e78b5d70",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 0,
					"blockchain": "ethereum",
					"symbol": "TATSU",
					"token_address": "0x92f419fb7a750aed295b0ddf536276bf5a40124f",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 0,
					"blockchain": "ethereum",
					"symbol": "TRUMP",
					"token_address": "0x576e2bed8f7b46d34016198911cdf9886f78bea7",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 1.4042121466536308,
					"blockchain": "arbitrum",
					"symbol": "USDC",
					"token_address": "0xaf88d065e77c8cc2239327c5edb3a432268e5831",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 1.000052,
					"blockchain": "polygon",
					"symbol": "USDC",
					"token_address": "0x2791bca1f2de4661ed88a30c99a7a9449aa84174",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 0.9588890191734135,
					"blockchain": "polygon",
					"symbol": "USDC",
					"token_address": "0x3c499c542cef5e3811e1192ce70d8cc03d5c3359",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 1.00191,
					"blockchain": "bnb",
					"symbol": "USDT",
					"token_address": "0x55d398326f99059ff775485246999027b3197955",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 0.45053651656007465,
					"blockchain": "ethereum",
					"symbol": "VIRTU",
					"token_address": "0x102dc1840f0c3c179670f21fa63597e82df34e60",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 10327.317321948814,
					"blockchain": "arbitrum",
					"symbol": "WETH",
					"token_address": "0x82af49447d8a07e3bd95bd0d56f35241523fbab1",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 3747.7899999999995,
					"blockchain": "polygon",
					"symbol": "WETH",
					"token_address": "0x7ceb23fd6bc0add59e62ac25578270cff1b9f619",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 65.18790621906469,
					"blockchain": "ethereum",
					"symbol": "WMINIMA",
					"token_address": "0x669c01caf0edcad7c2b8dc771474ad937a7ca4af",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 3754.01,
					"blockchain": "ethereum",
					"symbol": "stETH",
					"token_address": "0xae7ab96520de3a18e5e111b5eaab095312d7fe84",
					"wallet_address": "0xce1974637f19bfb8ff2ebf3f4c891612e9f61c9e"
				},
				{
					"average_cost": 0.09009309859368093,
					"blockchain": "base",
					"symbol": "$mfer",
					"token_address": "0xe3086852a4b125803c815a158249ae468a3254ca",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 3.0995330113591143e-9,
					"blockchain": "bnb",
					"symbol": "ANONGATE",
					"token_address": "0xd7efe3fbeaa68b317bc72937f52dcbe5bb67f5fc",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 443.08208905783835,
					"blockchain": "base",
					"symbol": "APEIN",
					"token_address": "0x280df118566625cb6678c036d0f9027069b73f46",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0,
					"blockchain": "bnb",
					"symbol": "AQUAGOAT",
					"token_address": "0x07af67b392b7a202fad8e0fbc64c34f33102165b",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0.8217836963390361,
					"blockchain": "fantom",
					"symbol": "BEETS",
					"token_address": "0xf24bcf4d1e507740041c9cfd2dddb29585adce1e",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 363.1869295565394,
					"blockchain": "ethereum",
					"symbol": "BTRFLY",
					"token_address": "0xc55126051b22ebb829d00368f4b12bde432de5da",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0.00029459387697123126,
					"blockchain": "fantom",
					"symbol": "BUTT",
					"token_address": "0xf42cc7284389fbf749590f26539002ca931323d0",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0,
					"blockchain": "fantom",
					"symbol": "CHILL",
					"token_address": "0xe47d957f83f8887063150aaf7187411351643392",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 3.41058912801327e-9,
					"blockchain": "fantom",
					"symbol": "CONK",
					"token_address": "0xb715f8dce2f0e9b894c753711bd55ee3c04dca4e",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 18.22249762150428,
					"blockchain": "bnb",
					"symbol": "Cake",
					"token_address": "0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0.2854190529605514,
					"blockchain": "bnb",
					"symbol": "DBZ",
					"token_address": "0x7a983559e130723b70e45bd637773dbdfd3f71db",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 303.8253165104987,
					"blockchain": "ethereum",
					"symbol": "DEFROGS",
					"token_address": "0xd555498a524612c67f286df0e0a9a64a73a7cdc7",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0.01924725469264706,
					"blockchain": "base",
					"symbol": "DEGEN",
					"token_address": "0x4ed4e862860bed51a9570b96d89af5e1b0efefed",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0.000024746553294649436,
					"blockchain": "ethereum",
					"symbol": "FADE",
					"token_address": "0xdb83becd16164b81b5eee04a10ae8537d1f327e5",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0.06472655893586712,
					"blockchain": "fantom",
					"symbol": "FOO",
					"token_address": "0xfbc3c04845162f067a0b6f8934383e63899c3524",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0.00006012190634214629,
					"blockchain": "ethereum",
					"symbol": "FOUR",
					"token_address": "0x244b797d622d4dee8b188b03546acaabd0cf91a0",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 61.5838070876973,
					"blockchain": "fantom",
					"symbol": "FOXY",
					"token_address": "0x11cc291fddcd17a969bc4608618601bfe5af13bd",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0.00000521,
					"blockchain": "base",
					"symbol": "FRAME",
					"token_address": "0x91f45aa2bde7393e0af1cc674ffe75d746b93567",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0.1322315310253598,
					"blockchain": "fantom",
					"symbol": "FS",
					"token_address": "0xc758295cd1a564cdb020a78a681a838cf8e0627d",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0.8079869915891194,
					"blockchain": "fantom",
					"symbol": "FUSD",
					"token_address": "0xad84341756bf337f5a0164515b1f6f993d194e1f",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0.0007504291814284083,
					"blockchain": "ethereum",
					"symbol": "JPEG",
					"token_address": "0xe80c0cd204d654cebe8dd64a4857cab6be8345a3",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0,
					"blockchain": "fantom",
					"symbol": "JUST",
					"token_address": "0x37c045be4641328dfeb625f1dde610d061613497",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 4.291150285537211,
					"blockchain": "fantom",
					"symbol": "KINS",
					"token_address": "0x471d83bdb3fbc1b67dbb9f6758feeab56a20268f",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 2.949727050994211e-16,
					"blockchain": "bnb",
					"symbol": "LTRBT",
					"token_address": "0x17d749d3e2ac204a07e19d8096d9a05c423ea3af",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0.9971590000000001,
					"blockchain": "avalanche_c",
					"symbol": "MIM",
					"token_address": "0x130966628846bfd36ff31a822705796e8cb8c18d",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 1.0602609371127576,
					"blockchain": "fantom",
					"symbol": "MIM",
					"token_address": "0x82f0b8b456c1a451378467398982d4834b6829c1",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0,
					"blockchain": "fantom",
					"symbol": "NIPS",
					"token_address": "0x667afbb7d558c3dfd20fabd295d31221dab9dbc2",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0,
					"blockchain": "fantom",
					"symbol": "ONYX",
					"token_address": "0x57b3c627a2e9688e788108df4b56f3d09c46d16e",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 5.616274386686069e-7,
					"blockchain": "fantom",
					"symbol": "PPR",
					"token_address": "0x1d09150dc45b10fde93da0ccb2d0342830a78aa3",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0.00024345844707235502,
					"blockchain": "ethereum",
					"symbol": "RUME",
					"token_address": "0x431e365aa04e8d098c563f6f2c4b572bf5cce357",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 2.3635003397489193e-8,
					"blockchain": "ethereum",
					"symbol": "SPKI",
					"token_address": "0x0f3debf94483beecbfd20167c946a61ea62d000f",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 9.415203152501646e-10,
					"blockchain": "bnb",
					"symbol": "SPORE",
					"token_address": "0x33a3d962955a3862c8093d1273344719f03ca17c",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0,
					"blockchain": "fantom",
					"symbol": "SUN",
					"token_address": "0x60e91f89a2986975822de3bfe50df002ef46eaad",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 6.672292092295325,
					"blockchain": "fantom",
					"symbol": "SUSHI",
					"token_address": "0xae75a438b2e0cb8bb01ec1e1e376de11d44477cc",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0.0005306545273743223,
					"blockchain": "base",
					"symbol": "Sendit",
					"token_address": "0xba5b9b2d2d06a9021eb3190ea5fb0e02160839a4",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0,
					"blockchain": "fantom",
					"symbol": "TCS",
					"token_address": "0xfbfae0dd49882e503982f8eb4b8b1e464eca0b91",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 17538.326909527386,
					"blockchain": "avalanche_c",
					"symbol": "TIME",
					"token_address": "0xb54f16fb19478766a268f172c9480f8da1a7c9c3",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0.001859565275731802,
					"blockchain": "base",
					"symbol": "TN100x",
					"token_address": "0x5b5dee44552546ecea05edea01dcd7be7aa6144a",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0.0011457652560454628,
					"blockchain": "ethereum",
					"symbol": "UNLEASH",
					"token_address": "0x0e9cc0f7e550bd43bd2af2214563c47699f96479",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 1.0002620399194186,
					"blockchain": "arbitrum",
					"symbol": "USDC",
					"token_address": "0xaf88d065e77c8cc2239327c5edb3a432268e5831",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 1.0009625048892632,
					"blockchain": "base",
					"symbol": "USDC",
					"token_address": "0x833589fcd6edb6e08f4c7c32d4f71b54bda02913",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0.9265925540737833,
					"blockchain": "optimism",
					"symbol": "USDC",
					"token_address": "0x7f5c764cbc14f9669b88837ca1490cca17c31607",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 1.001372905914898,
					"blockchain": "polygon",
					"symbol": "USDC",
					"token_address": "0x3c499c542cef5e3811e1192ce70d8cc03d5c3359",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 1.0006678372240874,
					"blockchain": "arbitrum",
					"symbol": "USDT",
					"token_address": "0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 1.000618,
					"blockchain": "linea",
					"symbol": "USDT",
					"token_address": "0xa219439258ca9da29e9cc4ce5596924745e12b93",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 1.0012720616705977,
					"blockchain": "optimism",
					"symbol": "USDT",
					"token_address": "0x94b008aa00579c1307b0ef2c499ad98a8ce58e58",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 1.003630323109184,
					"blockchain": "ethereum",
					"symbol": "USDe",
					"token_address": "0x4c9edd5852cd905f086c759e8383e09bff1e68b3",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 2.8209927136762167e-8,
					"blockchain": "fantom",
					"symbol": "Voodoo",
					"token_address": "0x95ce7b991cfc7e3ad8466ac20746b9bed7713b0a",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 975.3893801120616,
					"blockchain": "bnb",
					"symbol": "WBNB",
					"token_address": "0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 34366.25,
					"blockchain": "arbitrum",
					"symbol": "WBTC",
					"token_address": "0x2f2a2543b76a4166549f7aab2e75bef0aefc5b0f",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 2444.420990698853,
					"blockchain": "arbitrum",
					"symbol": "WETH",
					"token_address": "0x82af49447d8a07e3bd95bd0d56f35241523fbab1",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 4581.493919697668,
					"blockchain": "ethereum",
					"symbol": "WETH",
					"token_address": "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 2.5411525544823927,
					"blockchain": "fantom",
					"symbol": "WFTM",
					"token_address": "0x21be370d5312f44cb42ce377bc9b8a0cef1a4c83",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 135.34820521177357,
					"blockchain": "fantom",
					"symbol": "ZON",
					"token_address": "0xbf1f3d7266a1080bf448e428daa37eec6b05a8ed",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0.035346424435378755,
					"blockchain": "fantom",
					"symbol": "ZOO",
					"token_address": "0x09e145a1d53c0045f41aeef25d8ff982ae74dd56",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0,
					"blockchain": "fantom",
					"symbol": "bb-yv-FTM",
					"token_address": "0xc3bf643799237588b7a6b407b3fc028dd4e037d2",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0,
					"blockchain": "fantom",
					"symbol": "bb-yv-USD",
					"token_address": "0x5ddb92a5340fd0ead3987d3661afcd6104c3b757",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 1.0502478847633163,
					"blockchain": "fantom",
					"symbol": "fUSDT",
					"token_address": "0x049d68029688eabf473097a2fc38ef61633a3c7a",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0.00043619170478691996,
					"blockchain": "base",
					"symbol": "member",
					"token_address": "0x7d89e05c0b93b24b5cb23a073e60d008fed1acf9",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				},
				{
					"average_cost": 0.7415189533557974,
					"blockchain": "fantom",
					"symbol": "sGOAT",
					"token_address": "0x43f9a13675e352154f745d6402e853fecc388aa5",
					"wallet_address": "0xf3940eb383853d364f82b76f5d513fef73b06e0f"
				}
			],
			"metadata": {
				"column_names": [
					"wallet_address",
					"symbol",
					"blockchain",
					"token_address",
					"average_cost"
				],
				"column_types": [
					"varbinary",
					"varchar",
					"varchar",
					"varbinary",
					"double"
				],
				"row_count": 180,
				"result_set_bytes": 22479,
				"total_row_count": 180,
				"total_result_set_bytes": 18535,
				"datapoint_count": 900,
				"pending_time_millis": 461,
				"execution_time_millis": 7368
			}
		}
	}
	`
	res := dadapter.GetExecutionResultResponse{}
	err := json.Unmarshal([]byte(jsonData), &res)
	if err != nil {
		return nil, err
	}
	for _, row := range res.Result.Rows {
		// Parse the data
		walletAddr, wOk := row["wallet_address"]
		blockchain, bOk := row["blockchain"]
		tokenAddress, tOk := row["token_address"]
		symbol, sOk := row["symbol"]
		averageCost, cOk := row["average_cost"]
		if !wOk || !bOk || !tOk || !sOk || !cOk {
			continue
		}

		// Convert the data to the correct type
		walletAddrStr, wOk := walletAddr.(string)
		blockchainStr, bOk := blockchain.(string)
		tokenAddressStr, tOk := tokenAddress.(string)
		symbolStr, sOk := symbol.(string)
		averageCostFloat, cOk := averageCost.(float64)
		if !wOk || !bOk || !tOk || !sOk || !cOk {
			continue
		}

		avgCosts = append(avgCosts, model.OnchainAssetAvgCost{
			WalletAddress: walletAddrStr,
			Blockchain:    blockchainStr,
			TokenAddress:  tokenAddressStr,
			Symbol:        symbolStr,
			AverageCost:   averageCostFloat,
		})
	}

	// var (
	// 	limit  int64 = 500
	// 	offset int64 = 0
	// )
	// for {
	// 	res, err := j.service.Dune.GetExecutionResult(executionId, limit, offset)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to get execution result: %w", err)
	// 	}
	// 	for _, row := range res.Result.Rows {
	// 		// Parse the data
	// 		walletAddr, wOk := row["wallet_address"]
	// 		blockchain, bOk := row["blockchain"]
	// 		tokenAddress, tOk := row["token_address"]
	// 		symbol, sOk := row["symbol"]
	// 		averageCost, cOk := row["average_cost"]
	// 		if !wOk || !bOk || !tOk || !sOk || !cOk {
	// 			continue
	// 		}

	// 		// Convert the data to the correct type
	// 		walletAddrStr, wOk := walletAddr.(string)
	// 		blockchainStr, bOk := blockchain.(string)
	// 		tokenAddressStr, tOk := tokenAddress.(string)
	// 		symbolStr, sOk := symbol.(string)
	// 		averageCostFloat, cOk := averageCost.(float64)
	// 		if !wOk || !bOk || !tOk || !sOk || !cOk {
	// 			continue
	// 		}

	// 		avgCosts = append(avgCosts, model.OnchainAssetAvgCost{
	// 			WalletAddress: walletAddrStr,
	// 			Blockchain:    blockchainStr,
	// 			TokenAddress:  tokenAddressStr,
	// 			Symbol:        symbolStr,
	// 			AverageCost:   averageCostFloat,
	// 		})
	// 	}

	// 	// If there is no more data, break the loop
	// 	if res.NextOffset == nil || *res.NextOffset == 0 {
	// 		break
	// 	}
	// 	// Otherwise update the offset
	// 	offset = *res.NextOffset
	// }

	j.log.Info("finished getting execution result")
	return avgCosts, nil
}

func (j *updateUserOnchainAssetAvgCostJob) updateEvmAssetAvgCosts(assets []model.OnchainAssetAvgCost) error {
	if err := j.entity.UpsertManyOnchainAssetAvgCost(assets); err != nil {
		return err
	}

	return nil
}

func (j *updateUserOnchainAssetAvgCostJob) logSentry(err error, extra map[string]interface{}) {
	sentryTags := map[string]string{
		"type": "system",
	}
	j.service.Sentry.CaptureErrorEvent(sentrygo.SentryCapturePayload{
		Message: fmt.Sprintf("[CJ prod mochi] - update_user_asset_average_cost failed - %v", err),
		Tags:    sentryTags,
		Extra:   extra,
	})
}
