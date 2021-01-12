package config

type TypeCommonConfig interface {
	GetType() string
}

type CommonConfig struct {
	Type string `json:"type"`
}

func(t CommonConfig) GetType() string {
	return t.Type
}

type ConfigRaw map[string]interface{}