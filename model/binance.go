package model

//SymbolFilter 심볼 필터, 추후 여러종류 필요함
type SymbolFilter struct {
	FilterType string `json:"filterType"`
	MinPrice   string `json:"minPrice"`
	MaxPrice   string `json:"maxPrice"`
	TickSize   string `json:"tickSize"`
}

//SymbolInfo 바이낸스 심볼 정보
type SymbolInfo struct {
	// "ETHBTC"
	Symbol string `json:"symbol"`
	//"TRADING"
	Status string `json:"status"`
	//"ETH"
	BaseAsset string `json:"baseAsset"`
	// "BTC",
	QuoteAsset string `json:"quoteAsset"`

	// baseAssetPrecision         int
	// quotePrecision             int
	// quoteAssetPrecision        int
	// baseCommissionPrecision    int
	// quoteCommissionPrecision   int
	// orderTypes                 []string
	// icebergAllowed             bool
	// ocoAllowed                 bool
	// quoteOrderQtyMarketAllowed bool
	// isSpotTradingAllowed       bool
	// isMarginTradingAllowed     bool
	Filters []SymbolFilter `json:"filters"`
}

//ExchangeInfo 거래소 정보
type ExchangeInfo struct {
	Timezone   string `json:"timezone"`
	ServerTime int64  `json:"serverTime"`
	//rateLimits
	//exchangefilters
	Symbols []SymbolInfo `json:"symbols"`
}

//Kline 바이낸스 캔들
type Kline struct {
	OpenTime         int64
	Open             float64
	High             float64
	Low              float64
	Close            float64
	Volume           float64
	CloseTime        int64
	QuoteVolume      float64
	NumberOfTrades   int64
	TakerBaseVolume  float64
	TakerQuoteVolume float64
	Ignore           float64
}
