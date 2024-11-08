package proxy

import (
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
	for _, cf := range configs {
		lb.nodes = append(lb.nodes, &Node{
			URL:           cf.URL,
			MaxBPM:        cf.MaxBPM,
			MaxRPM:        cf.MaxRPM,
			lastResetTime: time.Now(),
		})
	}
	return lb
}

func (lb *LoadBalancer) GetNextAvailableNode(bodySize int32) *Node {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	for _, node := range lb.nodes {
		if node.CanHandle(bodySize) {
			return node
		}
	}
	return nil
}
