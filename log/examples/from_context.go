package main

import (
	"context"
	"github.com/baishan-development-guizhou/golang-library/log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	logger := log.Configure().WithOutputEncoder(log.ConsoleOutputEncoder).WithCallerEncoder(log.ShortRoutineCallerEncoder).
		WithStacktrace(false).WithLevel(log.DebugLevel).
		Init()
	logger = logger.With("dsd", "sdd")
	ctx, _ := log.AssociateC(context.Background(), logger)
	go func() {
		iLogger := log.C(ctx)
		iLogger.Info("second")
		wg.Done()
	}()
	wg.Wait()
}
