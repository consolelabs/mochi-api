## [6.27.1](https://github.com/consolelabs/mochi-api/compare/v6.27.0...v6.27.1) (2024-01-22)


### Bug Fixes

* **changelog:** update unique value for product changelog view ([25cfa20](https://github.com/consolelabs/mochi-api/commit/25cfa20aaafe20542f258cc5536ac0f2da7fd847))
* payment settings - init prioritized tokens got null ([e9a4f79](https://github.com/consolelabs/mochi-api/commit/e9a4f79797b7581751130c0c2ed8351d122013c0))
* payment settings - use balances as prioritized tokens ([f700459](https://github.com/consolelabs/mochi-api/commit/f70045967b40c3da7ae1636fa613f2b63a82e2e3))

# [6.27.0](https://github.com/consolelabs/mochi-api/compare/v6.26.7...v6.27.0) (2024-01-17)


### Bug Fixes

* change response chain name of api get assets ([3c2ccc2](https://github.com/consolelabs/mochi-api/commit/3c2ccc2ab852195e23265d3683240646d8d1d86d))
* handle nil token when not supported ([be9956a](https://github.com/consolelabs/mochi-api/commit/be9956a719f57a62836ab7101fdf91b42eac531d))
* not filter native token in asset ([#1319](https://github.com/consolelabs/mochi-api/issues/1319)) ([ba23f17](https://github.com/consolelabs/mochi-api/commit/ba23f1741fbe4146b67a94b85a107d754a809885))
* **product-changelog:** add version field ([a14d7fc](https://github.com/consolelabs/mochi-api/commit/a14d7fcbb617c2d500969e8423d90546296576d2))
* **product-changelog:** get changelog version ([7247bb6](https://github.com/consolelabs/mochi-api/commit/7247bb64b98caaed929d2cbad9034e2cc56d3437))
* revert logic of SearchCoins ([a4218f8](https://github.com/consolelabs/mochi-api/commit/a4218f8c7bb632b4069283d7de5ddfa067f16292))
* update payment setting API ([731cbc7](https://github.com/consolelabs/mochi-api/commit/731cbc72bb2b5d40d08597c45cfcdb4670bced6a))
* update setting swagger ([61e2b9e](https://github.com/consolelabs/mochi-api/commit/61e2b9ece6d2c3262610b86840f4d783ca6f9301))


### Features

* wallet sol render lending ([#1318](https://github.com/consolelabs/mochi-api/issues/1318)) ([0a9240b](https://github.com/consolelabs/mochi-api/commit/0a9240bede70644c5900684db3f73f7a1c86c2a1))

## [6.26.7](https://github.com/consolelabs/mochi-api/compare/v6.26.6...v6.26.7) (2024-01-12)


### Bug Fixes

* add more data for user balances api ([2e29b7c](https://github.com/consolelabs/mochi-api/commit/2e29b7c6ac5de4f190cd9841fa80197d860eab21))
* add more data for user balances api ([2ac99ab](https://github.com/consolelabs/mochi-api/commit/2ac99abebf0f090d70308335f61ba036c9b4980c))
* add swagger for get user total balances ([7c2147f](https://github.com/consolelabs/mochi-api/commit/7c2147f58bb062dd5388bfb660853d6363ae9794))
* must wait stable when searching pair ([68084ce](https://github.com/consolelabs/mochi-api/commit/68084ce45979de7dddec1714978e47668b131b81))
* privacy settings - update structure ([ccfe54b](https://github.com/consolelabs/mochi-api/commit/ccfe54bb5f8da3ed49661b45324e6ecfff8ce61d))
* privacy settings - update structure ([4f46398](https://github.com/consolelabs/mochi-api/commit/4f463989491f291ccbf8aa5b35069edffb25d771))
* transfer-v2 - overwrited metadata ([4ed1e06](https://github.com/consolelabs/mochi-api/commit/4ed1e0623e2a0d3d0ea8bf9bbd496aea5dd62989))
* user privacy settings - update initial configuration ([7c49483](https://github.com/consolelabs/mochi-api/commit/7c494833d3629640fb697f8cc4172943b29d3c6c))
* validate money source ([0a6613e](https://github.com/consolelabs/mochi-api/commit/0a6613e6f8214739bb43a080651ff24b30608af6))
* validate privacy custom settings ([18e2fd7](https://github.com/consolelabs/mochi-api/commit/18e2fd78a7a93ee8fcdd4af4587f9a68c2681672))

## [6.26.6](https://github.com/consolelabs/mochi-api/compare/v6.26.5...v6.26.6) (2024-01-05)


### Bug Fixes

* temp disable changelog notify ([#1306](https://github.com/consolelabs/mochi-api/issues/1306)) ([673df6f](https://github.com/consolelabs/mochi-api/commit/673df6f8f75b16c92015eb004e0b73c032e30c76))

## [6.26.5](https://github.com/consolelabs/mochi-api/compare/v6.26.4...v6.26.5) (2024-01-05)


### Bug Fixes

* multiple ticker ([#1305](https://github.com/consolelabs/mochi-api/issues/1305)) ([65f5695](https://github.com/consolelabs/mochi-api/commit/65f56955b7f1eeb9fa6774def68e485642a0306e))
* remove unused cache binance ([#1304](https://github.com/consolelabs/mochi-api/issues/1304)) ([6326b41](https://github.com/consolelabs/mochi-api/commit/6326b41726f698ba8c0e96c1256bb38340874c20))
* setting validation ([c8ccb24](https://github.com/consolelabs/mochi-api/commit/c8ccb24116a6b0a155ba239fdccb70c3d67f4602))
* update get vaults API ([1ec71cd](https://github.com/consolelabs/mochi-api/commit/1ec71cdb1def684ef15ad6320e009e4ef5cf6f46))
* update setting web platform ([ddefac5](https://github.com/consolelabs/mochi-api/commit/ddefac5916fa9531da6f4c1a6c07843084f00757))

## [6.26.4](https://github.com/consolelabs/mochi-api/compare/v6.26.3...v6.26.4) (2024-01-04)


### Bug Fixes

* add logger to debug geckoterminal data crawling ([92064c2](https://github.com/consolelabs/mochi-api/commit/92064c263387308acad8e0e5691d35da7986abe4))
* **emojis:** remove pagination when query emojis list with codes ([ae91faf](https://github.com/consolelabs/mochi-api/commit/ae91faf448ca10b6b19ba0c5e0c5b8241bdebca3))
* notification settings - add validation ([487e9ac](https://github.com/consolelabs/mochi-api/commit/487e9ace4d6ccfd5af95dad46ac31eb8dd48ff87))

## [6.26.3](https://github.com/consolelabs/mochi-api/compare/v6.26.2...v6.26.3) (2024-01-03)


### Bug Fixes

* remove cache total amt binance ([#1298](https://github.com/consolelabs/mochi-api/issues/1298)) ([7c88fd9](https://github.com/consolelabs/mochi-api/commit/7c88fd9edbb0a218813d57a5dc2989ccfe38f516))

## [6.26.2](https://github.com/consolelabs/mochi-api/compare/v6.26.1...v6.26.2) (2024-01-03)


### Bug Fixes

* update binance query api ([#1297](https://github.com/consolelabs/mochi-api/issues/1297)) ([6b38970](https://github.com/consolelabs/mochi-api/commit/6b38970bc91ebaf3ea64427cc22a4d9816045c98))

## [6.26.1](https://github.com/consolelabs/mochi-api/compare/v6.26.0...v6.26.1) (2023-12-28)


### Bug Fixes

* user watchlist token remove pagination ([#1294](https://github.com/consolelabs/mochi-api/issues/1294)) ([2c49ecd](https://github.com/consolelabs/mochi-api/commit/2c49ecdaec252045ed4f1fba1a1fc9bf899214cf))

# [6.26.0](https://github.com/consolelabs/mochi-api/compare/v6.25.4...v6.26.0) (2023-12-28)


### Bug Fixes

* add icon wallet asset ([#1293](https://github.com/consolelabs/mochi-api/issues/1293)) ([179dfea](https://github.com/consolelabs/mochi-api/commit/179dfea37d9a7bddfcae606ef1a7d73d1766ab4e))
* init notification flags ([3e7829b](https://github.com/consolelabs/mochi-api/commit/3e7829b8f2c274189a5cdbdf07dc3f880a2214d7))
* init user notification flags ([f75b56b](https://github.com/consolelabs/mochi-api/commit/f75b56b13915ac9444fab5b684976da1f7f22819))
* payment settings - remove default token ([16e654c](https://github.com/consolelabs/mochi-api/commit/16e654cbe5a759801a9425276b98597a4841aa2b))
* render sol wallet ([#1292](https://github.com/consolelabs/mochi-api/issues/1292)) ([fba2e74](https://github.com/consolelabs/mochi-api/commit/fba2e74e9fdbfe532d5ee304e63871e223662355))
* small tweak - update noti settings API ([65e8e1c](https://github.com/consolelabs/mochi-api/commit/65e8e1c9ec87355c46c67f67febf6a02814a2e3d))
* update get coingecko market cache time to 5 minutes ([#1295](https://github.com/consolelabs/mochi-api/issues/1295)) ([683ffe3](https://github.com/consolelabs/mochi-api/commit/683ffe3700ab2ee113f71c42b727d947ea1541eb))
* update noti settings ([8df26ff](https://github.com/consolelabs/mochi-api/commit/8df26ff1e01f707517e304590761f688618abca6))


### Features

* implement read/write mochi setting ([55bcf6b](https://github.com/consolelabs/mochi-api/commit/55bcf6b4d3602f1bb16152e4a540adcaa5fd29b4))
* mochi settings ([086943d](https://github.com/consolelabs/mochi-api/commit/086943d930b9bbbb6d29cf09ec4df1c3c1f340f1))

## [6.25.4](https://github.com/consolelabs/mochi-api/compare/v6.25.3...v6.25.4) (2023-12-20)


### Bug Fixes

* remove base when get balance vault ([#1284](https://github.com/consolelabs/mochi-api/issues/1284)) ([965c7b8](https://github.com/consolelabs/mochi-api/commit/965c7b8d33a277d53fde8649aa1292e662251ec4))

## [6.25.3](https://github.com/consolelabs/mochi-api/compare/v6.25.2...v6.25.3) (2023-12-19)


### Bug Fixes

* init wrong friendtech api endpoint ([145981e](https://github.com/consolelabs/mochi-api/commit/145981e5f64d8613de203c7ebc7f6665091b2262))

## [6.25.2](https://github.com/consolelabs/mochi-api/compare/v6.25.1...v6.25.2) (2023-12-19)


### Bug Fixes

* remove default friend scan api ([1ebfacc](https://github.com/consolelabs/mochi-api/commit/1ebfacc6f7045405f6bd449d22d0ccd6f22fc475))

## [6.25.1](https://github.com/consolelabs/mochi-api/compare/v6.25.0...v6.25.1) (2023-12-19)


### Bug Fixes

* friendtech svc is missing initiating sentry ([ea22a5d](https://github.com/consolelabs/mochi-api/commit/ea22a5d3d92087b098ce49da5ba1a70d9c2c8cdc))
* temp solution handle exceed covalent quote ([#1279](https://github.com/consolelabs/mochi-api/issues/1279)) ([7fa2411](https://github.com/consolelabs/mochi-api/commit/7fa24119a4596a1aef8b2cbb630061ce76e8c4ee))

# [6.25.0](https://github.com/consolelabs/mochi-api/compare/v6.24.4...v6.25.0) (2023-12-13)


### Bug Fixes

* add retry nghenhan service ([#1278](https://github.com/consolelabs/mochi-api/issues/1278)) ([c5649c3](https://github.com/consolelabs/mochi-api/commit/c5649c30b0648ee98707b0cbe13c4e1453e167ce))


### Features

* mock API - global profile info ([833d7d3](https://github.com/consolelabs/mochi-api/commit/833d7d3fb5aebc7a9fb2792f9d5c83483210d153))

## [6.24.4](https://github.com/consolelabs/mochi-api/compare/v6.24.3...v6.24.4) (2023-12-11)


### Bug Fixes

* cache market data coingecko ([#1276](https://github.com/consolelabs/mochi-api/issues/1276)) ([19a11e8](https://github.com/consolelabs/mochi-api/commit/19a11e8b6f335d6eb30da6f1050466482d80d1e9))
* transfer-v2 validation ([1c693ea](https://github.com/consolelabs/mochi-api/commit/1c693ea149a1a2b89f0812229b577a6c8c8b12df))

## [6.24.3](https://github.com/consolelabs/mochi-api/compare/v6.24.2...v6.24.3) (2023-12-08)


### Bug Fixes

* enrich template tx ([#1274](https://github.com/consolelabs/mochi-api/issues/1274)) ([a7adeec](https://github.com/consolelabs/mochi-api/commit/a7adeec51a642a47febc2d54d2ab1384816ab894))

## [6.24.2](https://github.com/consolelabs/mochi-api/compare/v6.24.1...v6.24.2) (2023-12-08)


### Bug Fixes

* nghenhan - init sentry ([64b63c6](https://github.com/consolelabs/mochi-api/commit/64b63c66f55ae07faa972e9fb8dc3c94b1aea5de))

## [6.24.1](https://github.com/consolelabs/mochi-api/compare/v6.24.0...v6.24.1) (2023-12-07)


### Bug Fixes

* **vault:** update correct key to query evm assets ([#1272](https://github.com/consolelabs/mochi-api/issues/1272)) ([cf3bf4f](https://github.com/consolelabs/mochi-api/commit/cf3bf4ff14c723024c49a703bc1b00c6e87d612b))

# [6.24.0](https://github.com/consolelabs/mochi-api/compare/v6.23.3...v6.24.0) (2023-12-06)


### Bug Fixes

* extend timeout url ([#1271](https://github.com/consolelabs/mochi-api/issues/1271)) ([9cfed45](https://github.com/consolelabs/mochi-api/commit/9cfed4529308adfe050e7161c3e16121b72cba17))


### Features

* add redis sentiel url env ([#1270](https://github.com/consolelabs/mochi-api/issues/1270)) ([245d437](https://github.com/consolelabs/mochi-api/commit/245d4370afd237b69f28f8f48a2d665b2496fcaa))

## [6.23.3](https://github.com/consolelabs/mochi-api/compare/v6.23.2...v6.23.3) (2023-12-06)


### Bug Fixes

* add chain type ([#1267](https://github.com/consolelabs/mochi-api/issues/1267)) ([426ba32](https://github.com/consolelabs/mochi-api/commit/426ba32b5edddbca6f7bf963f7155248bf6f77ef))
* add chain type asset ([#1269](https://github.com/consolelabs/mochi-api/issues/1269)) ([75a115e](https://github.com/consolelabs/mochi-api/commit/75a115e40781054b8cc71e9e089b7d3a3f66788a))
* add more log service ([#1263](https://github.com/consolelabs/mochi-api/issues/1263)) ([c5567ad](https://github.com/consolelabs/mochi-api/commit/c5567ade5b5f4606754a460961f3c396d814ecc0))
* handle case vault not exist ([#1265](https://github.com/consolelabs/mochi-api/issues/1265)) ([38e0e0c](https://github.com/consolelabs/mochi-api/commit/38e0e0c7258fd4d39a5cb582de01dfdcdc352843))
* handle non supported wallet ([#1262](https://github.com/consolelabs/mochi-api/issues/1262)) ([fef0443](https://github.com/consolelabs/mochi-api/commit/fef04436ff521b0279ad3a7335c667c02e436a6d))
* only send internal coingecko error to sentry ([#1266](https://github.com/consolelabs/mochi-api/issues/1266)) ([206fced](https://github.com/consolelabs/mochi-api/commit/206fced1317bcbe3c842d1b8a4f99f7ee9b35272))
* resolve crash ([#1264](https://github.com/consolelabs/mochi-api/issues/1264)) ([4e524a6](https://github.com/consolelabs/mochi-api/commit/4e524a626315f4c03f75796977ce0b29ca288b69))
* rm wrong log ([16b4813](https://github.com/consolelabs/mochi-api/commit/16b4813768282006454794e4e79b3bf124fa0383))

## [6.23.2](https://github.com/consolelabs/mochi-api/compare/v6.23.1...v6.23.2) (2023-12-01)


### Bug Fixes

* add decimal moniker ([#1260](https://github.com/consolelabs/mochi-api/issues/1260)) ([9f98b65](https://github.com/consolelabs/mochi-api/commit/9f98b65398ae969867b964b943e9703eccefd9db))
* add fields moniker default ([#1259](https://github.com/consolelabs/mochi-api/issues/1259)) ([2813d3a](https://github.com/consolelabs/mochi-api/commit/2813d3a54f06305613ca56c2e3ea3eb3c2826c8e))
* birdeye requests - update error msg ([bd4db3e](https://github.com/consolelabs/mochi-api/commit/bd4db3e0b5e062f6704c35d59732dfaeba33514a))
* handle rate limit skymavis ([#1258](https://github.com/consolelabs/mochi-api/issues/1258)) ([2f54a6f](https://github.com/consolelabs/mochi-api/commit/2f54a6fd30341561ec44a39e4fb3b0163eddb9b6))
* handle timeout 3rd party ([#1257](https://github.com/consolelabs/mochi-api/issues/1257)) ([2688797](https://github.com/consolelabs/mochi-api/commit/268879775b2d67d6177733d7881610b035be33d8))

## [6.23.1](https://github.com/consolelabs/mochi-api/compare/v6.23.0...v6.23.1) (2023-11-30)


### Bug Fixes

* restore api create guild ([#1255](https://github.com/consolelabs/mochi-api/issues/1255)) ([312c81c](https://github.com/consolelabs/mochi-api/commit/312c81c0fa0a9f7fe3f9a2e95edf634b808d01b8))
* slow response gm webhook ([#1256](https://github.com/consolelabs/mochi-api/issues/1256)) ([c4e12e0](https://github.com/consolelabs/mochi-api/commit/c4e12e0a8adfc6f8cad59c03373b4175176210c5))

# [6.23.0](https://github.com/consolelabs/mochi-api/compare/v6.22.12...v6.23.0) (2023-11-28)


### Bug Fixes

* add sentry service layer ([#1247](https://github.com/consolelabs/mochi-api/issues/1247)) ([3242d19](https://github.com/consolelabs/mochi-api/commit/3242d195f20ee03507778cc39e0470198add0e9e))
* add theme transaction ([#1248](https://github.com/consolelabs/mochi-api/issues/1248)) ([d17f4a3](https://github.com/consolelabs/mochi-api/commit/d17f4a3aca85ae06eea71fb8a0b7c83d92b9e310))
* add theme transaction v2 ([#1249](https://github.com/consolelabs/mochi-api/issues/1249)) ([4c7a6f2](https://github.com/consolelabs/mochi-api/commit/4c7a6f23295d95b8b0a1803fb5dedabaa144cf3f))
* handle coingecko id not found ([#1252](https://github.com/consolelabs/mochi-api/issues/1252)) ([7668a5b](https://github.com/consolelabs/mochi-api/commit/7668a5be8b9c7453912afa09cd47d3107fdc1aef))
* handle retry getting gas tracker ([#1254](https://github.com/consolelabs/mochi-api/issues/1254)) ([1a46c54](https://github.com/consolelabs/mochi-api/commit/1a46c54ab381a36709d7eabf861bbef26e90c95e))
* ignore search coingecko if query empty ([#1253](https://github.com/consolelabs/mochi-api/issues/1253)) ([f69c62f](https://github.com/consolelabs/mochi-api/commit/f69c62f0231ab1c3787b7c33a11b0918cd04a84b))


### Features

* add method to get jetton balance ([5b0bf50](https://github.com/consolelabs/mochi-api/commit/5b0bf5002dd047b3bd3577f5d90be51e4299f838))

## [6.22.12](https://github.com/consolelabs/mochi-api/compare/v6.22.11...v6.22.12) (2023-11-24)


### Bug Fixes

* add sentry for cj invalidate emoji and fetch coingecko token ([#1246](https://github.com/consolelabs/mochi-api/issues/1246)) ([de3443f](https://github.com/consolelabs/mochi-api/commit/de3443fcd0803a72ca4bf7ad450ffcebc37be9d5))

## [6.22.11](https://github.com/consolelabs/mochi-api/compare/v6.22.10...v6.22.11) (2023-11-21)


### Bug Fixes

* handle case create vault req ([#1244](https://github.com/consolelabs/mochi-api/issues/1244)) ([4f976d4](https://github.com/consolelabs/mochi-api/commit/4f976d425fafd6429d16f582f15b79da9b4e6733))

## [6.22.10](https://github.com/consolelabs/mochi-api/compare/v6.22.9...v6.22.10) (2023-11-21)


### Bug Fixes

* logic vault ([#1243](https://github.com/consolelabs/mochi-api/issues/1243)) ([259feb5](https://github.com/consolelabs/mochi-api/commit/259feb5251497bd22974303e2078d5d3410c682f))

## [6.22.9](https://github.com/consolelabs/mochi-api/compare/v6.22.8...v6.22.9) (2023-11-21)


### Bug Fixes

* get custom token null ([#1241](https://github.com/consolelabs/mochi-api/issues/1241)) ([205d589](https://github.com/consolelabs/mochi-api/commit/205d589670b9b900f0a422e3b7977ab2620d7a90))
* query geckoterminal not found return empty ([#1242](https://github.com/consolelabs/mochi-api/issues/1242)) ([4804da3](https://github.com/consolelabs/mochi-api/commit/4804da30afbe44ff7066f7146d56229869931fcd))

## [6.22.8](https://github.com/consolelabs/mochi-api/compare/v6.22.7...v6.22.8) (2023-11-20)


### Bug Fixes

* add auth query profile ([#1239](https://github.com/consolelabs/mochi-api/issues/1239)) ([7871d75](https://github.com/consolelabs/mochi-api/commit/7871d75597f78bbebb58d90e04b682a0f7b50294))
* bump toolkit typeset version ([#1240](https://github.com/consolelabs/mochi-api/issues/1240)) ([a05b25a](https://github.com/consolelabs/mochi-api/commit/a05b25a3d23e5586811fbac5175edac2aa5d9658))
* enrich token price moniker ([#1236](https://github.com/consolelabs/mochi-api/issues/1236)) ([88f3b89](https://github.com/consolelabs/mochi-api/commit/88f3b89d1b4642827aa5696d4aabdd31048f58fb))
* return id wallet onchain ([#1237](https://github.com/consolelabs/mochi-api/issues/1237)) ([108e6e4](https://github.com/consolelabs/mochi-api/commit/108e6e4285f91f72751162a7ad5077e8d359f236))
* wrong compare address onchain asset ([#1238](https://github.com/consolelabs/mochi-api/issues/1238)) ([c0b36fd](https://github.com/consolelabs/mochi-api/commit/c0b36fdcc559f27356d3e8abf2778b004ccdac5c))

## [6.22.7](https://github.com/consolelabs/mochi-api/compare/v6.22.6...v6.22.7) (2023-11-17)


### Bug Fixes

* mochi secret env ([53b6326](https://github.com/consolelabs/mochi-api/commit/53b6326b268fbdec3440d28f693c7dc47e8a1b31))

## [6.22.6](https://github.com/consolelabs/mochi-api/compare/v6.22.5...v6.22.6) (2023-11-17)


### Bug Fixes

* update token role get profile by 50 batch ([#1234](https://github.com/consolelabs/mochi-api/issues/1234)) ([27f7795](https://github.com/consolelabs/mochi-api/commit/27f7795216de8ac191f513973ce43400cc651694))

## [6.22.5](https://github.com/consolelabs/mochi-api/compare/v6.22.4...v6.22.5) (2023-11-17)


### Bug Fixes

* update token role log ([#1233](https://github.com/consolelabs/mochi-api/issues/1233)) ([52fde58](https://github.com/consolelabs/mochi-api/commit/52fde58a1d3ceab0520eeca2a2fc71c8adcabe09))

## [6.22.4](https://github.com/consolelabs/mochi-api/compare/v6.22.3...v6.22.4) (2023-11-15)


### Bug Fixes

* track channel tip tx ([#1231](https://github.com/consolelabs/mochi-api/issues/1231)) ([a6cecdb](https://github.com/consolelabs/mochi-api/commit/a6cecdb75d948fd8b2a5714f6dcaa57805d81e14))

## [6.22.3](https://github.com/consolelabs/mochi-api/compare/v6.22.2...v6.22.3) (2023-11-13)


### Bug Fixes

* apply get list profile for token role update ([#1229](https://github.com/consolelabs/mochi-api/issues/1229)) ([0522e74](https://github.com/consolelabs/mochi-api/commit/0522e74b96b60cca2c4c914d3df08395116c473e))

## [6.22.2](https://github.com/consolelabs/mochi-api/compare/v6.22.1...v6.22.2) (2023-11-09)


### Bug Fixes

* support mnt for update user role ([8d0835a](https://github.com/consolelabs/mochi-api/commit/8d0835af575d47f76fb696b2d11a6d7ffe312bd1))

## [6.22.1](https://github.com/consolelabs/mochi-api/compare/v6.22.0...v6.22.1) (2023-11-09)


### Bug Fixes

* rename env name application id ([#1227](https://github.com/consolelabs/mochi-api/issues/1227)) ([bc60715](https://github.com/consolelabs/mochi-api/commit/bc607151cf4315faefc30e911d3db101cb0674f9))

# [6.22.0](https://github.com/consolelabs/mochi-api/compare/v6.21.0...v6.22.0) (2023-11-09)


### Features

* command permission install link ([#1226](https://github.com/consolelabs/mochi-api/issues/1226)) ([8fcd8ca](https://github.com/consolelabs/mochi-api/commit/8fcd8cadc80391a02a27b67c54a765ec5d896056))

# [6.21.0](https://github.com/consolelabs/mochi-api/compare/v6.20.0...v6.21.0) (2023-11-08)


### Bug Fixes

* whitelist missing chains to serve getting list asset ([ec0b700](https://github.com/consolelabs/mochi-api/commit/ec0b7009cc598a6a598c0bf8d10c57a5cd7c08ac))


### Features

* **cmd:** add available cmds to discord guild ([#1219](https://github.com/consolelabs/mochi-api/issues/1219)) ([a808a6a](https://github.com/consolelabs/mochi-api/commit/a808a6a3e702ecd025027ba1993c23b062266676))

# [6.20.0](https://github.com/consolelabs/mochi-api/compare/v6.19.0...v6.20.0) (2023-11-06)


### Features

* integrate chain Mantle ([ef15e89](https://github.com/consolelabs/mochi-api/commit/ef15e895d2a021a899e1bfcfd9a7a46fa4518062))

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
