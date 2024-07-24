package config

type PaymentGateway struct {
	Payment Payment `mapstructure:"payment"`
}

type Payment struct {
	ApiKey       string `mapstructure:"api-key"`
	MerchantCode string `mapstructure:"merchant-code"`
}
