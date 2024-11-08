package proxy

import (
	"sync"
	"time"
)

type Node struct {
	URL           string
	MaxBPM        int32
	MaxRPM        int32
	currentBPM    int32
	currentRPM    int32
	lastResetTime time.Time
	mu            sync.Mutex
}

func (n *Node) CanHandle(bodySize int32) bool {
	n.mu.Lock()
	defer n.mu.Unlock()

	now := time.Now()
	if now.Sub(n.lastResetTime) >= time.Minute {
		n.currentBPM = 0
		n.currentRPM = 0
		n.lastResetTime = now
	}

	if n.currentBPM+bodySize > n.MaxBPM ||
		n.currentRPM+1 > n.MaxRPM {
		return false
	}
	n.currentBPM += bodySize
	n.currentRPM++
	return true
}
