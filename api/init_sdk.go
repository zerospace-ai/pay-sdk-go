package api

type SDKConfig struct {
	ApiKey             string
	ApiSecret          string
	PlatformPubKey     string
	PlatformRiskPubKey string
	RsaPrivateKey      string
}

type Sdk struct {
	config SDKConfig
}

func NewSDK(conf SDKConfig) *Sdk {
	return &Sdk{
		config: conf,
	}
}

func (s *Sdk) InitSDK() error {
	return nil
}

func (s *Sdk) GetApiKey() string {
	return s.config.ApiKey
}
