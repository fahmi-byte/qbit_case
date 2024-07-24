package config

type PaymentConfig struct {
	ApiKey       string
	MerchantCode string
}

func (c *PaymentConfig) AccessApiKey() string {
	return c.ApiKey
}

func (c *PaymentConfig) AccessMerchantCode() string {
	return c.MerchantCode
}
