// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package provider

import (
	"log"
	"sync"
	"time"

	"github.com/aws/amazon-cloudwatch-agent/extension/agenthealth/handler/stats/agent"
)

// intervalStats restricts the Stats get function to once
// per interval.
type intervalStats struct {
	interval time.Duration

	lastGet  time.Time
	once     *sync.Once
	onceLock sync.RWMutex

	stats     agent.Stats
	statsLock sync.Mutex
}

var _ agent.StatsProvider = (*intervalStats)(nil)

func (p *intervalStats) Stats(string) agent.Stats {
	p.onceLock.RLock()
	defer p.onceLock.RUnlock()
	var stats agent.Stats
	p.once.Do(func() {
		stats = p.getStats()
		p.lastGet = time.Now()
		time.AfterFunc(p.interval, p.resetOnce)
	})
	return stats
}

func (p *intervalStats) getStats() agent.Stats {
	p.statsLock.Lock()
	defer p.statsLock.Unlock()
	return p.stats
}

func (p *intervalStats) setStats(stats agent.Stats) {
	p.statsLock.Lock()
	defer p.statsLock.Unlock()
	p.stats = stats
}

func (p *intervalStats) resetOnce() {
	log.Printf("time taken to reset: %v | %s", time.Since(p.lastGet), p.lastGet)
	p.onceLock.Lock()
	defer p.onceLock.Unlock()
	p.once = new(sync.Once)
}

func newIntervalStats(interval time.Duration) *intervalStats {
	return &intervalStats{
		once:     new(sync.Once),
		interval: interval,
	}
}
