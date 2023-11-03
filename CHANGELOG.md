# [6.19.0](https://github.com/consolelabs/mochi-api/compare/v6.18.0...v6.19.0) (2023-11-03)


### Bug Fixes

* cannot bind uri remove wl ([#1221](https://github.com/consolelabs/mochi-api/issues/1221)) ([74c91e5](https://github.com/consolelabs/mochi-api/commit/74c91e5b20f338a1eaf39ba43035d10cf737d3c1))
* friend tech key api - add client id ([1be14b5](https://github.com/consolelabs/mochi-api/commit/1be14b5c24df3435807402fa1813ec8f4dbdc8ec))
* query ticker rune ([#1220](https://github.com/consolelabs/mochi-api/issues/1220)) ([6fd19f8](https://github.com/consolelabs/mochi-api/commit/6fd19f89fbfea2bd5267ec95e4d1cec0f110677b))
* transfer-v2 auth ([f055d94](https://github.com/consolelabs/mochi-api/commit/f055d9463d9d8c742dbd8c372bc93c6a98f95295))
* transfer-v2 authorization ([f3b4f53](https://github.com/consolelabs/mochi-api/commit/f3b4f533861141848705140d5b0bd40a19cee2ad))


### Features

* product themes ([#1218](https://github.com/consolelabs/mochi-api/issues/1218)) ([31b858b](https://github.com/consolelabs/mochi-api/commit/31b858b4ebcbbf76a2478cbab24c3432c767aa7a))

# [6.18.0](https://github.com/consolelabs/mochi-api/compare/v6.17.3...v6.18.0) (2023-10-11)


### Bug Fixes

* wrong format invalid emojis ([57cfbf6](https://github.com/consolelabs/mochi-api/commit/57cfbf6cc3294ba5498a61297bf80041f440b6df))


### Features

* integrate chain zkSync ([a768890](https://github.com/consolelabs/mochi-api/commit/a768890fe5b980dd6cb0a9405f376251b4054e11))
* tono command permission manage ([#1214](https://github.com/consolelabs/mochi-api/issues/1214)) ([eed50a3](https://github.com/consolelabs/mochi-api/commit/eed50a3f395eb825d1ab3cd199fc0541cabac4f0))

## [6.17.3](https://github.com/consolelabs/mochi-api/compare/v6.17.2...v6.17.3) (2023-10-09)


### Bug Fixes

* allow all origin cors ([#1211](https://github.com/consolelabs/mochi-api/issues/1211)) ([3005165](https://github.com/consolelabs/mochi-api/commit/30051656d76b2e0ce4c82a45d32b53403204bf7f))
* allow request id header cors ([#1210](https://github.com/consolelabs/mochi-api/issues/1210)) ([c0015bd](https://github.com/consolelabs/mochi-api/commit/c0015bd7a17e69797167153c59f4a82ef87ad93c))
* handle error when get emojis ([5c713e6](https://github.com/consolelabs/mochi-api/commit/5c713e6863bb508b3f5c2dc9010f067b518f61fa))

## [6.17.2](https://github.com/consolelabs/mochi-api/compare/v6.17.1...v6.17.2) (2023-10-02)


### Bug Fixes

* format message in validate emoji ([244f556](https://github.com/consolelabs/mochi-api/commit/244f55643b651348568c541d7d4b0f3684ed6793))

## [6.17.1](https://github.com/consolelabs/mochi-api/compare/v6.17.0...v6.17.1) (2023-10-02)


### Bug Fixes

* guild id list to validate emoji ([fee5ae7](https://github.com/consolelabs/mochi-api/commit/fee5ae7ef080f8d3492dea44968dd26c3ba2541b))

# [6.17.0](https://github.com/consolelabs/mochi-api/compare/v6.16.1...v6.17.0) (2023-10-02)


### Features

* alert invalidate emoji ([baae8a1](https://github.com/consolelabs/mochi-api/commit/baae8a1576ebd076f4e81892fe0223383b883315))

## [6.16.1](https://github.com/consolelabs/mochi-api/compare/v6.16.0...v6.16.1) (2023-10-02)


### Bug Fixes

* **compare-token:** handle case token not found return 404 ([#1203](https://github.com/consolelabs/mochi-api/issues/1203)) ([610ec5e](https://github.com/consolelabs/mochi-api/commit/610ec5e8d869af6cc54b359bb04bbdf1f3939611))
* **get-market-chart:** handle app panic when geckoterminal return err or empty ([#1202](https://github.com/consolelabs/mochi-api/issues/1202)) ([2af08d0](https://github.com/consolelabs/mochi-api/commit/2af08d0ca1c22506f00bf3f6a162b514364d3187))

# [6.16.0](https://github.com/consolelabs/mochi-api/compare/v6.15.0...v6.16.0) (2023-09-22)


### Bug Fixes

* cache profile cmd - grouped by platform ([323ffff](https://github.com/consolelabs/mochi-api/commit/323ffff263525f51ea97dd6485252c267b35da85))
* convert friend tech key from camelCase to snake_case ([9affaac](https://github.com/consolelabs/mochi-api/commit/9affaacb10ca6821a945876e921467b85c61bcb3))


### Features

* support emoji query native token ([#1200](https://github.com/consolelabs/mochi-api/issues/1200)) ([b84de28](https://github.com/consolelabs/mochi-api/commit/b84de289335d988db3b4fb1ab0a112e20be987be))

# [6.15.0](https://github.com/consolelabs/mochi-api/compare/v6.14.1...v6.15.0) (2023-09-20)


### Bug Fixes

* cache - find by contract ([d347262](https://github.com/consolelabs/mochi-api/commit/d347262a9b57cca2cd6f8cf622caf422b0674a17))
* **tip:** add channel url to transferv2 payload ([#1196](https://github.com/consolelabs/mochi-api/issues/1196)) ([b0796eb](https://github.com/consolelabs/mochi-api/commit/b0796ebf40d381d54958b2e1213459c0691e61b6))
* update key price alert cache ([#1191](https://github.com/consolelabs/mochi-api/issues/1191)) ([8d9a7d6](https://github.com/consolelabs/mochi-api/commit/8d9a7d60dde50fefc7b07127da52268534123b21))
* use snake_case in friend metadata response ([44def76](https://github.com/consolelabs/mochi-api/commit/44def767a42b6b5b5c554bcbdb024083ba727c6a))


### Features

* add tracking key price change percentage ([03d955c](https://github.com/consolelabs/mochi-api/commit/03d955cfb24a2e4e2e6b7148f40d56d2e035e64f))
* job - cache top profile cmd usage ([b8f832b](https://github.com/consolelabs/mochi-api/commit/b8f832be62fe0342431cd2218945ff3f66159034))
* support hashtag with dictionary ([#1195](https://github.com/consolelabs/mochi-api/issues/1195)) ([48ccc09](https://github.com/consolelabs/mochi-api/commit/48ccc099d03f253e066f9c79f2b9f5dcb0441183))
* validate profile is member of server ([#1197](https://github.com/consolelabs/mochi-api/issues/1197)) ([6667f2e](https://github.com/consolelabs/mochi-api/commit/6667f2e0edac58a1166325ebd589697d0cc0b52d))

## [6.14.1](https://github.com/consolelabs/mochi-api/compare/v6.14.0...v6.14.1) (2023-09-15)


### Bug Fixes

* support insensitive hashtag template ([#1190](https://github.com/consolelabs/mochi-api/issues/1190)) ([c700b48](https://github.com/consolelabs/mochi-api/commit/c700b4887b35e8d0fa6311df5288a297680e5f38))

# [6.14.0](https://github.com/consolelabs/mochi-api/compare/v6.13.0...v6.14.0) (2023-09-15)


### Bug Fixes

* change friend scan api ([b161773](https://github.com/consolelabs/mochi-api/commit/b161773463547f5a385462342295758a61b509c6))
* count total emojis in pagination ([#1183](https://github.com/consolelabs/mochi-api/issues/1183)) ([eaabbdf](https://github.com/consolelabs/mochi-api/commit/eaabbdf39c7805cf41da9f82928aaab6006127a3))
* handle not found hashtag etmplate ([#1187](https://github.com/consolelabs/mochi-api/issues/1187)) ([d80bda8](https://github.com/consolelabs/mochi-api/commit/d80bda8a14b01806e9e214a032f0019f1a815b8a))
* order list tx by time asc ([43e4703](https://github.com/consolelabs/mochi-api/commit/43e4703779caf2a85a62118510738b388f46a3c9))
* produce kafka message using buildin utility ([#1181](https://github.com/consolelabs/mochi-api/issues/1181)) ([3c28fec](https://github.com/consolelabs/mochi-api/commit/3c28fecc2ffa31387051a406668a5c3457afb392))
* store hashtag template in metadata ([#1178](https://github.com/consolelabs/mochi-api/issues/1178)) ([02a997d](https://github.com/consolelabs/mochi-api/commit/02a997dde3c4db7e2a8268ef75913bbdd3b4e5ef))
* **tip:** add channel name for tip notification ([#1185](https://github.com/consolelabs/mochi-api/issues/1185)) ([3d3424a](https://github.com/consolelabs/mochi-api/commit/3d3424a0066cb69572b443ce91449d56302047cb))


### Features

* job to track friend tech keys new txs come ([893010c](https://github.com/consolelabs/mochi-api/commit/893010c739e23c45e3bb59df35c6e20e6efa3bd5))
* paginate emojis ([#1182](https://github.com/consolelabs/mochi-api/issues/1182)) ([7df5247](https://github.com/consolelabs/mochi-api/commit/7df5247c4d3b88966bfba3eeca1579b3308b0505))
* push key price change alert to ([52ba788](https://github.com/consolelabs/mochi-api/commit/52ba7880add43338394b813e105891d82783cd2d))

# [6.13.0](https://github.com/consolelabs/mochi-api/compare/v6.12.1...v6.13.0) (2023-09-13)


### Features

* check new changelog ([298fa63](https://github.com/consolelabs/mochi-api/commit/298fa63c98b45e4af1c40b98c108ca226227c515))
* data hashtag and api render ([#1174](https://github.com/consolelabs/mochi-api/issues/1174)) ([cabf66b](https://github.com/consolelabs/mochi-api/commit/cabf66b76ad0b28e873f551aaf4937d64bd585a2))
* friend tech key tracking apis ([#1173](https://github.com/consolelabs/mochi-api/issues/1173)) ([efad45b](https://github.com/consolelabs/mochi-api/commit/efad45b029b309adb091429160923edc9aa8801e))
* simple search for friend scan keys ([#1172](https://github.com/consolelabs/mochi-api/issues/1172)) ([35395f1](https://github.com/consolelabs/mochi-api/commit/35395f125492fd05900d527b5a12808caf0713f1))

## [6.12.1](https://github.com/consolelabs/mochi-api/compare/v6.12.0...v6.12.1) (2023-09-12)


### Bug Fixes

* store original tip amount ([#1168](https://github.com/consolelabs/mochi-api/issues/1168)) ([b38e787](https://github.com/consolelabs/mochi-api/commit/b38e787c26736adea203cca94acb93cb5fc7c239))
* **tip:** include channel id to notification payload ([#1170](https://github.com/consolelabs/mochi-api/issues/1170)) ([d671040](https://github.com/consolelabs/mochi-api/commit/d671040611c68fa04447b67b0e4a87848ca15f1c))
* update response get onboarding status ([#1171](https://github.com/consolelabs/mochi-api/issues/1171)) ([e6eb3ad](https://github.com/consolelabs/mochi-api/commit/e6eb3addf487664b95556543eba0a0e0e53b1bca))

# [6.12.0](https://github.com/consolelabs/mochi-api/compare/v6.11.0...v6.12.0) (2023-09-08)


### Features

* add command table ([bf2cbf6](https://github.com/consolelabs/mochi-api/commit/bf2cbf6be11ccbcde1a885d803e16aaa504879e4))
* alter field for product changelogs ([3dfca7c](https://github.com/consolelabs/mochi-api/commit/3dfca7c1ef95ee3046b6c978a8aac9b2e434b924))

# [6.11.0](https://github.com/consolelabs/mochi-api/compare/v6.10.1...v6.11.0) (2023-09-06)


### Bug Fixes

* request take to long for 1st time using token roles in discord ([17ae056](https://github.com/consolelabs/mochi-api/commit/17ae056f4b9cc6d16235f469d2e681ed3ec034db))
* update ci test ([de5ebfc](https://github.com/consolelabs/mochi-api/commit/de5ebfcf20a3a2185a8abc34b9b35d046b80874a))
* update function names to match convention ([8d362ff](https://github.com/consolelabs/mochi-api/commit/8d362fffd0930571fe02a97515df5d857dd06c13))


### Features

* tip with original tx id ([#1164](https://github.com/consolelabs/mochi-api/issues/1164)) ([c9a69ac](https://github.com/consolelabs/mochi-api/commit/c9a69acc63b2e3ded10b6ca6386165aebe62edc5))

## [6.10.1](https://github.com/consolelabs/mochi-api/compare/v6.10.0...v6.10.1) (2023-08-31)


### Bug Fixes

* improve performance profile sol ([#1161](https://github.com/consolelabs/mochi-api/issues/1161)) ([8f4c8fa](https://github.com/consolelabs/mochi-api/commit/8f4c8fa2919b191f312671dad2e8facd45a2a2f4))
* rm retry query coingecko sol token ([#1162](https://github.com/consolelabs/mochi-api/issues/1162)) ([c74efc1](https://github.com/consolelabs/mochi-api/commit/c74efc16b4d82b795296c4391de8e058c60a7412))

# [6.10.0](https://github.com/consolelabs/mochi-api/compare/v6.9.0...v6.10.0) (2023-08-30)


### Features

* asset base chain ([#1160](https://github.com/consolelabs/mochi-api/issues/1160)) ([4e7ca22](https://github.com/consolelabs/mochi-api/commit/4e7ca2212d1b79af762e687d236e936c8ef66083))

# [6.9.0](https://github.com/consolelabs/mochi-api/compare/v6.8.0...v6.9.0) (2023-08-29)


### Bug Fixes

* allow tip equal bal ([#1159](https://github.com/consolelabs/mochi-api/issues/1159)) ([8e008f0](https://github.com/consolelabs/mochi-api/commit/8e008f0b101313cb99c52440af0abf9a6511de7c))
* send more params tip moniker ([#1158](https://github.com/consolelabs/mochi-api/issues/1158)) ([97fcc13](https://github.com/consolelabs/mochi-api/commit/97fcc13ca0122d103405be6036f54c03f853872b))
* store message transfer token ([#1157](https://github.com/consolelabs/mochi-api/issues/1157)) ([8b70d6f](https://github.com/consolelabs/mochi-api/commit/8b70d6f0812fe4b804e910868366b7ce59dc517a))
* update checkbox template pr ([#1150](https://github.com/consolelabs/mochi-api/issues/1150)) ([2a8f15e](https://github.com/consolelabs/mochi-api/commit/2a8f15e4171ec6bf7f097680b4864bbaea1736ce))
* update content template pr ([#1151](https://github.com/consolelabs/mochi-api/issues/1151)) ([19f9b5a](https://github.com/consolelabs/mochi-api/commit/19f9b5a035019b971c4a1e6150bc7d296b9da9f9))
* update template pr ([#1149](https://github.com/consolelabs/mochi-api/issues/1149)) ([33bdad9](https://github.com/consolelabs/mochi-api/commit/33bdad97579ce7601128a7f57426f99a2ce0fb28))


### Features

* add product changelogs view ([a6c8a22](https://github.com/consolelabs/mochi-api/commit/a6c8a22b883222471ee0827f5853754e556db4bf))
* get changelog from github ([5927e4b](https://github.com/consolelabs/mochi-api/commit/5927e4b9b8f2df73082691d0475429e5f66237dc))
* onboarding start ([#1148](https://github.com/consolelabs/mochi-api/issues/1148)) ([86c4c62](https://github.com/consolelabs/mochi-api/commit/86c4c62da52c75f10db5651dfb1af1a097af37b1))
* token verbose add ethploer data ([ca7da53](https://github.com/consolelabs/mochi-api/commit/ca7da530c8fdb55f19cf9a6b0efec80fff1f63fb))

# [6.8.0](https://github.com/consolelabs/mochi-api/compare/v6.7.1...v6.8.0) (2023-08-23)


### Bug Fixes

* get emoji url ([b8f86ed](https://github.com/consolelabs/mochi-api/commit/b8f86edf0ef28ec05511f9ecddfe12653584e520))
* remove validate profile id ([90bef2b](https://github.com/consolelabs/mochi-api/commit/90bef2b02fdc2a8cbb273ce957577974bd5384d8))
* top cmd does not display some unknown users ([f85882e](https://github.com/consolelabs/mochi-api/commit/f85882ebfd2a8ae9c487e3ac7355ce5e23c6f221))
* vault transfer address ([401e615](https://github.com/consolelabs/mochi-api/commit/401e61550a737fc32ea64a0c3671fdc7393d795e))


### Features

* add product metadata changelog ([036808d](https://github.com/consolelabs/mochi-api/commit/036808da1c63bef09712d9855b46358ba57d8e29))
* alias product command ([#1140](https://github.com/consolelabs/mochi-api/issues/1140)) ([225a763](https://github.com/consolelabs/mochi-api/commit/225a76322e204441be2a5583a62b3aa0de19582a))
* get ault by id ([#1139](https://github.com/consolelabs/mochi-api/issues/1139)) ([790d9b8](https://github.com/consolelabs/mochi-api/commit/790d9b8ec1e59d4ae2835963e856c51aaec809e8))
* token verbose ([340c35b](https://github.com/consolelabs/mochi-api/commit/340c35bd632b3088a3944e3cf8ae18c6cea55ec4))

## [6.7.1](https://github.com/consolelabs/mochi-api/compare/v6.7.0...v6.7.1) (2023-08-21)


### Bug Fixes

* top cmd does not display some unknown users ([#1142](https://github.com/consolelabs/mochi-api/issues/1142)) ([43c0d82](https://github.com/consolelabs/mochi-api/commit/43c0d824e178bc918c8072209dc3da4d394f9efc))

# [6.7.0](https://github.com/consolelabs/mochi-api/compare/v6.6.0...v6.7.0) (2023-08-18)


### Features

* remove validate profile id ([64b196b](https://github.com/consolelabs/mochi-api/commit/64b196ba0213b37c12a98d42a46baade118710d7))

# [6.6.0](https://github.com/consolelabs/mochi-api/compare/v6.5.10...v6.6.0) (2023-08-17)


### Bug Fixes

* [sentry-296] swap route missing param ([#1129](https://github.com/consolelabs/mochi-api/issues/1129)) ([76f6292](https://github.com/consolelabs/mochi-api/commit/76f629215590c7dd1c9d396e244862b1f03c258f))
* cannot render changelog ([#1124](https://github.com/consolelabs/mochi-api/issues/1124)) ([1dacfc2](https://github.com/consolelabs/mochi-api/commit/1dacfc2248df39204caf578455418d1c416f93ed))


### Features

* add cd step to notify changelog ([#1131](https://github.com/consolelabs/mochi-api/issues/1131)) ([ee91925](https://github.com/consolelabs/mochi-api/commit/ee91925a11f9f5a14825a8c79ffd30bcae3101a7))
* product bot commands ([#1125](https://github.com/consolelabs/mochi-api/issues/1125)) ([7cc6b13](https://github.com/consolelabs/mochi-api/commit/7cc6b132f4e5615a4e691a1ca55758c299ad9da9))
