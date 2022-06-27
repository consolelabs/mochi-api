package consts

type Constant struct {
	ERC721 struct {
		TransferTopic string
		InterfaceId   [4]byte
	}
}

// TODO: beautify
var C = Constant{
	ERC721: struct {
		TransferTopic string
		InterfaceId   [4]byte
	}{
		TransferTopic: "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
		InterfaceId:   [4]byte{128, 172, 88, 205}, // stand for erc721-id
	},
}
