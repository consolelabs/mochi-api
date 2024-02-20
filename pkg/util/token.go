package util

import (
	"fmt"
	"strings"
)

var emojis = map[string]string{
	"FTM":          "967285237686108212",
	"SPIRIT":       "967285237962924163",
	"TOMB":         "967285237904179211",
	"REAPER":       "967285238306857063",
	"BOO":          "967285238042599434",
	"SPELL":        "967285238063587358",
	"BTC":          "967285237879013388",
	"ETH":          "991657409082830858",
	"BNB":          "972205674715054090",
	"CAKE":         "972205674371117126",
	"OP":           "1002151912403107930",
	"USDT":         "1005010747308396544",
	"USDC":         "1005010675342520382",
	"ADA":          "1005010608443359272",
	"XRP":          "1005010559856554086",
	"BUSD":         "1005010097535197264",
	"DOT":          "1005009972716908554",
	"DOGE":         "1004962950756454441",
	"DAI":          "1005009904433647646",
	"MATIC":        "1037985931816349746",
	"AVAX":         "1005009817523474492",
	"UNI":          "1005012087967334443",
	"SHIB":         "1005009723277463703",
	"TRX":          "1005009394209128560",
	"WBTC":         "1005009348956790864",
	"ETC":          "1005009314802569277",
	"LEO":          "1005009244187263047",
	"LTC":          "1005009185940963380",
	"FTT":          "1005009144044064779",
	"CRO":          "1005009127937949797",
	"LINK":         "1005008904205385759",
	"NEAR":         "1005008870038589460",
	"ATOM":         "1005008855111049216",
	"XLM":          "1005008839139151913",
	"XMR":          "1005008819866312724",
	"BCH":          "1005008800106942525",
	"APE":          "1005008782486675536",
	"DFG":          "1007157463256145970",
	"BUTT":         "1007247521468403744",
	"WDOGE":        "1010512669448605756",
	"REN":          "1037985602202779690",
	"MANA":         "1037985604010508360",
	"COMP":         "1037985570724528178",
	"YFI":          "1037985592971116564",
	"BAT":          "1037985578341371964",
	"AAVE":         "1037985567146774538",
	"BNT":          "1037985589355626517",
	"MKR":          "1037985596964081696",
	"ANC":          "1037985575334051901",
	"BRUSH":        "1037985582162378783",
	"ICY":          ":ice_cube:",
	"MOCHI_CIRCLE": "1021636928094883864",
	"WALLET":       "1077631121614970992",
	"PENCIL":       "1078633895500722187",
	"PLUS":         "1078633897513992202",
	"POINTINGDOWN": "1058304350650384434",
	"APPROVE":      "1077631110047080478",
}

func GetEmoji(key string) string {
	if strings.ToUpper(key) == "ICY" {
		return emojis["ICY"]
	}
	return fmt.Sprintf("<:%s:%s>", key, emojis[strings.ToUpper(key)])
}

func GetEmojiID(key string) string {
	return emojis[strings.ToUpper(key)]
}
