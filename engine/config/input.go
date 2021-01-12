package config

import "context"

type TypeInputConfig interface {
	TypeCommonConfig
	Start(ctx context.Context) (err error)
}

type InputConfig struct {
	CommonConfig
}