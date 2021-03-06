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
		return nil, fmt.Errorf("unable to init consul client: %v", err)
	}
	return &Discovery{
		client: client,
	}, nil
}

// GetServices returns map of registered services
func (d *Discovery) GetServices() (map[string][]string, error) {
	return d.getServices()
}

// returns map of services
func (d *Discovery) getServices() (map[string][]string, error) {
	data, _, err := d.client.Catalog().Services(nil)
	if err != nil {
		return nil, fmt.Errorf("unable to get services: %v", err)
	}

	return data, nil
}
