package handlers

import (
	"fmt"
	"github.com/heptiolabs/healthcheck"
	"github.com/space-devops/mountebank-sidecar/pkg/config"
	"github.com/space-devops/mountebank-sidecar/pkg/logger"
	"github.com/space-devops/mountebank-sidecar/pkg/utils"
	"time"
)

func HealthcheckHandler() *healthcheck.Handler {
	// Create a Handler that we can use to register liveness and readiness checks.
	health := healthcheck.NewHandler()

	// Make sure we can connect to an upstream dependency over TCP in less than
	// 50ms. Run this check asynchronously in the background every 10 seconds
	// instead of every time the /ready or /live endpoints are hit.
	//
	// Async is useful whenever a check is expensive (especially if it causes
	// load on upstream services).
	host := config.GetConfig().Mountebank.Host
	port := config.GetConfig().Mountebank.Health.Port
	_ = config.GetConfig().Mountebank.Health.Path

	upstreamAddr := fmt.Sprintf("%s:%d", host, port)

	msg := fmt.Sprintf("Readiness upstream service %s", upstreamAddr)
	logger.LogInfo(msg, utils.NoCorrelationId)

	health.AddReadinessCheck(
		"upstream-dep-tcp",
		healthcheck.Timeout(func() error {
			// Simulate some work that could take a long time
			time.Sleep(time.Millisecond * 10)
			return nil
		}, 5*time.Millisecond))

	// Implement a custom check with a 50 millisecond timeout.
	health.AddLivenessCheck("custom-check-with-timeout", healthcheck.Timeout(func() error {
		// Simulate some work that could take a long time
		time.Sleep(time.Millisecond * 100)
		return nil
	}, 50*time.Millisecond))

	// Sleep for just a moment to make sure our Async handler had a chance to run
	time.Sleep(500 * time.Millisecond)

	return &health
}
