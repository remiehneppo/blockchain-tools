package service

type BinanceAPI interface {
	Ping() (string, error)
	GetAllCoins() ([]interface{}, error)
}

type BinanceAPIService struct {
	apiKey   string
	endpoint string
}

func NewBinanceAPIService(apiKey, endpoint string) *BinanceAPIService {
	return &BinanceAPIService{
		apiKey:   apiKey,
		endpoint: endpoint,
	}
}

func (b *BinanceAPIService) Ping() (string, error) {
	// Implement the ping logic here
	return "pong", nil
}

func (b *BinanceAPIService) GetAllCoins() ([]interface{}, error) {
	// Implement the logic to get all coins here
	return []interface{}{"BTC", "ETH", "BNB"}, nil
}
