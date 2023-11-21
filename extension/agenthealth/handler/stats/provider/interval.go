// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package provider

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/aws/amazon-cloudwatch-agent/extension/agenthealth/handler/stats/agent"
)

// intervalStats restricts the Stats get function to once
// per interval.
type intervalStats struct {
	interval time.Duration

	once *sync.Once
	mu   sync.RWMutex

	stats atomic.Value
}

var _ agent.StatsProvider = (*intervalStats)(nil)

func (p *intervalStats) Stats(string) agent.Stats {
	p.mu.RLock()
	defer p.mu.RUnlock()
	var stats agent.Stats
	p.once.Do(func() {
		stats = p.getStats()
		time.AfterFunc(p.interval, p.resetOnce)
	})
	return stats
}

func (p *intervalStats) getStats() agent.Stats {
	var stats agent.Stats
	if value := p.stats.Load(); value != nil {
		stats = value.(agent.Stats)
	}
	return stats
}

func (p *intervalStats) setStats(stats agent.Stats) {
	p.stats.Store(stats)
}

func (p *intervalStats) resetOnce() {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.once = new(sync.Once)
}

func newIntervalStats(interval time.Duration) *intervalStats {
	return &intervalStats{
		once:     new(sync.Once),
		interval: interval,
	}
}
