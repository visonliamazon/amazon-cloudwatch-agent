// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT

package provider

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/assert"

	"github.com/aws/amazon-cloudwatch-agent/extension/agenthealth/handler/stats/agent"
)

func TestIntervalStats(t *testing.T) {
	interval := time.Millisecond
	s := newIntervalStats(interval)
	s.setStats(agent.Stats{
		ThreadCount: aws.Int32(2),
	})
	assert.NotNil(t, s.Stats("").ThreadCount)
	assert.Nil(t, s.Stats("").ThreadCount)
	time.Sleep(interval)
	assert.Eventually(t, func() bool {
		return s.Stats("").ThreadCount != nil
	}, 5*interval, interval)
}
