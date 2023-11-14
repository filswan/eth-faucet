package server

type Config struct {
	network         string
	symbol          string
	httpPort        int
	interval        int
	payout          int
	payoutString    string
	proxyCount      int
	hcaptchaSiteKey string
	hcaptchaSecret  string
}

func NewConfig(network, symbol string, payoutString string, httpPort, interval, payout, proxyCount int, hcaptchaSiteKey, hcaptchaSecret string) *Config {
	return &Config{
		network:         network,
		symbol:          symbol,
		httpPort:        httpPort,
		interval:        interval,
		payout:          payout,
		payoutString:    payoutString,
		proxyCount:      proxyCount,
		hcaptchaSiteKey: hcaptchaSiteKey,
		hcaptchaSecret:  hcaptchaSecret,
	}
}
