package main

import (
	"context"
	"github/baishan-development-guizhou/golang-library/log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	logger := log.Configure().WithCaller(true).WithStacktrace(true).WithLevel(log.InfoLevel).Init()
	logger = logger.With("X-Request-ID", "balabala")
	logger.Error("first")
	ctx, _ := log.Associate(context.Background(), logger)
	go func() {
		_, iLogger := log.Context(ctx)
		iLogger.Info("second")
		wg.Done()
	}()
	wg.Wait()
}
