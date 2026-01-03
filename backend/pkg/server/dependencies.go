package server

import (
	"github.com/valkey-io/valkey-go"
)

type Dependencies struct {
	ValkeyClient valkey.Client
}

func NewValkeyClient(addr string) (valkey.Client, error) {
	client, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{addr}})
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewDependencies(cfg Config) (Dependencies, error) {
	valkeyClient, err := NewValkeyClient(cfg.ValkeyAddr)
	if err != nil {
		return Dependencies{}, err
	}
	return Dependencies{
		ValkeyClient: valkeyClient,
	}, nil
}
