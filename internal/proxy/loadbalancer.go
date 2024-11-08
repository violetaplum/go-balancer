package proxy

import (
	"github.com/samber/lo"
	"github.com/violetaplum/go-balancer/config"
	"sync"
	"time"
)

type LoadBalancer struct {
	nodes []*Node
	mu    sync.Mutex
}

func NewLoadBalancer(configs []config.NodeConfig) *LoadBalancer {
	lb := &LoadBalancer{}
	lb.nodes = lo.FilterMap(configs, func(cfg config.NodeConfig, _ int) (*Node, bool) {
		return &Node{
			URL:           cfg.URL,
			MaxBPM:        cfg.MaxBPM,
			MaxRPM:        cfg.MaxRPM,
			lastResetTime: time.Now(),
		}, true
	})

	return lb
}

func (lb *LoadBalancer) GetNextNode(bodySize int32) *Node {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	for _, node := range lb.nodes {
		if node.GetNode(bodySize) {
			return node
		}
	}
	return nil
}
