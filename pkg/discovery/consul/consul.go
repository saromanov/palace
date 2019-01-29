package consul

import (
	"fmt"
	"time"

	"github.com/hashicorp/consul/api"
)

// Dicovery provides service discovery for Consul
type Discovery struct {
	client *api.Client
}

func New(address string)(*Discovery, error){
	conf := api.DefaultConfig()
	conf.Address = address
}