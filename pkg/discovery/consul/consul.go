package consul

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

// Dicovery provides service discovery for Consul
type Discovery struct {
	client *api.Client
}

// New provides initialization for consul discovery
func New(address string) (*Discovery, error) {
	conf := api.DefaultConfig()
	conf.Address = address
	client, err := api.NewClient(conf)
	if err != nil {
		return nil, fmt.Errorf("unablw to init consul client: %v", err)
	}
	return &Discovery{
		client: client,
	}, nil
}
