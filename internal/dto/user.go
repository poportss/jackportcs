package dto

type UserTradeLink struct {
	TradeLink string `json:"tradeLink"`
}

type UserAddress struct {
	Address    string `json:"address"`
	Number     string `json:"number"`
	Complement string `json:"complement"`
	ZipCode    string `json:"zipCode"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
}
